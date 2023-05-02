import { defineStore } from "pinia"
import { useQuizzesStore } from '../stores/QuizStore'

// used to store the currently logged in user
export const useUserStore = defineStore('user', {
    state: () => ({
        user: {
            name: "",
            surname: "",
            id: 0
        }
        //name: "example"
    }),
    getters: {
        isLoggedIn() {
            return this.user.id != 0
        },
        getName() {
            return this.user.name + " " + this.user.surname
        }
    },
    actions: {
        setUser(name, surname, id) {
            const quizzesStore = useQuizzesStore()
            this.user.name = name
            this.user.surname = surname
            this.user.id = id
            quizzesStore.addNewUser(id)
        }
    },
    persist: true

})