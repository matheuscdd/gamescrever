package questions

import "go.mongodb.org/mongo-driver/bson/primitive"

type ExactWritten struct {
	Id *primitive.ObjectID
	Content *Content
	Reason  *Content
	Answer  *string
}

func (q *ExactWritten) Statement() *Content {
	return q.Content
}

func (q *ExactWritten) InsertId() {
	id := primitive.NewObjectID()
	q.Id = &id
}
