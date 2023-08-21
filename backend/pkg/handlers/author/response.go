package author

import (
	"github.com/globalsign/mgo/bson"
)

type Author struct {
	ID               bson.ObjectId `json:"id" bson:"_id"`
	Email            string        `json:"email" bson:"email"`
	Fullname         string        `json:"fullname" bson:"fullname"`
	Bio              string        `json:"bio" bson:"bio"`
	Location         string        `json:"location" bson:"location"`
	Facebook         string        `json:"facebook" bson:"facebook"`
	Instagram        string        `json:"instagram" bson:"instagram"`
	Twitter          string        `json:"twitter" bson:"twitter"`
	ShortDescription string        `json:"shortdescription" bson:"shortdescription"`
	Features         string        `json:"features" bson:"features"`
	CreatedOn        string        `json:"createdon" bson:"createdon"`
	Version          int           `json:"version" bson:"version"`
}
