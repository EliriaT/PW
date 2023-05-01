<template>
    <div class="py-14 overflow-auto md:px-36 flex flex-col px-10 items-center">
        <h1 class="text-5xl font-bold text-center ">{{ quiz.title }}</h1>



        <div class="text-gray-800 font-semibold text-lg text-center py-4">
            Currently at question {{ currentQuestionIndex + 1 }} of {{ questions.length }}
        </div>

        <h2 class="text-center font-bold lg:text-5xl lg:leading-normal lg:px-48 md:text-4xl md:leading-normal md:px-20  sm:text-2xl sm:leading-normal sm:px-2 text-xl py-10  px-2 ">{{ currentQuestion.question }}</h2>

        <form v-if="currentQuestion" class="grid sm:grid-cols-1 md:grid-cols-1 lg:grid-cols-2 xl:grid-cols-2 gap-10 justify-items-center pb-10 pt-10">
            <button
                class=" bg-yellow-300 text-2xl p-8 m-4 rounded-lg font-bold min-w-full  min-h-full overflow-hidden shadow-md hover:shadow-lg px-6 py-4  "
                v-for="answer in currentQuestion.answers" :key="answer" v-html="answer"
                @click.prevent="handleClick"></button>
        </form>
        <button class="bg-yellow-500 rounded-xl  shadow-md hover:shadow-lg text-xl py-4 px-10 w-1/6 font-bold mt-6" >Next</button>
    </div>
</template>

<script>
export default {
    props: ['id'],
    data() {
        return {
            //   id: this.$route.params.id
            quiz: {},
            questions: [],
            currentQuestionIndex: 0,
            currentQuestion: {},
            selectedAnswer: ''
        }
    },
    // computed:{
    //     title(){
    //         return quiz.title
    //     } 
    // },
    mounted() {
        fetch('https://late-glitter-4431.fly.dev/api/v54/quizzes/' + this.id, {
            method: "GET", // *GET, POST, PUT, DELETE, etc.
            headers: {
                "X-Access-Token": "a68baa0fe20a17aea823776f987a2741395d24402430e4e2296bce48f56310ac"

            }
        })
            .then(res => res.json())
            .then(data => {
                this.quiz = data
                this.questions = data.questions
                this.currentQuestion = this.questions[this.currentQuestionIndex]
            })
            .catch(err => console.log(err.message))
    }
}
</script>

<style></style>