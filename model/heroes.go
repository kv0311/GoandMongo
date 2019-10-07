package model

type Hero struct {
	Name   string `json:"name" bson:"name"`
	Alias  string `json:"alias" bson:"alias"`
	Signed bool   `json:"signed" bson:"signed`
}
