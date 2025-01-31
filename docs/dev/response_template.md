# 接口返回格式（重要）
> https://aisuda.bce.baidu.com/amis/zh-CN/docs/types/api#%E6%8E%A5%E5%8F%A3%E8%BF%94%E5%9B%9E%E6%A0%BC%E5%BC%8F-%E9%87%8D%E8%A6%81-

所有配置在 amis 组件中的接口，都要符合下面的返回格式
```json
{
  "status": 0,
  "msg": "",
  "data": {
    ...其他字段
  }
}
```
+ status: 返回 0，表示当前接口正确返回，否则按错误请求处理； 
+ msg: 返回接口处理信息，主要用于表单提交或请求失败时的 toast 显示； 
+ data: 必须返回一个具有 key-value 结构的对象。 

status、msg 和 data 字段为接口返回的必要字段。
<br><br>

正确的格式
```json
{
  "status": 0,
  "msg": "",
  "data": {
    // 正确
    "text": "World!"
  }
}
```

错误的格式
直接返回字符串或者数组都是不推荐的
```json
{
  "status": 0,
  "msg": "",
  "data": "some string" // 错误，使用 key 包装
}
```
