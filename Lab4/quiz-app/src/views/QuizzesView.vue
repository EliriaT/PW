<template>
  <div class="mt-11 mx-11 grid md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-10 justify-items-center ">
    <!-- cards go here -->

    <div v-for="n in 20" :key="n">
      <router-link v-for="n in quizzes" :key="n.id" :to="{ name: 'quizView', params: { id: n.id } }">
        <Card :header="n.title" :description="'Questions ' + n.questions_count" />
      </router-link>
    </div>

  </div>
</template>

<script>
import { mapStores } from 'pinia'
import Card from '../components/Card.vue'
import { useErrorStore } from '../stores/ErrorStore'
import { useUserStore } from '../stores/UserStore'
import {useQuizzesStore} from '../stores/QuizStore'

export default {
  name: 'QuizzesView',
  components: {
    Card
  },
  data() {
    return {
      quizzes: []
    }

  },
  computed: {

    ...mapStores(useErrorStore)

  },
  mounted() {
    fetch('https://late-glitter-4431.fly.dev/api/v54/quizzes', {
      method: "GET", // *GET, POST, PUT, DELETE, etc.
      headers: {
        "X-Access-Token": "a68baa0fe20a17aea823776f987a2741395d24402430e4e2296bce48f56310ac"

      },

    }).then(response => {
      if (response.ok) {
        response.json().then(data => {
          this.quizzes = data
          console.log(data)
        })
      } else {
        response.json().then(json => {
          console.log(json)
          this.errorStore.setError(json.message)

        })
      }
    }).catch(err => {
      console.log(err.message)
      this.errorStore.setError(err.message)
    })


  }
}
</script>

<style></style>