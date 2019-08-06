## 解决国内go get被墙的问题

介绍如何使用gopm解决国内go get被墙的问题

# golang之旅 -- gopm

### 什么是gopm

在nodejs中我们有npm，可以通过npm来下载安装一些依赖包。在go中也开发了类似的东西，那就是gopm。这玩意儿是七牛开发的。在这里说下，七牛公司大部分程序都是用go语言编写的，所以开发出这么一个方便的东西肯定也是合情合理的。

### gopm地址

git地址：[https://github.com/gpmgo/gopm](https://github.com/gpmgo/gopm)  
官方地址：[https://gopm.io/gopm](https://gopm.io/gopm)  
文档路径：[https://github.com/gpmgo/docs/tree/master/zh-CN](https://github.com/gpmgo/docs/tree/master/zh-CN)  

### gopm安装

```bash
go get -u github.com/gpmgo/gopm
```
通过这个命令来安装插件，默认的会存放到GOBIN，如果没有配置%GOBIN%环境变量，那么会默认安装到%GOPATH%下的bin目录

### gopm使用

- `gopm help get` 或 `gopm get -h`  //查看gopm get帮助信息
- `gopm get xxx`  //可以将指定的包下载到gopm的本地仓库`~/.gopm/repos`
- `gopm get -g xxx` //可以将指定的包下载到GOPATH下
- `gopm get -u xxx` //可以将指定包下载到GOPATH下，并同时下载包的依赖
- `gopm get -l xxx` //可以将指定的包下载到当前所在目录(不常用)
- `gopm get -v xxx` //显示下载包的进度信息

例如:
```bash
gopm get -g -u -v golang.org/x/net
```
> 用gopm get -g代替go get ； 不采用-g参数，会把依赖包下载.vendor目录下面 ； 采用-g 参数，可以把依赖包下载到GOPATH目录中；
