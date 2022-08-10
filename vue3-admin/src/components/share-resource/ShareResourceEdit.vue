<script setup>
import {inject, ref, reactive, watch, onMounted} from 'vue'
import {ElLoading, ElMessage} from 'element-plus'
import {getQiniuToken, getShareResourceTypeList, postShareResourceEdit} from '../../common/api'
import localStorageCache from "../../common/localStorage"
import axios from "axios"
import {IsLockArray} from "../../common/constant"

const props = defineProps(['rowInfo'])
const emits = defineEmits(['changeRow'])
const showEditDialog = ref(inject('showEditDialog'))
const fileList = ref([])
const scene = localStorageCache.get('x-scene')
const qiniuTokenRes = reactive({
  info: {
    domain: '',
    key: '',
    scene: '',
    token: '',
    up_serve: '',
  }
})

const initEditFormData = {
  id: 0,
  index: -1,
  tid: 0,
  img_url: "",
  content: "",
  sort: 0,
  is_lock: 1,
}
const editFormData = reactive({
  info: Object.assign({}, initEditFormData)
})

const typeList = ref([])

onMounted(() => {
  getShareResourceTypeList({page_size: 100}).then(res => {
    typeList.value = res.data.list
    typeList.value.unshift({"id": 0, "name": "请选择"})
  })
})

watch(props.rowInfo, () => {
  if (props.rowInfo.row.id > 0) {
    editFormData.info = Object.assign(editFormData.info, props.rowInfo.row)
  } else {
    editFormData.info = Object.assign({}, initEditFormData)
  }
})

const onChange = (uploadFile, uploadFiles) => {
  editFormData.info.img_url = URL.createObjectURL(uploadFile.raw)
  fileList.value = [uploadFile]
}

const onEdit = async () => {
  let isLoading = ElLoading.service({
    lock: true,
    text: 'Loading',
  })

  if (fileList.value.length > 0) {
    await getQiniuToken(scene.upload.customer_service_wx, 'png').then(res => {
      qiniuTokenRes.info = res.data[0]
    })

    let qiniuFormData = new FormData()
    qiniuFormData.append("key", qiniuTokenRes.info.key)
    qiniuFormData.append("token", qiniuTokenRes.info.token)
    qiniuFormData.append("file", fileList.value[0].raw, fileList.value[0].raw.name)
    let axiosObj = axios.create({
      baseURL: "", // 所有的请求地址前缀部分
      timeout: 60000, // 请求超时时间毫秒
      withCredentials: false, // 异步请求携带cookie
      headers: {},
      data: {},
    })
    await axiosObj.post(qiniuTokenRes.info.up_serve, qiniuFormData).then(res => {
      editFormData.info.img_url = res.data.key
    })
  }

  postShareResourceEdit(editFormData.info).then(res => {
    if (res.code === 1) {
      let row = Object.assign({}, editFormData.info, res.data)
      typeList.value.forEach((item) => {
        if (item.id === editFormData.info.tid) {
          row.share_resource_type = item
          return
        }
      })
      row.img_url = editFormData.info.img_url
      showEditDialog.value = false
      emits('changeRow', row)

      isLoading.close()
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
      <el-form-item label="分类:">
        <el-select v-model="editFormData.info.tid" placeholder="Please Select is lock">
          <el-option v-for="item in typeList" :label="item.name" :value="item.id">{{ item.name }}</el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="客服二维码：">
        <el-upload
            class="my-avatar-uploader"
            :auto-upload="false"
            :show-file-list="false"
            :on-change="onChange"
            accept="image/*"
        >
          <img v-if="editFormData.info.img_url" :src="editFormData.info.img_url" class="my-avatar"/>
          <el-icon v-else class="my-avatar-uploader-icon">
            <Plus/>
          </el-icon>
        </el-upload>
      </el-form-item>
      <el-form-item label="是否锁定:">
        <el-select v-model="editFormData.info.is_lock" placeholder="Please Select is lock">
          <el-option v-for="item in IsLockArray" :value="item.value" :label="item.label"/>
        </el-select>
      </el-form-item>
      <el-form-item label="内容:">
        <el-input
            v-model="editFormData.info.content"
            :rows="4"
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
