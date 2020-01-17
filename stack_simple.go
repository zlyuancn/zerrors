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
    "path"
    "runtime"
)

type stackSimple struct {
    pc   uintptr
    file string
    line int
}

func (m *stackSimple) Format(s fmt.State, verb rune) {
    switch verb {
    case 's':
        if m.pc == 0 {
            fmt.Fprintf(s, "unknown caller")
            return
        }

        switch {
        case s.Flag('+'):
            fn := runtime.FuncForPC(m.pc)
            if fn == nil {
                io.WriteString(s, "unknown")
            } else {
                file, _ := fn.FileLine(m.pc)
                fmt.Fprintf(s, "%s\n\t%s", fn.Name(), file)
            }
        default:
            io.WriteString(s, path.Base(m.file))
        }
    case 'd':
        fmt.Fprintf(s, "%d", m.line)
    case 'n':
        name := runtime.FuncForPC(m.pc).Name()
        io.WriteString(s, funcname(name))
    case 'v':
        io.WriteString(s, "\n")
        m.Format(s, 's')
        io.WriteString(s, ":")
        m.Format(s, 'd')
    }
}

func callerSimple() *stackSimple {
    pc, file, line, _ := runtime.Caller(2)
    return &stackSimple{
        pc:   pc,
        file: file,
        line: line,
    }
}
