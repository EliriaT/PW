<template>

<div class="bg-green-100 h-screen w-screen overflow-auto pb-11 ">
  <Navbar/>
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

<style>
/* #app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
} */
</style>
