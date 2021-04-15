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

type OssResponse struct {
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	SecurityToken   string `json:"security_token"`
	Error           string `json:"error"`
}
