package gerror

import (
	"encoding/json"
	"path/filepath"
	"runtime"
	"strings"
)

type Error struct {
	File     string `json:"file"`
	Function string `json:"function"`
	Detail   any    `json:"detail"`
	Line     int    `json:"line"`
	Code     string `json:"code"`
	Err      error  `json:"-"`
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
	print(objectJSON, "\n\n")
}

func GetError(err error) *Error {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	function := runtime.FuncForPC(pc[0])
	file, line := function.FileLine(pc[0])
	return &Error{
		Code:     "internal",
		File:     strings.Replace(filepath.Base(file), ".go", "", 1),
		Function: filepath.Base(function.Name()),
		Line:     line,
		Detail:   err.Error(),
		Err:      err,
	}
}

func UserError(msg string, code string) *Error {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	function := runtime.FuncForPC(pc[0])
	file, line := function.FileLine(pc[0])
	return &Error{
		Code:     code,
		File:     strings.Replace(filepath.Base(file), ".go", "", 1),
		Function: function.Name(),
		Line:     line,
		Detail:   msg,
	}
}
