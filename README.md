# 能保留调用链的错误包

---

# 获得

`go get -u github.com/zlyuancn/zerrors`

# 此代码非原创

+ 参考 [github.com/pkg/errors](github.com/pkg/errors)

# 文档

[godoc](https://godoc.org/github.com/zlyuancn/zerrors)

# 示例

```go
package main

import (
    "fmt"
    "github.com/zlyuancn/zerrors"
)

func f2() error {
    return zerrors.New("f2")
}

func f1() error {
    err := f2()
    err = zerrors.Wrap(err, "f1")
    return err
}

func main() {
    err := f1()
    fmt.Println(zerrors.ToString(err))
    fmt.Println(zerrors.ToDetailString(err))
}
```

# 方法描述

## 完整调用链

```
创建一个包含完整调用链的错误
    New(message string) error 
    Newf(format string, args ...interface{}) error
让一个error持有当前完整调用链
    WithStack(err error) error
对一个错误进行描述并包含当前完整调用链
    Wrap(err error, message string) error
    Wrapf(err error, format string, args ...interface{}) error
```

## 简单调用信息

```
创建一个包含当前调用信息的错误
    NewSimple(message string) error
    NewSimplef(format string, args ...interface{}) error
让一个error持有当前调用信息
    WithSimple(err error) error
对一个错误进行描述并包含当前调用信息
    WrapSimple(err error, message string) error
    WrapSimplef(err error, format string, args ...interface{}) error
```

## 不包含调用信息

```
对一个错误进行描述, 不包含调用信息
    WithMessage(err error, message string) error
    WithMessagef(err error, format string, args ...interface{}) error
```

## 工具

```
获取错误原因, 它会透过zerrors模块的包装获取最开始的错误
    Cause(err error) error
获取详细的错误描述
    ToDetailString(err error) string
获取简要的错误描述
    ToString(err error) string
```
