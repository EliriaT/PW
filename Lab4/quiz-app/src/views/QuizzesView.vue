<template>
  <div class="mt-11 mx-11 grid md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-10 justify-items-center ">
    <!-- cards go here -->

    <!-- <div v-for="n in 20" :key="n"> -->
    <router-link class="min-w-full min-h-full" v-for="n in quizzes" :key="n.id"
      :to="{ name: 'quizView', params: { id: n.id } }"
      >
      <!-- :class="{ 'bg-teal-400': getQuizStateInNum(n.id, n.questions_count) == 3, 'bg-yellow-200': getQuizStateInNum(n.id, n.questions_count) == 2 }"  -->
      <Card class="min-w-full min-h-full" @click="addQuizToStore(n.id, n.questions_count)" :header="n.title"
        :description="'Questions ' + n.questions_count" :quizState="getQuizStateInString(n.id, n.questions_count)"
        :class="{ 'bg-teal-600': getQuizStateInNum(n.id, n.questions_count) == 3, 'bg-yellow-200': getQuizStateInNum(n.id, n.questions_count) == 2 }"
       
      />
    </router-link>
    <!-- </div> -->
    <!-- //getCurrentQuestionIndex(n.id) -->
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