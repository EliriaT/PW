import { createRouter, createWebHistory } from 'vue-router'
import QuizzesView from '../views/QuizzesView.vue'

const routes = [
  {
    path: '/',
    name: 'quizzes',
    component: QuizzesView
  },
  {
    path: '/register',
    name: 'register',
  },
  {
    path: '/logout',
    name: 'logout',
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
