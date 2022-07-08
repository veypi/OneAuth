import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { store, key } from './store'
import OneIcon from '@veypi/one-icon'
import naive from 'naive-ui'
import { Api } from './api'

import '@veypi/msg/index.css'
import './index.css'
import './assets/icon.js'
import 'animate.css'

const app = createApp(App)

app.use(Api)
app.use(naive)
app.use(OneIcon)
app.use(router)
app.use(store, key)
app.mount('#app')
