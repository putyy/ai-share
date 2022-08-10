import request from './request'

export function login(username, password) {
    return request.get("login", {username: username, password: password})
}

export function adminUserInfo() {
    return request.get("admin-user-info")
}

export function getSystemInfo() {
    return request.get("system-info/detail")
}

export function postSystemInfoEdit(data) {
    return request.post("system-info/edit", data)
}

export function postDel(id, m_mark) {
    return request.post("__delete", {id, m_mark})
}

export function postLock(id, is_lock, m_mark) {
    return request.post("__lock", {id, is_lock, m_mark})
}


export function getScene() {
    return request.get("public/scene")
}

export function getQiniuToken(scene, ext) {
    return request.get("public/qiniu-token", {scene, ext})
}

export function getVipConfig() {
    return request.get("public/vip-config")
}

export function getUserList(data) {
    return request.get("user/list", data || {})
}

export function postUserLock(id, is_lock) {
    return request.post("user/lock", {id, is_lock})
}

export function postUserDel(id) {
    return request.post("user/deleted", {id})
}

export function postUserOpenVip(data) {
    return request.post("user/open-vip", data)
}

export function postUserRechargeAiDou(data) {
    return request.post("user/recharge-ai-dou", data)
}

export function getMenuGroupList(data) {
    return request.get("menu-group/list", data)
}

export function postMenuGroupEdit(data) {
    return request.post("menu-group/edit", data)
}

export function getMenuList(data) {
    return request.get("menu/list", data)
}

export function postMenuEdit(data) {
    return request.post("menu/edit", data)
}

export function getContentList() {
    return request.get("content/list")
}

export function postContentEdit(data) {
    return request.post("content/edit", data)
}

export function getFeedbackList(data) {
    return request.get("feedback/list", data)
}

export function postFeedbackEdit(data) {
    return request.post("feedback/remark", data)
}

export function getShareResourceTypeList(data) {
    return request.get("share-resource-type/list", data)
}

export function postShareResourceTypeEdit(data) {
    return request.post("share-resource-type/edit", data)
}

export function getShareResourceList(data) {
    return request.get("share-resource/list", data)
}

export function postShareResourceEdit(data) {
    return request.post("share-resource/edit", data)
}

export function getPptTypeList(data) {
    return request.get("ppt-type/list", data)
}

export function postPptTypeEdit(data) {
    return request.post("ppt-type/edit", data)
}

export function getPptList(data) {
    return request.get("ppt/list", data)
}

export function postPptEdit(data) {
    return request.post("ppt/edit", data)
}

export function getPptContent(id) {
    return request.get("ppt/content", {id})
}

export function postPptContentEdit(data) {
    return request.axiosObj().post("ppt/content-edit", data)
}

