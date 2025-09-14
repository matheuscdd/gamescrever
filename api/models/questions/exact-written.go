package questions

import "go.mongodb.org/mongo-driver/bson/primitive"

type ExactWritten struct {
	Id *primitive.ObjectID
	Content *Content
	Reason *Content
	Answers *[]string
}

func (q *ExactWritten) Statement() *Content {
	return q.Content
}

func (q *ExactWritten) InsertId() {
	id := primitive.NewObjectID()
	q.Id = &id
}

func (q *ExactWritten) Validate(response string) bool {
	return false
}
