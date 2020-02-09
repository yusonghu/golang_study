package main

import (
	"fmt"
	"reflect"
)

//在Go语言中，使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	//type name类型（Type）和type kind种类（Kind）
	fmt.Printf("type:%v kind:%v\n", v.Name(), v.Kind())
}

type myInt int64

//通过反射获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

//通过反射设置变量的值
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	reflectType("字符串类型")
	var m myInt
	reflectType(m)
	reflectValue(24.3)
	var a int64 = 100
	reflectSetValue2(&a)
	fmt.Println(a)
}
