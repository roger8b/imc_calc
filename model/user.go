package model

type User struct {
	Sex              string `json:"sex"`
	Age              int `json:"age"`
	PhysicalActivity string `json:"physical_activity"`
	Weight           string `json:"weight"`
	Height           string `json:"height"`
}
