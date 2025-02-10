# 修改用户密码

## 一、targets：
1. 修改用户密码

<br><br><br>

## 二、路由
```
/user/change_password/
```
request (json):
```json
{
    "old_password": "123456",
    "new_password": "1234567",
    "again_new_password": "1234567"
}
```
response (json):<br>

success
```json
{
    "status": 0,
    "msg": "密码更新成功",
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