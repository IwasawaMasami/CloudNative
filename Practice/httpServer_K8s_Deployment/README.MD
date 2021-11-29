**要求：**
编写 Kubernetes 部署脚本将 httpserver 部署到 kubernetes 集群，以下是你可以思考的维度

优雅启动 ==> ok
优雅终止 ==> ok
资源需求和 QoS 保证 ==> ok
探活 ==> ok
日常运维需求，日志等级
配置和代码分离

将httpServer应用优雅的运行在Kubernetes 之上，我们还应该考虑如何将服务发布给对外的调用方。
来尝试用Service, Ingress将你的服务发布给集群外部的调用方吧
在第一部分的基础.上提供更加完备的部署spec，包括(不限于) 
●Service
●Ingress
可以考虑的细节
●如何确保整个应用的高可用
●如何通过证书保证httpServer的通讯安全
