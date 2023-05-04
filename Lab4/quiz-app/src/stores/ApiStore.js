import { defineStore } from "pinia"

export const useApiStore = defineStore('apiKey', {
    state: () => ({
        apiKey: process.env.VUE_APP_API_KEY
    }),
  
    actions: {
        setApiKey(key) {
            this.apiKey = key

        }
    },
    persist: {
        storage: sessionStorage,
      },

})