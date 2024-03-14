# Go unsafe

## unsafe的使用

**本教程基于`go1.17`，为了简化使用所以不会引入`uniptr`**

所以大致的方案就是`unsafe.Pointer <==> ptr`,不安全指针和安全指针之间的转换

### 修改未导出的结构体

```go
// 假设有这么个结构体
type person struct{
	name string
	age int
	status string
}

// 拿到了一个结构体的指针p
*p.name = "小明"
*p.age = 18
*p.status = "原始状态"

// 修改第一个字段
name := (*string)(unsafe.Poniter(p))
*name = "小红"

// 修改status
// Add 入参 一个不安全指针的地址 一个是位置偏移量 返回一个距离偏移量的不安全指针
// Add 类似于 C语言 里面的指针运算
// Pointer 类似于 C语言 里面的 *void
status := (*string)(unsafe.Add(unsafe.Pointer(p),unsafe.Sizeof(int(0) + unsafe.Sizeof(string("")))))
*status = "状态已经被修改"

```

### 字符串转数组

#### 以前的的写法

```go
s := "我是需要被转的字符串"
ByteSliceHeader := (*reflrct.SliceHeader)(unsafe.Pointer(&s))
ByteSliceHeader.Len = 4 // 可以自定义切片结构的数据
ByteSlice := (*[]byte)(unsafe.Pointer(ByteSliceHeader))
```

#### 现在的写法

```go
s := "我是需要被转的字符串"
// StringData 返回字符串的数据不安全指针
// Slice 入参第一个是数据不安全指针 第二个是容量和长度 返回的是与指针同类型的一个切片
// Slice 等同于 (*[len]Type)(unsafe.Pionter(ptr))[:]
ByteSlice := unsafe.Slice(unsafe.StringData(s),4)
```

#### 数组转字符串

```go
// 这里使用刚才的ByteSlice
s := (*string)(unsafe.Pointer(*ByteSlice))
```

### 什么时候使用unsafe

1. 追求高性能的时候
2. 需要强制修改数据内存的时候

过程文件在[这里](https://github.com/log4gin/log4gin.github.io/tree/main/src/use_unsafe),和上面的示列略有不同，但内容一致