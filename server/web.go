package server

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
)

var permittedUsers string

func setWeb() {
	permittedUsers = os.Getenv("DISCORD_PERMITTED_USERS")

	gothic.Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	goth.UseProviders(discord.New(os.Getenv("DISCORD_KEY"), os.Getenv("DISCORD_SECRET"), os.Getenv("CALLBACK_URL")))

	p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		if user, err := gothic.CompleteUserAuth(res, req); err == nil {
			permittedCheck(user, res)
		} else {
			gothic.BeginAuthHandler(res, req)
		}
	})

	p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {
		user, err := gothic.CompleteUserAuth(res, req)
		chk(err)

		permittedCheck(user, res)
	})

	p.Get("/", func(res http.ResponseWriter, req *http.Request) {
		http.Redirect(res, req, "/auth/discord", http.StatusPermanentRedirect)
	})
}

func sendIndex(res http.ResponseWriter) {
	page, err := ioutil.ReadFile("./views/index.html")
	chk(err)

	res.Write(page)
}

func permittedCheck(user goth.User, res http.ResponseWriter) {
	for _, v := range strings.Split(permittedUsers, ",") {
		if v == user.UserID {
			sendIndex(res)
			return
		}
	}

	res.WriteHeader(http.StatusForbidden)
	res.Write([]byte("Forbidden"))
}
