### 仓库相关
* 乙方仓库地址：https://dev-gitlab.wanxingrowth.com/
* 我们仓库地址：http://git.zk020.cn/youmi-wx-apps/

### 开发参考：
* 定义rpc接口 **common-protos项目 foo分支**
* 注意包名 protos.项目名

```
message FooRequest{
  string name = 1;
  int64 age = 2;
  int64 address = 3;
}

message FooReply{
    int64 id = 1;
    int64 name = 2;
    int64 age = 3;
    int64 address = 4;
}

service FooService {
  rpc getFooList (FooRequest) returns (FooReply) {
  }
  rpc saveFoo (FooRequest) returns (FooReply) {
  }
  rpc delFoo (FooRequest) returns (FooReply) {
  }
  rpc getFoo (FooRequest) returns (FooReply) {
  }
}
```

### 通用语言代码生成
```
#!/bin/bash

ORIGINDIR=$(pwd)

docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc -I . --go_out=plugins=grpc:. *.proto

cd ${ORIGINDIR}
```


### 接口的实现 接口实现（rpc服务端代码） **user项目**

* git subtree add --prefix=protos/foo http://git.zk020.cn/youmi-wx-apps/common-protos.git foo
* pkg/rpc/foo/FooServiceImpl.go

```
package foo

import (
	"context"
	"user/pkg/utils/log"
	protos_foo "user/protos/foo"
)

type FooServiceImpl struct {
}

func (f FooServiceImpl) GetFooList(context context.Context, request *protos_foo.FooRequest) (*protos_foo.FooReply, error) {
	log.GetLogger().Infof(`FooServiceImpl: %s`, "GetFooList")
	// 数据库查询等业务逻辑

	reply := &protos_foo.FooReply{
		Id:      0,
		Name:    0,
		Age:     0,
		Address: 0,
	}

	return reply, nil
}

func (f FooServiceImpl) SaveFoo(context context.Context, request *protos_foo.FooRequest) (*protos_foo.FooReply, error) {
	log.GetLogger().Infof(`FooServiceImpl: %s`, "SaveFoo")
	// 数据库查询等业务逻辑
	reply := &protos_foo.FooReply{
		Id:      0,
		Name:    0,
		Age:     0,
		Address: 0,
	}

	return reply, nil
}

func (f FooServiceImpl) DelFoo(context context.Context, request *protos_foo.FooRequest) (*protos_foo.FooReply, error) {
	log.GetLogger().Infof(`FooServiceImpl: %s`, "DelFoo")
	// 数据库查询等业务逻辑
	reply := &protos_foo.FooReply{
		Id:      0,
		Name:    0,
		Age:     0,
		Address: 0,
	}

	return reply, nil
}

func (f FooServiceImpl) GetFoo(context context.Context, request *protos_foo.FooRequest) (*protos_foo.FooReply, error) {
	log.GetLogger().Infof(`FooServiceImpl: %s`, "GetFoo")
	// 数据库查询等业务逻辑
	reply := &protos_foo.FooReply{
		Id:      0,
		Name:    0,
		Age:     0,
		Address: 0,
	}

	return reply, nil
}

```

### 接口的注册
```
protos_foo.RegisterFooServiceServer(rpcService.GetRPCConnection(), &foo.FooServiceImpl{})
```

### 客户端的调用（controller）
```
  conn, err := grpc.Dial("127.0.0.1:8016", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Errorf(`conn err: %v`, err)
	}
	RpcFooService := protos_foo.NewFooServiceClient(conn)

	request := &protos_foo.FooRequest{
		Name:    "",
		Age:     "",
		Address: "",
	}

	context, _ := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	list, err := RpcFooService.GetFooList(context, request)
```

### 账号相关
* 测试环境：root@112.126.60.67
* MySQL：用户名：root 密码：Consul@zkcm@001
* Redis：密码：Consul@zkcm@

### 项目结构：
| 项目 | 端口 | 说明 |  |  |
| --- | --- | --- | --- | --- |
| banner | 8000 | 轮播图相关 |  |  |
| card | 8001 | 卡 |  |  |
| coupon | 8002 | 优惠券 |  |  |
| distribution | 8003 | 分销 |  |  |
| fuyou-payment-gateway | 8004 | 富有支付 |  |  |
| goods | 8005 | 商品 |  |  |
| merchant | 8006 | 商户 |  |  |
| order | 8011 | 订单 |  |  |
| oss-interface | 8012 | oss |  |  |
| payment | 8013 | 支付网关 | |  |
| rebate | 8014 | 返利 |  |  |
| sms | 8015 | sms |  |  |
| user | 8016 | 用户 |  |  |
| wechat-interface | 8017 | 微信调用 |  |  |
| mini-program-interface | 9000 | 小程序接口 |  |  |
| merchant-interface | 9001 | 后台管理接口 |  |  |

**重要：common-protos是公共的依赖模块：git subtree管理**

### 测试环境部署
参考builds/build.sh, publish.sh部署脚本 目前先支持supervisor, 后续k8s

### 目前项目的缺陷
* **服务端panic没处理, 级别严重: 直接导致服务不可用**
* rpc service端的封装
* rpc client端的封装
* **服务网关: 服务注册和服务发现, 用consul改造**

### 接口文档地址
* 项目swagger建议更换
* 文档地址：http://112.126.60.67:4999/web/#/  账号密码：showdoc/123456

