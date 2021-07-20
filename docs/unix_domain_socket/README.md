# Unix Domain Socket

## 运行

### Server

```bash
cd ./server

go run .
# 2021/07/20 12:02:56 Client connected [unix]
```

### Client by cli

```bash
nc -U /tmp/echo.sock

# hello world
```

## 参考资料

- [Unix domain sockets in Go](https://eli.thegreenplace.net/2019/unix-domain-sockets-in-go/) - Go示例