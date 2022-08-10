import { getUserInfo } from '../common/api'

export default {
  data: {
    user: {
      id: 0,
      nick_name: '',
      head_img: '',
      user_name: ''
    },
    more: {
      register_day: 0,
      vip: 0,
      wallet: {
        balance: 0,
        total_balance: 0
      },
      friends: 0,
      dou: {
        ai_dou: 0,
        total_ai_dou: 0
      },
      bg_img: ''
    }
  },
  methods: {
    userInfo() {
      let that = this
      getUserInfo().then(res => {
        that.user = res.data.user
        that.more = res.data.more
      })
    },
    userInfoUpdate() {
      let that = this
      getUserInfo().then(res => {
        that.user = res.data.user
        that.more = res.data.more
      })
    }
  },
  created() {
    console.log('created in mixin userInfo')
  }
}
