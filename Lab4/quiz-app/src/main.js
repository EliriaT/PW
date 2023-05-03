import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/tailwind.css'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
// import * as dotenv from 'dotenv'
// dotenv.config()

createApp(App).use(router).use(createPinia().use(piniaPluginPersistedstate)).mount('#app')
