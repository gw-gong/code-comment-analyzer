# 获取用户项目上传操作记录

## 一、targets：
1. 获取用户项目上传操作记录

<br><br><br>

## 二、路由
```
/user/get_project_upload_record/
```
request (json):
```json
/user/get_project_upload_record/?operating_record_id=10
```
response (json):<br>

success
```json
{
    "status": 0,
    "msg": "获取项目上传记录成功",
    "data": {
      "label": "test",
      "children": [
        {
          "label": "c.c",
          "value": "D:\\_GGW2021\\Desktop\\毕业\\ccanalyzer\\code_comment_analyzer\\media\\projects\\user3\\45041369-6d8a-4672-a2a4-2f1ae7e90cf9\\extracted\\test\\c.c"
        },
        {
          "label": "cpp.cpp",
          "value": "D:\\_GGW2021\\Desktop\\毕业\\ccanalyzer\\code_comment_analyzer\\media\\projects\\user3\\45041369-6d8a-4672-a2a4-2f1ae7e90cf9\\extracted\\test\\cpp.cpp"
        },
        {
          "label": "css.css",
          "value": "D:\\_GGW2021\\Desktop\\毕业\\ccanalyzer\\code_comment_analyzer\\media\\projects\\user3\\45041369-6d8a-4672-a2a4-2f1ae7e90cf9\\extracted\\test\\css.css"
        },
        {
          "label": "golang.go",
          "value": "D:\\_GGW2021\\Desktop\\毕业\\ccanalyzer\\code_comment_analyzer\\media\\projects\\user3\\45041369-6d8a-4672-a2a4-2f1ae7e90cf9\\extracted\\test\\golang.go"
        },
        {
          "label": "html.html",
          "value": "D:\\_GGW2021\\Desktop\\毕业\\ccanalyzer\\code_comment_analyzer\\media\\projects\\user3\\45041369-6d8a-4672-a2a4-2f1ae7e90cf9\\extracted\\test\\html.html"
        },
        {
          "label": "javascript.js",
          "value": "D:\\_GGW2021\\Desktop\\毕业\\ccanalyzer\\code_comment_analyzer\\media\\projects\\user3\\45041369-6d8a-4672-a2a4-2f1ae7e90cf9\\extracted\\test\\javascript.js"
        },
        {
          "label": "Main.java",
          "value": "D:\\_GGW2021\\Desktop\\毕业\\ccanalyzer\\code_comment_analyzer\\media\\projects\\user3\\45041369-6d8a-4672-a2a4-2f1ae7e90cf9\\extracted\\test\\Main.java"
        },
        {
          "label": "main.txt",
          "value": "D:\\_GGW2021\\Desktop\\毕业\\ccanalyzer\\code_comment_analyzer\\media\\projects\\user3\\45041369-6d8a-4672-a2a4-2f1ae7e90cf9\\extracted\\test\\main.txt"
        },
        {
          "label": "python.py",
          "value": "D:\\_GGW2021\\Desktop\\毕业\\ccanalyzer\\code_comment_analyzer\\media\\projects\\user3\\45041369-6d8a-4672-a2a4-2f1ae7e90cf9\\extracted\\test\\python.py"
        }
      ]
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