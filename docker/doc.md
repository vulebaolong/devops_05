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

- giống mạng LAN
- để cho các container nói chuyện với nhau

- bridge: network mặc định, cấu nối giữa 2 container
- host:
    - dùng luôn mạng của máy tính cá nhân/ server host
    - dễ trùng port
- none: không sài mạng, cô lập
    - chơi 1 mình

```bash

# Tạo network
docker network ls
docker network create [NETWORK_NAME]
docker network remove [NETWORK_NAME]
docker network connect [NETWORK_NAME] [CONTAINER_NAME/ID]
docker network disconnect [NETWORK_NAME] [CONTAINER_NAME/ID]
```

## Volume

gắn volume có 2 loại: 

- volume: sử dụng tên volum 
- mount: sử dụng đường trên máy host

Mongodb - 27017: /data/db
MySQL - 3306: /var/lib/mysql
SQL server - 1433: /vat/opt/mssql
PostgreSQL - 5432: - 17 trở xuống: /var/lib/postgres/data - 18 trở lên: /var/lib/postgresql

```bash
docker volume ls
docker volume create [NETWORK_NAME]
docker volume remove [NETWORK_NAME]

docker volume create database_volumne
docker run --name database -e POSTGRES_PASSWORD=12345 -d -v database_volume:/var/lib/postgresql postgres:18

# tạo dữ liệu
psql -U postgres

# list database
\l

# list table
\dt

SELECT current_database();

CREATE TABLE users (
    name VARCHAR(255)
);
INSERT INTO users (name)  VALUES
('THẦY SANG'),
('THẦY SANG'),
('THẦY SANG');

SELECT * FROM users;

docker run --name u_bind_mount -d -it -v ./bind_mount_host:/bind_mount_container ubuntu:26.04

# backup
# lấy volume nén vào file
# https://docs.docker.com/engine/storage/volumes/#back-up-restore-or-migrate-data-volumes

docker run --rm --volumes-from dbstore -v $(pwd):/backup ubuntu tar cvf /backup/backup.tar /dbdata
# docker run --rm: container tạm, sẽ tự động xoá khi lệnh end

docker run --rm -v database_volume:/backup_volume -v $(pwd):/backup_mount ubuntu  tar cvf /backup_mount/backup.tar -C /backup_volume .

docker run -v database_volume:/backup_volume -v $(pwd):/backup_mount -d -it ubuntu
cd backup_volume
cd backup_mount
tar cvf /backup_mount/backup.tar -C /backup_volume .

# restore
# lấy file giải nén vào volume
docker volume create database_restore

docker run --rm -v database_restore:/restore_volume -v $(pwd):/restore_mount ubuntu  tar xvf /restore_mount/backup.tar -C /restore_volume

docker run --name database_new -e POSTGRES_PASSWORD=12345 -d -v database_restore:/var/lib/postgresql postgres:18
```

## Tối ưu image
### Tối ưu tốc độ build image
- tận dụng cache của docker
### Tối ưu dung lượng build image
- dùng stage
- tạo ra 1 môi trường máy mới
- và chỉ mang qua những file cần thiết để run

## Clear rác, dọn dẹp
```bash
docker image prune -f
docker builder prune -f
```

## Compose


