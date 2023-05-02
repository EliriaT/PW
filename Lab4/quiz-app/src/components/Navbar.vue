<template>
  <div class="bg-gray-900 text-gray-100 py-3.5 px-6 shadow-md  md:flex justify-between items-center">

    <router-link :to="{ name: 'quizzes' }">
      <div class="flex items-center cursor-pointer">


        <span class="text-green-500 text-xl mr-1 hover:text-white">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
            class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M4.26 10.147a60.436 60.436 0 00-.491 6.347A48.627 48.627 0 0112 20.904a48.627 48.627 0 018.232-4.41 60.46 60.46 0 00-.491-6.347m-15.482 0a50.57 50.57 0 00-2.658-.813A59.905 59.905 0 0112 3.493a59.902 59.902 0 0110.399 5.84c-.896.248-1.783.52-2.658.814m-15.482 0A50.697 50.697 0 0112 13.489a50.702 50.702 0 017.74-3.342M6.75 15a.75.75 0 100-1.5.75.75 0 000 1.5zm0 0v-3.675A55.378 55.378 0 0112 8.443m-7.007 11.55A5.981 5.981 0 006.75 15.75v-1.5" />
          </svg>
        </span>


        <h1 class="text-2xl  hover:text-green-500" id="title">Let's quiz!</h1>

      </div>
    </router-link>

    <span @click="toggleMenu()" class="absolute md:hidden right-6 top-2 cursor-pointer text-3xl p-2">

      <svg v-if="!isOpen" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
        stroke="currentColor" class="w-6 h-6">
        <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
      </svg>

      <svg v-else xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
        stroke="currentColor" class="w-6 h-6">
        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
      </svg>

    </span>


    <ul class="md:flex md:items-center md-px-0 px-6 md:pb-0 pb-10
    md:static absolute bg-gray-900 md:w-auto w-full top-14 duration-100  ease-linear"
      :class="[isOpen ? 'left-0' : 'left-[-100%]']">

      <router-link :to="{ name: 'quizzes' }">
        <li class="md:mx-4 md:my-0 my-4">

          <a href="quizzes" class="text-xl hover:text-green-500">Quizzes</a>
        </li>
      </router-link>

      <router-link :to="{ name: 'login' }" v-if="!userStore.user.id">
        <li class="md:mx-4 md:my-0 my-4">

          <a href="register" class="text-xl hover:text-green-500">Login</a>
        </li>
      </router-link>

      <router-link :to="{ name: 'register' }" v-if="!userStore.user.id">
        <li class="md:mx-4 md:my-0 my-4">

          <a href="register" class="text-xl hover:text-green-500">Register</a>
        </li>
      </router-link>

      <router-link :to="{ name: 'quizzes' }" v-else>
        <li class="md:mx-4 md:my-0 my-4">

          <a href="quizzes" class="text-xl hover:text-green-500">{{ "Hi " + userStore.getName + "!" }}</a>
        </li>
      </router-link>

      <router-link v-if="userStore.user.id" :to="{ name: 'logout' }">
        <li class="md:mx-4 md:my-0 my-4" @click="userStore.$reset">

          <a href="logout" class="text-xl hover:text-green-500">Sign Out</a>
        </li>
      </router-link>


      <Button v-if="userStore.user.id">Reset Scores</Button>
      <Button v-if="userStore.user.id" @click="deleteUser">Delete account</Button>

    </ul>

  </div>
</template>

<script>
import { ref } from '@vue/reactivity'
import Button from './Button.vue'
import { useUserStore } from '../stores/UserStore'
import { useErrorStore } from '../stores/ErrorStore'
import { useQuizzesStore } from '../stores/QuizStore'
import { useRouter } from 'vue-router'

export default {
  components: { Button },
  setup() {
    const router = useRouter()
    const userStore = useUserStore()
    const errorStore = useErrorStore()
    const quizStore = useQuizzesStore()

    const isOpen = ref(false)
    let links = [
      { name: "Quizzes", link: "quizzes" },
      { name: "Register", link: "register" },
      { name: "Log out", link: "logout" },

    ]

    function toggleMenu() {
      isOpen.value = !isOpen.value
    }

    function deleteUser() {

      fetch('https://late-glitter-4431.fly.dev/api/v54/users/' + userStore.user.id, {
        method: "DELETE",
        headers: {
          "X-Access-Token": "a68baa0fe20a17aea823776f987a2741395d24402430e4e2296bce48f56310ac",
        },
      }).then(response => {
        if (response.ok) {
          response.json().then(json => {
            quizStore.deleteUserInformation(userStore.user.id)
            userStore.$reset()
            router.push({ name: 'register' })

          })
        } else {
          response.json().then(json => {
            errorStore.setError(json.message)
            console.log(json.message)
          })
        }
      }).catch(err => {
        console.log(err.message)
        this.errorStore.setError(err.message)
      })

    }

    return { links, isOpen, toggleMenu, userStore, deleteUser, errorStore, quizStore }
  }
}
</script>

<style>
.router-link-exact-active {
  color: #42b983;
}

#title {
  color: #f6fcf9;
}

#title:hover {
  color: #42b983;
}
</style>