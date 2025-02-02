# 单文件分析

## 一、targets：
1. 分析单个文件的文本内容

<br><br><br>

## 二、路由
```
/public/analyze_file/
```
request (json):
```json
{
    "language":    "Python", 
    "fileContent": "# 文件页面 get\n@require_GET\ndef file(request):\n    indexPage = os.path.join('static', 'file.html')\n    return FileResponse(open(indexPage, 'rb'))\n\n# 这块返回的是处理的逻辑\n# ----------------------------------------------------------------\n# ----------------------------------------------------------------\n# 处理单个文件 post\n\"\"\"\n    上传的内容：{\n        \"fileType\":    C C++ Python Java Golang HTML CSS JavaScript\n        \"fileContent\": 文件内容文本\n    }\n\n    返回内容：{\n        \"status\": \n        \"analyzeData\":\n    }\n\"\"\"\n@csrf_exempt\n@require_POST\ndef analyze_file(request):\n    try:\n        # 获取请求的JSON数据\n        body_unicode = request.body.decode('utf-8')\n        data = json.loads(body_unicode)\n        file_type = data.get('fileType')\n        file_content = data.get('fileContent')\n\n        # 检查请求数据的有效性\n        if not file_type or not file_content:\n            return JsonResponse({'status': 'error', 'message': 'Invalid input'}, status=400)\n\n    except (UnicodeDecodeError, json.JSONDecodeError) as e:\n        return JsonResponse({'status': 'error', 'message': 'Invalid input format'}, status=400)\n\n    # 根据文件类型和内容进行相应的分析\n    analyze_data = {\n        'dsad': \"dasdad\"\n    }\n    # 返回处理后的数据\n    return JsonResponse({\n        'status': 'success',\n        'analyzeData': analyze_data\n    })\n"
}
```
response (json):<br>
success
```json
{
    "status": 0,
    "msg": "已分析",
    "language": "Python",
    "data": {
        "analyzed_table": [
            {
                "key": "代码注释行数",
                "value": 20,
                "badgeText": "图",
                "badgeLevel": "info"
            },
            {
                "key": "代码行数",
                "value": 48,
                "badgeText": "图",
                "badgeLevel": "info"
            },
            {
                "key": "注释覆盖比",
                "value": "41.67%",
                "badgeText": "重点",
                "badgeLevel": "success"
            },
            {
                "key": "单行注释个数",
                "value": 9,
                "badgeText": "图",
                "badgeLevel": "info"
            },
            {
                "key": "单行注释平均长度",
                "value": "23.56个字符"
            },
            {
                "key": "多行注释个数",
                "value": 1,
                "badgeText": "图",
                "badgeLevel": "info"
            },
            {
                "key": "多行注释平均长度",
                "value": "177.00个字符"
            },
            {
                "key": "词数<2 的注释",
                "value": 2,
                "badgeText": "图",
                "badgeLevel": "info"
            },
            {
                "key": "2<=词数<=30 的注释",
                "value": 8,
                "badgeText": "图",
                "badgeLevel": "info"
            },
            {
                "key": "词数>30 的注释",
                "value": 0,
                "badgeText": "图",
                "badgeLevel": "info"
            },
            {
                "key": "注释的最大词数",
                "value": 20
            }
        ],
        "commentLines": 20,
        "codeLines": 48,
        "commentDensity": "41.67%",
        "lenWordsBefore": 10,
        "wordsBefore": [
            {
                "key": "的",
                "value": 7
            },
            {
                "key": "内容",
                "value": 4
            },
            {
                "key": "文件",
                "value": 3
            },
            {
                "key": "返回",
                "value": 3
            },
            {
                "key": "处理",
                "value": 3
            },
            {
                "key": "数据",
                "value": 3
            },
            {
                "key": "请求",
                "value": 2
            },
            {
                "key": "页面",
                "value": 1
            },
            {
                "key": "get",
                "value": 1
            },
            {
                "key": "这块",
                "value": 1
            }
        ],
        "lenWordsAfter": 10,
        "wordsAfter": [
            {
                "key": "内容",
                "value": 4
            },
            {
                "key": "文件",
                "value": 3
            },
            {
                "key": "返回",
                "value": 3
            },
            {
                "key": "数据",
                "value": 3
            },
            {
                "key": "请求",
                "value": 2
            },
            {
                "key": "页面",
                "value": 1
            },
            {
                "key": "get",
                "value": 1
            },
            {
                "key": "这块",
                "value": 1
            },
            {
                "key": "逻辑",
                "value": 1
            },
            {
                "key": "单个",
                "value": 1
            }
        ],
        "singleLineCommentsLength": 9,
        "singleLineComments": [
            " 文件页面 get",
            " 这块返回的是处理的逻辑",
            " ----------------------------------------------------------------",
            " ----------------------------------------------------------------",
            " 处理单个文件 post",
            " 获取请求的JSON数据",
            " 检查请求数据的有效性",
            " 根据文件类型和内容进行相应的分析",
            " 返回处理后的数据"
        ],
        "multiLineCommentsLength": 1,
        "multiLineComments": [
            "<br>&nbsp;&nbsp;&nbsp;&nbsp;上传的内容：{<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;\"fileType\":&nbsp;&nbsp;&nbsp;&nbsp;C&nbsp;C++&nbsp;Python&nbsp;Java&nbsp;Golang&nbsp;HTML&nbsp;CSS&nbsp;JavaScript<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;\"fileContent\":&nbsp;文件内容文本<br>&nbsp;&nbsp;&nbsp;&nbsp;}<br><br>&nbsp;&nbsp;&nbsp;&nbsp;返回内容：{<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;\"status\":&nbsp;<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;\"analyzeData\":<br>&nbsp;&nbsp;&nbsp;&nbsp;}<br>"
        ]
    }
}
```
failure
```json

```