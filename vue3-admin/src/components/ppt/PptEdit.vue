<script setup>
import {inject, ref, reactive, watch} from 'vue'
import {ElLoading, ElMessage} from 'element-plus'
import {getQiniuToken, postPptEdit} from '../../common/api'
import {IsLockArray} from "../../common/constant"
import axios from "axios";
import localStorageCache from "../../common/localStorage";

const props = defineProps(['rowInfo', 'typeList'])
const emits = defineEmits(['changeRow'])
const showEditDialog = ref(inject('showEditDialog'))

const initEditFormData = {
  id: 0,
  index: -1,
  tid: 0,
  name: "",
  img_url: "",
  desc_content: "",
  sort: 0,
  file_url: "",
  ai_dou: 0,
  is_lock: 1
}
const editFormData = reactive({
  info: Object.assign({}, initEditFormData)
})

const scene = localStorageCache.get('x-scene')
const coverFile = ref([])
const pptFile = ref([])

watch(props.rowInfo, () => {
  if (props.rowInfo.row.id > 0) {
    editFormData.info = Object.assign(editFormData.info, props.rowInfo.row)
  } else {
    editFormData.info = Object.assign({}, initEditFormData)
  }
})

const onCoverFileChange = (uploadFile, uploadFiles) => {
  editFormData.info.img_url = URL.createObjectURL(uploadFile.raw)
  coverFile.value = [uploadFile]
}

const onPptFileChange = (uploadFile, uploadFiles) => {
  editFormData.info.file_url = URL.createObjectURL(uploadFile.raw)
  pptFile.value = [uploadFile]
}

const onEdit = async () => {
  let isLoading = ElLoading.service({
    lock: true,
    text: 'Loading',
  })

  let qiniuToken = null
  if (coverFile.value.length > 0) {
    await getQiniuToken(scene.upload.ppt_cover, 'png').then(res => {
      qiniuToken = res.data[0]
    })

    let qiniuFormData = new FormData()
    qiniuFormData.append("key", qiniuToken.key)
    qiniuFormData.append("token", qiniuToken.token)
    qiniuFormData.append("file", coverFile.value[0].raw, coverFile.value[0].raw.name)
    let axiosObj = axios.create({
      baseURL: "", // 所有的请求地址前缀部分
      timeout: 60000, // 请求超时时间毫秒
      withCredentials: false, // 异步请求携带cookie
      headers: {},
      data: {},
    })
    await axiosObj.post(qiniuToken.up_serve, qiniuFormData).then(res => {
      coverFile.value = []
      editFormData.info.img_url = res.data.key
    })
  }

  if (pptFile.value.length > 0) {
    await getQiniuToken(scene.upload.ppt_file, 'pptx').then(res => {
      qiniuToken = res.data[0]
    })

    let qiniuFormData = new FormData()
    qiniuFormData.append("key", qiniuToken.key)
    qiniuFormData.append("token", qiniuToken.token)
    qiniuFormData.append("file", pptFile.value[0].raw, pptFile.value[0].raw.name)
    let axiosObj = axios.create({
      baseURL: "", // 所有的请求地址前缀部分
      timeout: 60000, // 请求超时时间毫秒
      withCredentials: false, // 异步请求携带cookie
      headers: {},
      data: {},
    })
    await axiosObj.post(qiniuToken.up_serve, qiniuFormData).then(res => {
      pptFile.value = []
      editFormData.info.file_url = res.data.key
    })
  }

  postPptEdit(editFormData.info).then(res => {
    isLoading.close()
    if (res.code === 1) {
      let row = Object.assign({}, editFormData.info, res.data)
      props.typeList.forEach((item) => {
        if (item.id === editFormData.info.tid) {
          row.ppt_type = item
          return
        }
      })
      row.img_url = editFormData.info.img_url
      row.file_url = editFormData.info.file_url
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
      <el-form-item label="标题:">
        <el-input v-model="editFormData.info.name" placeholder="Please input name"/>
      </el-form-item>
      <el-form-item label="分类:">
        <el-select v-model="editFormData.info.tid" placeholder="Please Select is type">
          <el-option v-for="item in props.typeList" :value="item.id" :label="item.name"/>
        </el-select>
      </el-form-item>
      <el-form-item label="所需爱享豆:">
        <el-input v-model="editFormData.info.ai_dou" placeholder="Please input ai_dou"/>
      </el-form-item>
      <el-form-item label="封面：">
        <el-upload
            class="my-avatar-uploader"
            :auto-upload="false"
            :show-file-list="false"
            :on-change="onCoverFileChange"
            accept="image/*"
        >
          <img v-if="editFormData.info.img_url" :src="editFormData.info.img_url" class="my-avatar"/>
          <el-icon v-else class="my-avatar-uploader-icon">
            <Plus/>
          </el-icon>
        </el-upload>
      </el-form-item>
      <el-form-item label="ppt文件：">
        <el-upload
            :auto-upload="false"
            :show-file-list="false"
            :on-change="onPptFileChange"
        >
          <template #trigger>
            <el-input type="primary" v-model="editFormData.info.file_url" disabled/>
          </template>
        </el-upload>
      </el-form-item>
      <el-form-item label="内容简介:">
        <el-input
            v-model="editFormData.info.desc_content"
            :rows="3"
            type="textarea"
            placeholder="Please input desc_content"
        />
      </el-form-item>
      <el-form-item label="排序值:">
        <el-input v-model="editFormData.info.sort" placeholder="Please input sort"/>
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
