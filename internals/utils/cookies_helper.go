package utils

import (
	"log"
	"net/http"
	"time"

	cookiemonster "github.com/MercuryEngineering/CookieMonster"
	"github.com/go-rod/rod/lib/proto"
)

func GetCookieJar(cookiePath string) (cookies []*http.Cookie) {

	cookies, err := cookiemonster.ParseFile(cookiePath)
	if err != nil {
		log.Panic(err)
	}

	return cookies

}

func ConvertHTTPCookieToRodCookie(cookies []*http.Cookie) []proto.NetworkCookieParam {
	var rodCookies []proto.NetworkCookieParam

	for _, cookie := range cookies {
		rodCookie := proto.NetworkCookieParam{
			Name:     cookie.Name,
			Value:    cookie.Value,
			Path:     cookie.Path,
			Domain:   cookie.Domain,
			Secure:   cookie.Secure,
			HTTPOnly: cookie.HttpOnly,
			// URL, SameSite, Expires, Priority, SameParty, SourceScheme, SourcePort, PartitionKey need to be set based on your specific requirements
		}

		// Convert standard cookie expiration to TimeSinceEpoch if needed
		if cookie.Expires.IsZero() {
			// Handle zero time if necessary
		} else {
			rodCookie.Expires = proto.TimeSinceEpoch(cookie.Expires.UnixNano() / int64(time.Millisecond))
		}

		// Handle other fields like SameSite, Priority, etc. according to your logic

		rodCookies = append(rodCookies, rodCookie)
	}

	return rodCookies
}

func ConvertRodCookieToHTTPCookie(cookies []proto.NetworkCookie) []*http.Cookie {
	var httpCookies []*http.Cookie

	for _, cookie := range cookies {
		httpCookie := &http.Cookie{
			Name:     cookie.Name,
			Value:    cookie.Value,
			Path:     cookie.Path,
			Domain:   cookie.Domain,
			Secure:   cookie.Secure,
			HttpOnly: cookie.HTTPOnly,
		}

		// Convert TimeSinceEpoch to standard cookie expiration if needed
		if cookie.Expires == 0 {
			// Handle zero time if necessary
		} else {
			httpCookie.Expires = time.Unix(0, int64(cookie.Expires)*int64(time.Millisecond))
		}

		// Handle other fields like SameSite, Priority, etc. according to your logic

		httpCookies = append(httpCookies, httpCookie)
	}

	return httpCookies
}
