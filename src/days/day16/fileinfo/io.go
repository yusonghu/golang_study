package fileinfo

import (
	"fmt"
	"io"
	"log"
	"os"
)

func FileRead() {
	/*
		读取数据：
			Reader接口：
				Read(p []byte)(n int, error)
	*/
	//读取本地aa.txt文件中的数据
	//step1：打开文件
	fileName := "log.log"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//step3：关闭文件
	defer file.Close()

	//step2：读取数据
	bs := make([]byte, 4, 4)

	////第一次读取
	//n, err := file.Read(bs)
	//fmt.Println(err)        //<nil>
	//fmt.Println(n)          //4
	//fmt.Println(bs)         //[97 98 99 100]
	//fmt.Println(string(bs)) //abcd
	//
	////第二次读取
	//n, err = file.Read(bs)
	//fmt.Println(err)        //<nil>
	//fmt.Println(n)          //4
	//fmt.Println(bs)         //[101 102 103 104]
	//fmt.Println(string(bs)) //efgh
	//
	////第三次读取
	//n, err = file.Read(bs)
	//fmt.Println(err)        //<nil>
	//fmt.Println(n)          //2
	//fmt.Println(bs)         //[105 106 103 104]
	//fmt.Println(string(bs)) //ijgh
	//
	////第四次读取
	//n, err = file.Read(bs)
	//fmt.Println(err) //EOF
	//fmt.Println(n)   //0
	//
	n := -1
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("读取到了文件的末尾，结束读取操作。。")
			break
		}
		fmt.Println(string(bs[:n]))
	}
}

func FileWrite() {
	/*
		写出数据：
	*/

	fileName := "iowrite.txt"
	//step1：打开文件
	//step2：写出数据
	//step3：关闭文件
	//file,err := os.Open(fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//写出数据
	//bs :=[]byte{65,66,67,68,69,70}//A,B,C,D,E,F
	bs := []byte{97, 98, 99, 100} //a,b,c,d
	n, err := file.Write(bs[:2])
	fmt.Println(n)
	HandleErr(err)
	file.WriteString("\n")

	//直接写出字符串
	n, err = file.WriteString("HelloWorld")
	fmt.Println(n)
	HandleErr(err)

	//	字符串字节数组
	file.WriteString("\n")
	n, err = file.Write([]byte("today"))
	fmt.Println(n)
	HandleErr(err)
}

func FileCopy1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if HandleErr(err) != nil {
		return -1, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if HandleErr(err) != nil {
		return -1, err
	}
	defer file1.Close()
	defer file2.Close()
	//拷贝数据
	bs := make([]byte, 1024, 1024)
	read := -1 //读取的数据量
	total := 0
	for {
		read, err = file1.Read(bs)
		if err != nil && err != io.EOF {
			if HandleErr(err) != nil {
				return total, err
			}
			break
		} else if err == io.EOF || read == 0 {
			fmt.Println("拷贝完毕。。。")
			return total, nil
		}
		total += read
		file2.Write(bs[:read])
	}
	return total, nil
}

func HandleErr(err error) error {
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func FileCopy2(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	buffer := make([]byte, 20)
	return io.CopyBuffer(file2, file1, buffer)
}
