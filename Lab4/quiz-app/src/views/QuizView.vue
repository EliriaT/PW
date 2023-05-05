<template>
    <div class="pt-14 overflow-auto md:px-36 flex flex-col px-10  items-center  ">

        <div v-if="!loaded" class="mt-28 ">
            <h1 class="text-green-800 text-5xl">Loading...</h1>
            <svg class="animate-spin h-36 w-36 text-green-800 ml-3 mt-6" xmlns="http://www.w3.org/2000/svg" fill="none"
                viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                <path strokeLinecap="round" strokeLinejoin="round"
                    d="M15.182 15.182a4.5 4.5 0 01-6.364 0M21 12a9 9 0 11-18 0 9 9 0 0118 0zM9.75 9.75c0 .414-.168.75-.375.75S9 10.164 9 9.75 9.168 9 9.375 9s.375.336.375.75zm-.375 0h.008v.015h-.008V9.75zm5.625 0c0 .414-.168.75-.375.75s-.375-.336-.375-.75.168-.75.375-.75.375.336.375.75zm-.375 0h.008v.015h-.008V9.75z" />
            </svg>
        </div>



        <h1 v-if="quiz.title" class="text-5xl font-bold text-center text-red-900 ">{{ quiz.title }}</h1>

        <div v-if="currentQuestion && quiz.title" class="text-gray-600 font-semibold md:text-2xl text-lgtext-center py-4">
            Currently at question {{ currentQuestionIndex + 1 }} of {{ questions.length }}
        </div>

        <h2 v-if="currentQuestion && quiz.title"
            class="text-center font-bold xl:text-4xl xl:px-40 xl:leading-normal  lg:text-3xl lg:leading-normal lg:px-2 md:text-2xl md:leading-normal md:px-2  sm:text-2xl sm:leading-normal sm:px-2 text-xl py-10  px-2 ">
            {{ htmlDecode(currentQuestion.question) }}</h2>

        <h2 v-if="!currentQuestion" class="text-center text-6xl p-9 text-red-800 mt-24">Score: <span
                class="block text-center text-5xl p-9 text-green-800"
                :class="{ 'text-red-600': (this.score / this.questions.length * 100) < 50 }"> {{ (this.score /
                    this.questions.length * 100).toFixed(2) }} % </span></h2>

        <form v-if="currentQuestion" @submit.prevent="handleSubmit"
            class="grid sm:grid-cols-1 md:grid-cols-1 lg:grid-cols-2 xl:grid-cols-2 gap-10 justify-items-center pb-16 pt-10">
            <button
                class=" transition duration-200 ease-out bg-white hover:bg-gray-100 border-2 border-green-700 text-2xl p-8 m-4 rounded-lg font-bold min-w-full  min-h-full overflow-hidden shadow-xl hover:shadow-lg px-6 py-4  hover:scale-110"
                v-for="answer in currentQuestion.answers" :key="answer"
                :class="{ 'bg-yellow-100 border-8 border-yellow-300  ': selectedAnswer == answer, 'shake border-red-400 bg-red-50': warnButton }"
                @click="handleClick">
                {{ htmlDecode(answer) }}
            </button>
        </form>
        <button v-if="currentQuestion && quiz.title" @click="nextQuestion"
            :class="{ 'cursor-not-allowed': selectedAnswer == '', 'hover:bg-green-300  hover:scale-105': selectedAnswer !== '', 'shake border-red-400 bg-red-200 text-red-900': warnButton }"
            class=" text-center bg-green-200  border-4 border-yellow-400 rounded-xl text-xl lg:text-3xl  w-1/6 h-12  sm:w-1/2 lg:w-1/6  mb-9 lg:h-16 font-semibold    shadow-lg hover:shadow-xl  ">
            Next
        </button>
    </div>
</template>

<script>
import { mapStores } from 'pinia'
import { useUserStore } from '../stores/UserStore'
import { useErrorStore } from '../stores/ErrorStore'
import { useQuizzesStore } from '../stores/QuizStore'
import clickAns from '../audio/cartoon-jump.mp3'
import notSelectedAns from '../audio/error.mp3'
import nextQuestion from '../audio/game-complete.mp3'

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
            score: 0,
            loaded: false

        }
    },
    // apiKey is provided from main app and injected in this component. Accessable on this.apiKey
    inject: ['apiKey'],
    computed: {
        // Pinia will add by the default the "Store" suffix to the id  of each store. Thus if the id is "userStore", I can access the store by this.userStoreStore. But I changed the id to simply store.
        ...mapStores(useUserStore, useErrorStore, useQuizzesStore)

    },
    methods: {
        warnSelectAnswer() {
            const audio = new Audio(notSelectedAns)
            audio.play()
            this.warnButton = true
            setTimeout(() => {
                this.warnButton = false
            }, 1500)
        },
        nextQuestion() {

            if (this.selectedAnswer == '') {
                this.warnSelectAnswer()
            } else {
                const audio = new Audio(nextQuestion)
                audio.play()
                this.constructRequestBody()
                this.submitAnswer()
                this.resetData()

            }
        },
        handleClick(e) {
            const audio = new Audio(clickAns)
            audio.play()

            if (this.selectedAnswer == e.target.innerText) {
                this.selectedAnswer = ''
            } else {
                this.selectedAnswer = e.target.innerText
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
        submitAnswer() {

            fetch('https://late-glitter-4431.fly.dev/api/v54/quizzes/' + this.quiz.id + "/submit", {
                method: "POST",
                headers: {
                    "X-Access-Token": this.apiKey,
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
                        console.log(this.score)
                        this.quizzesStore.saveAnswer(this.userStore.user.id, this.id, this.score)
                        // console.log(json)
                        // console.log(this.score)
                    })
                } else {
                    response.json().then(json => {
                        console.log(json)
                        this.error = json.message
                        this.errorStore.setError(json.message[0])

                    })
                }
            }).catch(err => {
                console.log(err.message)
                this.errorStore.setError(err.message)
            })
        },
        // TODO I should go next question only if there are no errors from fetch()
        resetData() {
            this.currentQuestionIndex += 1
            this.currentQuestion = this.questions[this.currentQuestionIndex]
            this.selectedAnswer = ''
        },
        htmlDecode(input) {
            var doc = new DOMParser().parseFromString(input, "text/html");
            return doc.documentElement.textContent;
        },

    },

    mounted() {
        fetch('https://late-glitter-4431.fly.dev/api/v54/quizzes/' + this.id, {
            method: "GET", // *GET, POST, PUT, DELETE, etc.
            headers: {
                "X-Access-Token": this.apiKey

            }
        }).then(response => {
            if (response.ok) {
                response.json().then(data => {

                    this.quiz = data
                    this.questions = data.questions

                    const quiz = this.quizzesStore.getQuizState(this.userStore.user.id, this.id)


                    if (!quiz.lastAnsweredIndex) {
                        this.currentQuestionIndex = 0
                    } else {
                        this.currentQuestionIndex = quiz.lastAnsweredIndex
                    }

                    this.score = quiz.score

                    this.currentQuestion = this.questions[this.currentQuestionIndex]

                    this.loaded = true
                  
                })
            } else {
                response.json().then(json => {
                    console.log(json)
                    this.error = json.message
                    this.errorStore.setError(json.message)
                    console.log(this.apiKey)
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
}
</style>