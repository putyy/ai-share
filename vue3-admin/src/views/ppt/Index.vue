<script setup>
import {ref, onMounted, watch, reactive, provide} from "vue"
import {getPptList, getPptTypeList} from "../../common/api"
import PptEdit from '../../components/ppt/PptEdit.vue'
import PptContentEdit from '../../components/ppt/PptContentEdit.vue'
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
  tid: "",
  total: 0,
  page: currentPage.value,
  page_size: pageSize.value,
})

const showEditDialog = ref(false)
provide('showEditDialog', showEditDialog)

const showContentEditDialog = ref(false)
provide('showContentEditDialog', showContentEditDialog)

const pid = ref(false)
provide('pid', pid)

const typeList = ref([])

const rowInfo = reactive({
  row: {}
})

onMounted(() => {
  tableList()
  getPptTypeList({page_size: 100}).then(res => {
    typeList.value = res.data.list
  })
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
  getPptList(searchFormData).then((res) => {
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

const handleContentEdit = (index, row) => {
  pid.value = row.id
  showContentEditDialog.value = true
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
    <div class="my-table-top">
      <el-form class="my-form">
        <el-form-item label="关键字:">
          <el-input v-model="searchFormData.keywords" placeholder="Please input"/>
        </el-form-item>
        <el-form-item label="分类:">
          <el-select v-model="searchFormData.tid" placeholder="Please Select is type">
            <el-option v-for="item in typeList" :value="item.id" :label="item.name"/>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSearch">搜索</el-button>
        </el-form-item>
      </el-form>
      <div class="actions">
        <el-button type="primary" @click="handleAdd">新增</el-button>
      </div>
    </div>

    <el-table :data="tableData" height="70vh" style="width: 100%">
      <el-table-column prop="id" label="ID"/>
      <el-table-column prop="ppt_type.name" label="分类"/>
      <el-table-column label="封面">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <el-avatar shape="square" :size="100" fit="fill" :src="scope.row.img_url"/>
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="desc_content" label="内容"/>
      <el-table-column prop="ai_dou" label="所需爱享豆"/>
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
          <el-row>
            <el-button size="small" type="primary" @click="handleContentEdit(scope.$index, scope.row)">编辑内容</el-button>
          </el-row>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination background layout="total, sizes, prev, pager, next, jumper" v-model:page-size="pageSize"
                   v-model:currentPage="currentPage" :page-sizes="[10, 30, 50, 100]" :total="total"/>

  </el-container>

  <PptEdit :rowInfo="rowInfo" :typeList="typeList" @changeRow="onChangeRow"/>
  <PptContentEdit />
</template>
