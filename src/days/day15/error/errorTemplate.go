package error

import (
	"fmt"
	"github.com/pkg/errors"
	"net"
	"os"
)

func ToError() error {
	f, err := os.Open("/test.txt")
	if err != nil {
		// 自定义错误
		return errors.New("未找到文件")
	}
	fmt.Println(f.Name())
	return nil
}

//	自定义错误类型
type error interface {
	Error() string
}

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func ToNet() {
	addr, err := net.LookupHost("golangbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}

type AgentContext struct {
	name string
}

//	错误处理
func (self *AgentContext) IsValidHostType(hostType string) bool {
	return hostType == "virtual_machine" || hostType == "bare_metal"
}
