# 登出逻辑

## 一、targets：
1. 用户登出，清除redis中的token

<br><br><br>

## 二、路由
```
/user/logout/
```
request (json):
```json
none
```
response (json):<br>
success
```json
{
    "status": 0,
    "msg": "已成功退出登录",
    "data": {}
}
```
failure
```json
重复登出无所谓，返回成功即可
没有token，也返回成功就行
```