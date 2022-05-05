package test_dto

type User struct {
	Name string
	Age  int
}

var UserDB = map[int]User{
	1: {"AlexHan", 29},
	2: {"AllenIverson", 47},
	3: {"VinceCarter", 45},
}

var TestQueryUserFunc func(id int) (User, error)
