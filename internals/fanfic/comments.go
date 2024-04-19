package fanfic

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/capoverflow/ao3api-rod/internals/models"
	"github.com/go-rod/rod"
)

func GetFanficComments(page *rod.Page, work models.Work) models.Work {

	// elm := page.MustElementByJS(`document.querySelector("#show_comments_link_top > a"`), nil)

	page.MustElementByJS("() => document.querySelector('#show_comments_link_top > a')", nil).MustClick()

	commentsPageNumber := 1

	commentsPage := page.MustElementByJS("() => document.querySelector('#comments_placeholder > ol.pagination.actions')")

	commentsPageText := commentsPage.MustText()

	// Define a regular expression pattern to match numbers
	pattern := "[0-9]+"

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find all matches in the input string
	matches := re.FindAllString(commentsPageText, -1)

	if len(matches) == 0 {
		fmt.Println("No numbers found in the input string.")

	}

	// Convert the matched strings to integers and find the maximum
	max := 0
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			fmt.Printf("Error converting %s to an integer: %v\n", match, err)
			continue
		}
		if num > max {
			max = num
		}
	}
	// Assume the we are on the first page of comments
	// Get the html of the comments

	var comments models.Comments

	_ = comments

	max = 3

	// create a slice of with max as the length
	commentsSlice := make([]models.Comment, max)

	for i := range commentsSlice {

		log.Println(i, max)

		// document.querySelector("#comments_placeholder > ol.thread")
		commentsHtml := page.MustElementByJS("() => document.querySelector('#comments_placeholder > ol.thread')").MustHTML()

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(commentsHtml))
		if err != nil {
			log.Println(err)
		}

		doc.Find("li").Each(func(i int, s *goquery.Selection) {

			icon := s.Find("h4.heading.byline > div.icon")
			icon.Find("a").Each(func(i int, el *goquery.Selection) {
				// For each item found, get the band and title
				iconText := el.Text()
				log.Println(iconText)
				href := el.AttrOr("href", "")
				log.Println(href)
			})

			log.Println(icon.Text())

			author := s.Find("h4.heading.byline > a").Text()
			if len(author) != 0 {
				log.Println(author)
			}

			parent := s.Find("h4.heading.byline > span.parent").Text()
			if len(parent) != 0 {
				parent = strings.TrimSpace(parent)
				log.Println(parent)
			}

			datetime := s.Find("h4.heading.byline > span.posted.datetime").Text()
			if len(datetime) != 0 {
				datetime = strings.ReplaceAll(datetime, "\n", "")
				re := regexp.MustCompile(`\s{3,}`)
				cleaned := re.ReplaceAllString(datetime, "  ")
				cleaned = strings.TrimSpace(cleaned)
				log.Println(cleaned)
			}
		})

	}
	log.Println("Comments page number: ", commentsPageNumber)

	fmt.Scanln()

	return work
}
