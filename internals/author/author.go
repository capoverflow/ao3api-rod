package author

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/capoverflow/ao3api-rod/internals/models"
	"github.com/go-rod/rod"
)

func GetAuthorDashboard(author models.Author, page *rod.Page) models.Author {
	params := author.AuthorParams
	log.Println(len(params.Author))
	if params.Addr == "" {
		log.Println("No address provided, using default")
		params.Addr = "archiveofourown.org"
	}
	if params.Author == "" {
		log.Println("No author provided exiting")
		return author
	}

	if params.Addr != "" {
		u, err := url.Parse(params.Addr)
		if err != nil {
			log.Println(err)
		}
		params.Addr = u.Host
	}
	author_url := fmt.Sprintf("https://%s/users/%s", params.Addr, params.Author)

	page.MustNavigate(author_url).MustWaitLoad()

	html := page.MustHTML()

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println(err)
	}

	doc.Find("li.pseud > ul > li").Each(func(i int, el *goquery.Selection) {
		if !strings.Contains(el.Text(), "All Pseuds (2)") {
			author.Profile.Pseuds = append(author.Profile.Pseuds, el.Text())
		}
	})

	doc.Find("#user-fandoms > ol > li").Each(func(i int, el *goquery.Selection) {
		title := strings.TrimSpace(el.Text())
		re := regexp.MustCompile(`\((\d+)\)$`)
		matches := re.FindStringSubmatch(title)
		var NumberOfWorks int
		if len(matches) > 0 {
			NumberOfWorks, err = strconv.Atoi(matches[1])
			if err != nil {
				log.Println(err)
			}

		}

		fandom_url, _ := el.Find("a").Attr("href")

		author.Fandoms = append(author.Fandoms, models.AuthorFandom{
			Fandom:        title,
			URL:           fandom_url,
			NumberOfWorks: NumberOfWorks,
		})
	})

	doc.Find("#user-works.work.listbox.group > ul.index.group > li").Each(func(i int, el *goquery.Selection) {

		log.Printf("Work %d", i)

		var work models.Work

		el.Find("div.header.module > h5.fandoms.heading > a").Each(func(_ int, a *goquery.Selection) {
			title := a.Text()

			work.Fandom = append(work.Fandom, strings.TrimSpace(title))
		})

		el.Find("h4.heading > a").Each(func(_ int, b *goquery.Selection) {
			title := b.Text()
			href, _ := b.Attr("href")
			switch {
			case strings.Contains(href, fmt.Sprintf("/users/%s", params.Author)):
				work.Author = append(work.Author, strings.TrimSpace(title))
			case strings.Contains(href, "/works/"):
				work.Title = title
				work.URL = href
			}
		})

		summary := strings.TrimSpace(el.Find("blockquote.userstuff.summary").Text())

		work.Summary = append(work.Summary, summary)

		dateTxt := el.Find("p.datetime")
		if dateTxt.Length() > 0 {

			date, err := time.Parse("_2 Jan 2006", dateTxt.Text())
			if err != nil {
				log.Println(err)
			}
			work.Updated = date
		}

		el.Find("ul.tags.commas > li").Each(func(i int, c *goquery.Selection) {
			tag := strings.TrimSpace(c.Text())
			hrf, _ := c.Find("a").Attr("href")
			// url escape the hrf
			hrf, err = url.QueryUnescape(hrf)
			if err != nil {
				log.Println(err)
			}

			switch {
			case strings.Contains(hrf, "*a*"):
				work.Relationship = append(work.Relationship, tag)
			case strings.Contains(hrf, "*s*"):
				work.Relationship = append(work.Relationship, tag)
			default:
				work.AlternativeTags = append(work.AlternativeTags, tag)
			}
		})

		el.Find("ul.required-tags > li").Each(func(i int, s *goquery.Selection) {
			work.RequiredTags = append(work.RequiredTags, strings.TrimSpace(s.Text()))
		})

		doc.Find("dl.stats dt").Each(func(i int, s *goquery.Selection) {
			label := s.Text()
			value := s.Next().Text()

			switch label {
			case "Language:":
				work.Language = value
			case "Words:":
				work.Words = value
			case "Chapters:":
				// Assuming you want the total number of chapters
				work.NBChapters = strings.Split(value, "/")[1]
			case "Comments:":
				work.Comments = value
			case "Kudos:":
				work.Kudos = value
			case "Bookmarks:":
				work.Bookmarks = value
			case "Hits:":
				work.Hits = value
			}
		})

		author.Works = append(author.Works, work)

	})

	return author

}

func GetAuthorWorks() {}

func GetAuthorWorksByFandom() {}

func GetAuthorBookmarks() {}

func GetAuthorGifts() {}

func GetAuthorSeries() {}

func GetAuthorProfile() {
	url := fmt.Sprintf("https://archiveofourown.org/users/%s/profile")
	log.Println(url)
}
