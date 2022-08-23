mockserver yaml client
=================
基于 https://www.mock-server.com/ 实现的简单 MockServer 服务构建工具。

## 使用方法
#### 1. 启动 mock-server
```bash
docker pull mockserver/mockserver
docker run --rm -ti -p 1080:1080 mockserver/mockserver
```

#### 2. 导入示例 mock-api
```bash
go run main.go examples
```

#### 3. 测试示例 api
```bash
curl -v -H 'Content-Type: application/json' 'http://127.0.0.1:1080/view/cart?cartid=055ca455-1df7-45bb-8535-4f83e7266092'
```

正常情况下接口应该返回如下数据：
```json
{
  "code" : 0,
  "data" : {
    "id" : "055CA455-1DF7-45BB-8535-4F83E7266092",
    "image" : "055CA455-1DF7-45BB-8535-4F83E7266092.png"
  },
  "message" : ""
}
```

#### 4. 添加 mock-api
上一步测试通过后就可以参考 examples 目录下的 yaml 文件配置开始编写需要的 mock-api。

通过 `go run main.go {your-path}` 命令可以将指定目录下的所有 yaml 文件中定义的 mock-api 导入到 mockserver。

> 也可以通过 go install 将 mockyaml 安装为可执行程序，然后直接执行 `mockyaml ./examples` 进行 mock-api 导入。

## 参考
- mockserver 文档： https://www.mock-server.com/#what-is-mockserver
- mockserver api mock 规则配置文档： https://www.mock-server.com/mock_server/creating_expectations.html#request_matchers

