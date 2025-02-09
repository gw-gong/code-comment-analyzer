# 获取说明文档

## 一、targets：
1. 获取说明文档的内容

<br><br><br>

## 二、路由
```
/public/get_readme/
```
request
```
null
```
response (json):<br>
success
```json
{
    "status": 0,
    "msg": "success",
    "data": "文件内容"
}
```
failure
```json
{
  "status": 0,
  "msg": "获取说明文档失败",
  "data": {}
}
```
