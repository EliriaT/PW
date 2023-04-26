const app = Vue.createApp({
    // data, functions
    // the data is only visible in the #app component
    data() {
        return {
            url: "http://www.google.com",
            showBooks: true,
            books: [
                { title: 'name of the wind', author: 'patrick rothfuss', img: 'assets/1.jpg', isFav: true },
                { title: 'the way of kings', author: 'brandon sanderson', img: 'assets/2.jpg', isFav: false },
                { title: 'the final empire', author: 'brandon sanderson', img: 'assets/3.jpg', isFav: true },
            ],
        }
    },
    // 'this' references the component itself
    methods: {
        changeTitle(title) {
            this.title = title
        },
        toggleShowBooks() {
            this.showBooks = !this.showBooks
        },
        toggleFav(book) {
            book.isFav = !book.isFav
        }
    },
    //almost just like data, but depends on the some field from data.  Thus when data changes, computed data also changes.
    computed: {
        filteredBooks() {
            return this.books.filter(
                (book) => book.isFav
            )
        }
    }

})

app.mount('#app')