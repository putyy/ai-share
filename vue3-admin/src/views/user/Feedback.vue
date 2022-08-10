<script setup>
import {ref, onMounted, watch, reactive} from "vue"
import {getFeedbackList, postFeedbackEdit} from "../../common/api"
import {ElLoading} from "element-plus";

const total = ref(0)
const tableData = ref([])
const currentPage = ref(1)
const pageSize = ref(10)
const searchFormData = reactive({
  keywords: "",
  total: 0,
  page: currentPage.value,
  page_size: pageSize.value,
})
const showEditDialog = ref(false)
const rowInfo = reactive({
  row: {
    id: 0,
    remark: ""
  }
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
  getFeedbackList(searchFormData).then((res) => {
    total.value = res.data.total
    tableData.value = res.data.list
  })
}

const onSearch = (res) => {
  total.value = 0
  currentPage.value = 1
  tableList()
}

const handleEdit = (index, row) => {
  row.index = index
  rowInfo.row = row
  showEditDialog.value = true
}

const onEdit = () => {
  let loading = ElLoading.service({
    lock: true,
    text: 'Loading',
  })
  postFeedbackEdit(rowInfo.row).then(res => {
    tableData.value[rowInfo.row.index] = rowInfo.row
    showEditDialog.value = false
    loading.close()
  })
}

</script>

<template>
  <el-container class="my-table-container">
    <el-form class="my-form">
      <el-form-item label="关键字:">
        <el-input v-model="searchFormData.keywords" placeholder="Please input keywords"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSearch">搜索</el-button>
      </el-form-item>
    </el-form>

    <el-table :data="tableData" style="width: 100%" height="70vh">
      <el-table-column prop="id" label="ID"/>
      <el-table-column prop="uid" label="Uid"/>
      <el-table-column prop="content" label="内容"/>
      <el-table-column prop="remark" label="备注"/>
      <el-table-column prop="created_at" label="创建时间"/>
      <el-table-column prop="updated_at" label="更新时间"/>
      <el-table-column prop="" label="操作">
        <template #default="scope">
          <el-row>
            <el-button size="small" type="primary" @click="handleEdit(scope.$index, scope.row)">备注</el-button>
          </el-row>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination background layout="total, sizes, prev, pager, next, jumper" v-model:page-size="pageSize"
                   v-model:currentPage="currentPage" :page-sizes="[10, 30, 50, 100]" :total="total"/>
  </el-container>
  <el-dialog v-model="showEditDialog" title="备注">
    <el-form label-width="5rem">
      <el-form-item label="备注:">
        <el-input
            v-model="rowInfo.row.remark"
            :rows="2"
            type="textarea"
            placeholder="Please input"
        />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onEdit()">提交</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>
