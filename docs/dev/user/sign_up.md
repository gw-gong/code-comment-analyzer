# 注册逻辑

## 一、targets：
1. 用户输入对应数据进行注册
2. 注册成功后，需要登录才可以获取token
3. 需要给用户设置默认nickname

<br><br><br>

## 二、路由
```
/user/sign_up/
```
request (json):
```json
{
    "email": "ggw@qq.com",
    "password": "123456",
    "password_again": "123456"
}

```
response (json):<br>
success
```json
{
    "status": 0,
    "msg": "注册成功",
    "data": {
        "uid": 9,
        "email": "ggwxxxxx@qq.com",
        "nickname": "Anonymous"
    }
}

```
failure
```json
{
    "status": 1,
    "msg": "邮箱已被注册",
    "data": {}
}
```