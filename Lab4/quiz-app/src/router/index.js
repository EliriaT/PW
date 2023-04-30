import { createRouter, createWebHistory } from 'vue-router'
import QuizzesView from '../views/QuizzesView.vue'
import QuizView from '../views/QuizView.vue'
import NotFound from '../views/NotFound.vue'
import RegisterView from '../views/RegisterView.vue'
import LoginView from '../views/LoginView.vue'

const routes = [
  {
    path: '/',
    name: 'quizzes',
    component: QuizzesView
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterView
  },
  {
    path: '/logout',
    name: 'logout'
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView
  },
  {
    path: '/quiz/:id',
    name: 'quizView',
    component: QuizView,
    props: true // accept route parameters as props
  },
  //catch 404
  {
    path: '/:catchAll(.*)',
    component: NotFound
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
