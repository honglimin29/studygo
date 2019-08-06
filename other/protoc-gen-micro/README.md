## protoc-gen-micro

本篇介绍如何使用protobuf的protoc-gen-micro工具快速生成go-micro服务代码

## 安装

### 安装protoc

#### Linux安装

github下载地址:[https://github.com/google/protobuf/releases](https://github.com/google/protobuf/releases)下载一个cpp包
```bash
./configure
make && make install
```

#### Windows安装
github地址:[https://github.com/google/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases/download/v3.9.0/protoc-3.9.0-win64.zip)下载win版本zip包,  
默认zip包为绿色版本，解压配置`PATH`路径指向`bin`目录即可使用，或直接将`bin`目录中的`protoc.exe`二进制文件拷贝到`$GOROOT/bin`目录下

#### 验证安装  

```
[root@docker ~]# protoc --version
libprotoc 3.9.0
```

### 安装protoc-gen-micro

```bash
go get github.com/micro/protoc-gen-micro
```

## 使用




## 官网

[protoc-gen-micro](https://github.com/micro/protoc-gen-micro)

## 扩展学习

[protobuf](https://github.com/google/protobuf)
[protoc-gen-go](https://github.com/golang/protobuf)



