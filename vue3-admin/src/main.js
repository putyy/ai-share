import {createApp} from 'vue'
import './style.css'
import AppContainer from './App.vue'
import router from './route'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import 'element-plus/dist/index.css'

const app = createApp(AppContainer)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.use(router).mount('#app')
