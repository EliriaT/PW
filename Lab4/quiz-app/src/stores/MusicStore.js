import { defineStore } from "pinia"
import firstMusic from '../audio/music/1.mp3'
import secondMusic from '../audio/music/2.mp3'
import thirdMusic from '../audio/music/3.mp3'
import fourthMusic from '../audio/music/4.mp3'

export const useMusicStore = defineStore('music', {
    state: () => ({
        songs: [firstMusic, secondMusic, thirdMusic, fourthMusic],
        randomIndex: Math.floor(Math.random() * 4),


    }),
    getters: {
        getCurrentSong() {
            return new Audio(this.songs[this.randomIndex])
        },
        changeSong() {

           
            return new Audio(this.songs[this.randomIndex])
        }
    },
    actions: {
        changeMusicIndex() {
          
            this.randomIndex = this.randomIndex == 3 ? 0 : this.randomIndex + 1
        }

    },

})