<script setup>
import {inject, ref, reactive} from 'vue'
import {ElLoading, ElMessage} from 'element-plus';
import {postUserRechargeAiDou} from '../../common/api';

const props = defineProps(['userRow'])
const emits = defineEmits(['changeRow'])
const showRechargeAiDouDialog = ref(inject('showRechargeAiDouDialog'))
const oData = reactive({
  uid: 0,
  show_price: 0,
  price: 0,
  ai_dou: 0,
  remark: '',
});

const onOpenVip = res => {
  oData.uid = props.userRow.info.id
  oData.price = parseInt(oData.show_price * 100)
  let loading = ElLoading.service({
    lock: true,
    text: 'Loading',
  })

  postUserRechargeAiDou(oData).then(res => {
    loading.close()
    if (res.code === 1) {
      props.userRow.info.ai_dou.total_ai_dou = parseInt(props.userRow.info.ai_dou.total_ai_dou) + parseInt(oData.ai_dou)
      props.userRow.info.ai_dou.ai_dou = parseInt(props.userRow.info.ai_dou.ai_dou) + parseInt(oData.ai_dou)
      emits('changeRow', props.userRow.info)
      showRechargeAiDouDialog.value = false
      ElMessage({
        message: '开通成功',
        type: 'success',
      })
    }
  })
}

</script>

<template>
  <el-dialog v-model="showRechargeAiDouDialog" title="充值爱享豆">
    <el-form label-width="5rem" class="my-edit">
      <el-form-item label="用户信息:">
        {{ userRow.info.nick_name }}: {{ userRow.info.user_name }}
      </el-form-item>
      <el-form-item label="支付金额:">
        <el-input v-model="oData.show_price"/>
      </el-form-item>
      <el-form-item label="充值数量:">
        <el-input v-model="oData.ai_dou"/>
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
