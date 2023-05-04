
<template>
  <Transition name="modal">
    <div v-if="errorStore.error.isError" class="modal-mask">
      <div class="modal-container">
        <div class="modal-header">
          <slot name="header">default header</slot>
        </div>

        <div class="modal-body">
          <slot name="body">default body</slot>
        </div>

        <Button class="bg-red-300 hover:bg-red-400 w-1/2  h-10" @click="updatePage">Okay (ツ)</Button>

      </div>
    </div>
  </Transition>
</template>

<script>
import Button from './Button.vue'
import { mapStores } from 'pinia'
import { useErrorStore } from '../stores/ErrorStore'

export default {
  props: {
    show: Boolean
  },
  computed: {
    ...mapStores(useErrorStore)
  },
  methods:{
    updatePage(){
      this.errorStore.$reset()
      this.$router.go(0)
    }
  },
  components: { Button },
}
</script>

<style>
.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  transition: opacity 0.3s ease;

}

.modal-container {
  width: 30%;
  margin: auto;
  padding: 20px 30px;
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-items: center;
}

.modal-header h3 {
  text-align: center;
  margin-top: 0;
  color: #ca3123;
  font-size: 2rem;
}

.modal-body {
  margin: 20px 0;
  text-align: center;
  color: #ca3123;
}


/*
 * The following styles are auto-applied to elements with
 * transition="modal" when their visibility is toggled
 * by Vue.js.
 *
 * You can easily play with the modal transition by editing
 * these styles.
 */

.modal-enter-from {
  opacity: 0;
}

.modal-leave-to {
  opacity: 0;
}

.modal-enter-from .modal-container,
.modal-leave-to .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
</style>