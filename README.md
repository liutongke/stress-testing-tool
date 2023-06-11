docker run -i -t -d -p 9600:9500 --name go-1 -v C:\Users\keke\dev\docker:/var/local 276895edf967


go run main.go -c 1 -n 1 -u http://192.168.0.105:9500 -t application/x-www-form-urlencoded -p C:\Users\keke\dev\docker\stress-testing-tool/post.txt

go run main.go -c 500 -n 1000000 -u http://192.168.0.105:9500/chat/register -t multipart/form-data -p C:\Users\keke\dev\docker\stress-testing-tool/post.txt

# **(一).使用**

*   直接使用源码运行，clone 项目源码到本地 **$GOPATH** 目录中。

```
λ go run main.go
示例: go run main.go -c 100 -n 10000 -u http://192.168.0.105:9500/
压测地址必填
当前请求参数: -c 0 -n 0 -u url
Usage of C:\Users\keke\AppData\Local\Temp\go-build67071217\b001\exe\main.exe:
  -c int
        用户数量
  -k int
        复用连接
  -n int
        发起请求数量
  -p string
        postfile，发送POST请求时需要上传的文件
  -t string
        即content-type，用于设置Content-Type请求头信息,post请求必选项
  -u string
        压测地址 (default "url")
```

