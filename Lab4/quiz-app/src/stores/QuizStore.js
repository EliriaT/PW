import { defineStore } from "pinia"

// I should store the result by user
// quizzesResultID

export const useQuizzesStore = defineStore('quizzes', {
    state: () => ({
        allResults: [{
            userId: 0,
            userResults: [{
                quizId: 0,
                questionsCount: 0,
                lastAnsweredIndex: 0,
                score: 0 
            }]
        }]
    }),
    // get a quiz state of the currently logged in user
    getters: {
        getUserResults() {
            return (userId) => {
                let allUserResults = this.allResults.filter((r) => r.userId == userId)
                if (allUserResults == undefined) {
                    return {}
                }
                return allUserResults[0]
            }

        },
        getQuizState() {
            return (userId, quizId) => {
                let allUserResults = this.getUserResults(userId)
                if (allUserResults == undefined) {
                    return {}
                }
                let quiz = allUserResults.userResults.filter((q) => q.quizId == quizId)
                if (quiz.length == 0) {
                    return {}
                }
                return quiz[0]
            }

        },
        // if the quiz is present in the store, it means it was started by the user and maybe finished
        isQuizStarted() {
            return (userId, quizId) => {
                let allUserResults = this.getUserResults(userId)
                if (allUserResults == undefined) {
                    return false
                }
                let quiz = allUserResults.userResults.filter((q) => q.quizId == quizId)
                if (quiz.length == 0) {
                    return false
                }
                if(quiz[0].lastAnsweredIndex>0){
                    return true
                }
                return false
            }


        },
        isQuizFinished() {
            return (userId, quizId) => {
                let allUserResults = this.getUserResults(userId)
                if (allUserResults == undefined) {
                    return [false, 0]
                }
                let quiz = allUserResults.userResults.filter((q) => q.quizId == quizId)
                if (quiz.length == 0) {
                    return [false, 0]
                }
                quiz = quiz[0]
                if (quiz.questionsCount == quiz.lastAnsweredIndex) {
                    return [true, quiz.score]
                }
                return [false, 0]
            }

        }
    },
    actions: {
        // this should be called on click of a quiz
        // this will mark a quiz as started for a user
        newQuiz(userId, quizId, questionsCount) {
            let allUserResults = this.getUserResults(userId)


            let quiz = this.getQuizState(userId, quizId)
            if (!quiz.quizId){
                 allUserResults.userResults.push({
                    quizId: quizId,
                    questionsCount: questionsCount,
                    lastAnsweredIndex: 0,
                    score: 0
                })
            }
               


        },
        saveAnswer(userId, quizId, currentScore) {
            let allUserResults = this.getUserResults(userId)
            let userResults = allUserResults.userResults
            let quiz = userResults.filter((q) => q.quizId == quizId)
            // console.log(quiz)
            // console.log(quizId)
            // console.log(userResults)

            quiz = quiz[0]
            quiz.score = currentScore
            quiz.lastAnsweredIndex += 1
        },
        addNewUser(userId) {
            let allUserResults = this.getUserResults(userId)
            if (allUserResults == undefined) {
                this.allResults.push({
                    userId: userId,
                    userResults: []
                })
            }
        },
        deleteUserInformation(userId){
            console.log(this.allResults)
            console.log("deleted!")
            this.allResults = this.allResults.filter((r) => r.userId !== userId)
            console.log(this.allResults)
        }
    },
    persist: true

})