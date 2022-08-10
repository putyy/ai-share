import Vuex from '@wepy/x'
import { wxLogin, getScene } from '../common/api'

const { globalData } = require('../common/globalData').default

export default new Vuex.Store({
  state: {
    token: '',
    uid: 0,
    systemInfo: null,
    scene: {
      text: {
        about: ''
      },
      upload: {}
    }
  },
  getters: {
    isReady (state) {
      return !!(state.token)
    },
    uid (state) {
      return state.uid
    },
    winWidth(state) {
      return state.systemInfo.windowWidth
    },
    winHeight(state) {
      return state.systemInfo.windowHeight
    },
    scene(state) {
      return state.scene
    }
  },
  // 计算属性
  mutations: {
    setToken(state, payload) {
      state.token = payload
    },
    setUid(state, payload) {
      state.uid = payload
    },
    setSystemInfo(state, payload) {
      state.systemInfo = payload
    },
    setScene(state, payload) {
      if (payload) {
        state.scene = payload
      }
    }
  },
  actions: {
    login({ commit, dispatch }) {
      if (globalData.isLoggingIn) {
        return
      }
      globalData.isLoggingIn = true
      wx.login({
        success: res => {
          let pages = getCurrentPages()
          let fromUid = 1
          if (pages.length >= 1) {
            let options = pages[pages.length - 1].options
            fromUid = options.from_uid || 1
          }
          wxLogin(res.code, fromUid).then(res => {
            globalData.isLoggingIn = false
            if (res.code === 0) {
              // 失败 进入失败页面
              wx.reLaunch({
                url: '/pages/err?s=1&from_uid=' + fromUid
              })
              return
            }
            wx.clearStorage()
            commit('setToken', res.data.token)
            commit('setUid', res.data.uid)
            wx.setStorageSync('x-data', res.data)
            dispatch('scene')
          })
        }
      })
    },
    scene({ commit, dispatch }) {
      let sceneInfo = wx.getStorageSync('x-scene')
      if (sceneInfo) {
        commit('setScene', JSON.parse(sceneInfo))
        return Promise.resolve()
      } else {
        getScene()
          .then(res => {
            if (res.code === 1) {
              wx.setStorageSync('x-scene', JSON.stringify(res.data))
              commit('setScene', res.data)
            }
          })
      }
    }
  }
})
