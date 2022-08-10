import { updateToken } from './helper'
import { hexMD5 } from './md5'
const { globalData } = require('../common/globalData').default

class Request {
  /**
   * constructor.
   *
   * @param {Object} wx 微信wx对象
   */
  constructor(wx) {
    this.wx = wx
    this.before = []
    this.after = []
    this.requestList = []
  }

  static handleIntercept(handles, data) {
    return handles.reduce((old, current) => {
      return current(old)
    }, data)
  }

  use(before, after) {
    typeof before === 'function' && this.before.push(before)
    typeof after === 'function' && this.after.push(after)
  }

  get(url, data, hideLoading = false, header = {}) {
    return this.request({
      url,
      data,
      method: 'GET',
      header,
      hideLoading
    })
  }

  post(url, data, hideLoading = false, header = {}) {
    return this.request({
      url,
      data: data,
      method: 'POST',
      header,
      hideLoading
    })
  }

  request(config) {
    // 部分拦截可能导致小程序白屏 可以取消拦截
    if (globalData.isLoggingIn && !config.url.includes('wx-auth')) {
      // 登录中且请求非登录Api则拦截
      this.requestList = []
      return new Promise((resolve, reject) => {
      })
    }
    let urlMd5 = hexMD5(config.url)
    if (this.requestList.includes(urlMd5)) {
      return new Promise((resolve, reject) => {
      })
    }

    this.requestList.push(urlMd5)
    if (!config.hideLoading) {
      wx.showLoading({
        title: '...'
      })
    }

    let _config = Request.handleIntercept(this.before, config)
    return new Promise((resolve, reject) => {
      this.wx.request({
        ..._config,
        ...{
          success: res => {
            this.requestList.splice(this.requestList.indexOf(urlMd5), 1)
            if (!config.hideLoading) {
              wx.hideLoading()
            }

            let response = Request.handleIntercept(this.after, res)
            if (response && response.statusCode === 200) {
              resolve(response.data)
            } else {
              reject(response)
            }
          },
          fail: err => {
            this.requestList.splice(this.requestList.indexOf(urlMd5), 1)
            wx.hideLoading()
            reject(err)
          }
        }
      })
    })
  }
}

const request = new Request(wx)

request.use(function(config) {
  if (!config.header) {
    config.header = {}
  }
  // 拦截所有请求添加header头认证
  config.header['x-token'] = wx.getStorageSync('x-data').token
  config.header['content-type'] = 'application/x-www-form-urlencoded'
  return config
}, function(response) {
  switch (response.data.code) {
    case 20001: // token无效
      updateToken()
      break
    case 30001: // 不能频繁操作
      wx.showToast({
        title: '不能频繁操作',
        icon: 'none'
      })
      break
    case 30002: // 小程序审核中
      wx.reLaunch({ url: '/pages/gallery' })
      break
    case 10002: // 账号被锁定
    case 30003: // 关站维护
      wx.setStorageSync('err-data', response.data)
      wx.reLaunch({ url: '/pages/err' })
      break
    case 30004: // 版本号变化
      wx.clearStorageSync()
      updateToken()
      break
    default:
      return response
  }
})

// 处理baseURL
request.use(function(config) {
  if (!/^https:\/\/.*/.test(config.url)) {
    config.url = API_URL + '/api/' + config.url
  }
  return config
}, null)

export default request
