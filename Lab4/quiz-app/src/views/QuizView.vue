<template>
    <h1>Quiz view {{ id }}</h1>
    <p>{{ quiz }}</p>
</template>

<script>
export default {
    props: ['id'],
    data() {
        return {
            //   id: this.$route.params.id
            quiz: null
        }
    },
    mounted() {
        fetch('https://late-glitter-4431.fly.dev/api/v54/quizzes/' + this.id, {
            method: "GET", // *GET, POST, PUT, DELETE, etc.
            headers: {
                "X-Access-Token": "a68baa0fe20a17aea823776f987a2741395d24402430e4e2296bce48f56310ac"

            }
        })
            .then(res => res.json())
            .then(data => this.quiz = data)
            .catch(err => console.log(err.message))
    }
}
</script>

<style></style>