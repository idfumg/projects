package db

import (
	"html/template"
	"log"
	"myapp/data"
)

func GetPageData(pageGUID string, UserSession data.Session) (*data.Page, error) {
	page := data.Page{}
	rawContent := ""

	query :=
		`SELECT id, page_title, page_content, page_date, page_guid
	  		FROM pages WHERE page_guid=?`

	err := Database.QueryRow(query, pageGUID).Scan(&page.Id, &page.Title, &rawContent, &page.Date, &page.GUID)
	if err != nil {
		return nil, err
	}

	page.Content = template.HTML(rawContent)
	page.Session = UserSession

	comments, err := GetCommentsData(page.Id)
	if err != nil {
		log.Println(err)
	}
	page.Comments = comments

	return &page, nil
}

func GetPagesData() ([]*data.Page, error) {
	query :=
		`SELECT id, page_title, page_content, page_date, page_guid
			FROM pages 
				ORDER BY ? DESC`

	rows, err := Database.Query(query, "page_date")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pages []*data.Page

	for rows.Next() {
		page := data.Page{}
		rawContent := ""
		rows.Scan(&page.Id, &page.Title, &rawContent, &page.Date, &page.GUID)
		page.Content = template.HTML(rawContent)

		comments, err := GetCommentsData(page.Id)
		if err != nil {
			log.Println(err)
		}
		page.Comments = comments

		pages = append(pages, &page)
	}

	return pages, nil
}
