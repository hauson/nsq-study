# nsq-study

# nsq 官网

https://nsq.io/overview/design.html

## 学习的模块

生产者 消费者 持久化 admin 服务发现 - 这两个是重点 分布式 - 这两个是重点

## 线索型的数据, 方法, 主要是外部的

# 1.建立起消费者的链接
# 2.server端：建立消息驱动机制, msgtype -> msghandler
# 3.建立起大根堆 和 小根堆

# ----- 稳住心态， 一点点的做

# 0. lookupd 用用于后期的服务发现, 整个项目尽量使用 channel来解决并发的问题

# 3. 维持长连接的存活, 在服务端发心跳，还是在client端发送心跳包

# 4. 建立消息的封装和解封装的协议层 protocol

# 6. client端：建立消息的驱动机制， mystype ->msghandler

# 服务发现， admin 本质上都是消息驱动机制

# 关键点， 处理掉并发, 通过加锁和其他方式


