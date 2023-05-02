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
                score: -1 //-1 means that the quiz is not finished
            }]
        }]
    }),
    // get a quiz state of the currently logged in user
    getters: {
        getUserResults(userId){
            allUserResults = this.allResults.filter((r) => r.userId === userId)
            return allUserResults
        },
        getQuizState(userId, quizId) {
            allUserResults = this.getUserResults(userId)
            quiz = allUserResults.userResults.filter((q) => q.quizId === quizId)
            return quiz
        }
    },
    actions: {
        // this should be called on click of a quiz
        // this will mark a quiz as started for a user
        newQuiz(userId, quizId, questionsCount) {
            allUserResults = this.getUserResults(userId)
            // Is it passed by reference??? to do: test it
            // because if not, I ought to substitute the allUserResults of that user in allResults
            allUserResults.userResults.push({
                quizId: quizId,
                questionsCount: questionsCount,
                lastAnsweredIndex: 0,
                score: -1
            })
           

        },
        saveAnswer(userId,quizId, currentScore) {
            allUserResults = this.getUserResults(userId)
            // again is it by reference?
            quiz = allUserResults.userResults.filter((q) => q.quizId === quizId)

            // quiz = this.quizzes.filter((q) => q.quizId === quizId)
            // this.quizzes = this.quizzes.filter((q) => q.quizId !== quizId)
            quiz.currentScore = currentScore
            quiz.lastAnsweredIndex += 1
            // this.quizzes.push(quiz)
        }
    },
    persist: true

})