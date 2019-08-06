## 从零开始搭建Go语言开发环境

一步一步，从零开始搭建Go语言开发环境。

# 安装Go语言及搭建Go语言开发环境
![](/static/image/development/install_go_dev_1.jpg)

### 下载

#### 下载地址go

Go官网下载地址：[https://golang.org/dl/](https://golang.org/dl/)

Go官方镜像站（推荐）：[https://golang.google.cn/dl/](https://golang.google.cn/dl/)

#### 版本的选择

Windows平台和Mac平台推荐下载可执行文件版，Linux平台下载压缩文件版。
![](/static/image/development/install_go_dev_2.png)

### 安装

#### Windows安装

此安装实例以 `64位Win10`系统安装 `Go1.11.5可执行文件版本`为例。

将上一步选好的安装包下载到本地。  
![](/static/image/development/install_go_dev_3.png)
  
双击下载好的文件

![](/static/image/development/install_go_dev_4.png)

![](/static/image/development/install_go_dev_5.png)

![](/static/image/development/install_go_dev_6.png)

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
![](/static/image/development/mac_install_go.png)

#### 检查
上一步安装过程执行完毕后，可以打开终端窗口，输入`go version`命令，查看安装的Go版本。
