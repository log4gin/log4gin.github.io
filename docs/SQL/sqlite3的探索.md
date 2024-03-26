# 记录对`Sqlite3`的探索

## 历史

原美国海军军方使用，为了简化数据库配置，2000年8月`sqlite 1.0`，2004年`sqlite3`有较大改动

## 仅有的数据类型

1.  NULL
2.  INTERGER
3.  REAL
4.  TEXT
5.  BLOB

## 探索过程

```sqlite
-- 打开数据库
-- 如果没有对数据库进行写入，文件不会被创建
sqlite3 ./test.db

-- 创建一个临时表
create temp table test(
	name text,
    age number
);

-- 保留一个表
create table save(
	name text,
    age number
);

-- 查看当前存在的表
.tables

-- 插入点数据
insert into save values("gin",18);
insert into test values("gin",18);


-- 备份数据库
-- 配置导出格式
.output backup.sql
-- 导出 save,留空就是全部
.dump save

-- 退出
.q

-- 导入备份
sqlite3 ./test.db < ./backup.sql

-- 格式化输出
.h on 
.mode column
.timer on

-- 查看save表结构
.schema save

-- 属性亲和
insert into save values(19,"big-gin");
insert into save values(x'20',20);

-- 查看一下类型
name  typeof(name)  age      typeof(age)
----  ------------  -------  -----------
gin   text          18       integer    
19    text          big-gin  text       
      blob          20       integer 
-- sqlite3 仅仅支持 之前的数据类型
-- 会将其他的亲和类型隐式转换，如果转换失败就保留数据原来的类型直接插入
```



