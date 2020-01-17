package zerrors

import (
    "fmt"
    "io"
)

// 创建一个包含完整调用链的错误
func New(message string) error {
    return &fundamental{
        msg:   message,
        stack: callers(),
    }
}

// 创建一个包含完整调用链的错误
func Newf(format string, args ...interface{}) error {
    return &fundamental{
        msg:   fmt.Sprintf(format, args...),
        stack: callers(),
    }
}

// 创建一个包含完整调用链的错误
func Errorf(format string, args ...interface{}) error {
    return &fundamental{
        msg:   fmt.Sprintf(format, args...),
        stack: callers(),
    }
}

type fundamental struct {
    msg string
    *stack
}

func (f *fundamental) Error() string { return f.msg }

func (f *fundamental) Format(s fmt.State, verb rune) {
    switch verb {
    case 'v':
        if s.Flag('+') {
            io.WriteString(s, f.msg)
            f.stack.Format(s, verb)
            return
        }
        fallthrough
    case 's':
        io.WriteString(s, f.msg)
    case 'q':
        fmt.Fprintf(s, "%q", f.msg)
    }
}

// 让一个error持有当前完整调用链
func WithStack(err error) error {
    if err == nil {
        return nil
    }
    return &withStack{
        err,
        callers(),
    }
}

type withStack struct {
    error
    *stack
}

func (w *withStack) Cause() error { return w.error }

func (w *withStack) Format(s fmt.State, verb rune) {
    switch verb {
    case 'v':
        if s.Flag('+') {
            fmt.Fprintf(s, "%+v", w.Cause())
            w.stack.Format(s, verb)
            return
        }
        fallthrough
    case 's':
        io.WriteString(s, w.Error())
    case 'q':
        fmt.Fprintf(s, "%q", w.Error())
    }
}

// 对一个错误进行描述并包含当前完整调用链
func Wrap(err error, message string) error {
    if err == nil {
        return nil
    }
    err = &withMessage{
        cause: err,
        msg:   message,
    }
    return &withStack{
        err,
        callers(),
    }
}

// 对一个错误进行描述并包含当前完整调用链
func Wrapf(err error, format string, args ...interface{}) error {
    if err == nil {
        return nil
    }
    err = &withMessage{
        cause: err,
        msg:   fmt.Sprintf(format, args...),
    }
    return &withStack{
        err,
        callers(),
    }
}

// 对一个错误进行描述, 不包含调用信息
func WithMessage(err error, message string) error {
    if err == nil {
        return nil
    }
    return &withMessage{
        cause: err,
        msg:   message,
    }
}

// 对一个错误进行描述, 不包含调用信息
func WithMessagef(err error, format string, args ...interface{}) error {
    if err == nil {
        return nil
    }
    return &withMessage{
        cause: err,
        msg:   fmt.Sprintf(format, args...),
    }
}

type withMessage struct {
    cause error
    msg   string
}

func (w *withMessage) Error() string { return w.msg + ": " + w.cause.Error() }
func (w *withMessage) Cause() error  { return w.cause }

func (w *withMessage) Format(s fmt.State, verb rune) {
    switch verb {
    case 'v':
        if s.Flag('+') {
            fmt.Fprintf(s, "%+v\n", w.Cause())
            io.WriteString(s, w.msg)
            return
        }
        fallthrough
    case 's', 'q':
        io.WriteString(s, w.Error())
    }
}

// 获取错误原因, 它会透过zerrors模块的包装获取最开始的错误
func Cause(err error) error {
    type causer interface {
        Cause() error
    }

    for err != nil {
        cause, ok := err.(causer)
        if !ok {
            break
        }
        err = cause.Cause()
    }
    return err
}

// 获取详细的错误描述
func ToDetailString(err error) string {
    return fmt.Sprintf("%+v", err)
}

// 获取简要的错误描述
func ToString(err error) string {
    return err.Error()
}
