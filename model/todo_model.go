package model

type Todo struct {
	Base  `bson:",inline"`
	Title string `json:"title" bson:"title"`
	Done  bool   `json:"done" bson:"done"`
}

func (dest *Todo) Copy(src Todo) {
	dest.Title = src.Title
	dest.Done = src.Done
}
