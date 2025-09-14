package questions

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Content struct {
	Text *string `json:"text,omitempty"`
	Audio *string `json:"audio,omitempty"`
	Image *string `json:"image,omitempty"`
}

type Question interface {
	Statement() *Content
	InsertId()
}

type Option struct {
	Id *primitive.ObjectID
	IsCorrect *bool
	Content *Content
	Reason *Content
}

func (o *Option) InsertId() {
	id := primitive.NewObjectID()
	o.Id = &id
}