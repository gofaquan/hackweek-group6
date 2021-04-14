package service

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"mime"
	"os"
	"path/filepath"
)

// UploadTokenService 获得上传oss token的服务
type UploadTokenService struct {
	Filename string `form:"filename" json:"filename"`
}

type Response struct {
	Key   string `json:"key"`
	Put   string `json:"put"`
	Get   string `json:"get"`
	Error string `json:"error"`
}

// Post 创建token
func (service *UploadTokenService) Post() *Response {
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		return &Response{
			Error: err.Error(),
		}
	}

	// 获取存储空间。
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		return &Response{
			Error: err.Error(),
		}
	}

	// 获取扩展名
	ext := filepath.Ext(service.Filename)

	// 带可选参数的签名直传。
	options := []oss.Option{
		oss.ContentType(mime.TypeByExtension(ext)),
	}

	key := "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ext
	// 签名直传。
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		return &Response{
			Error: err.Error(),
		}
	}
	// 查看图片
	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		return &Response{
			Error: err.Error(),
		}
	}

	return &Response{
		Key: key,
		Put: signedPutURL,
		Get: signedGetURL,
	}
}
