# 上传项目并返回项目树形结构

## 一、targets：
1. 上传.zip文件，解压并返回项目树形结构

<br><br><br>

## 二、路由
```
/public/upload_and_get_tree/
```
request (form-data):<br>
key: file, <br>value: 上传的.zip文件

response (json):<br>
success
```json
{
    "status": 0,
    "message": "文件已解压",
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

```
