const localStorageCache = {
    /**
     * 总容量5M
     * 存入缓存，支持字符串类型、json对象的存储
     * @param key 缓存key
     * @param value
     * @param time
     * @time 数字 缓存有效时间（秒） 默认1200s
     * */
    set: function set(key, value, time) {
        if (!localStorage) {
            return false
        }
        if (!time || isNaN(time)) {
            time = 1200
        }
        try {
            let expireDate = (new Date() - 1) + time * 1000;
            localStorage.setItem(key, JSON.stringify({val: value, exp: expireDate}))
        } catch (e) {
        }
        return true
    },
    get: function get(key) {
        try {
            if (!localStorage) {
                return null
            }
            let value = localStorage.getItem(key)
            let result = JSON.parse(value)
            let now = new Date() - 1
            if (!result) {
                return null
            }// 缓存不存在
            if (now > result.exp) { // 缓存过期
                this.del(key)
                return null
            }
            return result.val
        } catch (e) {
            this.del(key)
            return null
        }
    },
    del: function del(key) {
        if (!localStorage) {
            return false
        }
        localStorage.removeItem(key)
        return true
    },
    // 清空所有缓存
    delAll: function delAll() {
        if (!localStorage) {
            return false
        }
        localStorage.clear()
        return true
    }
}

export default localStorageCache
