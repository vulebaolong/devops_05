## Image

chứa code đã được build (ngôn ngữ của source)

```bash
docker build -t js_nextjs_img .


docker image ls

docker image remove
```
## Container

Container là gì?
https://docs.docker.com/get-started/docker-concepts/the-basics/what-is-a-container/

là process được cô lập.
https://docs.docker.com/get-started/docker-overview#containers

```bash
docker run --name js_nextjs_container -d -p 12345:3000 js_nextjs_img
# -d: detached chạy ngầm, không chiếm dụng terminal
# -p: publish

docker container list

docker ps
# p: process
# s: status
# liệt kê những container đang chạy

docker ps -A
# liệt kê tất cả container kể cả container đang off

docker container stop [CONTAINER_NAME]
docker container start [CONTAINER_NAME]
docker container remove [CONTAINER_NAME]
docker container inspect [CONTAINER_NAME]
docker container logs [CONTAINER_NAME]
docker container stats [CONTAINER_NAME]
```

## Network
## Volume
## Compose

<!-- Tối ưu tốc độ build image -->
<!-- Tối ưu dung lượng build image -->