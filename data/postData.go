package data

import (
	"time"
)

type InPost struct {
	UserId          string    `json:"userId" bson:"userId"`
	Id              string    `json:"id" bson:"_id"`
	Caption         string    `json:"caption" bson:"caption"`
	ImgUrl          string    `json:"imgUrl" bson:"imgUrl"`
	PostedTimestamp time.Time `json:"-" bson:"postedTimestamp"`
}

type OutPost struct {
	UserId          string    `json:"userId" bson:"userId"`
	Id              string    `json:"id" bson:"_id"`
	Caption         string    `json:"caption" bson:"caption"`
	ImgUrl          string    `json:"imgUrl" bson:"imgUrl"`
	PostedTimestamp time.Time `json:"postedTimeStamp" bson:"postedTimestamp"`
}
