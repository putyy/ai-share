import {createRouter, createWebHistory} from 'vue-router'
import localStorageCache from "./common/localStorage"

const routes = [
    {
        path: '/',
        component: () => import('./components/layout/Index.vue'),
        // 重定向
        redirect: {name: 'Home'},
        // 子路由
        children: [
            {
                path: '/home',
                name: 'Home',
                component: () => import('./views/Home.vue'),
            },
            {
                path: '/user',
                name: 'User',
                component: () => import('./views/user/Index.vue'),
            },
            {
                path: '/user/feedback',
                name: 'Feedback',
                component: () => import('./views/user/Feedback.vue'),
            },

            {
                path: '/user/menu',
                name: 'UserMenu',
                component: () => import('./views/menu/Index.vue'),
            },
            {
                path: '/user/menu-group',
                name: 'UserMenuGroup',
                component: () => import('./views/menu/Group.vue'),
            },
            {
                path: '/system',
                name: 'System',
                component: () => import('./views/System.vue'),
            },
            {
                path: '/share-resource',
                name: 'ShareResource',
                component: () => import('./views/share-resource/Index.vue'),
            },
            {
                path: '/share-resource-type',
                name: 'ShareResourceType',
                component: () => import('./views/share-resource/Type.vue'),
            },
            {
                path: '/ppt',
                name: 'Ppt',
                component: () => import('./views/ppt/Index.vue'),
            },
            {
                path: '/ppt-type',
                name: 'PptType',
                component: () => import('./views/ppt/Type.vue'),
            },
            {
                path: '/content',
                name: 'Content',
                component: () => import('./views/Content.vue'),
            },
        ],
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('./views/Login.vue')
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes: routes,
})

router.beforeEach((to, from, next) => {
    let token = localStorageCache.get('x-token')
    if ((!token || !token.token) && !to.path.includes("login")) {
        next({
            path: '/login'
        })
        return
    }
    next()
})

export default router
