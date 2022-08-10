<script setup>
import {inject, reactive, ref} from "vue";
import localStorageCache from "../../common/localStorage";
import {ElLoading, ElMessage} from "element-plus";
import {postUserOpenVip} from "../../common/api";

const props = defineProps(['userRow'])
const emits = defineEmits(['changeRow'])
const showOpenVipDialog = ref(inject('showOpenVipDialog'))
const vipConfig = localStorageCache.get("x-vip-config")
const oData = reactive({
  uid: 0,
  level: 0,
  show_price: 0,
  price: 0,
  remark: "",
});

const onOpenVip = res => {
  oData.uid = props.userRow.info.id
  oData.price = parseInt(oData.show_price * 100)
  let loading = ElLoading.service({
    lock: true,
    text: 'Loading',
  })

  postUserOpenVip(oData).then(res => {
    loading.close()
    if (res.code === 1) {
      props.userRow.info.vip = oData.level
      props.userRow.info.vip_end_at = res.data.vip_end_at
      emits('changeRow', props.userRow.info)
      showOpenVipDialog.value = false
      ElMessage({
        message: '开通成功',
        type: 'success',
      })
    }
  })
}
</script>

<template>
  <el-dialog v-model="showOpenVipDialog" title="开通VIP">
    <el-form label-width="5rem" class="my-edit">
      <el-form-item label="用户信息:">
        {{ userRow.info.nick_name }}: {{ userRow.info.user_name }}
      </el-form-item>
      <el-form-item label="当前等级:">
        {{ vipConfig[userRow.info.vip].Name }}
      </el-form-item>
      <el-form-item label="开通等级:">
        <el-select v-model="oData.level" placeholder="选择VIP等级">
          <el-option v-for="item in vipConfig" :label="item.Name" :value="item.Level">
            <span>等级:{{ item.Name }} (价格:{{ item.Price }} 收益:{{
                item.Profit
              }}% 时长:{{ item.Length === -1 ? "无限制" : item.Length + "月" }} )</span>
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="支付金额:">
        <el-input v-model="oData.show_price"/>
      </el-form-item>
      <el-form-item label="备注:">
        <el-input
            v-model="oData.remark"
            :rows="2"
            type="textarea"
            placeholder="Please input"
        />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onOpenVip">提交</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>
