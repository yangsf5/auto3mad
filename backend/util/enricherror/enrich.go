package enricherror

import (
	"fmt"
	"runtime"
	"strings"
)

func ErrorPosition(tierParam ...int) string {
	tier := 2
	if len(tierParam) > 0 {
		tier = tierParam[0]
	}
	_, file, line, _ := runtime.Caller(tier)
	arr := strings.Split(file, "/")
	return fmt.Sprintf("[%v:%v] ", arr[len(arr)-1], line)
}

func GetErrorContent(err interface{}) string {
	if err == nil {
		return ""
	}

	if content, ok := err.(string); ok {
		return content
	} else if e, ok := err.(error); ok {
		return e.Error()
	}

	return ""
}
