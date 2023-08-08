package utils

import (
	"GIK_Web/database"
)

type post struct {
	PostName string
}

func GetItemNameByID(id uint) (name string, err error) {

	var postData post
	database.Database.Raw("SELECT * FROM `AP6_posts` WHERE `ID` = ?", id).Scan(&postData)

	name = postData.PostName
	err = nil

	return
}
