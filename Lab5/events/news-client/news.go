package news

type NewsClient interface {
	SearchNews(topic string) (News, error)
	NewsToString(news News) string
}

type News struct {
	Amount    int
	NewsLinks []NewsLinks
}

type NewsLinks struct {
	Id     string
	Title  string
	WebUrl string
}
