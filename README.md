### 特性
1. 分布式，可拓扑的架构
2. 支持单个，房间推送
3. 心跳支持（gorilla/websocket内置）
4. 基于redis 做消息推送
5. 轻量级
6.

### 部署
拉取
```
go get -u github.com/Terry-Ye/im
mv $GOPATH/src/github.com/Terry-Ye/im $GOPATH/src/im
go get ./...



```

golang.org 包拉不下来的情况，例
```
package golang.org/x/net/ipv4: unrecognized import path "golang.org/x/net/ipv4" (https fetch: Get https://golang.org/x/net/ipv4?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)
```

从github 拉下来，再移动位置
```

git clone https://github.com/golang/net.git
mkdir -p golang.org/x/

mv net $GOPATH/src/golang.org/x/
```

### demo


### 继续完善点
1. 在线列表
2. 支持wss