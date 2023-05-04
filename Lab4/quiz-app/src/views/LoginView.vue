<template>
    <div class=" w-1/4 absolute  bg-green-100 top-1/3 left-1/3 ml-24 ">
        <form class="bg-white shadow-md rounded px-8 pt-6 pb-8 " @submit.prevent="handleSubmit">

            <div class="mb-4">
                <label class="block text-gray-700 text-lg font-bold mb-2" for="Name">
                    Name
                </label>
                <input
                    class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    id="Name" type="text" placeholder="Name" v-model="name">
            </div>

            <div class="mb-4">
                <label class="block text-gray-700 text-lg font-bold mb-2" for="Surname">
                    Surname
                </label>
                <input
                    class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    id="Surname" type="text" placeholder="Surname" v-model="surname">
            </div>
            <div v-if="error">
                <p v-for="err in error" class="text-red-500 font-bold bg-red-100 rounded-md text-center p-2 mb-4"
                    :key="err">{{
                        err }}</p>
            </div>

            <div class="flex items-center justify-between">

                <button
                    class="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
                    Sign in
                </button>
            </div>

        </form>

    </div>
</template>

<script>

import { useUserStore } from '../stores/UserStore'
import { useErrorStore } from '../stores/ErrorStore'
import clickSound from '../audio/button.mp3'

export default {
    components: {

    },
    inject: ['apiKey'],
    name: 'RegisterView',
    data() {
        return {
            name: '',
            surname: '',
            error: [],
            users: [],
            loggedUser: {}

        }
    },
    watch: {

        name(newName, oldName) {
            this.error = this.error.filter((el) => el !== 'Name must be at least 3 characters long')
            this.name.length > 2 ?
                this.error : this.error.push('Name must be at least 3 characters long')


        },
        surname(newSurname, oldSurname) {
            this.error = this.error.filter((el) => el !== 'Surname must be at least 3 characters long')
            this.surname.length > 2 ?
                this.error : this.error.push('Surname must be at least 3 characters long')
        }
    },
    mounted() {
        fetch('https://late-glitter-4431.fly.dev/api/v54/users', {
            method: "GET",
            headers: {
                "X-Access-Token": this.apiKey,

            },
        }).then(response => {
            if (response.ok) {
                response.json().then(json => {
                    this.users = json
                })
            } else {
                response.json().then(json => {

                    this.error.push(json.message)
                    this.errorStore.setError(json.message)
                })
            }
        }).catch(err => {
            console.log(err.message)
            this.errorStore.setError(err.message)
        })
    },
    methods: {
        playSound() {
            const audio = new Audio(clickSound)
            audio.play()
        },
        handleSubmit() {
            this.playSound()
            if (this.error.length == 0 &&  this.surname!='' &&  this.name!='') {
                this.loggedUser = this.users.filter(user => user.name == this.name && user.surname == this.surname)
                if (this.loggedUser.length == 0) {
                    this.error.push('No such user!')
                } else {
                   
                    this.loggedUser = this.loggedUser[0]
                    this.userStore.setUser(this.loggedUser.name, this.loggedUser.surname, this.loggedUser.id)
                    this.name = ''
                    this.surname = ''
                    this.error = []
                    this.$router.push({ name: 'quizzes' })
                }

            }

        }
    }, setup() {
        const userStore = useUserStore()
        const errorStore = useErrorStore()
        return {
            userStore, errorStore
        }

    }
}

</script>
    
<style></style>