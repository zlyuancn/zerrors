# zerrors
能保留堆栈链的错误包

## 获得zerrors
`go get -u github.com/zlyuancn/zerrors`

## 此代码非原创
### 参考 `github.com/pkg/errors`

## 示例
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
