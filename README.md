## rocket
go相关工具和学习示例

##1. 
- XXX
- XXX
- XXX

## 2. go命令规范
- XXX
- XXX
- XXX

## 3. Beego web框架
- 官网地址：
- XXX
- XXX

## 4. Short link模块
- 介绍：<br/>
URL缩短服务在base62中编码URL并将它们存储在redis中<br/>
特点：缩短、信息网址缩短和重定向

- **API信息**<br/>

  > URL地址：<br/>
  > POST /api/shorten
  
  ***参数：***
  
  | Name | Type | Description |
  | ---- | ---- | ----------- |
  | `url`  | `string` | `Required. e.g. http://www.baidu.com` |
  | `expiration_in_minutes` | `int` | `Required. e.g. value 0 represents permanent` |
  
  ***响应：***
  
  <pre>
  {
      "shortlink": "A"
  }
  </pre>
  
  > 获取短链接URL<br/>
  > `GET /api/info?shortlink=shortlink`
  
  ***参数：***
  
  | Name | Type | Description |
  | ---- | ---- | ----------- |
  | `shortlink`  | `string` | `Required. id of shortened. e.g. A` |
  
  ***响应：***
  
  <pre>
  {
      "url": "http://www.baidu.com",
      "created_at": "2020-01-10 10:01:08.128660011 +0800",
      "expiration_in_minutes": 60
  }
  </pre>
  
    ***重定向：***
  
    > `GET /{shortlink}`
  
    Expand the short link and return a **temporary redirect** HTTP status 

## .注意事项
- XXX
- XXX
- XXX

