# 修改用户信息

## 一、targets：

1. 修改用户信息

`<br><br>``<br>`

## 二、路由

```
/user/update_user_info/
```

request (json):

```json
{
    "nickname": "xxx"
}
```

form-data:
key: "profile_picture"
value: file

头像存储按照 `avatar.md`说明

response (json):`<br>`

success

```json
{
    "status": 0,
    "msg": "用户信息更新成功",
    "data": {}
}
```

failure

```json
{
    "status": 1,
    "msg": "error message",
    "data": {}
}

```
