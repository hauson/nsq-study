#/bin/sh

# define var
image=chx
tag=v1
container=shanshi1

# deocker
echo ${image}:${tag}
docker ps -a | grep ${image}:${tag} | awk '{print $1}' | xargs docker rm
docker rmi ${image}:${tag}
docker build -f dockerfile -t ${image}:${tag} .
docker images |grep ${image} | grep ${tag}
docker run --name ${container} -d ${image}:${tag}
