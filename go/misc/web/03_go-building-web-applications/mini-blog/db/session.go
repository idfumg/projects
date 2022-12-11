package db

import (
	"fmt"
	"myapp/data"
	"time"
)

func GetSessionUID(sid string) int {
	user := data.User{}

	err := Database.QueryRow("SELECT user_id FROM sessions WHERE session_id=?", sid).Scan(&user.Id)
	if err != nil {
		return 0
	}

	return user.Id
}

func UpdateSession(sid string, uid int) {
	const timeFmt = "2006-01-02T15:04:05.999999999"
	tstamp := time.Now().Format(timeFmt)
	_, err := Database.Exec("INSERT INTO sessions SET session_id=?, user_id=?, session_active=1, session_update=? ON DUPLICATE KEY UPDATE user_id=?, session_update=?", sid, uid, tstamp, uid, tstamp)
	if err != nil {
		fmt.Println(err)
		return
	}
}
