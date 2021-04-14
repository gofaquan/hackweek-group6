package model

type Temperalor struct {
	Url         string
	Temperature int
}

type PostPicture struct {
	Username string `json:"username"`
	Url      string `json:"url"`
	Time     string
}
