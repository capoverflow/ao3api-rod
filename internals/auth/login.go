package auth

import (
	"log"

	"github.com/capoverflow/ao3api-rod/internals/models"
	"github.com/capoverflow/ao3api-rod/internals/utils"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func IsLoggedIn() bool {
	return false
}

func IsLoginNeeded() bool {
	return false
}

func Login(Page *rod.Page, login models.Login) error {

	switch {
	case len(login.CookiesPath) > 0:
		log.Println("Using cookies")
		cookies := utils.GetCookieJar(login.CookiesPath)
		login.Cookies = utils.ConvertCookies(cookies)
		err := LoginWithCookies(Page, login.Cookies)
		if err != nil {
			log.Panic(err)
		}
	case len(login.Username) > 0 && len(login.Password) > 0:
		log.Println("Using username and password")
		Page.MustNavigate("https://archiveofourown.org/users/login").MustWaitLoad()
		err := LoginWithCredentials(Page, login)
		if err != nil {
			log.Panic(err)
		}

	}

	return nil
}

func LoginWithCookies(page *rod.Page, cookies []proto.NetworkCookieParam) error {
	log.Println("Login using cookies...")
	for _, cookie := range cookies {
		page.MustSetCookies(&cookie)
	}

	return nil
}

func LoginWithCredentials(page *rod.Page, login models.Login) error {
	log.Println("Using username and password")
	username := login.Username
	password := login.Password

	page.MustElement("#user_login").MustInput(username)

	page.MustElement("#user_password").MustInput(password)

	page.MustElement("#new_user > dl > dd.submit.actions > input").MustClick()

	page.MustWaitLoad()

	return nil
}

func Logout() error {
	return nil
}
