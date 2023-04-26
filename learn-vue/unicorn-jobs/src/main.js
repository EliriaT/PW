import { createApp } from 'vue'
import App from './App.vue'
// automatically imports index.js from the router folder
import router from './router'

createApp(App).use(router).mount('#app')
