package db

import (
	"log"
	"myapp/data"
)

func GetCommentsData(pageId int64) ([]*data.Comment, error) {
	query :=
		`SELECT id, comment_name, comment_email, comment_text FROM comments WHERE page_id=?`

	rows, err := Database.Query(query, pageId)
	if err != nil {
		return nil, err
	}

	var comments []*data.Comment
	for rows.Next() {
		var comment data.Comment
		rows.Scan(&comment.Id, &comment.Name, &comment.Email, &comment.Text)
		comments = append(comments, &comment)
	}
	return comments, nil
}

func UpdateCommentData(id, name, email, comments string) (int64, bool, error) {
	query :=
		`UPDATE comments 
	  		SET comment_name=?, comment_email=?, comment_text=? 
	    		WHERE id=?`

	res, err := Database.Exec(query, name, email, comments, id)
	if err != nil {
		return 0, false, err
	}

	cnt, err := res.RowsAffected()
	if err != nil {
		return 0, false, err
	}

	return cnt, cnt > 0, nil
}

func CreateNewCommentData(pageId, name, email, comments string) (int64, bool, error) {
	query :=
		`INSERT INTO comments 
			(page_id, comment_name, comment_email, comment_text, comment_date) 
			VALUES 
			(?, ?, ?, ?, CURRENT_TIMESTAMP)`

	res, err := Database.Exec(query, pageId, name, email, comments)
	if err != nil {
		log.Println(err)
		return -1, false, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, false, err
	}

	return id, true, nil
}
