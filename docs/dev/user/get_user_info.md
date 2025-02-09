# 获取用户信息

## 一、targets：
1. 获取用户信息

<br><br><br>

## 二、路由
```
/user/get_user_info/
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
        "tableInfo": [
            {
                "nickname": "xxxxxxxxxxxxx",
                "email": "xxx@qq.com",
                "date_joined": "2024-05-06"
            }
        ]
    }
}
```
failure
```json
{
    "status": 1,
    "msg": "未登录",
    "data": {}
}

```