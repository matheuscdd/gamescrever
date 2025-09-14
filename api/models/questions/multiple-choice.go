package questions

import "go.mongodb.org/mongo-driver/bson/primitive"

type MultipleChoice struct {
	Id *primitive.ObjectID
	Content   *Content
	Options   *[]Option
	Reason    *Content
	LimitTime *int
	Points    *int
	Tag       *string
}

func (q *MultipleChoice) Statement() *Content {
	return q.Content
}

func (q *MultipleChoice) InsertId() {
	id := primitive.NewObjectID()
	q.Id = &id
}