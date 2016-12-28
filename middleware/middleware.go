package middleware

import (
	"net/http"
	"os"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET_KEY")))

// AuthMiddleware is run before all the handler to make sure user is logged in
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		session, err := store.Get(r, "session")
		if session.IsNew || session.Values["user"] == nil || err != nil {
			//not logged in, show login page
			log.Errorf(ctx, "AuthMiddleware : Error in retrieving session", err, session)
			http.Redirect(w, r, "/", 302)
			return
		}
		h.ServeHTTP(w, r)
	})

}
