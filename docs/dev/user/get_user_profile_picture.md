# 获取用户头像

## 一、targets：
1. 获取用户头像信息

<br><br><br>

## 二、路由
```
/user/get_user_profile_picture/
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
    "msg": "获取用户信息成功",
    "data": {
        "profile_picture": "https://gxblogs.com/static/imgs/touxiang.png",
        "text": "未设置"
    }
}

```
failure
```json
{
  "status": 0,
  "msg": "未登录",
  "data": {
    "profile_picture": null,
    "text": "?"
  }
}
```