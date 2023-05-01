import { defineStore } from "pinia"

export const useQuizzesStore = defineStore('quizzes', {
    state: () => ({
        quizzes: [{
            quizId: 0,
            questionsCount: 0,
            lastAnsweredIndex: 0,
            score: -1 //-1 means that the quiz is not finished
        }]

    }),
    getters: {
        getQuizState(quizId) {
            return this.quizzes.filter((q) => q.quizId === quizId)
        },
        // getName() {
        //     return this.user.name + " " + this.user.surname
        // }
    },
    actions: {
        newQuiz(quizId, questionsCount) {
            this.quizzes.push({
                quizId: quizId,
                questionsCount: questionsCount,
                lastAnsweredIndex: 0,
                score: -1
            })

        },
        saveAnswer(quizId, currentScore) {
            quiz = this.quizzes.filter((q) => q.quizId === quizId)
            this.quizzes = this.quizzes.filter((q) => q.quizId !== quizId)
            quiz.currentScore = currentScore
            quiz.lastAnsweredIndex += 1
            this.quizzes.push(quiz)
        }
    },
    persist: true

})