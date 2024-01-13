package base

import (
	"log"

	"github.com/capoverflow/ao3api-rod/internals/auth"
	"github.com/capoverflow/ao3api-rod/internals/models"
	"github.com/capoverflow/ao3api-rod/internals/utils"
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
	browser := rod.New().ControlURL(url).MustConnect()

	Page = browser.MustPage()

	switch {
	case len(config.Login.CookiesPath) > 0:
		log.Println("Using cookies")
		cookies := utils.GetCookieJar(config.Login.CookiesPath)
		config.Login.Cookies = utils.ConvertCookies(cookies)
		err := auth.LoginWithCookies(Page, config.Login.Cookies)
		if err != nil {
			log.Panic(err)
		}
	case len(config.Login.Username) > 0 && len(config.Login.Password) > 0:
		log.Println("Using username and password")
		Page.MustNavigate("https://archiveofourown.org/users/login").MustWaitLoad()
		err := auth.LoginWithCredentials(Page, config.Login)
		if err != nil {
			log.Panic(err)
		}

	}

	Page.Eval(`localStorage.setItem("accepted_tos", "20180523");`)

	// reload the Page to make the localStorage work
	Page.MustReload().MustWaitLoad()

	Page.MustWaitLoad()
}
