import { defineStore } from "pinia"

export const useErrorStore = defineStore('error', {
    state: () => ({
        error: {
            message: "",
            isError: false,
           
        }
    }),
  
    actions: {
        setError(message) {
            this.error.message = message
            this.error.isError = true
        },
        clearError(){
            this.error.message=""
            this.error.isError=false
        }
    },

})