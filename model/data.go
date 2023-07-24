package model

import "database/sql"

type Data struct {
	ID    int
	Name  string
	Field sql.NullString
	//
	// input your data
	//
}

type DataResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Field string `json:"field"`
	//
	// input your data
	//
}
