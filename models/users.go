package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Avatar   string             `bson:"avatar" json:"avatar,omitempty"`
	Banner   string             `bson:"banner" json:"banner,omitempty"`
	Name     string             `bson:"name" json:"name,omitempty"`
	Nick     string             `bson:"nick" json:"nick,omitempty"`
	Bio      string             `bson:"bio" json:"bio,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password,omitempty"`
	Site     string             `bson:"site" json:"site,omitempty"`
	DateNasc time.Time          `bson:"dateTime" json:"date_time,omitempty"`
}
