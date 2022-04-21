package models

type Student struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
