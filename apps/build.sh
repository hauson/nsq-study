#/bin/sh

# define var
image=chx
tag=v1
container=s1

# deocker
echo ${image}:${tag}
docker ps -a | grep ${image}:${tag} | awk '{print $1}' | xargs docker rm
docker rmi ${image}:${tag}
docker build -f dockerfile -t ${image}:${tag} .
docker images |grep ${image} | grep ${tag}
docker run --name ${container} -d ${image}:${tag}

# docker-compose up -d
# docker-compose down
# docker-compose ps #查看各个组件的运行详情
# docker-compose logs #查看组件日志
