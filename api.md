---
title: 点赞排行网站（Gin）
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"

---


Base URLs:

* <a href="http://127.0.0.1:8080">开发环境: http://127.0.0.1:8080</a>

# 点赞排行网站（Gin）

## PUT 登陆接口

PUT /user/login

> Body 请求参数

```yaml
username: user123
password: "12345"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» username|body|string| 是 |用户账户|
|» password|body|string| 是 |用户密码|

> 返回示例

> 成功

```json
{
  "code": 200,
  "msg": "登录成功",
  "data": "user123",
  "count": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|


## POST 注册接口

POST /user/register

> Body 请求参数

```yaml
username: test
password: "123"
confirmPassword: "123"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» username|body|string| 是 |用户账号|
|» password|body|string| 是 |用户密码|
|» confirmPassword|body|string| 是 |密码确认|

> 返回示例

> 成功

```json
{
  "code": 200,
  "msg": "注册成功",
  "data": null,
  "count": 0
}
```

```json
{
  "code": 4004,
  "msg": "用户已存在",
  "data": null,
  "count": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

## POST 获取选手列表接口

POST /player/list

> Body 请求参数

```yaml
aid: "1"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» aid|body|string| 是 |活动id|

> 返回示例

> 成功

```json
{
  "code": 200,
  "msg": "success",
  "data": [
    {
      "id": 1,
      "aid": 1,
      "ref": "0001",
      "nickname": "下雨le",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/11.png",
      "score": 3,
      "add_time": 1686105642,
      "update_time": 0
    },
    {
      "id": 2,
      "aid": 1,
      "ref": "0002",
      "nickname": "栀夏",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/12.png",
      "score": 3,
      "add_time": 1686105642,
      "update_time": 0
    },
    {
      "id": 3,
      "aid": 1,
      "ref": "0003",
      "nickname": "闷油瓶",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/13.png",
      "score": 3,
      "add_time": 1686105642,
      "update_time": 0
    },
    {
      "id": 4,
      "aid": 1,
      "ref": "0004",
      "nickname": "喵喵兽",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/14.png",
      "score": 0,
      "add_time": 1686105642,
      "update_time": 0
    },
    {
      "id": 5,
      "aid": 1,
      "ref": "0005",
      "nickname": "弃殇",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/15.png",
      "score": 3,
      "add_time": 1686105642,
      "update_time": 0
    },
    {
      "id": 6,
      "aid": 1,
      "ref": "0006",
      "nickname": "冷巷",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/16.png",
      "score": 0,
      "add_time": 1686105642,
      "update_time": 0
    },
    {
      "id": 7,
      "aid": 1,
      "ref": "0007",
      "nickname": "栀子香气",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/17.png",
      "score": 7,
      "add_time": 1686105642,
      "update_time": 0
    },
    {
      "id": 8,
      "aid": 1,
      "ref": "0008",
      "nickname": "失与倦",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/18.png",
      "score": 0,
      "add_time": 1686105642,
      "update_time": 0
    }
  ],
  "count": 8
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|


## POST 投票接口

POST /vote/add

> Body 请求参数

```yaml
userid: "1"
playerid: "1"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» userid|body|string| 是 |用户id|
|» playerid|body|string| 是 |选手id|

> 返回示例

> 成功

```json
{
  "code": 200,
  "msg": "投票成功",
  "data": 14,
  "count": 1
}
```

```json
{
  "code": 4001,
  "msg": "您已经投过票了",
  "data": null,
  "count": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

## POST 排行榜接口

POST /ranking

> Body 请求参数

```yaml
aid: "1"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» aid|body|string| 是 |活动id|

> 返回示例

> 成功

```json
{
  "code": 200,
  "msg": "success",
  "data": [
    {
      "id": 7,
      "aid": 1,
      "ref": "0007",
      "nickname": "栀子香气",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/17.png",
      "score": 7,
      "add_time": 1686105642,
      "update_time": 0
    },
    {
      "id": 5,
      "aid": 1,
      "ref": "0005",
      "nickname": "弃殇",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/15.png",
      "score": 4,
      "add_time": 1686105642,
      "update_time": 0
    },
    {
      "id": 3,
      "aid": 1,
      "ref": "0003",
      "nickname": "闷油瓶",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/13.png",
      "score": 3,
      "add_time": 1686105642,
      "update_time": 0
    },
    {
      "id": 2,
      "aid": 1,
      "ref": "0002",
      "nickname": "栀夏",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/12.png",
      "score": 3,
      "add_time": 1686105642,
      "update_time": 0
    },
    {
      "id": 1,
      "aid": 1,
      "ref": "0001",
      "nickname": "下雨le",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/11.png",
      "score": 3,
      "add_time": 1686105642,
      "update_time": 0
    },
    {
      "id": 8,
      "aid": 1,
      "ref": "0008",
      "nickname": "失与倦",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/18.png",
      "score": 0,
      "add_time": 1686105642,
      "update_time": 0
    },
    {
      "id": 6,
      "aid": 1,
      "ref": "0006",
      "nickname": "冷巷",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/16.png",
      "score": 0,
      "add_time": 1686105642,
      "update_time": 0
    },
    {
      "id": 4,
      "aid": 1,
      "ref": "0004",
      "nickname": "喵喵兽",
      "declaration": "很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！",
      "avatar": "/images/14.png",
      "score": 1,
      "add_time": 1686105642,
      "update_time": 0
    }
  ],
  "count": 1
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|


