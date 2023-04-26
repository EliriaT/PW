<template>
    <p>Reaction time - {{ score }} ms</p>
    <p class="rank">{{ rank }}</p>
</template>
  
<script>
export default {
    props: ['score'],
    data() {
        return {
            rank: null
        }
    },
    mounted() {
        if (this.score < 250) {
            this.rank = "Ninja skills"
        } else if (this.score < 500) {
            this.rank = "Wave skills"
        } else if (this.score < 600) {
            this.rank = "Turtle skills"
        }
    }
}

</script>

<style></style>