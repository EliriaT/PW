<template>
  <div class=" w-1/4 absolute  bg-green-100 top-1/3 left-1/3 ml-24 ">
    <form class="bg-white shadow-md rounded px-8 pt-6 pb-8 " @submit.prevent="handleSubmit">

      <div class="mb-4">
        <label class="block text-gray-700 text-lg font-bold mb-2" for="Name">
          Name
        </label>
        <input
          class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          id="Name" type="text" placeholder="Name" v-model="name">
      </div>

      <div class="mb-4">
        <label class="block text-gray-700 text-lg font-bold mb-2" for="Surname">
          Surname
        </label>
        <input
          class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          id="Surname" type="text" placeholder="Surname" v-model="surname">
      </div>
      <div v-if="error">
        <p v-for="err in error" class="text-red-500 font-bold bg-red-100 rounded-md text-center p-2 mb-4" :key="err">{{
          err }}</p>
      </div>

      <div class="flex items-center justify-between">

        <button
          class="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
          Sign Up
        </button>
      </div>

    </form>

  </div>
</template>

<script>

import { useUserStore } from '../stores/UserStore'
import { useErrorStore } from '../stores/ErrorStore'

export default {
  components: {

  },
  name: 'RegisterView',
  data() {
    return {
      name: '',
      surname: '',
      error: [],
      body: {},
      response: {}

    }
  },
  watch: {

    name(newName, oldName) {
      this.error = this.error.filter((el) => el !== 'Name must be at least 3 characters long')
      this.name.length > 2 ?
        this.error : this.error.push('Name must be at least 3 characters long')


    },
    surname(newSurname, oldSurname) {
      this.error = this.error.filter((el) => el !== 'Surname must be at least 3 characters long')
      this.surname.length > 2 ?
        this.error : this.error.push('Surname must be at least 3 characters long')
    }
  },
  methods: {
    handleSubmit() {
      // this.error = []
      // this.name.length > 2 ?
      //   [] : this.error.push('Name must be at least 3 characters long')


      // this.surname.length > 2 ?
      //   this.error : this.error.push('Surname must be at least 3 characters long')


      if (this.error.length == 0) {
        this.body = {
          "data": {
            "name": this.name,
            "surname": this.surname
          }
        }

        fetch('https://late-glitter-4431.fly.dev/api/v54/users', {
          method: "POST",
          headers: {
            "X-Access-Token": "a68baa0fe20a17aea823776f987a2741395d24402430e4e2296bce48f56310ac",
            "Content-Type": "application/json"

          },
          body: JSON.stringify(this.body),
        }).then(response => {
          if (response.ok) {
            response.json().then(json => {
              this.response = json
              this.userStore.setUser(json.name, json.surname, json.id)
              this.name = ''
              this.surname = ''
              this.$router.push({ name: 'quizzes' })

            })
          } else {
            response.json().then(json => {

              this.error = [json.message]
              this.errorStore.setError(json.message)
            })
          }
        }).catch(err => {
          console.log(err.message)
          this.errorStore.setError(err.message)
        })

      }

    }
  }, setup() {
    const userStore = useUserStore()
    const errorStore = useErrorStore()
    return {
      userStore, errorStore
    }

  }
}

</script>

<style></style>