package types

type SystemUser struct {
	Id       string  `json:"id"`
	Phone    string  `json:"phone"`
	Name     string  `json:"name"`
	Nickname string  `json:"nickname,omitempty"`
	Duty     string  `json:"duty,omitempty"`
	Avatar   string  `json:"avatar,omitempty"`
	Roles    []int64 `json:"roles"`
}
