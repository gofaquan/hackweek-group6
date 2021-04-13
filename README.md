# hackweek-group6

# API 接口文档

## 1.1 API V1 接口说明

- 接口基准地址：`121.196.103.173:8080`
- 服务端已开启 CORS 跨域支持
- API V1 认证统一使用 Token 认证
- 需要授权的 API ，必须在请求头中使用 `Authorization` 字段提供 `token` 令牌
- 使用 HTTP Status Code 标识状态
- 数据返回格式统一使用 JSON

### 1.1.1 支持的请求方法

- POST（CREATE）：在服务器新建一个资源。
- PUT（UPDATE）：在服务器更新资源（客户端提供改变后的完整资源）。

### 1.1.2 通用返回状态说明

| *状态码* | *含义*                 | *说明*                                              |
| -------- | ---------------------- | --------------------------------------------------- |
| 200      | Success                     | 请求成功                                            |
| 201      | CREATED                | 创建成功                                            |
| 204      | DELETED                | 删除成功                                            |
| 400      | BAD REQUEST            | 请求的地址不存在或者包含不支持的参数                |
| 401      | UNAUTHORIZED           | 未授权                                              |
| 403      | FORBIDDEN              | 被禁止访问                                          |
| 404      | NOT FOUND              | 请求的资源不存在                                    |
| 422      | Unprocesable entity    | [POST/PUT] 当创建一个对象时，发生一个验证错误 |
| 500      | INTERNAL SERVER ERROR  | 内部错误                                            |
| 1001     | ERROR_USERNAME_USED    | 用户名已存在                                        |
| 1002     | ERROR_PASSWORD_WRONG   | 密码错误                                            |
| 1003     | ERROR_USER_NOT_EXIST   | 用户不存在                                          |
| 1004     | ERROR_TOKEN_EXIST      | TOKEN不存在                                         |
| 1005     | ERROR_TOKEN_RUNTIME    | TOKEN已过期                                         |
| 1006     | ERROR_TOKEN_WRONG      | TOKEN错误                                           |
| 1007     | ERROR_TOKEN_TYPE_WRONG | TOKEN格式错误                                       |
| 1008     | ERROR_USER_NO_RIGHT    | 用户无权限                                          |
| 1009     | ERROR_PASSWORD_EDIT    | 密码修改成功             
| 1010     |ERROR_PASSWORD_NOTSAME  |输入的密码不一致
|1011 |ERROR_USER_FAIL  |注册失败
1012|ERROR_BOTH_WRONG|帐号或密码错误|
|1013|ERROR_PASSWORD_EDFL|密码修改失败|

## 1.2. 登录

### 1.2.1. 登录验证接口

- 请求路径：/login
- 请求方法：POST
- 请求参数

| 参数名   | 参数说明 | 备注       |
| -------- | -------- | ---------- |
| username | 用户名   | 5~18个字符 |
| password | 密码     | 8~18个字符 |

- 响应参数

| 参数名 | 参数说明 | 备注            |
| ------ | -------- | --------------- |
| token  | 令牌     | 基于 jwt 的令牌 |
| msg    | 信息     |                 |
| status | 状态码   |                 |
| data   | 用户信息 |                 |

- 响应数据

```json
{
    "Data":{
    "username": "1234567890",
    "password": "$2a$10$GilaAfEvj.RZGh7cVh7oAe8OzDZeiQiGqlzgz4XgjE7Trkb8JYFEG",
    "msg": "OK"
    },
    "status": 200,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InF3ZSIsImV4cCI6MTYwNzc4MTY2NiwiaXNzIjoiaGFja3dlZWsyIn0.5M-FcjnCOCqk6QCbXyQZ4u12ifWppHxGbA1XTSNdgAg"
}
```

### 1.2.2. 注册用户

- 请求路径：/user/registratioon
- 请求方法：POST
- 请求参数
- 响应参数（JSON传,前端没有相应填入框，目前只填两个参数）

| 参数名   | 参数说明    | 备注 |
| -------- | ----------- | ---- |
| username | 用户名      |      |
| password | 用户角色 ID |      |

- 响应数据

```json
{
    "msg": "OK",
    "status": 200
}
```

### 1.3.3 修改密码

- 请求路径：/user/transformation
- 请求方法：PUT
- 请求参数

| 参数名         | 参数说明   | 备注           |
| -------------- | ---------- | -------------- |
| user           | 用户信息   | 包含用户名密码 |
| new_pwd        | 新密码     | 8~16个字符     |
| confirm_new_pw | 确认新密码 | 8~16个字符     |
| password       | 旧密码     | 8~16个字符     |
|                |            |                |

- 响应参数
- 响应数据

```json
{
    	"status" : 200,
		"msg" :  "密码修改成功"
}
```