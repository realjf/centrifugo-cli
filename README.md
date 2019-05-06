# centrifugo-cli
centrifugo v2.2.1 command line app

```
# 首先运行程序
./centrifugo-cli
# 建立链接
connect --address=localhost --user=123 --port=8000
# ping命令
ping
# 订阅频道
subscribe -c=public
# 发布信息
publish -c=public -d=hello
# 取消订阅
unsubscribe -c=public
# 广播信息
broadcast -c=public -d=hello
# 显示帮助信息
help [command]
```
