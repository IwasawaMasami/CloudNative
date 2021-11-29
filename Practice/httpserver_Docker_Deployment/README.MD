**运行命令：**
```
docker run --rm -d  -p 8080:8090/tcp tookizhang/myhttpserver:v1
```
**Dockerfile：**
```Dockerfile
FROM golang:alpine
 WORKDIR /go/src/app
 ADD . .
 RUN go mod init
 RUN go build -o myHttpServer main.go 
 EXPOSE 8090
 CMD ["./myHttpServer"]
```
**容器内查看结果**：
```
➜  ~ docker inspect 9a8273a8871d | grep -i pid
            "Pid": 2107789,
            "PidMode": "",
            "PidsLimit": null,
➜  ~ nsenter -t 2107789 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
356: eth0@if357: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:12:00:03 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.18.0.3/16 brd 172.18.255.255 scope global eth0
       valid_lft forever preferred_lft forever
➜  ~ 
```
**镜像测试访问：**
http://lazytoki.cn:8080