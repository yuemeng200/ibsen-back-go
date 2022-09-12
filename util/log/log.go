package log

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
)

func Logger(format string, texts ...interface{}) {
	prefix := "%s:%d "
	_, file, line, _ := runtime.Caller(1)
	str := fmt.Sprintf(prefix+format, file, line, texts)
	fmt.Fprintf(gin.DefaultWriter, str)
}
