<script setup>
import {inject, ref} from 'vue'

const sidebarCollapse = ref(inject('sidebarCollapse'))
const menus = [
  {
    id: 11,
    name: '用户管理',
    icon: 'user',
    children: [
      {id: 12, name: '用户列表', path: '/user'},
      {id: 13, name: '用户反馈', path: '/user/feedback'},
      {id: 14, name: 'center菜单', path: '/user/menu'},
      {id: 15, name: 'center菜单分组', path: '/user/menu-group'},
    ],
  },
  {
    id: 21,
    name: '推广管理',
    icon: 'share',
    children: [
      {id: 22, name: 'PPT列表', path: '/ppt'},
      {id: 22, name: 'PPT分类', path: '/ppt-type'},
      {id: 23, name: '图库列表', path: '/share-resource'},
      {id: 23, name: '图库分类', path: '/share-resource-type'},
    ],
  },
  {
    id: 31,
    name: '系统管理',
    icon: 'setting',
    children: [
      {id: 32, name: '系统设置', path: '/system'},
      {id: 33, name: '文本素材', path: '/content'},
    ],
  },
]

</script>

<template>
  <div class="sidebar">
    <el-menu default-active="1-1" class="menu" :collapse="sidebarCollapse" router>
      <div class="logo">
        <img src="../../assets/logo.png" width="32" height="32"/>
        <span v-show="!sidebarCollapse"> 爱享素材</span>
      </div>

      <el-sub-menu v-for="menu in menus" :key="menu.id" :index="menu.id + ''">
        <template #title>
          <component class="menu-icon" :is="menu.icon"></component>
          <span>{{ menu.name }}</span>
        </template>
        <el-menu-item v-for="sub in menu.children" :key="sub.id" :index="sub.path + ''">{{ sub.name }}</el-menu-item>
      </el-sub-menu>
    </el-menu>
  </div>
</template>

<style lang="less" scoped>
.menu-icon {
  max-width: 1rem;
}

.sidebar {
  height: 100vh;
  box-shadow: .2rem 0 .6rem 0 rgba(0, 0, 0, 0.1);
  .menu {
    border-right: unset;
    .logo {
      height: 5rem;
      line-height: 5rem;
      padding: 0 1rem;
      overflow: hidden;
      img {
        vertical-align: middle;
        margin-left: .6rem;
      }
      span {
        font-weight: bold;
        font-size: .16rem;
      }
    }
  }
  .menu:not(.el-menu--collapse) {
    width: 10rem;
  }
}
</style>
