import request from './request'

// ------------ public api start-------------
export function wxLogin(code, fromUid) {
  return request.get('wx-auth', { code: code, from_uid: fromUid })
}

export function getScene() {
  return request.get('public/scene', {})
}

export function getQiniuToken(scene, ext) {
  return request.get('public/qiniu-token', { scene: scene, ext: ext || '' })
}

export function getContent(scene) {
  return request.get('public/content', { scene: scene })
}

export function postVideoParse(content) {
  return request.post('video-parse/short', { content: content })
}
// ------------ public api end-------------

// ------------ member-center api end-------------
export function getCenterInfo() {
  return request.get('member-center/center-info', {})
}

export function getUserInfo() {
  return request.get('member-center/user', {})
}

export function postUserEdit(data) {
  return request.post('member-center/user-edit', data)
}

export function getFansList(lastId, vip, keyword) {
  return request.get('member-center/fans-list', { vip: vip || '-1', keyword: keyword || '', last_id: lastId || 0 })
}

export function getAiDouLog(lastId, type, keyword) {
  return request.get('member-center/ai-dou-log', { type: type || '', keyword: keyword || '', last_id: lastId || 0 })
}

export function getWalletLog(lastId, type, keyword) {
  return request.get('member-center/wallet-log', { type: type || '', keyword: keyword || '', last_id: lastId || 0 })
}

export function getFeedbackList(lastId) {
  return request.get('member-center/feedback-list', { last_id: lastId || 0 })
}

export function postFeedbackCreate(content) {
  return request.post('member-center/feedback-create', { content: content })
}

// ------------ member-center api end-------------

// ------------ ppt api start-------------
export function getPptType() {
  return request.get('ppt/type', {})
}

export function getPptList(typeId, lastId, keyword) {
  return request.get('ppt/list', { type_id: typeId || '', keyword: keyword || '', last_id: lastId || 0 })
}

export function getPptContent(id) {
  return request.get('ppt/detail', { id: id })
}

export function postPptBuy(id) {
  return request.post('ppt/buy', { id: id })
}
// ------------ ppt api end-------------

// ------------ resource api start-------------
export function getResourceType() {
  return request.get('share-resource/type', {})
}

export function getResourceList(lastId, tid) {
  return request.get('share-resource/list', { last_id: lastId || 0, tid: tid || 0 })
}
// ------------ resource api end-------------
