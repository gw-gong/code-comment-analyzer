# 删除用户一条上传记录

## 一、targets：
1. 根据操作ID删除用户一条上传记录

<br><br><br>

## 二、路由
```
/user/delete_operating_record/
```
request (json):
```json
{
    "id": 1
}
```
response (json):<br>

success
```json
{
    "status": 0,
    "msg": "删除操作记录成功",
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