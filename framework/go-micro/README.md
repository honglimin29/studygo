### go-micro微服务框架学习

本篇主要对go-micro微服务开发框架的个人学习整理，不涉及其它工具和库，如需系统的学习micro，请访问官方地址[https://micro.mu](https://micro.mu)

## 什么是 micro ?

Micro 是一个着眼于分布式系统开发的微服务生态系统，它是一个框架也是一个工具集，有着自己的社区和强大的生态系统。

Micro由开源的库与工具组成，旨在辅助微服务开发。
- go-micro - 基于Go语言的可插拔RPC微服务开发框架；包含服务发现、RPC客户/服务端、广播/订阅机制等等。
- go-plugins - go-micro的插件有etcd、kubernetes、nats、rabbitmq、grpc等等。
- micro - 微服务工具集包含传统的入口点（entry point）；API 网关、CLI、Slack Bot、代理及Web UI。  

还有其它相关的库和服务可以参考 [github.com/micro](https://github.com/micro)



