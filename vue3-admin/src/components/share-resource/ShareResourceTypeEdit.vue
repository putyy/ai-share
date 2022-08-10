<script setup>
import {inject, ref, reactive, watch} from 'vue'
import {ElLoading, ElMessage} from 'element-plus'
import {postShareResourceTypeEdit} from '../../common/api'
import {IsLockArray} from "../../common/constant"

const props = defineProps(['rowInfo'])
const emits = defineEmits(['changeRow'])
const showEditDialog = ref(inject('showEditDialog'))

const initEditFormData = {
  id: 0,
  index: -1,
  name: "",
  sort: 0,
  is_lock: 1,
}
const editFormData = reactive({
  info: Object.assign({}, initEditFormData)
})

watch(props.rowInfo, () => {
  if (props.rowInfo.row.id > 0) {
    editFormData.info = Object.assign(editFormData.info, props.rowInfo.row)
  } else {
    editFormData.info = Object.assign({}, initEditFormData)
  }
})

const onEdit = async () => {
  let isLoading = ElLoading.service({
    lock: true,
    text: 'Loading',
  })

  postShareResourceTypeEdit(editFormData.info).then(res => {
    isLoading.close()
    if (res.code === 1) {
      let row = Object.assign({}, editFormData.info, res.data)
      emits('changeRow', row)
      showEditDialog.value = false
      ElMessage({
        message: 'success',
        type: 'success',
      })
    }
  })
}

</script>

<template>
  <el-dialog v-model="showEditDialog" title="编辑">
    <el-form label-width="5rem" class="my-edit">
      <el-form-item label="分类名:">
        <el-input v-model="editFormData.info.name" placeholder="Please name"/>
      </el-form-item>
      <el-form-item label="排序值:">
        <el-input v-model="editFormData.info.sort" placeholder="Please sort"/>
      </el-form-item>
      <el-form-item label="是否锁定:">
        <el-select v-model="editFormData.info.is_lock" placeholder="Please Select is lock">
          <el-option v-for="item in IsLockArray" :value="item.value" :label="item.label"/>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onEdit()">提交</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>
