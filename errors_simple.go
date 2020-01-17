/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/1/17
   Description :
-------------------------------------------------
*/

package zerrors

import (
    "fmt"
    "io"
)

// 创建一个包含当前调用信息的错误
func NewSimple(message string) error {
    return &fundamentalSimple{
        msg:         message,
        stackSimple: callerSimple(),
    }
}

// 创建一个包含当前调用信息的错误
func NewSimplef(format string, args ...interface{}) error {
    return &fundamentalSimple{
        msg:         fmt.Sprintf(format, args...),
        stackSimple: callerSimple(),
    }
}

type fundamentalSimple struct {
    msg string
    *stackSimple
}

func (f *fundamentalSimple) Error() string { return f.msg }

func (f *fundamentalSimple) Format(s fmt.State, verb rune) {
    switch verb {
    case 'v':
        if s.Flag('+') {
            io.WriteString(s, f.msg)
            f.stackSimple.Format(s, verb)
            return
        }
        fallthrough
    case 's':
        io.WriteString(s, f.msg)
    case 'q':
        fmt.Fprintf(s, "%q", f.msg)
    }
}

// 让一个error持有当前调用信息
func WithSimple(err error) error {
    if err == nil {
        return nil
    }
    return &withStackSimple{
        err,
        callerSimple(),
    }
}

type withStackSimple struct {
    error
    *stackSimple
}

func (w *withStackSimple) Cause() error { return w.error }

func (w *withStackSimple) Format(s fmt.State, verb rune) {
    switch verb {
    case 'v':
        if s.Flag('+') {
            fmt.Fprintf(s, "%+v", w.Cause())
            w.stackSimple.Format(s, verb)
            return
        }
        fallthrough
    case 's':
        io.WriteString(s, w.Error())
    case 'q':
        fmt.Fprintf(s, "%q", w.Error())
    }
}

// 对一个错误进行描述并包含当前调用信息
func WrapSimple(err error, message string) error {
    if err == nil {
        return nil
    }
    err = &withMessage{
        cause: err,
        msg:   message,
    }
    return &withStackSimple{
        err,
        callerSimple(),
    }
}

// 对一个错误进行描述并包含当前调用信息
func WrapSimplef(err error, format string, args ...interface{}) error {
    if err == nil {
        return nil
    }
    err = &withMessage{
        cause: err,
        msg:   fmt.Sprintf(format, args...),
    }
    return &withStackSimple{
        err,
        callerSimple(),
    }
}
