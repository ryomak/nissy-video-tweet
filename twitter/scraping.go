package twitter

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var nissyPicNum int = 10
var AAAPicNum int = 10

type News struct {
	ID    string
	Text  string
	Date  string
	URL   string
	Image string
	Tags  string
}

func ScrapeNissy() ([]News, error) {
	url := "http://avex.jp/nissy/news/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return []News{}, err
	}
	articles := []News{}
	doc.Find("dd").Each(func(_ int, s *goquery.Selection) {
		article := News{ID: "Nissy"}
		article.Text = s.Find("a").Text()
		newsURL, _ := s.Find("a").Attr("href")
		article.URL = url + newsURL
		article.Tags = " #Nissy #AAA #NissyEntertainment2ndLIVE"
		articles = append(articles, article)
	})
	doc.Find("dt").Each(func(index int, s *goquery.Selection) {
		articles[index].Date = s.Find("time").Text()
		articles[index].Image = "./media/images/Nissy/" + strconv.Itoa(index%nissyPicNum) + ".jpeg"
	})
	return articles, nil
}

func ScrapeAAA() ([]News, error) {
	url := "http://avex.jp/aaa/news/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return []News{}, err
	}
	articles := []News{}
	doc.Find("dd").Each(func(_ int, s *goquery.Selection) {
		article := News{ID: "AAA"}
		article.Text = s.Find("a").Text()
		newsURL, _ := s.Find("a").Attr("href")
		article.URL = url + newsURL
		article.Tags = " #AAA"
		articles = append(articles, article)
	})
	doc.Find("dt").Each(func(index int, s *goquery.Selection) {
		articles[index].Date = s.Find("span").Text() + "-" + s.Find("time").Text()
		articles[index].Image = "./media/images/AAA/" + strconv.Itoa(index%AAAPicNum) + ".jpeg"
	})
	return articles, nil
}

func ScrapeAtae() ([]News, error) {
	url := "http://avex.jp/atae/news/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return []News{}, err
	}
	articles := []News{}
	doc.Find(".item").Each(func(index int, s *goquery.Selection) {
		article := News{ID: "Atae"}
		article.Text = s.Find(".main > .title > a").Text()
		newsURL, _ := s.Find(".main > .title > a").Attr("href")
		article.URL = url + newsURL
		article.Date = s.Find(".main > .time").Text()
		article.Image, _ = s.Find(".sub > a > img").Attr("src")
		article.Tags = " #AAA #atae #shinjiro"
		articles = append(articles, article)
	})
	return articles, nil
}
