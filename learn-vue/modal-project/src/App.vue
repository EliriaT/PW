<template>
  <h1>Hi {{ title }}</h1>
  <input type="text" ref="name">
  <button @click="handleClick">click me</button>
<!-- .modals . because it is a class -->
  <teleport to=".modals" v-if="showModal">
    <!-- props for data, slots for templates -->
    <Modal_m :heading="heading" :text="text" theme="sale" @close="toggleModel">
      <h1>{{ heading }}</h1>
      <p>{{ text }}</p>
      <template v-slot:links>
        <a href="#">Sign up now</a>
        <a href="#">More info</a>
      </template>

    </Modal_m>
  </teleport>


  <div v-if="showModal2">
    <!-- props for data, slots for templates -->
    <Modal_m :heading="heading" :text="text" theme="sale" @close="toggleModel2">
      <h1>{{ heading }}</h1>
      <p>{{ text }}</p>
      <a href="#">Sign up now</a>
      <a href="#">More info</a>


    </Modal_m>
  </div>
  <button @click.left="toggleModel">Show model</button>
  <button @click.left="toggleModel2">Show model2</button>
</template>

<script>
import Modal_m from './components/Modal-m'
export default {
  name: 'App', //optional
  components: { Modal_m },
  data() {
    return {
      title: "My first vue little project",
      heading: 'Sign up for the Givaway!',
      text: 'Grab your ninja swag for half price!',
      showModal: false,
      showModal2: false
    }
  },
  methods: {
    handleClick() {
      console.log(this.$refs.name)
      this.$refs.name.classList.add('active')
      this.$refs.name.focus()
    },
    toggleModel() {
      this.showModal = !this.showModal
    },
    toggleModel2() {
      this.showModal2 = !this.showModal2
    }
  }

}
</script>

<!-- styles are global for every global element -->
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

h1 {
  border-bottom: 1px solid #ddd;
  display: inline-block;
  padding-bottom: 10px;
}
</style>
