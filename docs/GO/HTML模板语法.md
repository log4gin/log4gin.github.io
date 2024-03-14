# GO的HTML模板语法

#### 所有的操作都在`{{ }}`里面使用

#### `.`代表的是当前作用域的当前作用对象

```
{{- . -}}

{{- /* 上面是去除 . 左右的空白，我是不留行注释 */}}
```

#### 变量

```
定义
{{- $how_long :=(len "output")}}
使用
{{- $how_long = 7 }}

顶级变量
{{ $ }} 表示传入模板的变量，不可变更

```

#### 判断
```
{{if x}}
 T0
{{else}}

    {{if y}} T1
    {{else if z }}  T2
    {{end}}
{{end}}

x y z 数据为零值时作为false
```
#### 比较函数
```
eq arg1 arg2：
    arg1 == arg2时为true
ne arg1 arg2：
    arg1 != arg2时为true
lt arg1 arg2：
    arg1 < arg2时为true
le arg1 arg2：
    arg1 <= arg2时为true
gt arg1 arg2：
    arg1 > arg2时为true
ge arg1 arg2：
    arg1 >= arg2时为true
```

#### 迭代
```
{{range $x := . -}}
{{$y := 333}}
{{- if (gt $x 33)}}
{{println $x $y ($z := 444)}}{{- end}}


{{else}}
如果 . 为零值 就能看到我了
{{- end}}

. = []int{11,22,33,44,55}
```

#### 选择渲染
```
{{with pipeline}} T1
{{else}} T0
{{end}}

pipeline为零值渲染T0，否则渲染T1
```

#### 定义模板
```
{{- define "T1"}}
我是T1模板，打印传入参数 {{println .}}
{{end}}
```
#### 使用模板
```
{{template "T1" "haha"}}
```

#### 默认模板
```
{{block "T1" .}} one {{end}}

如果T1存在，则执行找到的T1
如果没找到T1，则临时定义一个{{define "T1"}} one {{end}}，并执行它
```

---

 [Go-template用法](https://www.cnblogs.com/f-ck-need-u/p/10053124.html)

 [配合Gin使用](https://gin-gonic.com/zh-cn/docs/examples/html-rendering/)

 [配合embed打包](https://zhuanlan.zhihu.com/p/372800639?utm_id=0)