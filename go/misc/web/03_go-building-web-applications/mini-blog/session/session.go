package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"myapp/data"
	"myapp/db"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	SessionStore = sessions.NewCookieStore([]byte("our-social-network-application"))
	UserSession  = data.Session{
		Authenticated:   false,
		Unauthenticated: true,
	}
)

func GenerateSessionId() string {
	sid := make([]byte, 24)
	_, err := io.ReadFull(rand.Reader, sid)
	if err != nil {
		log.Fatal("Couldn't generate a session id because of:", err)
	}
	return base64.URLEncoding.EncodeToString(sid)
}

func ValidateSession(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, "app-session")
	if sid, valid := session.Values["sid"]; valid {
		currentUID := db.GetSessionUID(sid.(string))
		db.UpdateSession(sid.(string), currentUID)
		UserSession.Id = sid.(string)
	} else {
		newSID := GenerateSessionId()
		session.Values["sid"] = newSID
		session.Save(r, w)
		UserSession.Id = newSID
		db.UpdateSession(newSID, 0)
	}
	fmt.Println(UserSession.Id, UserSession.User)
}
