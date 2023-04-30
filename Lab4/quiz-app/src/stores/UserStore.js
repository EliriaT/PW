import { defineStore } from "pinia"

export const useUserStore = defineStore('userStore', {
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
            this.user.name = name
            this.user.surname = surname
            this.user.id = id

        }
    },
    persist: true

})