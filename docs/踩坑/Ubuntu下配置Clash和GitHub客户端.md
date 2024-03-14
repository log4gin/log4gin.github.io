# [Ubuntu下配置 Clash和GitHub客户端](https://github.com/log4gin/log4gin/issues/13)


[Clash下载](https://github.com/Fndroid/clash_for_windows_pkg/releases)
Ubuntu下Clash正常配置完成后还需要系统这边开启代理

配置完成后就可以正常科学上网了

然后我去GitHub桌面版Push代码的时候出现了

```
OpenSSL SSL_connect: SSL_ERROR_SYSCALL in connection to github.com:443
```

原因是Git的Http代理的问题，Git支持三种协议：git、ssh和http，本来push的时候应该走ssh隧道的，但是因为设置了http代理，所以就走了http的代理

在终端下输入取消http代理

```
git config --global --unset http.proxy1
```

Push成功 ：）