package service

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/kashimashino/hackweek-group6/model"
)

type OssService struct{}

func (oss *OssService) UploadToken() *model.OssResponse {

	//构建一个阿里云客户端, 用于发起请求。
	//构建阿里云客户端时，需要设置AccessKey ID和AccessKey Secret。
	client, err := sts.NewClientWithAccessKey("cn-shenzhen", "LTAI5t86K6LPjU4x3wTTGMKj", "4b6upzrv91GnhBudYXnTQT49dAqlN9")

	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	//设置参数。关于参数含义和设置方法，请参见API参考。
	request.RoleArn = "acs:ram::1011107079365118:role/guofaquan"
	request.RoleSessionName = "guofaquan"

	//发起请求，并得到响应。
	response, err := client.AssumeRole(request)
	if err != nil {
		fmt.Printf("response is %#v\n", response.Credentials)
		return &model.OssResponse{Error: err.Error()}
	}
	fmt.Printf("response is %#v\n", response.Credentials)

	return &model.OssResponse{
		AccessKeyId:     response.Credentials.AccessKeyId,
		AccessKeySecret: response.Credentials.AccessKeySecret,
		SecurityToken:   response.Credentials.SecurityToken,
	}
}
