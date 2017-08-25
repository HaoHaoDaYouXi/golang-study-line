# Golang
Golang 入门学习
# 安装
## 手动安装新的稳定版本
1.  下载地址：http://www.golangtc.com/download
2.  把安装包解压到/opt/目录下
 
   ```bash
    sudo tar -C /opt/ -xzf go1.8.linux-amd64.tar.gz
    ```
3.  配置go的环境变量
    +   `vim /etc/profile`

        ```bash
        export GOPATH=/home/www/gopath 
        export GOROOT=/usr/local/go
        export GOARCH=386
        export GOOS=linux
        export GOBIN=$GOROOT/bin/
        export GOTOOLS=$GOROOT/pkg/tool/
        export PATH=$PATH:$GOBIN:$GOTOOLS
        ```
    +   重新加载 profile 文件：`source /etc/profile`    
4.  查看Go环境变量：`go env` 

    ```bash
    GOARCH="386"
    GOBIN="/opt/go/bin/"
    GOEXE=""
    GOHOSTARCH="amd64"
    GOHOSTOS="linux"
    GOOS="linux"
    GOPATH="/mnt/hgfs/Linux-Share/Go/GolangPath"
    GORACE=""
    GOROOT="/opt/go"
    GOTOOLDIR="/opt/go/pkg/tool/linux_amd64"
    GCCGO="gccgo"
    GO386=""
    CC="gcc"
    GOGCCFLAGS="-fPIC -m32 -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build922856580=/tmp/go-build -gno-record-gcc-switches"
    CXX="g++"
    CGO_ENABLED="0"
    PKG_CONFIG="pkg-config"
    CGO_CFLAGS="-g -O2"
    CGO_CPPFLAGS=""
    CGO_CXXFLAGS="-g -O2"
    CGO_FFLAGS="-g -O2"
    CGO_LDFLAGS="-g -O2"
    ```
5.  主要变量说明
    +   `GOROOT` 就是go的安装路径
    +   `GOPATH` 是作为编译后二进制的存放目的地和import包时的搜索路径
        +   GOPATH目录结构
6.  代码目录结构规划

    ```bash
    cd $GOPATH
    mkdir src
    cd src
    
    `````  
        
        
        
