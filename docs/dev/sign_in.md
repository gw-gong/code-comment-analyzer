# 登录逻辑

## 一、targets：
1. 用户输入邮箱和密码即可成功登录
2. 校验邮箱和密码

<br><br><br>

## 二、路由
```
/user/login/
```
request (json):
```json
{
    "email": "xxx@qq.com",
    "password": "123456"
}
```
response (json):<br>
success
```json
{
    "status": 0,
    "msg": "登录成功",
    "data": {
        "uid": 3,
        "email": "xxx@qq.com",
        "nickname": "xxxxxxxxxxxxx"
    }
}
```
failure
```json
{
    "status": 1,
    "msg": "邮箱或密码无效",
    "data": {}
}
```