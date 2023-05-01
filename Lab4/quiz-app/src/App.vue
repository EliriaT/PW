<template>
  <div @scroll="handleScroll" class="bg-green-100 h-screen w-screen overflow-auto pb-4 ">
    <Navbar :key="isOpen" />
    <router-view />
  </div>
</template>

<script>
import Navbar from "./components/Navbar.vue"

export default {
  name: 'Quiz-App',
  data() {
    return {
      isOpen: false
    }
  },
  components: {
    Navbar
  },
  methods: {

    handleScroll(e) {
      const { scrollTop, offsetHeight, scrollHeight } = e.target

      if (scrollTop <= 150) {
        this.isOpen = !this.isOpen
      }

    }
  },
}
</script>

<style></style>
