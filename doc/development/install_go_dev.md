## 从零开始搭建Go语言开发环境

一步一步，从零开始搭建Go语言开发环境。

# 安装Go语言及搭建Go语言开发环境
![](/doc/static/image/development/install_go_dev_1.jpg)

### 下载

#### 下载地址

Go官网下载地址：[https://golang.org/dl/](https://golang.org/dl/)

Go官方镜像站（推荐）：[https://golang.google.cn/dl/](https://golang.google.cn/dl/)

#### 版本的选择

Windows平台和Mac平台推荐下载可执行文件版，Linux平台下载压缩文件版。
![](/doc/static/image/development/install_go_dev_2.png)

### 安装

#### Windows安装

此安装实例以 `64位Win10`系统安装 `Go1.11.5可执行文件版本`为例。

将上一步选好的安装包下载到本地。  
![](/doc/static/image/development/install_go_dev_3.png)
  
双击下载好的文件

![](/doc/static/image/development/install_go_dev_4.png)

![](/doc/static/image/development/install_go_dev_5.png)

![](/doc/static/image/development/install_go_dev_6.png)

#### Linux下安装

我们在版本选择页面选择并下载好`go1.11.5.linux-amd64.tar.gz`文件：
```bash
wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
```
将下载好的文件解压到`/usr/local`目录下：
```bash
mkdir -p /usr/local/go  # 创建目录
tar -C /usr/lcoal/go zxvf go1.11.5.linux-amd64.tar.gz. # 解压
```
如果提示没有权限，加上`sudo`以root用户的身份再运行。执行完就可以在`/usr/local/`下看到go目录了。

配置环境变量： Linux下有两个文件可以配置环境变量，其中`/etc/profile`是对所有用户生效的；`$HOME/.profile`是对当前用户生效的，根据自己的情况自行选择一个文件打开，添加如下两行代码，保存退出。
```bash
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```
修改`/etc/profile`后要重启生效，修改`$HOME/.profile`后使用source命令加载`$HOME/.profile`文件即可生效。   
检查:
```bash
~ go version
go version go1.11.5 linux/amd64
```
#### Mac下安装

下载可执行文件版，直接点击**下一步**安装即可，默认会将go安装到`/usr/local/go`目录下。
![](/doc/static/image/development/mac_install_go.png)

#### 检查
上一步安装过程执行完毕后，可以打开终端窗口，输入`go version`命令，查看安装的Go版本。

### 配置GOPATH

`GOPATH`是一个环境变量，用来表明你写的go项目的存放路径（工作目录）。

`GOPATH`路径最好只设置一个，所有的项目代码都放到`GOPATH`的`src`目录下。

Linux和Mac平台就参照上面配置环境变量的方式将自己的工作目录添加到环境变量中即可。 Windows平台按下面的步骤将`D:\code\go`添加到环境变量：

![](/doc/static/image/development/install_go_dev_7.png)

![](/doc/static/image/development/install_go_dev_8.png)

![](/doc/static/image/development/install_go_dev_9.png)

![](/doc/static/image/development/install_go_dev_10.png)

![](/doc/static/image/development/install_go_dev_11.png)

![](/doc/static/image/development/install_go_dev_12.png)

![](/doc/static/image/development/install_go_dev_13.png)

在 Go 1.8 版本之前，`GOPATH`环境变量默认是空的。从 Go 1.8 版本开始，Go 开发包在安装完成后会为 `GOPATH`设置一个默认目录，参见下表。

**GOPATH在不同操作系统平台上的默认值**

| 平台 | GOPATH默认值	| 举例 |
| ----- | ----- | ----- |
| Windows | %USERPROFILE%/go | C:\Users\用户名\go |
| Unix | $HOME/go | /home/用户名/go |
同时，我们将 `GOROOT`下的bin目录及`GOPATH`下的bin目录都添加到环境变量中。

### Go项目结构 

在进行Go语言开发的时候，我们的代码总是会保存在`$GOPATH/src`目录下。在工程经过`go build`、`go install`或`go get`等指令后，会将下载的第三方包源代码文件放在`$GOPATH/src`目录下， 产生的二进制可执行文件放在 `$GOPATH/bin`目录下，生成的中间缓存文件会被保存在 `$GOPATH/pkg` 下。

如果我们使用版本管理工具（Version Control System，VCS。常用如Git）来管理我们的项目代码时，我们只需要添加`$GOPATH/src`目录的源代码即可。`bin` 和 `pkg` 目录的内容无需版本控制。

### 适合个人开发者


