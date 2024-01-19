package utils

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/capoverflow/ao3api-rod/internals/models"
)

func ContainsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func AppendIfNotExists(slice []string, str string) []string {
	if !ContainsString(slice, str) {
		log.Println("appending", str)
		slice = append(slice, str)
	}
	return slice
}

func IsWorkUrl(work string) bool {
	u, err := url.Parse(work)
	if err != nil {
		log.Println(err)
	}
	if u.Host == "archiveofourown.org" {
		if strings.Contains(u.Path, "works") {
			return true
		}
	}
	return false
}

func GetWorkID(val string) models.Work {

	var WorkID, ChapterID string

	// log.Println(val)

	re := regexp.MustCompile(`archiveofourown\.org\/works\/(?P<work_id>[0-9]+)(?:\/chapters\/(?P<chapter_number>[0-9]+))?`)

	match := re.FindStringSubmatch(val)
	// log.Println(match)
	if len(match) > 0 {
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				// fmt.Printf("%s: %s\n", name, match[i])
				if name == "work_id" {
					WorkID = match[i]
				} else if name == "chapter_number" {
					ChapterID = match[i]
				}
			}
		}
	}

	chapters := []models.Chapter{
		{ChapterID: ChapterID},
	}

	// log.Println(WorkID, ChapterID)
	return models.Work{
		WorkID:   WorkID,
		Chapters: chapters,
	}
	// return "", ""

}

// DeduplicateWorks takes a slice of Work and returns a new slice with duplicates removed
func DeduplicateWorks(works []models.Work) []models.Work {
	seen := make(map[string]bool)
	var uniqueWorks []models.Work

	for _, work := range works {
		if _, exists := seen[work.WorkID]; !exists {
			uniqueWorks = append(uniqueWorks, work)
			seen[work.WorkID] = true
		}
	}

	return uniqueWorks
}

func RandSleep(min, max int) {
	n := rand.Intn(max-min) + min
	log.Printf("Sleeping %d seconds...", n)
	for i := n; i > 0; i-- {
		fmt.Printf("\r%d...", i)
		time.Sleep(time.Second)
	}
	fmt.Println("\rDone")
}

func CountDown(n int) {
	log.Printf("Sleeping %d seconds...", n)
	for i := n; i > 0; i-- {
		fmt.Printf("\r%d...", i)
		time.Sleep(time.Second)
	}
	fmt.Println("\rDone")
}

// WaitEveryXIterations waits for a specified duration every x iterations and counts down
func WaitEveryXIterations(length, currentIndex, x, waitSeconds int) {
	if length <= 0 || x <= 0 {
		return // Avoid division by zero or negative values
	}
	if currentIndex%x == 0 {
		CountDown(waitSeconds)
	}
}
