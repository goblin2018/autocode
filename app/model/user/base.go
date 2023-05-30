package user

type User struct {
	Id       string `json:"id" bson:"_id"`
	UpdateAt int64  `json:"updateAt" bson:"updateAt"`
	CreateAt int64  `json:"createAt" bson:"createAt"`
	Phone    string `bson:"phone" json:"phone"`
	Name     string `bson:"name" json:"name"`
	Age      int64  `bson:"age" json:"age"`
}
