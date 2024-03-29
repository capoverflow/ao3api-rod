package fanfic

import (
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/capoverflow/ao3api-rod/internals/models"
	"github.com/go-rod/rod"
)

func GetFanficChapters(work models.Work, page *rod.Page) (models.Work, error) {

	html := page.MustWaitIdle().MustHTML()

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println(err)
	}

	var chapters []models.Chapter

	doc.Find("ol.chapter.index.group li").Each(func(i int, s *goquery.Selection) {
		chapter := models.Chapter{}

		str, _ := s.Find("a").Attr("href")
		splitStr := strings.Split(str, "/")
		chapter.ChapterID = splitStr[len(splitStr)-1]

		chapter.Name = strings.TrimSpace(s.Find("a").Text())

		dateStr := s.Find("span.datetime").Text()
		dateStr = strings.Trim(dateStr, "()")
		chapter.Date, _ = time.Parse("2006-01-02", dateStr)

		chapters = append(chapters, chapter)
	})

	// log.Println(work.Title)

	if work.Title == "" {
		log.Println("Title is empty")
		doc.Find("h2.heading").Each(func(i int, el *goquery.Selection) {
			el.Find("a").Each(func(z int, s *goquery.Selection) {
				attr, _ := s.Attr("href")
				if strings.Contains(attr, "works") {
					work.Title = strings.TrimSpace(s.Text())
				} else if strings.Contains(attr, "users") {
					work.Author = append(work.Author, strings.TrimSpace(s.Text()))
				}
			})
		})
	}

	work.Chapters = chapters

	return work, nil
}
