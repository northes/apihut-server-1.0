# APIHut Server 端

[APIHut](https://apihut.net/) 源码

网站上线时间：2021.06.01

## 文档
https://docs.apihut.net/


## 环境
Golang 1.6

Mysql 5.7.26

Redis 3.0.504


## Q&A

IP接口获取到的IP不是用户的真实IP（::1）？

如果是通过nginx反代的获取的就是nginx的反代地址，通过配置nginx配置文件可解决

```
proxy_set_header Host $host;
proxy_set_header X-Real-IP $remote_addr;
proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
```

反向代理的设置为
http://127.0.0.1:8080
不要写
http://localhost:8080
否则会出现 ::1 与 真实IP 地址间来回切换的情况