#1. 开启基本的nsqlookupd
docker run -d --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq /nsqlookupd

# 开启一个结点
docker run -d --name nsqd -p 4150:4150 -p 4151:4151  nsqio/nsq /nsqd  --broadcast-address=172.17.0.3  --lookupd-tcp-address=172.17.0.3:4160

