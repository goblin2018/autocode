package types

type House struct {
	Id    string  `json:"id"`
	OrgId string  `json:"orgId"`
	Phone string  `json:"phone"`
	Name  string  `json:"name"`
	Addr  string  `json:"addr"`
	Users []*User `json:"users"`
}

type ListHouseReq struct {
	OrgId string `form:"orgId,omitempty"`
	Page  int64  `form:"page,omitempty"`
	Size  int64  `form:"size,omitempty"`
}

type ListHouseResp struct {
	Items []*House `json:"items"`
	Total int64    `json:"total"`
}

type DelHouseReq struct {
	Id string `json:"id"`
}
