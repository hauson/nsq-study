# nsq-study

# nsq 官网
https://nsq.io/overview/design.html

## 学习的模块
 生产者
 消费者
 持久化
 admin
 服务发现 -  这两个是重点
 分布式   -  这两个是重点

## 线索型的数据, 方法, 主要是外部的

# tcpserver  -- 将这块的内容拆出来，学习下
# 形成一个容易测试的闭环

# 0. lookupd 用用于后期的服务发现, 整个项目尽量使用 channel来解决并发的问题
# 1. 建立一个server， client 两个进程, 可以使用docker-file 来启动, 以及使用docker-compose 来启动， 两个服务的启动有先后顺序， 先启动 server, 然后启动client
# 2. server 建立tcp的监听, client 建立tcp的发送
# 3. 维持长连接的存活, 在服务端发心跳，还是在client端发送心跳包
# 4. 建立消息的封装和解封装的协议层 protocol
# 5. server端：建立消息驱动机制, msgtype -> msghandler
# 6. client端：建立消息的驱动机制， mystype ->msghandler


# 服务发现， admin 本质上都是消息驱动机制
# 关键点， 处理掉并发, 通过加锁和其他方式

#1. 使用docker-compose 启动本地服务
#2. 编译两个bin: server, client