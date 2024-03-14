# nmap学习笔记


### ping ip

```shell
nmap -sP ip
```

### 扫描之前先不进行ping，防止主机禁用ping

```shell
nmap -P0 ip
```

### 查看开放了的端口

```shell
nmap --open ip
```

### 半开放扫描，不会记入日志，需要root权限

```shell
nmap -sS ip
```

### UDP扫描，一般结果不准

```shell
nmap -sU ip
```

### 穿过防火墙扫描

```shell
nmap -sA ip
```

### 查看端口服务及其版本

```shell
nmap -sV ip
```

### 探测主机的漏洞，容易误报，可以看到主机信息

```shell
nmap -O ip
```

### 全面扫描

```shell
nmap -A ip
```

### 显示扫描进程

```shell
nmap -v ip
```

### 控制扫描的速度，有0-5个等级

 ```shell
nmap -T4 ip
 ```

### 扫描整个子网

```shell
nmap 192.168.25.1/24
```

### 扫描多个主机

```shell
nmap ip0 ip1
```

### 扫描文件中的主机列表

```shell
nmap -iL ips.txt
```

### 扫描指定端口范围

```shell
nmap -p 80-8080
```

---

### 启动脚本扫描

```shell
nmap --script=脚本 ip
```

- auth：负责处理鉴权证书、绕开鉴权的脚本。
- broadcast：处理在局域网内探查更多服务开启的状况，如 dhcp / dns / sqlserver 等服务。
- brute：提供暴力破解方式，针对常见的应用如 http / snmp 等。
- default：使用 sC 或 A 选项时默认的脚本，提供基本脚本扫描能力。
- discovery：挖掘更多的网络服务信息，如 smb 枚举、snmp 查询等。
- dos：用于进行拒绝服务攻击。
- exploit：利用已知的漏洞入侵系统。
- external：利用第三方的数据库或资源，如进行 whois 解析。
- fuzzer：模糊测试脚本，发送异常的包到目标主机，探测出潜在的漏洞。
- malware：探测目标是否感染了病毒，是否开启了后门。
- safe：与 fuzzer 功能相反，属于安全性脚本。
- version：负责增强信性服务与版本扫描功能的脚本。
- vuln：负责检查目标主机是否有常见的漏洞，如 ms08_067。



> 学习于[初学kali之namp使用](https://www.nctry.com/878.html)