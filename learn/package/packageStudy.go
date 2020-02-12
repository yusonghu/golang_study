//	包（package）是多个Go源码的集合，是一种高级的代码复用方案，Go语言为我们提供了很多内置包，如fmt、os、io等。
//	定义包
package myPackage

//	导入外部包<多行导入>
import (
	"fmt"
	"log"
	//	匿名导入
	_ "github.com/go-sql-driver/mysql"
)

//	如果想在一个包中引用另外一个包里的标识符（如变量、常量、类型、函数等）时，该标识符必须是对外可见的（public）。
//	在Go语言中只需要将标识符的首字母大写就可以让标识符对外可见了。
//	首字母小写，外部包不可见，只能在当前包内使用
//	首字母大写，外部包可见,可在其他包中使用
var a = 100

const Mode = 1

type Person struct {
	name string
	age  int8
}
type Run interface {
	//	仅包内访问的方法
	init()
	//	可在包外访问
	Speed()
}

func add(a, b int) int {
	return a + b
}

//init()初始化函数
//在导包的时候会调用,Go语言包会从main包开始检查其导入的所有包，每个包中又可能导入了其他的包。
func init() {
	fmt.Println("调用了init函数")
}
func main() {
	fmt.Println("测试包的导入")
	log.Println("日志包导入")
}
