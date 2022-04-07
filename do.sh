# 启动 nsqlookupd
 ./nsqlookupd

# 启动nsqd1
./nsqd --tcp-address=0.0.0.0:4150 --http-address=0.0.0.0:4151 --https-address=0.0.0.0:4152 --lookupd-tcp-address=127.0.0.1:4160


# 启动nsqd2
./nsqd --tcp-address=0.0.0.0:4180 --http-address=0.0.0.0:4181 --https-address=0.0.0.0:4182 --lookupd-tcp-address=127.0.0.1:4160
