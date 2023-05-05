<template>
  <div class="mt-11 mx-11 grid md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-10 justify-items-center ">
    <!-- cards go here -->

    <div v-if="!loaded" class="mx-60 my-40  absolute">
      <h1 class="text-green-800 text-5xl">Loading...</h1>
      <svg class="animate-spin h-36 w-36 text-green-800 ml-3 mt-6" xmlns="http://www.w3.org/2000/svg" fill="none"
        viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
        <path strokeLinecap="round" strokeLinejoin="round"
          d="M15.182 15.182a4.5 4.5 0 01-6.364 0M21 12a9 9 0 11-18 0 9 9 0 0118 0zM9.75 9.75c0 .414-.168.75-.375.75S9 10.164 9 9.75 9.168 9 9.375 9s.375.336.375.75zm-.375 0h.008v.015h-.008V9.75zm5.625 0c0 .414-.168.75-.375.75s-.375-.336-.375-.75.168-.75.375-.75.375.336.375.75zm-.375 0h.008v.015h-.008V9.75z" />
      </svg>
    </div>

    <router-link class="min-w-full min-h-full" v-for="n in quizzes" :key="n.id"
      :to="{ name: 'quizView', params: { id: n.id } }">

      <Card class="min-w-full min-h-full" @click="addQuizToStore(n.id, n.questions_count)" :header="n.title"
        :description="'Questions ' + n.questions_count" :quizState="getQuizStateInString(n.id, n.questions_count)"
        :class="{ 'bg-emerald-500': getQuizStateInNum(n.id, n.questions_count) == 3, 'bg-yellow-200': getQuizStateInNum(n.id, n.questions_count) == 2, 'bg-white': getQuizStateInNum(n.id, n.questions_count) == 1 }" />
    </router-link>

  </div>
</template>

<script>
import { mapStores } from 'pinia'
import Card from '../components/Card.vue'
import { useErrorStore } from '../stores/ErrorStore'
import { useUserStore } from '../stores/UserStore'
import { useQuizzesStore } from '../stores/QuizStore'
import clickSound from '../audio/button.mp3'

export default {
  name: 'QuizzesView',
  components: {
    Card
  },
  inject: ['apiKey'],
  data() {
    return {
      quizzes: [],
      loaded: false

    }

  },
  computed: {

    ...mapStores(useErrorStore, useUserStore, useQuizzesStore)

  },
  methods: {
    playSound() {
      const audio = new Audio(clickSound)
      audio.play()
    },
    addQuizToStore(quizId, questionsCount) {
      this.playSound()
      this.quizzesStore.newQuiz(this.userStore.user.id, quizId, questionsCount)
    },

    getQuizStateInString(quizId, questionsCount) {
      if (this.quizzesStore.isQuizStarted(this.userStore.user.id, quizId)) {
        let isFinished, score
        [isFinished, score] = this.quizzesStore.isQuizFinished(this.userStore.user.id, quizId)
        if (isFinished) {
          let percentage = score / questionsCount * 100
          percentage.toFixed(2)
          return "Score: " + percentage.toFixed(2) + " %"
        }
        const quiz = this.quizzesStore.getQuizState(this.userStore.user.id, quizId)
        return "In progress: " + (quiz.lastAnsweredIndex) + "/10"
      } else {
        return "Try it!"
      }
    },
    getQuizStateInNum(quizId, questionsCount) {
      if (this.quizzesStore.isQuizStarted(this.userStore.user.id, quizId)) {
        let isFinished, score
        [isFinished, score] = this.quizzesStore.isQuizFinished(this.userStore.user.id, quizId)
        if (isFinished) {
          return 3
        }
        return 2
      } else {
        return 1
      }
    },
    sortQuizzes() {
      // console.log("started")
      const finishedQuizzes = this.quizzes.filter((q) => {
        let isFinished, score
        [isFinished, score] = this.quizzesStore.isQuizFinished(this.userStore.user.id, q.id)

        return isFinished
      })

      const startedQuizzes = this.quizzes.filter(q => {
        let isFinished, score
        [isFinished, score] = this.quizzesStore.isQuizFinished(this.userStore.user.id, q.id)
        return this.quizzesStore.isQuizStarted(this.userStore.user.id, q.id) && !isFinished
      })

      const otherQuizzes = this.quizzes.filter(q => {
        let isFinished, score
        let isStarted

        [isFinished, score] = this.quizzesStore.isQuizFinished(this.userStore.user.id, q.id)

        if (isFinished == false && this.quizzesStore.isQuizStarted(this.userStore.user.id, q.id) == false) {
          return true
        }
        return false
      })

      this.quizzes = finishedQuizzes.concat(startedQuizzes, otherQuizzes)
    }

  },
  mounted() {
    fetch('https://late-glitter-4431.fly.dev/api/v54/quizzes', {
      method: "GET", // *GET, POST, PUT, DELETE, etc.
      headers: {
        "X-Access-Token": this.apiKey

      },

    }).then(response => {
      if (response.ok) {
        response.json().then(data => {
          this.quizzes = data
          this.sortQuizzes()
          this.loaded = true

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

<style>
.someStrangeBug {
  background-color: red;
}
</style>