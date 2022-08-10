<script setup>
import {onMounted, reactive, ref} from "vue";
import {getQiniuToken, getSystemInfo, postSystemInfoEdit} from "../common/api";
import localStorageCache from "../common/localStorage";
import {ElLoading} from "element-plus";

const miniCheck = reactive({
  data: {
    mini_check: "1",
    system_close: "1",
    system_close_content: "",
    customer_service_wx: "",
  }
})
const fileObj = ref()
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
const isLoading = ref()
const isSubmitForm = ref(false)

onMounted(() => {
  getSystemInfo().then(res => {
    if (res.code === 1) {
      miniCheck.data = res.data
    }
  })
})

const submitForm = () => {
  if (isSubmitForm.value) {
    return
  }
  isLoading.value = ElLoading.service({
    lock: true,
    text: 'Loading',
  })
  isSubmitForm.value = true
  if (fileList.value.length > 0) {
    getQiniuToken(scene.upload.customer_service_wx, 'png').then(res => {
      qiniuTokenRes.info = res.data[0]
      fileObj.value.submit()
    })
  } else {
    postSystemInfoEdit(miniCheck.data).then(res => {
      isSubmitForm.value = false
      isLoading.value.close()
    })
  }
}

const onSuccess = (response, uploadFile) => {
  miniCheck.data.customer_service_wx = response.key
  if (isSubmitForm) {
    postSystemInfoEdit(miniCheck.data).then(res => {
      isSubmitForm.value = false
      isLoading.value.close()
      fileList.value = []
    })
  }
}

const onChange = (uploadFile, uploadFiles) => {
  miniCheck.data.customer_service_wx = URL.createObjectURL(uploadFile.raw)
  fileList.value = [uploadFile]
}

</script>

<template>
  <el-form>
    <el-form-item label="小程序是否审核中：">
      <el-radio-group v-model="miniCheck.data.mini_check" class="ml-4">
        <el-radio label="1" size="large">审核中</el-radio>
        <el-radio label="2" size="large">审核通过</el-radio>
      </el-radio-group>
    </el-form-item>
    <el-form-item label="是否关站：">
      <el-radio-group v-model="miniCheck.data.system_close" class="ml-4">
        <el-radio label="1" size="large">开站</el-radio>
        <el-radio label="2" size="large">关站</el-radio>
      </el-radio-group>
    </el-form-item>
    <el-form-item label="关站公告：" class="system_close_content">
      <el-input v-model="miniCheck.data.system_close_content" type="textarea"/>
    </el-form-item>
    <el-form-item label="客服二维码：">
      <el-upload
          ref="fileObj"
          class="my-avatar-uploader"
          :file-list="fileList"
          :action="qiniuTokenRes.info.up_serve"
          :data="qiniuTokenRes.info"
          :auto-upload="false"
          :show-file-list="false"
          :on-success="onSuccess"
          :on-change="onChange"
          accept="image/*"
      >
        <img v-if="miniCheck.data.customer_service_wx" :src="miniCheck.data.customer_service_wx" class="my-avatar"/>
        <el-icon v-else class="my-avatar-uploader-icon">
          <Plus/>
        </el-icon>
      </el-upload>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submitForm()">提交</el-button>
    </el-form-item>
  </el-form>
</template>
<style lang="less">
.system_close_content textarea {
  width: 15rem;
}
</style>
