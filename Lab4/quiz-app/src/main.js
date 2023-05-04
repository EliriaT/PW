import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/tailwind.css'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'


const app = createApp(App)

// auto unwrapping injected ref
app.config.unwrapInjectedRef = true

// app.provide(/* key */ 'message', /* value */ 'hello!')

app.use(router).use(createPinia().use(piniaPluginPersistedstate)).mount('#app')



