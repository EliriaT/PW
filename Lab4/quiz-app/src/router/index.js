import { createRouter, createWebHistory } from 'vue-router'
import QuizzesView from '../views/QuizzesView.vue'
import QuizView from '../views/QuizView.vue'
import NotFound from '../views/NotFound.vue'
import RegisterView from '../views/RegisterView.vue'
import LoginView from '../views/LoginView.vue'
import { useUserStore } from '../stores/UserStore'

const routes = [
  {
    path: '/',
    name: 'quizzes',
    component: QuizzesView
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterView,
    beforeEnter: (to, from) => {
      // reject the navigation
      const userStore = useUserStore()
      if (
        userStore.isLoggedIn
      ) {
        // console.log("aici")
        return from
      }

    },
  },
  {
    path: '/logout',
    name: 'logout'
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView,
    beforeEnter: (to, from) => {
      // reject the navigation
      const userStore = useUserStore()
      if (
        userStore.isLoggedIn
      ) {

        return from
      }

    },
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

router.beforeEach(async (to, from) => {
  const userStore = useUserStore()
  if (
    !userStore.isLoggedIn &&

    (to.name !== 'login' && to.name !== 'register')
  ) {
    // console.log("aici:((")
    // redirect the user to the login page
    return { name: 'login' }
  }
  // console.log("aici:((")

})

export default router
