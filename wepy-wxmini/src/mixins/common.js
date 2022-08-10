import store from '../store'

export default {
  store,
  data: {},
  onShareAppMessage() {
    // eslint-disable-next-line no-undef
    let pages = getCurrentPages()
    console.log(this.buildUrl(pages[pages.length - 1].route))
    return {
      path: this.buildUrl(pages[pages.length - 1].route),
      success: function(res) {
        console.log('share success:', res)
      },
      fail: function(res) {
        console.log('share fail:', res)
      }
    }
  },
  methods: {
    buildUrl(url) {
      let query = {}
      let path = url
      let l = url.indexOf('?')
      if (l !== -1) {
        path = url.substr(0, l)
        let str = url.substr(l + 1)
        let pairs = str.split('&')
        for (let i = 0; i < pairs.length; i++) {
          let pair = pairs[i].split('=')
          query[pair[0]] = pair[1]
        }
      }
      if (/^pages/.test(path)) {
        path = '/' + path
      }
      path = path + '?'
      query.from_uid = store.state.uid

      for (let k in query) {
        path = path + k + '=' + query[k] + '&'
      }

      return path.substring(0, path.length - 1)
    },
    jump(page) {
      wx.navigateTo({
        url: this.buildUrl(page)
      })
    }
  },
  created() {
    console.log('created in mixin common')
  }
}
