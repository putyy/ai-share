<script setup>
import {inject, ref, reactive, watch, onMounted} from 'vue'
import {ElLoading, ElMessage} from 'element-plus'
import {getMenuGroupList, postMenuEdit} from '../../common/api'
import localStorageCache from "../../common/localStorage"
import {IsLockArray, MenuClickType} from "../../common/constant"

const props = defineProps(['rowInfo'])
const emits = defineEmits(['changeRow'])
const showEditDialog = ref(inject('showEditDialog'))
const vipConfig = localStorageCache.get("x-vip-config")

const initEditFormData = {
  id: 0,
  gid: 0,
  index: -1,
  name: "",
  use_vip: 0,
  click_type: 1,
  click_func: "",
  path: "",
  app_id: "",
  extra_data: "",
  env_version: "",
  short_link: "",
  sort: 0,
  is_lock: 1,
}

const editFormData = reactive({
  info: Object.assign({}, initEditFormData)
})

const groupList = ref([])

onMounted(() => {
  getMenuGroupList({page_size: 100}).then(res => {
    groupList.value = res.data.list
    groupList.value.unshift({"id": 0, "name": "请选择"})
  })
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

  postMenuEdit(editFormData.info).then(res => {
    isLoading.close()
    if (res.code === 1) {
      let row = Object.assign({}, editFormData.info, res.data)
      groupList.value.forEach((item) => {
        if (item.id === editFormData.info.gid) {
          row.menu_group = item
          return
        }
      })
      showEditDialog.value = false
      emits('changeRow', row)

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
    <el-form class="my-edit">
      <el-form-item label="分类:">
        <el-select v-model="editFormData.info.gid" placeholder="Please Select is gid">
          <el-option v-for="item in groupList" :label="item.name" :value="item.id">{{ item.name }}</el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="菜单名称">
        <el-input v-model="editFormData.info.name" placeholder="Please input name"/>
      </el-form-item>
      <el-form-item label="使用权限限制:">
        <el-select v-model="editFormData.info.use_vip" placeholder="Please Select is use_vip">
          <el-option v-for="item in vipConfig" :value="item.Level" :label="item.Name"/>
        </el-select>
      </el-form-item>
      <el-form-item label="打开的页面路径">
        <el-input v-model="editFormData.info.path" placeholder="Please input path"/>
      </el-form-item>
      <el-form-item label="小程序appid">
        <el-input v-model="editFormData.info.app_id" placeholder="Please input app_id"/>
      </el-form-item>
      <el-form-item label="需要传递给目标小程序的数据(json格式):">
        <el-input
            v-model="editFormData.info.extra_data"
            :rows="3"
            type="textarea"
            placeholder="Please input"
        />
      </el-form-item>
      <el-form-item label="要打开的小程序版本">
        <el-input v-model="editFormData.info.env_version" placeholder="Please input env_version"/>
      </el-form-item>
      <el-form-item label="小程序链接">
        <el-input v-model="editFormData.info.short_link" placeholder="Please input short_link"/>
      </el-form-item>
      <el-form-item label="事件类型:">
        <el-select v-model="editFormData.info.click_type" placeholder="Please Select is click_type">
          <el-option v-for="item in MenuClickType" :value="item.value" :label="item.label"/>
        </el-select>
      </el-form-item>
      <el-form-item label="函数标识(事件类型为函数时用到):">
        <el-input v-model="editFormData.info.click_func" placeholder="Please input click_func"/>
      </el-form-item>
      <el-form-item label="排序">
        <el-input v-model="editFormData.info.sort" placeholder="Please input sort"/>
      </el-form-item>
      <el-form-item label="是否锁定">
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
