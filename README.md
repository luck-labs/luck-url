
:::info
💡  lucky-url：从不止于链接缩短，一个专业的营销推广工具
:::
|  | 后端 | 前端 |
| --- | --- | --- |
| **开发语言** | Go | React |
| **开发框架** | httprouter，RPC，sentinel-go | Ant Design、Umi |


---

## 架构图
## ![lucky-url-arch.png](https://cdn.nlark.com/yuque/0/2022/png/22901959/1651761642308-f6a8bf32-0917-4f84-bd2d-6a753e910016.png#clientId=uf7e3d32d-351c-4&crop=0&crop=0&crop=1&crop=1&from=paste&id=u7e491117&margin=%5Bobject%20Object%5D&name=lucky-url-arch.png&originHeight=514&originWidth=963&originalType=binary&ratio=1&rotation=0&showTitle=false&size=937159&status=done&style=none&taskId=uf01c0b8c-9163-45fa-a574-67ed27ccdbb&title=)
## 主要功能特性

- 支持短链域名、后缀长度、后缀字符集配置化
- lucky-url-go采用原生go rpc，底层存储基于redis，支持单机5W+ QPS
- 支持sentinel服务限流配置化
- id发射器采用snowflake算法，单前缀最多使用69年

---

## 模块介绍
### lucky-url-react 短链前端react服务
### luky-url-go 短链后端golang服务

---

## 技术方案
#### 随机数生成snowflake
随机数生成采用twitter snowflake技术方案，可以使用69年，如果希望短链后缀变短，可以调整ID的长度。通过压缩时间戳，工作机器，以及随机序列号。
时间戳：2 ^ 41 / 1000 / 3600 / 24 / 365 = 69.7306年
工作机器：2 ^ 10 = 1024台机器
随机序列号：2 ^ 12 = 4096 / ms ，相当于TPS 4096 * 1000 = 4,096,000 / s
![image.png](https://cdn.nlark.com/yuque/0/2022/png/22901959/1651756552991-0e3d8207-ad3b-404f-b9a8-c89e6e3d9f66.png#clientId=u2d721928-b4fa-4&crop=0&crop=0&crop=1&crop=1&from=paste&id=u8253f644&margin=%5Bobject%20Object%5D&name=image.png&originHeight=241&originWidth=789&originalType=binary&ratio=1&rotation=0&showTitle=false&size=68054&status=done&style=none&taskId=ucf9f1cb2-fb79-4fb6-85d1-4d4fb25820e&title=)
#### 短链后缀base58
短链后缀方案，通过随机数hash到对应的字符，整体采取base58的技术方案，去除如“? /” url不支持的等字符，如果短链需要定制字符，可以修改字符集。

---

## 使用示例

- 线上演示：[https://url.shetuankaoqin.com/#/](https://url.shetuankaoqin.com/#/)
- 产物示例：http://s.shetuankaoqin.com/LXvr9Q
- 截图：

---

## 快速开始
### 前端启动

- 安装依赖

```bash
$ yarn
```

- 启动服务

```bash
$ yarn start
```
### 后端启动
```json
$ sh build.sh
```
### 线上部署

- Nginx 配置
>     location /api/
>     {
>       proxy_pass [http://localhost:8801/;](http://localhost:8801/;)
>     }
>     location /
>     {
>       proxy_pass [http://localhost:8801/v1/jump/;](http://localhost:8801/v1/jump/;)
>     }

### 接口文档
#### 创建短链
**API** /v1/api/create
**Method** POST
**Request**

| Key | Name | Sample |
| --- | --- | --- |
| url | 长链 | http://www.baidu.com |

#### Response
```json
{
  "errno":0,
  "errmsg":"SUCCESS",
  "data":{
    "url":"x-url.cc/3kTMd"
  }
}
```
#### 短链跳转
**API** /v1/jump/:s
**Method** GET
**Request**

| Key | Name | Sample |
| --- | --- | --- |
| url | 短链 | x-url.cc/3kTMd |


