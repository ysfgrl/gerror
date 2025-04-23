package gerror

import (
	"encoding/json"
	"path/filepath"
	"runtime"
	"strings"
)

//type ErrorHandler = func(error Error)

var (
	LevelError = "error"
	LevelFatal = "fatal"
	LevelWarn  = "warn"
	LevelInfo  = "info"
)

//var errorPool *Pool
//
//func InitErrorHandler(handler ErrorHandler) {
//	errorPool = &Pool{
//		handler: handler,
//		eChan:   make(chan Error),
//	}
//	go errorPool.start()
//}
//
//type Pool struct {
//	eChan   chan Error
//	handler ErrorHandler
//}
//
//func (p *Pool) add(error *Error) {
//	p.eChan <- *error
//}
//
//func (p *Pool) start() {
//	for {
//		select {
//		case e := <-p.eChan:
//			p.handler(e)
//		}
//	}
//}

type Error struct {
	File     string `json:"file"`
	Function string `json:"function"`
	Detail   any    `json:"detail"`
	Line     int    `json:"line"`
	Code     string `json:"code"`
	Err      error  `json:"-"`
	Level    string `json:"level"`
}

func (e *Error) ToJsonByte() []byte {
	b, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		print("Error:" + err.Error())
		return nil
	}
	return b
}

func (e *Error) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	err := json.Unmarshal(e.ToJsonByte(), &result)
	if err != nil {
		print("Error:" + err.Error())
	}
	return result
}

func (e *Error) PrintConsole() {
	objectJSON := e.ToJsonByte()
	print(string(objectJSON), "\n\n")
}

func GetError(err error) *Error {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	function := runtime.FuncForPC(pc[0])
	file, line := function.FileLine(pc[0])
	e := &Error{
		Code:     "internal",
		File:     strings.Replace(filepath.Base(file), ".go", "", 1),
		Function: filepath.Base(function.Name()),
		Line:     line,
		Detail:   err.Error(),
		Err:      err,
		Level:    LevelFatal,
	}
	//if errorPool != nil {
	//	errorPool.add(e)
	//}
	return e
}

func GetErrorCode(err error, code string) *Error {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	function := runtime.FuncForPC(pc[0])
	file, line := function.FileLine(pc[0])
	e := &Error{
		Code:     code,
		File:     strings.Replace(filepath.Base(file), ".go", "", 1),
		Function: filepath.Base(function.Name()),
		Line:     line,
		Detail:   err.Error(),
		Err:      err,
	}
	//if errorPool != nil {
	//	errorPool.add(e)
	//}
	return e
}

func UserError(code string, level string) *Error {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	function := runtime.FuncForPC(pc[0])
	file, line := function.FileLine(pc[0])
	return &Error{
		Code:     code,
		File:     strings.Replace(filepath.Base(file), ".go", "", 1),
		Function: function.Name(),
		Line:     line,
		Detail:   "",
		Level:    level,
	}
}
