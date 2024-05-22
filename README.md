
# Gerror for error handling
[![Go Report Card](https://goreportcard.com/badge/github.com/ysfgrl/gerror)](https://goreportcard.com/report/github.com/ysfgrl/gerror)
[![GoDoc](https://godoc.org/github.com/ysfgrl/fibersocket?status.svg)](https://godoc.org/github.com/ysfgrl/gerror)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/ysfgrl/gerrorblob/master/LICENSE)



## ⚙️ Installation

```
go get -u github.com/ysfgrl/gerror
```


## ⚡️ [Examples](https://github.com/ysfgrl/gerror/tree/master/examples)

```go

func NewUserError() error {
    return errors.New("NewUserError")
}
func NewUserError2() *Error {
    if err := NewUserError(); err != nil {
        return GetError(err)
    }
    return nil
}

func main() {
    err := NewUserError()
    if err != nil {
        gerr := GetError(err)
        gerr.PrintConsole()
    }
    gerr := NewUserError2()
    if gerr != nil {
        gerr.PrintConsole()
    } else {
        t.Fatalf("%v", "Error TestGError")
    }
}
```
---

