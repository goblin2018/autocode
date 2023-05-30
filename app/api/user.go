package types

type User struct {
	Id    string `json:"id"`
	Phone string `json:"phone"`
	Name  string `json:"name"`
	Age   int64  `json:"age"`
}

type TestB struct {
	LaiLe string `json:"name"`
	Age   int64  `json:"age"`
}
