<script setup>
import {ref, onMounted, watch, reactive, provide} from "vue"
import {getMenuGroupList} from "../../common/api";
import MenuGroupEdit from '../../components/user/MenuGroupEdit.vue'
import Lock from '../../components/Lock.vue'
import Delete from '../../components/Delete.vue'

const total = ref(0)
const tableData = ref([])
provide('tableData', tableData)

const currentPage = ref(1)
const pageSize = ref(10)
const mMark = ref("")
provide('mMark', mMark)

const searchFormData = reactive({
  keywords: "",
  total: 0,
  page: currentPage.value,
  page_size: pageSize.value,
})

const showEditDialog = ref(false)
provide('showEditDialog', showEditDialog)

const rowInfo = reactive({
  row: {}
})

onMounted(() => {
  tableList()
})

watch(currentPage, () => {
  tableList()
})

watch(pageSize, () => {
  tableList()
})

const tableList = () => {
  searchFormData.page = currentPage.value
  searchFormData.page_size = pageSize.value
  searchFormData.total = total.value
  getMenuGroupList(searchFormData).then((res) => {
    total.value = res.data.total
    tableData.value = res.data.list
    mMark.value = res.data.m_mark
  })
}

const onSearch = (res) => {
  total.value = 0
  currentPage.value = 1
  tableList()
}

const handleAdd = (res) => {
  rowInfo.row = {}
  showEditDialog.value = true
}

const handleEdit = (index, row) => {
  row.index = index
  rowInfo.row = row
  showEditDialog.value = true
}

const onChangeRow = row => {
  if (row.index >= 0) {
    tableData.value[row.index] = row
  } else {
    tableData.value.unshift(row)
  }
}

</script>

<template>
  <el-container class="my-table-container">

    <el-table :data="tableData" height="70vh" style="width: 100%">
      <el-table-column prop="id" label="ID"/>
      <el-table-column prop="name" label="分类名"/>
      <el-table-column prop="sort" label="排序"/>
      <el-table-column prop="created_at" label="创建时间"/>
      <el-table-column prop="updated_at" label="更新时间"/>
      <el-table-column prop="" label="操作">
        <template #default="scope">
          <el-row>
            <el-button size="small" type="primary" @click="handleEdit(scope.$index, scope.row)">修改</el-button>
          </el-row>
          <el-row>
            <Lock :row="scope.row"/>
          </el-row>
          <el-row>
            <Delete :row="scope.row" :index="scope.$index"/>
          </el-row>
        </template>
      </el-table-column>
    </el-table>

<!--    <el-pagination background layout="total, sizes, prev, pager, next, jumper" v-model:page-size="pageSize"-->
<!--                   v-model:currentPage="currentPage" :page-sizes="[10, 30, 50, 100]" :total="total"/>-->

  </el-container>

  <MenuGroupEdit :rowInfo="rowInfo" @changeRow="onChangeRow"/>

</template>
