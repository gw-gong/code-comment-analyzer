# 获取用户操作记录

## 一、targets：
1. 获取用户操作记录

<br><br><br>

## 二、路由
```
/user/get_user_operating_records/
```
request (json): 分页参数
```json
/user/get_user_operating_records/?page=1&perPage=10
```
response (json):<br>

success
```json
{
    "status": 0,
    "msg": "获取上传操作记录成功",
    "data": {
        "data": [
            {
                "id": 1,
                "operation_type": "file",
                "created_at": "2024-05-06",
                "updated_at": "2024-05-06"
            },
            {
                "id": 2,
                "operation_type": "project",
                "created_at": "2024-05-06",
                "updated_at": "2024-05-06"
            }
        ],
        "total": 2,
        "page": 1
    }
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