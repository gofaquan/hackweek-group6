package model

type Temperalor struct {
	Username    string `json:"username"`
	Url         string `json:"url"`
	Temperature int    `json:"temperature"`
}

type GetPicture struct {
	Username string `json:"username"`
	Url      string `json:"url"`
}

type OssResponse struct {
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	SecurityToken   string `json:"security_token"`
	Error           string `json:"error"`
}
