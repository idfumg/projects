package endpoints

import (
	"fmt"
	"myapp/data"
	"myapp/db"
	"myapp/session"
	"net/http"
)

func LoginPOST(w http.ResponseWriter, r *http.Request) {
	session.ValidateSession(w, r)

	u := data.User{}
	name := r.FormValue("user_name")
	pass := r.FormValue("user_password")
	password := weakPasswordHash(pass)

	err := db.Database.QueryRow("SELECT user_id, user_name FROM users WHERE user_name=? and user_password=?", name, password).Scan(&u.Id, &u.Name)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	} else {
		db.UpdateSession(session.UserSession.Id, u.Id)
		fmt.Fprintln(w, u.Name)
		return
	}
}
