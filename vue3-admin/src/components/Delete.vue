<script setup>
import {inject, ref} from "vue";
import {postDel} from "../common/api";
import {ElMessageBox} from "element-plus";

const props = defineProps(['row'])
const mMark = ref(inject('mMark'))
const tableData = ref(inject('tableData'))

const handleDel = (row, index) => {
  ElMessageBox.confirm('确定删除该用户？').then(() => {
    postDel(row.id, mMark.value).then(res => {
      tableData.value.splice(index, 1)
    })
  })
}

</script>
<template>
  <el-button size="small" type="danger" @click="handleDel(props.row, props.index)">删除</el-button>
</template>
