import { createApp } from 'vue'
import App from './App.vue'
import OneIcon from '@veypi/one-icon'
import router from '@/router'
import { createPinia } from 'pinia'
import 'animate.css'
import './assets/icon.js'
import './index.css'

let app = createApp(App)
app.use(OneIcon)
app.use(router)
app.use(createPinia())
app.mount('#app')
