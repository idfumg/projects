package endpoints

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"myapp/db"
	"net/http"
	"regexp"
)

func weakPasswordHash(pass string) string {
	hash := sha1.New()
	io.WriteString(hash, pass)
	return string(hash.Sum(nil))
}

func RegisterPOST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}

	name := r.FormValue("user_name")
	email := r.FormValue("user_email")
	pass := r.FormValue("user_password")
	// pass2 := r.FormValue("user_password2")
	pageGUID := r.FormValue("referrer")

	gure := regexp.MustCompile("[^A-Za-z0-9]+")
	guid := gure.ReplaceAllString(name, "")

	password := weakPasswordHash(pass)

	query :=
		`INSERT INTO users
			(user_name, user_guid, user_email, user_password)
			VALUES
			(?, ?, ?, ?)`
	res, err := db.Database.Exec(query, name, guid, email, password)
	fmt.Println(res.RowsAffected())
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		http.Redirect(w, r, "/page/"+pageGUID, http.StatusMovedPermanently)
	}
}
