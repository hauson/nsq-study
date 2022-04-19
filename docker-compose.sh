-p , --project-name 执行项目名称
-f, --file 指定docker-compse 使用的文件
--verbose 输出更多调式信息

docker-compose up -d #启动nsq 所有组件
docker-compose down  #关闭nsq 所有组件
docker-compose exec 进入制定的容器
docker-compose kill 通过发送SIGKILL 信号来强制停止服务容器
docker-compose pause 暂停一个服务容器
docker-compose restart 重启项目中的服务
docker-compose rm 删除所有停止状态的服务容器

docker-compose scale 设置容器个数
docker-compose run 根据镜像启动容器
docker-compose start 启动容器
docker-compose stop 停止容器

docker-compose ps #查看各个组件的运行详情
docker-compose images 列出文件中包含的镜像
docker-compose logs #查看组件日志
docker-compose port 打印某个容器端口所映射的公共端口
docker-compose top

docker-compose 检查当前目录中docker-compose.yaml 文件格式是否正确
docker-compose -f aa.yaml 检查指定文件的格式

docker-compose build 构建所有组件的镜像

# 删除docker-compose目录下的所有镜像
docker-compose images | awk 'NR>=3 {print $4}' | xargs docker rmi

