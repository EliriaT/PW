package guardianNews

import (
	"encoding/json"
	"fmt"
	news "github.com/EliriaT/News-Tg-Bot/events/news-client"
	e "github.com/EliriaT/News-Tg-Bot/lib/error"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const (
	newsAmount = 5
	format     = "json"
)

var StatusNotOkError = fmt.Errorf("status is not ok.")

type GuardianClient struct {
	token    string
	host     string
	basePath string
	client   http.Client
}

func NewGuardianClient(host string, basePath string, token string) news.NewsClient {
	return GuardianClient{
		host:     host,
		basePath: basePath,
		client:   http.Client{},
		token:    token,
	}
}

func (g GuardianClient) SearchNews(topic string) (news.News, error) {
	topic = url.QueryEscape(topic)

	q := url.Values{}
	q.Add("api-key", g.token)
	if topic != "" {
		q.Add("q", topic)
	}
	q.Add("page-size", strconv.Itoa(newsAmount))
	q.Add("format", format)

	data, err := g.doRequest(q)

	if err != nil {
		return news.News{}, e.Wrap("could not search news", err)
	}

	var res response
	if err := json.Unmarshal(data, &res); err != nil {
		return news.News{}, e.Wrap("could not unmarshal news", err)
	}

	if res.Response.Status != "ok" {
		return news.News{}, e.Wrap("could not search news: ", StatusNotOkError)
	}

	newsSlice := make([]news.NewsLinks, 0, len(res.Response.Results))
	newsRes := news.News{
		Amount: newsAmount,
	}

	for _, v := range res.Response.Results {
		newsSlice = append(newsSlice, news.NewsLinks{
			Id:     v.ID,
			Title:  v.WebTitle,
			WebUrl: v.WebUrl,
		})
	}
	newsRes.NewsLinks = newsSlice
	return newsRes, nil
}

func (g GuardianClient) NewsToString(news news.News) string {
	stringNews := ""
	for _, n := range news.NewsLinks {
		stringNews = fmt.Sprintf("%s \n\n\n %s \n %s", stringNews, n.Title, n.WebUrl)

	}
	return stringNews
}

func (c GuardianClient) doRequest(query url.Values) (data []byte, err error) {
	defer func() { err = e.WrapIfErr("can't do request", err) }()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   c.basePath,
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
