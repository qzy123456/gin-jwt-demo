FROM golang

ENV GOPROXY=https://goproxy.cn,direct
   #开启go mod 模式
ENV CGO_ENABLED 0
    #必须配置为0，否则docker容器中编译失败，CGO_ENABLED=0的情况下，Go采用纯静态编译，避免各种动态链接库依赖的问题
RUN mkdir /app
WORKDIR  /app
  #切换到工作路径，建议到/go/src 路径下，曾在将项目文件拷贝至容器时，由于配置其他项目，导致一直拷贝不成功
ADD .  /app
   #将容器外项目文件拷贝至容器中
RUN go mod tidy
  #安装依赖
RUN go build main.go
#编译
EXPOSE 8888
#最终运行docker的命令
ENTRYPOINT  ["./main","-conf","cofigs/demo"]