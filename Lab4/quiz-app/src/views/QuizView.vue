<template>
    <div class="pt-14 overflow-auto md:px-36 flex flex-col px-10  items-center ">
        <h1 class="text-5xl font-bold text-center text-red-900 ">{{ quiz.title }}</h1>



        <div v-if="currentQuestion" class="text-gray-600 font-semibold md:text-2xl text-lgtext-center py-4">
            Currently at question {{ currentQuestionIndex + 1 }} of {{ questions.length }}
        </div>

        <h2 v-if="currentQuestion"
            class="text-center font-bold lg:text-5xl lg:leading-normal lg:px-48 md:text-4xl md:leading-normal md:px-20  sm:text-2xl sm:leading-normal sm:px-2 text-xl py-10  px-2 ">
            {{ currentQuestion.question }}</h2>

        <h2 v-else class="text-center text-6xl p-9 text-red-800 ">Score: <span
                class="block text-center text-5xl p-9 text-green-800"
                :class="{ 'text-red-600': (this.score / this.questions.length * 100) < 50 }"> {{ this.score /
                    this.questions.length * 100 }} % </span></h2>

        <form v-if="currentQuestion" @submit.prevent="handleSubmit"
            class="grid sm:grid-cols-1 md:grid-cols-1 lg:grid-cols-2 xl:grid-cols-2 gap-10 justify-items-center pb-16 pt-10">
            <button
                class=" transition duration-200 ease-out bg-white hover:bg-gray-100 border-2 border-green-700 text-2xl p-8 m-4 rounded-lg font-bold min-w-full  min-h-full overflow-hidden shadow-xl hover:shadow-lg px-6 py-4  hover:scale-110"
                v-for="answer in currentQuestion.answers" :key="answer"
                :class="{ 'bg-yellow-100 border-8 border-yellow-300  ': selectedAnswer == answer, 'shake border-red-400 bg-red-50': warnButton }"
                @click="handleClick">
                {{ answer }}
            </button>
        </form>
        <button v-if="currentQuestion" @click="nextQuestion"
            :class="{ 'cursor-not-allowed': selectedAnswer == '', 'hover:bg-green-300  hover:scale-105': selectedAnswer !== '', 'shake border-red-400 bg-red-200 text-red-900': warnButton }"
            class=" text-center bg-green-200  border-4 border-yellow-400 rounded-xl text-xl lg:text-3xl  w-1/6 h-12 mb-9 lg:h-16 font-semibold    shadow-lg hover:shadow-xl ">
            Next
        </button>
    </div>
</template>

<script>
import { mapStores } from 'pinia'
import { useUserStore } from '../stores/UserStore'
import { useErrorStore } from '../stores/ErrorStore'

export default {
    props: ['id'],
    data() {
        return {
            //   id: this.$route.params.id
            warnButton: false,
            quiz: {},
            questions: [],
            currentQuestionIndex: 0,
            currentQuestion: {},
            selectedAnswer: '',
            body: {},
            response: {},
            error: '', //or [] not to forget to output a div with the error,
            score: 0

        }
    },
    computed: {
        // Pinia will add by the default the "Store" suffix to the id  of each store. Thus if the id is "userStore", I can access the store by this.userStoreStore. But I changed the id to simply store.
        ...mapStores(useUserStore, useErrorStore)
        // title(){
        //     return quiz.title
        // } 
    },
    methods: {
        warnSelectAnswer() {
            this.warnButton = true
            setTimeout(() => {
                this.warnButton = false
            }, 1500)
        }
        ,
        nextQuestion() {
            if (this.selectedAnswer == '') {
                this.warnSelectAnswer()
            } else {
                this.constructRequestBody()
                this.submitAnswer()
                this.resetData()
            }
        },
        handleClick(e) {

            if (this.selectedAnswer == e.target.__vnode.key) {
                this.selectedAnswer = ''
            } else {
                this.selectedAnswer = e.target.__vnode.key
            }


        },
        constructRequestBody() {
            this.body = {
                "data": {
                    "question_id": this.currentQuestion.id,
                    "answer": this.selectedAnswer,
                    "user_id": this.userStore.user.id
                }
            }

        }
        ,
        // TODO i should save  the state of the passed quizz as started, passed, or init. If the user clicks on a started quizz, he should start it from the last unanswered question
        // I should disable going back from the browser in router here 
        submitAnswer() {

            fetch('https://late-glitter-4431.fly.dev/api/v54/quizzes/' + this.quiz.id + "/submit", {
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
                        if (this.response.correct === true) {
                            this.score += 1
                        }
                        console.log(json)
                        console.log(this.score)
                    })
                } else {
                    response.json().then(json => {
                        console.log(json)
                        this.error = json.message
                        this.errorStore.setError(json.message[0])

                    })
                }
            })
        },
        // TODO I should go next question only if there are no errors from fetch()
        resetData() {
            this.currentQuestionIndex += 1
            this.currentQuestion = this.questions[this.currentQuestionIndex]
            this.selectedAnswer = ''

        }
    },

    mounted() {
        fetch('https://late-glitter-4431.fly.dev/api/v54/quizzes/' + this.id, {
            method: "GET", // *GET, POST, PUT, DELETE, etc.
            headers: {
                "X-Access-Token": "a68baa0fe20a17aea823776f987a2741395d24402430e4e2296bce48f56310ac"

            }
        }).then(response => {
            if (response.ok) {
                response.json().then(data => {

                    this.quiz = data
                    this.questions = data.questions
                    this.currentQuestion = this.questions[this.currentQuestionIndex]
                    console.log(this.quiz)
                })
            } else {
                response.json().then(json => {
                    console.log(json)
                    this.error = json.message
                    this.errorStore.setError(json.message)

                })
            }
        })

    }
}
</script>

<style>
.shake {
    animation: shake 0.5s cubic-bezier(0.36, 0.07, 0.19, 0.97) both;
    transform: translate3d(0, 0, 0);
}

@keyframes shake {

    10%,
    90% {
        transform: translate3d(-1px, 0, 0);
    }

    20%,
    80% {
        transform: translate3d(2px, 0, 0);
    }

    30%,
    50%,
    70% {
        transform: translate3d(-4px, 0, 0);
    }

    40%,
    60% {
        transform: translate3d(4px, 0, 0);
    }
}</style>