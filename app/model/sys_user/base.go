package sys_user

type SystemUser struct {
	Id       string  `json:"id" bson:"_id"`
	UpdateAt int64   `json:"updateAt" bson:"updateAt"`
	CreateAt int64   `json:"createAt" bson:"createAt"`
	Phone    string  `bson:"phone" json:"phone"`
	Name     string  `bson:"name" json:"name"`
	Nickname string  `bson:"nickname,omitempty" json:"nickname,omitempty"`
	Duty     string  `bson:"duty,omitempty" json:"duty,omitempty"`
	Avatar   string  `bson:"avatar,omitempty" json:"avatar,omitempty"`
	Roles    []int64 `bson:"roles" json:"roles"`
}
