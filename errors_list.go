/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/3/14
   Description :
-------------------------------------------------
*/

package zerrors

import (
    "bytes"
    "errors"
    "fmt"
    "io"
)

// 强制要求实现error接口
var _ error = (*Errors)(nil)

// error列表
type Errors struct {
    errs []error
}

// 创建一个error列表
func NewErrors() *Errors {
    return &Errors{}
}

// 转为error
func (m *Errors) Err() error {
    if len(m.errs) == 0 {
        return nil
    }
    return m
}

// 实现error接口
func (m *Errors) Error() string {
    return m.String()
}

// 返回第一个error
func (m *Errors) FirstErr() error {
    if len(m.errs) == 0 {
        return nil
    }
    return m.errs[0]
}

// 获取所有的err
func (m *Errors) Errs() []error {
    return m.errs
}

// 添加一些错误
func (m *Errors) AddErrs(errs ...error) {
    m.errs = append(m.errs, errs...)
}

// 添加一些错误文本, 自动包含当前调用信息
func (m *Errors) Add(texts ...string) {
    errs := make([]error, len(texts))
    for i, t := range texts {
        errs[i] = &fundamentalSimple{
            msg:         t,
            stackSimple: callerSimple(),
        }
    }
    m.errs = append(m.errs, errs...)
}

// 添加一些错误文本, 不需要调用信息
func (m *Errors) AddNoStack(texts ...string) {
    errs := make([]error, len(texts))
    for i, t := range texts {
        errs[i] = errors.New(t)
    }
    m.errs = append(m.errs, errs...)
}

// 添加一个格式化错误文本, 自动包含当前调用信息
func (m *Errors) Addf(format string, a ...interface{}) {
    m.errs = append(m.errs, &fundamentalSimple{
        msg:         fmt.Sprintf(format, a...),
        stackSimple: callerSimple(),
    })
}

// 添加一个格式化错误文本, 不需要调用信息
func (m *Errors) AddfNoStack(format string, a ...interface{}) {
    m.errs = append(m.errs, fmt.Errorf(format, a...))
}

func (m *Errors) String() string {
    if len(m.errs) == 0 {
        return "<nil>"
    }

    return m.errs[0].Error()
}

func (m *Errors) Format(s fmt.State, verb rune) {
    if len(m.errs) == 0 {
        io.WriteString(s, "<nil>")
        return
    }

    switch verb {
    case 'v':
        var f string
        if s.Flag('+') {
            f = "%d: %+v\n"
        } else {
            f = "%d: %v\n"
        }

        var bs bytes.Buffer
        bs.WriteString("zerrors.Errors: {\n")
        for i, e := range m.errs {
            bs.WriteString(fmt.Sprintf(f, i, e))
        }
        bs.WriteString("}")

        io.WriteString(s, bs.String())
    case 's':
        io.WriteString(s, m.errs[0].Error())
    case 'q':
        fmt.Fprintf(s, "%q", m.errs[0])
    }
}
