package types

type User struct {
	Id        string `json:"id"`
	OpenId    string `json:"openId"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	AddressId string `json:"addressId"`
	Birthday  int64  `json:"birthday"`
	State     int64  `json:"state"`
	ExpireAt  int64  `json:"expireAt"`
}
