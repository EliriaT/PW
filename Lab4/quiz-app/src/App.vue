<template>
  <div @scroll="handleScroll" class="bg-green-100 h-screen w-screen overflow-auto pb-4 ">
    <Navbar :key="isOpen" />
    <Teleport to="body">
      <!-- use the modal component, pass in the prop -->
      <modal >
        <template #header>
          <h3>Sorry! An error happened! </h3>
        </template>
        <template #body>
          <h3>{{errorStore.error.message}} </h3>
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
  computed:{
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

    }
  },
}
</script>

<style></style>
