# ChatRoom
## 环境部署
### 安装git环境
首先安装git的安装包,下载[https://github.com/git-for-windows/git/releases/download/v2.33.1.windows.1/Git-2.33.1-64-bit.exe](https://github.com/git-for-windows/git/releases/download/v2.33.1.windows.1/Git-2.33.1-64-bit.exe)并执行安装  
需要将git的安装根目录下的bin路径添加到系统环境变量中  
### 安装go环境
然后安装go的安装包,下载[https://dl.google.com/go/go1.17.3.windows-amd64.msi](https://dl.google.com/go/go1.17.3.windows-amd64.msi)并执行安装  
需要将go的安装根目录下的bin路径添加到系统环境变量中  
### 安装protobuf插件
将Git目录中的bin目录添加到系统环境变量中  
### 生成二进制文件
windows下执行installclient.bat,installserver.bat分别编译安装客户端和服务器,安装地址默认go的安装根目录下显的bin目录中  
linux下执行goInstallLibs.sh分别编译安装客户端和服务器,安装地址默认go的安装根目录下显的bin目录中  
## 工程目录说明
### 根目录说明
bin目录: 存放生成的二进制程序  
log目录: 存放进程运行的日志文件  
proto目录: 存放客户端和服务器通信的protobuffer协议  
src目录: 存放golang源码  
### src目录说明
clientlogic目录: 存放客户端的逻辑代码  
serverlogic目录: 存放服务器的逻辑代码  
cmdid目录: 存放根据proto文件生成的胶水代码,包括协议号,协议指针转换,协议回调处理等  
lib目录: 存放封装的公共库文件,包括网络,日志,序列号,协议格式定义等  
protoout目录: 存放由protoc.exe生成的golang语言的序列化相关文件  
mainclient目录: 存放客户端执行的main函数入口文件  
mainserver目录: 存放服务器执行的main函数入口文件  
tools目录: 存放工具脚本  
## 实现原理
每一个连接由一个新创建网络协程处理,网络协程负责数据收发以及协议的序列化与反序列化,通过channel与主逻辑协程通信,主逻辑协程只有一个.  
tcptask是程序作为服务端接受的连接对象,而tcpclient是程序作为客户端发起建立的连接对象  
