package base

import (
	"github.com/capoverflow/ao3api-rod/internals/auth"
	"github.com/capoverflow/ao3api-rod/internals/models"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

var (
	Page *rod.Page
)

func Init(config models.RodConfig) {

	var Browser *rod.Browser

	if config.RemoteURL != "" {
		// connect to the remote browser
		Browser = rod.New().ControlURL(config.RemoteURL).MustConnect().NoDefaultDevice()

	} else {

		// 	// Create a launcher with headless option
		l := launcher.New().Headless(config.Headless)

		// Launch and connect to the browser
		url := l.MustLaunch()
		// log.Println("browser url:", url)
		Browser = rod.New().ControlURL(url).MustConnect()

	}

	// Set browser viewport size

	// Get Blank Page
	Page = Browser.MustPage("https://archiveofourown.org").MustWaitLoad()

	auth.Login(Page, config.Login)

	Page.MustWaitLoad().MustEval(`
	() => {
		var acceptedTOS = localStorage.getItem("accepted_tos");
	
		if (acceptedTOS === null || acceptedTOS === "") {
		  acceptedTOS = '20180523';
		  localStorage.setItem("accepted_tos", acceptedTOS);
		}
		
		console.log(acceptedTOS);
	};
	`)

	// reload the Page to make the localStorage work
	Page.MustReload().MustWaitLoad()

}
