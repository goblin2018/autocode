package house

type House struct {
	Id       string `json:"id" bson:"_id"`
	UpdateAt int64  `json:"updateAt" bson:"updateAt"`
	CreateAt int64  `json:"createAt" bson:"createAt"`
	OrgId    string `bson:"orgId" json:"orgId"`
	Phone    string `bson:"phone" json:"phone"`
	Name     string `bson:"name" json:"name"`
	Addr     string `bson:"addr" json:"addr"`
}

type ListHouseReq struct {
	OrgId string `json:"orgId"`
	Page  int64  `json:"page"`
	Size  int64  `json:"size"`
}
