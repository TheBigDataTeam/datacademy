package author

import "github.com/globalsign/mgo/bson"

type GetRequest struct {
	ID bson.ObjectId `json:"id" bson:"_id" validate:"required"`
}
type CreateRequest struct {
	Email            string `json:"email" bson:"email" validate:"required"`
	Fullname         string `json:"fullname" bson:"fullname" validate:"required"`
	Bio              string `json:"bio" bson:"bio" validate:"required"`
	Location         string `json:"location" bson:"location" validate:"required"`
	Facebook         string `json:"facebook" bson:"facebook" validate:"required"`
	Instagram        string `json:"instagram" bson:"instagram" validate:"required"`
	Twitter          string `json:"twitter" bson:"twitter" validate:"required"`
	ShortDescription string `json:"shortdescription" bson:"shortdescription" validate:"required"`
	Features         string `json:"features" bson:"features" validate:"required"`
}
type UpdateRequest struct {
	Email            string `json:"email" bson:"email" validate:"required"`
	Fullname         string `json:"fullname" bson:"fullname" validate:"required"`
	Bio              string `json:"bio" bson:"bio" validate:"required"`
	Location         string `json:"location" bson:"location" validate:"required"`
	Facebook         string `json:"facebook" bson:"facebook" validate:"required"`
	Instagram        string `json:"instagram" bson:"instagram" validate:"required"`
	Twitter          string `json:"twitter" bson:"twitter" validate:"required"`
	ShortDescription string `json:"shortdescription" bson:"shortdescription" validate:"required"`
	Features         string `json:"features" bson:"features" validate:"required"`
}
type DeleteRequest struct {
	ID bson.ObjectId `json:"id" bson:"_id" validate:"required"`
}
