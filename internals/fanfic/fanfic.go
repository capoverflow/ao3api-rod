package fanfic

import (
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/capoverflow/ao3api-rod/internals/models"
	"github.com/go-rod/rod"
)

func GetFanfic(page *rod.Page) (fanfic models.Work) {
	html := page.MustHTML()

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println(err)
	}

	doc.Find("div.chapter.preface.group > h3.title").Each(func(i int, el *goquery.Selection) {
		var splitStr []string
		val, exist := el.Find("a").Attr("href")
		if exist {
			splitStr = strings.Split(val, "/")
		}
		// log.Println()

		chapter := models.Chapter{
			// Date:      time.Time{},
			Name:      strings.TrimSpace(el.Text()),
			ChapterID: splitStr[len(splitStr)-1],
		}
		fanfic.Chapters = append(fanfic.Chapters, chapter)
	})

	doc.Find("dl.stats").Each(func(i int, s *goquery.Selection) {
		datestr := s.Find("dd.published").Text()
		layoutISO := "2006-01-02"
		fanfic.DatePublished, err = time.Parse(layoutISO, datestr)
		if err != nil {
			log.Println(err)
		}

		Updated := s.Find("dd.status").Text()
		if len(Updated) != 0 {
			fanfic.Updated, err = time.Parse(layoutISO, Updated)
			if err != nil {
				log.Println("error with time", err)
			}
		}
	})

	doc.Find("dl.stats > dd").Each(func(i int, el *goquery.Selection) {
		// log.Println(el.Text())
		n := el.Get(0) // Retrieves the internal *html.Node
		for _, a := range n.Attr {
			// log.Println(a.Val)
			switch a.Val {
			case "language":
				fanfic.Language = el.Text()
			case "words":
				fanfic.Words = el.Text()
			case "chapters":
				fanfic.NBChapters = el.Text()
			case "comments":
				fanfic.Kudos = el.Text()

			case "kudos":
				fanfic.Kudos = el.Text()

			case "bookmarks":
				fanfic.Kudos = el.Text()

			case "hits":
				fanfic.Hits = el.Text()
			}
		}

	})

	fanfic.Title = strings.TrimSpace(doc.Find("h2.title.heading").Text())

	doc.Find("h3.byline.heading").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, el *goquery.Selection) {
			fanfic.Author = append(fanfic.Author, strings.TrimSpace(el.Text()))
		})
	})

	doc.Find("div.summary.module").Each(func(i int, s *goquery.Selection) {
		s.Find("p").Each(func(i int, el *goquery.Selection) {
			fanfic.Summary = append(fanfic.Summary, el.Text())
		})
	})

	doc.Find("dd.fandom.tags").Each(func(i int, s *goquery.Selection) {
		s.Find("a.tag").Each(func(i int, el *goquery.Selection) {
			fanfic.Fandom = append(fanfic.Fandom, el.Text())
		})
	})

	doc.Find("dd.relationship.tags").Each(func(i int, s *goquery.Selection) {
		s.Find("a.tag").Each(func(i int, el *goquery.Selection) {
			fanfic.Relationship = append(fanfic.Relationship, el.Text())
		})
	})

	doc.Find("dd.freeform.tags").Each(func(i int, s *goquery.Selection) {
		s.Find("a.tag").Each(func(i int, el *goquery.Selection) {
			if len(el.Text()) != 0 {
				fanfic.AlternativeTags = append(fanfic.AlternativeTags, el.Text())
			}
		})
	})

	doc.Find("li.download").Each(func(_ int, s *goquery.Selection) {
		// log.Println(s.Find("a").Attr("href"))
		s.Find("a").Each(func(_ int, s *goquery.Selection) {
			href, ok := s.Attr("href")
			if ok {
				_, ok := s.Attr("class")
				if !ok {
					for p := s.Parent(); p.Size() > 0 && !ok; p = p.Parent() {
						_, ok = p.Attr("class")
					}
				}
				// log.Printf("Link #%d:\nhref: %s\ntext: %s\nclass: %s\n\n", i, href, s.Text(), classes)
				if !strings.Contains(href, "#") {
					// fanfic.Downloads = append(fanfic.Downloads, fmt.Sprintf("%s://download.%s%s", Params.Scheme, Params.Addr, href))

					// fanfic.Downloads = append(fanfic.Downloads, models.Download{
					// 	URL: href,
					// })

					var download models.Download

					// log.Println(href)
					u := url.URL{}
					u.Scheme = "https"
					u.Host = "archiveofourown.org"
					u.Path = href
					// log.Println(u.String())

					// use filetype and trimmed to get the filetype of the file

					path := u.Path

					// remove ? and everything after
					if strings.Contains(path, "?") {
						path = strings.Split(path, "?")[0]
					}
					filetype := strings.Split(path, ".")[1]

					// urls = append(urls, map[string]string{
					// 	"url":      u.String(),
					// 	"title":    page.MustInfo().Title,
					// 	"filetype": filetype,
					// })

					download.URL = u.String()
					download.Format = filetype

					fanfic.Downloads = append(fanfic.Downloads, download)

				}
			}
		})

	})

	// add additional code here

	return fanfic
}
