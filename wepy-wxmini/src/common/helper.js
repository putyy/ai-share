import store from '../store'
import { hexMD5 } from './md5'

export function updateToken() {
  // eslint-disable-next-line no-undef
  let pages = getCurrentPages()
  if (pages.length >= 1) {
    let route = pages[pages.length - 1].route
    let options = pages[pages.length - 1].options
    if (/^pages/.test(route)) {
      route = '/' + route
    }
    store.dispatch('login')
      .then(() => {
        setTimeout(function() {
          if (JSON.stringify(options) !== '{}') { // 页面有参数
            let ginseng = ''
            let oLength = Object.keys(options).length
            let oValues = Object.values(options)
            Object.keys(options).forEach((item, index) => {
              ginseng += `${item}=${oValues[index]}${index < oLength - 1 ? '&' : ''}`
            })
            wx.reLaunch({
              url: `${route}?${ginseng}`
            })
            return
          }
          wx.reLaunch({ url: route })
        }, 200)
      })
  } else {
    wx.reLaunch({ url: '/pages/home' })
  }
}

export function downAuthCheck(func) {
  wx.getSetting({
    success(res) {
      if (!res.authSetting['scope.writePhotosAlbum']) { // 未授权
        wx.authorize({
          scope: 'scope.writePhotosAlbum',
          success: () => { // 授权成功
            func()
          },
          fail: () => { // 拒绝授权
            wx.showToast({
              title: '你不授权，下载个锤子哦',
              icon: 'none'
            })
          }
        })
      } else {
        func()
      }
    }
  })
}

export function downVideo(url) {
  downAuthCheck(function() {
    wx.showLoading({
      title: '保存中'
    })
    let fileName = wx.env.USER_DATA_PATH + '/ai_share_' + hexMD5(url) + '.mp4'
    let fileManager = wx.getFileSystemManager()
    wx.downloadFile({
      url: url,
      filePath: fileName,
      success: res => {
        wx.saveVideoToPhotosAlbum({
          filePath: fileName,
          success: () => {
            wx.hideLoading()
            fileManager.unlink({
              filePath: fileName
            })
            wx.showToast({
              title: '保存成功',
              icon: 'none'
            })
          },
          fail: (err) => {
            console.log(err)
            wx.hideLoading()
            fileManager.unlink({
              filePath: fileName
            })
            wx.showToast({
              title: '保存失败',
              icon: 'none'
            })
          }
        })
      },
      fail: (err) => {
        console.log(err)
        wx.hideLoading()
        fileManager.unlink({
          filePath: fileName
        })
        wx.showToast({
          title: '保存失败，可能是文件太大哦'+ err.errMsg,
          icon: 'none'
        })
      }
    })
  })
}

export function downImage(url) {
  downAuthCheck(function() {
    wx.showLoading({
      title: '保存中'
    })
    let fileName = wx.env.USER_DATA_PATH + '/ai_share_' + hexMD5(url) + '.png'
    let fileManager = wx.getFileSystemManager()
    wx.downloadFile({
      url: url,
      filePath: fileName,
      success: res => {
        wx.saveImageToPhotosAlbum({
          filePath: fileName,
          success: () => {
            wx.hideLoading()
            fileManager.unlink({
              filePath: fileName
            })
            wx.showToast({
              title: '保存成功',
              icon: 'none'
            })
          },
          fail: (err) => {
            console.log(err)
            wx.hideLoading()
            fileManager.unlink({
              filePath: fileName
            })
            wx.showToast({
              title: '保存失败',
              icon: 'none'
            })
          }
        })
      },
      fail: (err) => {
        console.log(err)
        wx.hideLoading()
        fileManager.unlink({
          filePath: fileName
        })
        wx.showToast({
          title: '保存失败，可能是文件太大哦' + err.errMsg,
          icon: 'none'
        })
      }
    })
  })
}
