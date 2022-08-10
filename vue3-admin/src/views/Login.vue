<script setup>
import {ref, onMounted, onUnmounted} from 'vue'
import {login, adminUserInfo, getScene, getVipConfig} from "../common/api";
import {ElMessage} from "element-plus";
import localStorageCache from "../common/localStorage";
import {useRouter} from 'vue-router'

const user_name = ref("")
const password = ref("")
const router = useRouter()

onMounted(() => {
  let width = window.innerWidth,
      height = window.innerHeight,
      canvas = document.getElementById("canvas"),
      ctx = canvas.getContext("2d"),
      // 数量
      arc = 100,
      time = 0,
      size = 7,
      speed = 20,
      // 坐标数组
      parts = [],
      // 圆圈数量
      colors = ["#DA1212FF", "#f57900", "yellow", "#ce5c00", "#5c3566", "#fff", "#87C5DCFF"]

  let mouse = {x: 0, y: 0}

  canvas.setAttribute("width", width)
  canvas.setAttribute("height", height)

  // 创建arc个坐标数组
  for (let a = 0; a < arc; a++) {
    parts[a] = {
      x: Math.ceil(Math.random() * width),
      y: Math.ceil(Math.random() * height),
      toX: Math.random() * 5 - 1,
      toY: Math.random() * 2 - 1,
      c: colors[Math.floor(Math.random() * colors.length)],
      size: Math.random() * size
    }
  }

  function MouseMove(a) {
    mouse.x = a.layerX;
    mouse.y = a.layerY
  }

  function DistanceBetween(m, d) {
    let a = d.x - m.x;
    let b = d.y - m.y;
    return Math.sqrt(a * a + b * b)
  }

  // 监听元素触摸事件
  canvas.addEventListener("mousemove", MouseMove, false);
  (function particles() {
    ctx.clearRect(0, 0, width, height);
    for (let b = 0; b < arc; b++) {
      let c = parts[b];
      let a = DistanceBetween(mouse, parts[b]);
      a = Math.max(Math.min(15 - (a / 10), 10), 1);
      ctx.beginPath();
      ctx.arc(c.x, c.y, c.size * a, 0, Math.PI * 2, false);
      // 设置元素颜色
      ctx.fillStyle = c.c;
      ctx.strokeStyle = c.c;
      if (b % 2 === 0) {
        // 空心
        ctx.stroke()
      } else {
        // 实心
        ctx.fill()
      }
      c.x = c.x + c.toX * (time * 0.05);
      c.y = c.y + c.toY * (time * 0.05);
      if (c.x > width) {
        c.x = 0
      }
      if (c.y > height) {
        c.y = 0
      }
      if (c.x < 0) {
        c.x = width
      }
      if (c.y < 0) {
        c.y = height
      }
    }
    if (time < speed) {
      time++
    }
    setTimeout(particles, 1000 / 60)
  })()
  window.addEventListener('keydown', keyDown)
})

onUnmounted(() => {
  window.removeEventListener('keydown', keyDown, false)
})

const keyDown = (e) => {
  if (e.keyCode === 13) {
    onSubmit()
  }
}

const onSubmit = () => {
  if (!user_name.value || !password.value) {
    ElMessage({
      message: "信息有误",
      type: 'warning',
    })
    return
  }
  login(user_name.value, password.value).then(res => {
    localStorageCache.set('x-token', res.data, 86400)
    adminUserInfo().then(res => {
      localStorageCache.set('x-user', res.data, 86400)
    })
    getScene().then(res => {
      localStorageCache.set('x-scene', res.data, 86400)
    })
    getVipConfig().then(res => {
      localStorageCache.set('x-vip-config', res.data, 86400)
    })
    setTimeout(res => {
      router.push("/")
    }, 300)
  })
}

</script>

<template>
  <div class="container">
    <canvas id="canvas" width="1129" height="946"></canvas>
    <form class="form">
      <div class="item">
        <label>账号：</label>
        <div class="input">
          <el-input v-model="user_name" placeholder="UserName input"/>
        </div>
      </div>
      <div class="item">
        <label>密码：</label>
        <div class="input">
          <el-input v-model="password" type="password" placeholder="Please input"/>
        </div>
      </div>
      <div class="item">
        <el-button type="primary" @click="onSubmit">登录</el-button>
      </div>
    </form>
  </div>
</template>

<style scoped lang="less">
.container {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.form {
  position: absolute;
  background-color: #ffffff;
  flex-direction: column;
  display: flex;
  justify-content: center;
  padding: 5rem;
  z-index: 111;
  opacity: 0.9;

  .item {
    display: flex;
    flex-direction: row;
    margin-bottom: 1rem;
    justify-content: center;
    align-items: center;
  }
}

#canvas {
  overflow: hidden;
  width: 100vw;
  height: 100vh;
  background-color: black
}
</style>
