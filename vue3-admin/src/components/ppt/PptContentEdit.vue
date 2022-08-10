<script setup>
import {inject, ref, reactive, watch} from 'vue'
import {ElLoading, ElMessage, ElMessageBox} from 'element-plus'
import {getPptContent, getQiniuToken, postPptContentEdit, postUserDel} from '../../common/api'
import axios from "axios";
import localStorageCache from "../../common/localStorage";
import {hexMD5} from "../../common/md5";

const showContentEditDialog = ref(inject('showContentEditDialog'))
const pid = ref(inject('pid'))
const scene = localStorageCache.get('x-scene')
const initEditFormData = {
  id: 0,
  pid: 0,
  img_url: "",
  content: "",
  imgIndex: -1,
}
const contents = ref([])
const coverFile = ref([])

watch(pid, () => {
  contents.value = []
  coverFile.value = []
  getPptContent(pid.value).then(res=>{
    if (res.data.content.length > 0) {
      contents.value = res.data.content
    }
  })
})

const onCoverFileChange = (uploadFile, index) => {
  contents.value[index].img_url = URL.createObjectURL(uploadFile.raw)
  contents.value[index].imgIndex = coverFile.value.push(uploadFile) - 1
}

const addContent = () => {
  contents.value.push(Object.assign({}, initEditFormData))
}

const moveItem = (index, mark) => {
  let vIndex = 0
  switch (mark) {
    case "del":
      ElMessageBox.confirm('确定删除？')
          .then(() => {
            contents.value.splice(index, 1)
          })
          .catch(() => {
          })
      return
    case "up":
      vIndex = index - 1
      break
    case "down":
      vIndex = index + 1
      break
    default:
      return
  }
  if (contents.value[vIndex] === undefined) {
    return;
  }
  let temp = contents.value[index]
  contents.value[index] = contents.value[vIndex]
  contents.value[vIndex] = temp
}

const onEdit = async () => {

  if (contents.value.length < 0) {
    ElMessage({
      message: "内容必须",
      type: 'warning',
    })
    return
  }

  let isLoading = ElLoading.service({
    lock: true,
    text: 'Loading',
  })

  let qiniuToken = null
  let scenes = []
  contents.value.forEach(v => {
      if (v.imgIndex !== undefined && v.imgIndex !== -1 && coverFile.value[v.imgIndex]) {
        scenes.push(scene.upload.ppt_cover)
      }
  })

  if (scenes.length > 0) {
    await getQiniuToken(scenes.join('-'), 'png').then(res => {
      qiniuToken = res.data
    })
    let i = 0
    let axiosObj = axios.create({
      baseURL: "", // 所有的请求地址前缀部分
      timeout: 60000, // 请求超时时间毫秒
      withCredentials: false, // 异步请求携带cookie
      headers: {},
      data: {},
    })
    for (let v of contents.value) {
      if (v.imgIndex !== undefined && v.imgIndex !== -1 && coverFile.value[v.imgIndex]) {
        let qiniuFormData = new FormData()
        qiniuFormData.append("key", qiniuToken[i].key)
        qiniuFormData.append("token", qiniuToken[i].token)
        qiniuFormData.append("file", coverFile.value[v.imgIndex].raw, coverFile.value[v.imgIndex].raw.name)
        await axiosObj.post(qiniuToken[i].up_serve, qiniuFormData).then(res => {
          v.img_url = res.data.key
        })
        i++
      }
    }
    coverFile.value = []
  }

  let formData = new FormData()
  formData.append("id", pid.value)

  for (let v of contents.value) {
    formData.append("images[]", v.img_url)
    formData.append("contents[]", v.content)
  }
  postPptContentEdit(formData).then(res=>{
    showContentEditDialog.value = false
    isLoading.close()
  })
}

</script>

<template>
  <el-dialog v-model="showContentEditDialog" title="编辑">
    <el-form label-width="5rem" class="my-edit">
      <el-form-item v-for="(item, index) in contents" label="内容段" class="content-item">
        <el-upload
            class="my-avatar-uploader"
            :auto-upload="false"
            :show-file-list="false"
            :on-change="(file, fileList)=>{onCoverFileChange(file, index)}"
            :data="item"
            v-model:file-list="coverFile"
            accept="image/*"
        >
          <img v-if="item.img_url" :src="item.img_url" class="c-cover"/>
          <el-icon v-else class="c-cover-icon">
            <Plus/>
          </el-icon>
        </el-upload>
        <el-input
            v-model="item.content"
            :rows="4"
            type="textarea"
            placeholder="Please desc_content"
        />
        <div class="icons">
          <el-icon @click="moveItem(index, 'del')">
            <CloseBold/>
          </el-icon>
          <el-icon @click="moveItem(index, 'up')">
            <ArrowUpBold/>
          </el-icon>
          <el-icon @click="moveItem(index, 'down')">
            <ArrowDownBold/>
          </el-icon>
        </div>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="addContent()">添加内容段</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onEdit()">提交</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>
<style lang="less">
.content-item {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  //border: 1px solid #ccc;
  .el-textarea {
    max-width: 70%;
  }
}

.c-cover {
  width: 6rem;
  height: 6rem;
  display: block;
}

.c-cover-icon {
  font-size: 1.8rem !important;
  color: #8c939d !important;
  width: 6rem !important;
  height: 6rem !important;
  text-align: center !important;
}

.icons {
  display: flex;
  flex-direction: column;
  color: red;

  .el-icon {
    padding: 0.5rem;
  }
}
</style>
