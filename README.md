#slimim

#### slimim is a im server written by golang,owned by dylenfu
```
		  ____  _     ___ __  __   ___ __  __ 
		 / ___|| |   |_ _|  \/  | |_ _|  \/  |
		 \___ \| |    | || |\/| |  | || |\/| |
		  ___) | |___ | || |  | |  | || |  | |
		 |____/|_____|___|_|  |_| |___|_|  |_|
```                                      
                                      
## Features
 * Light weight
 * High performance
 * Pure Golang
 * Supports single push, multiple push and broadcasting
 * Supports one key to multiple subscribers (Configurable maximum subscribers count)
 * Supports heartbeats (Application heartbeats, TCP, KeepAlive, HTTP long pulling)
 * Supports authentication (Unauthenticated user can't subscribe)
 * Supports multiple protocols (WebSocket，TCP，HTTP）
 * Scalable architecture (Unlimited dynamic job and logic modules)
 * Asynchronous push notification based on Kafka
 
#### 项目结构
* gateway 网关,定位在线用户地址
* comet   接入层,客户端接入
* logic   逻辑层
* proto   数据结构
* rpc     http rest服务

#### 服务治理
* 使用kite微服务框架
* 服务之间使用grpc通讯
* 使用docker部署
* 使用kafka消息队列
