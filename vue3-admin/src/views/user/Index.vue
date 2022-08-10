<script setup>
import {ref, onMounted, watch, provide, reactive} from "vue"
import {getUserList, postUserDel, postUserLock} from "../../common/api"
import {ElMessageBox} from "element-plus"
import localStorageCache from "../../common/localStorage";
import OpenVip from '../../components/user/OpenVip.vue'
import RechargeAiDou from '../../components/user/RechargeAiDou.vue'
import {IsLockArray} from "../../common/constant";

const total = ref(0)
const tableData = ref([])
const currentPage = ref(1)
const pageSize = ref(10)
const searchFormData = reactive({
  id: "",
  user_name: "",
  nick_name: "",
  is_lock: "",
  created_at_start: "",
  created_at_end: "",
  total: 0,
  page: currentPage.value,
  page_size: pageSize.value,
})

const showOpenVipDialog = ref(false)
provide('showOpenVipDialog', showOpenVipDialog)

const showRechargeAiDouDialog = ref(false)
provide('showRechargeAiDouDialog', showRechargeAiDouDialog)

const userRow = reactive({
  info: {
    id: "",
    user_name: "",
    nick_name: "",
  }
})

const vipConfig = localStorageCache.get("x-vip-config")

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
  getUserList(searchFormData).then((res) => {
    total.value = res.data.total
    tableData.value = res.data.list
  })
}

const onSearch = (res) => {
  total.value = 0
  currentPage.value = 1
  tableList()
}

const handleLock = (index, row) => {
  let is_lock = row.is_lock === 1 ? 2 : 1
  postUserLock(row.id, is_lock).then(res => {
    row.is_lock = is_lock
  })
}

const handleDelete = (index, row) => {
  ElMessageBox.confirm('确定删除该用户？').then(() => {
    postUserDel(row.id).then(res => {
      tableData.value.splice(index, 1)
    })
  })
}

const handleOpenVip = (index, row) => {
  row.index = index
  userRow.info = row
  showOpenVipDialog.value = true
}

const handleRechargeAiDou = (index, row) => {
  row.index = index
  userRow.info = row
  showRechargeAiDouDialog.value = true
}

const onChangeRow = row => {
  tableData.value[row.index] = row
}

</script>

<template>
  <el-container class="my-table-container">
    <el-form class="my-form">
      <el-form-item label="Uid:">
        <el-input v-model="searchFormData.id" placeholder="Please input uid"/>
      </el-form-item>
      <el-form-item label="账号:">
        <el-input v-model="searchFormData.user_name" placeholder="Please input user_name"/>
      </el-form-item>
      <el-form-item label="昵称:">
        <el-input v-model="searchFormData.nick_name" placeholder="Please input user_name"/>
      </el-form-item>
      <el-form-item label="是否锁定:">
        <el-select v-model="searchFormData.is_lock" placeholder="Please Select is lock">
          <el-option v-for="item in IsLockArray" :value="item.value" :label="item.label"/>
        </el-select>
      </el-form-item>
      <el-form-item label="注册时间范围:">
        <el-date-picker
            v-model="searchFormData.created_at_start"
            type="datetime"
            placeholder="Please Select start time"
            value-format="YYYY-MM-DD HH:mm:ss"
        />
        <el-date-picker
            v-model="searchFormData.created_at_end"
            type="datetime"
            placeholder="Please Select end time"
            value-format="YYYY-MM-DD HH:mm:ss"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSearch">搜索</el-button>
      </el-form-item>
    </el-form>

    <el-table :data="tableData" height="70vh" style="width: 100%">
      <el-table-column prop="id" label="ID"/>
      <el-table-column label="头像">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <el-avatar shape="square" :size="100" fit="fill" :src="scope.row.head_img"/>
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="user_name" label="账号"/>
      <el-table-column prop="nick_name" label="昵称"/>
      <el-table-column prop="nick_name" label="昵称"/>
      <el-table-column prop="vip" label="vip等级">
        <template #default="scope">
          <div>{{ vipConfig[scope.row.vip].Name }}</div>
        </template>
      </el-table-column>
      <el-table-column prop="user_contact.superior" label="上级UID"/>
      <el-table-column prop="wallet" label="钱包">
        <template #default="scope">
          <div>总收益: {{ scope.row.wallet.total_balance }}</div>
          <div>余额: {{ scope.row.wallet.balance }}</div>
        </template>
      </el-table-column>
      <el-table-column prop="wallet" label="钱包">
        <template #default="scope">
          <div>总爱豆: {{ scope.row.ai_dou.total_ai_dou }}</div>
          <div>余额: {{ scope.row.ai_dou.ai_dou }}</div>
        </template>
      </el-table-column>
      <el-table-column prop="vip_end_at" label="VIP到期时间"/>
      <el-table-column prop="created_at" label="注册时间"/>
      <el-table-column prop="" label="操作">
        <template #default="scope">
          <el-row>
            <el-button size="small" type="primary" @click="handleOpenVip(scope.$index, scope.row)">开通VIP</el-button>
          </el-row>
          <el-row>
            <el-button size="small" type="primary" @click="handleRechargeAiDou(scope.$index, scope.row)">充值爱豆
            </el-button>
          </el-row>
          <el-row>
            <el-button size="small" type="primary" @click="handleLock(scope.$index, scope.row)">
              {{ scope.row.is_lock === 1 ? "锁定" : "解除锁定" }}
            </el-button>
          </el-row>
          <el-row>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
          </el-row>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination background layout="total, sizes, prev, pager, next, jumper" v-model:page-size="pageSize"
                   v-model:currentPage="currentPage" :page-sizes="[10, 30, 50, 100]" :total="total"/>
  </el-container>
  <OpenVip :userRow="userRow" @changeRow="onChangeRow"/>
  <RechargeAiDou :userRow="userRow" @changeRow="onChangeRow"/>
</template>
