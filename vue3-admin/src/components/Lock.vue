<script setup>
import {inject, ref} from "vue";
import {postLock} from "../common/api";
import {ElMessage} from "element-plus";

const props = defineProps(['row'])
const mMark = ref(inject('mMark'))

const handleLock = (row) => {
  let is_lock = row.is_lock === 1 ? 2 : 1
  postLock(row.id, is_lock, mMark.value).then(res => {
    row.is_lock = is_lock
    ElMessage({
      message: 'success',
      type: 'success',
    })
  })
}

</script>
<template>
  <el-button size="small" type="primary" @click="handleLock(props.row)">{{ props.row.is_lock === 1 ? "锁定" : "解除锁定" }}</el-button>
</template>
