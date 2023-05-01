<template>
  <div class="bg-green-100 h-screen w-screen overflow-auto pb-11 ">
    <Navbar />
    <router-view />
  </div>
</template>

<script>
import Navbar from "./components/Navbar.vue"

export default {
  name: 'Quiz-App',
  components: {
    Navbar
  }
}
</script>

<style></style>
