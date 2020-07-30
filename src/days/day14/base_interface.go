package main

import "fmt"

type pad interface {
	music()
}

//	接口
type phone interface {
	//	这里则相当接口的继承
	pad
	call()
	play(name string)
}

type nokiaPhone struct {
}

func (n nokiaPhone) call() {
	fmt.Println("诺基亚打电话")
}

func (n nokiaPhone) play(name string) {
	fmt.Println("诺基亚可以玩" + name)
}

func (n nokiaPhone) music() {
	fmt.Println("诺基亚播放音乐")
}

type iphone struct {
}

func (i iphone) call() {
	fmt.Println("iphone 打电话")
}

func main() {
	var np phone = nokiaPhone{}
	//	iphone 没有实现phone的play方法，所以不能用phone类型接收
	var ip iphone = iphone{}
	np.music()
	np.call()
	np.play("马里奥")
	ip.call()
}
