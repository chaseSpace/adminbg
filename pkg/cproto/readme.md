# cproto包

简介
-   项目前后端采用HTTP Over JSON的方式进行数据交互
-	此目录是专门存放API协议的文件，每个API的相关信息都会注释到对应Req/Rsp的上方，比如API功能注释，路由信息等
-   前端同事可直接查看此目录下的`xxx_proto.go`文件来进行接口对接（**无需拉取到本地**）

标准Response结构如下
```text
type Response struct {
	Code      int         `json:"code"` // 自定义code，0表示OK
	Data      interface{} `json:"data"` // 数据
	Msg       string      `json:"msg"`  // 文案提示
}
```

请求示例
```text
POST http://localhost:8000/api/v1/GetChannelList
// Header
Content-Type: application/json
Authorization: Bearer TOKEN...

// JSON Body
{
  "ps": 5,
  "pn": 1
}
// JSON Response
{
  "code": 0,
  "data": {
    "list": [
      {
        "chan_id": 59768272,
        "chan_name": "渠道A-1",
        "parent_chan_name": "渠道A",
        "chan_level": 3,
        "invite_code": "",
        "room_id": "1001",
        "remark": "",
        "create_at": 1606977437,
        "status": 1
      }
    ],
    "total_count": 71
  },
  "msg": "操作成功",
}
```
**需要注意的是，data字段在没有数据时可能为null**

## @前端人员需知：
-   `xxx_proto.go`文件内定义的xxxRsp结构体将设置到标准Response.Data字段，并非直接呈现到前端
-   标准Response内的Code字段正常返回0，否则就是异常，前端可直接提示Msg字段内容

## @后端人员需知：
-	按照固定的xxxReq & xxxRsp 命名风格定义（嵌套struct除外）
-	遵循已有的注释风格
-	Req&Rsp成对存在，即使无字段
-	定义的Req-struct都是由gin的bind方法来操作，所以可以使用gin支持的tag
-	因为proto go文件会给前端人员阅读，所以尽量不要使用语法糖，比如 iota
