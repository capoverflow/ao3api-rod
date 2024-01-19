package base

import (
	"log"

	"github.com/capoverflow/ao3api-rod/internals/auth"
	"github.com/capoverflow/ao3api-rod/internals/models"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

var (
	Page *rod.Page
)

func Init(config models.RodConfig) {
	// 	// Create a launcher with headless option
	l := launcher.New().Headless(config.Headless)

	// Launch and connect to the browser
	url := l.MustLaunch()
	log.Println("browser url:", url)
	browser := rod.New().ControlURL(url).MustConnect()

	Page = browser.MustPage()

	auth.Login(Page, config.Login)

	Page.Eval(`localStorage.setItem("accepted_tos", "20180523");`)

	// reload the Page to make the localStorage work
	Page.MustReload().MustWaitLoad()

	Page.MustWaitLoad()
}
