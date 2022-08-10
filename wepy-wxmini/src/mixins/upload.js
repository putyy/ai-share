import { getQiniuToken } from '../common/api'

export default {
  data: {},
  methods: {
    uploadQiNiu(filePath, scene, ext) {
      return new Promise((resolve, reject) => {
        if (!filePath.includes('//tmp/')) {
          wx.showToast({
            title: '文件格式有误',
            icon: 'none',
            duration: 1000
          })
          return
        }
        getQiniuToken(scene, ext).then(res => {
          let tokens = res.data
          wx.uploadFile({
            url: tokens[0].up_serve,
            filePath,
            name: 'file',
            formData: {
              token: tokens[0].token,
              key: tokens[0].key
            },
            success(res) {
              resolve(JSON.parse(res.data))
            },
            fail(err) {
              reject(err)
            }
          })
        })
      })
    }
  },
  created() {
    console.log('created in mixin qiniu')
  }
}
