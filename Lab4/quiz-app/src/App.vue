<template>
  <div @scroll="handleScroll" class="bg-green-100 h-screen w-screen overflow-auto pb-4 ">
    <Navbar :key="isOpen" />
    <Teleport to="body">
      <!-- use the modal component, pass in the prop -->
      <modal>
        <template #header>
          <h3>Sorry! An error happened! </h3>
        </template>
        <template #body>
          <h3>{{ errorStore.error.message }} </h3>
        </template>
      </modal>
    </Teleport>
    <router-view />
  </div>
</template>

<script>
import Navbar from "./components/Navbar.vue"
import Modal from "./components/Modal.vue"
import { mapStores } from 'pinia'
import { useErrorStore } from './stores/ErrorStore'

export default {
  name: 'Quiz-App',
  data() {
    return {
      isOpen: false,

    }
  },
  computed: {
    ...mapStores(useErrorStore)
  },
  components: {
    Navbar, Modal
  },
  methods: {

    handleScroll(e) {
      const { scrollTop, offsetHeight, scrollHeight } = e.target

      if (scrollTop <= 150) {
        this.isOpen = !this.isOpen
      }
      // asfdsg sfdf
    },

    htmlDecode(input) {
      var doc = new DOMParser().parseFromString(input, "text/html");
      return doc.documentElement.textContent;
    },
    populateQuizzes() {
      let category = [22]
      for (let i = 0; i < category.length; i++) {
        fetch('https://opentdb.com/api.php?amount=14&category=' + category[i], {
          method: "GET", // *GET, POST, PUT, DELETE, etc.

        }).then(response => {
          if (response.ok) {
            response.json().then(data => {
              let quiz = data.results
              console.log(data)
              let firstQuiz = quiz[0]
              let body = { "data": { "title": firstQuiz.category, "questions": [] } }
              let mainBody = body.data
              for (let i = 0; i < quiz.length; i++) {
                const question = quiz[i]
                mainBody.questions.push({
                  "question": question.question,
                  "answers": [
                    question.correct_answer,
                    question.incorrect_answers[0],
                    question.incorrect_answers[1],
                    question.incorrect_answers[2]
                  ],
                  "correct_answer": question.correct_answer
                })
              }
              // in order to shufle the answers
              mainBody.questions = mainBody.questions.sort((a, b) => 0.5 - Math.random())


              fetch('https://late-glitter-4431.fly.dev/api/v54/quizzes', {
                method: "POST",
                headers: {
                  "X-Access-Token": process.env.VUE_APP_API_KEY,
                  "Content-Type": "application/json"

                },
                body: JSON.stringify(body),
              }).then(response => {
                if (response.ok) {
                  response.json().then(json => {
                    console.log("Succesfully created")

                  })
                } else {
                  response.json().then(json => {
                    console.log(json.message)
                  })
                }
              })

            })
          } else {
            response.json().then(json => {
              console.log(json)

            })
          }
        }).catch(err => {
          console.log(err.message)
        })
      }
    }
  },
  mounted() {
    // this.populateQuizzes()
  }
}
</script>

<style></style>
