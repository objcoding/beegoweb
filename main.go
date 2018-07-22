package main

import (
	"beegoweb/controller"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"os"
	"bufio"
)

type People struct {
	Name string
	Age  int
	Sex  string
}

func (p *People) Eat(thing string) string {
	return "eat" + thing
}

func Fuck()  {
	fmt.Println("fuck you")
}

func TestRflect() {

	// 函数反射
	f := reflect.ValueOf(Fuck)
	f.Call(nil)

	var man *People
	man = &People{Name: "zhangchenghui", Age: 18, Sex: "男"}
	man.Eat("香蕉")


	fmt.Println("-----------基本类型--------------")
	str := "我唱出心里话时，眼泪会泪"

	st := reflect.TypeOf(str)
	println(st.Name())

	sv := reflect.ValueOf(str)
	println(sv.String())

	t := reflect.TypeOf(man)
	v := reflect.ValueOf(man)

	fmt.Println("-----------Server--------------")
	var s http.Server
	ss := reflect.TypeOf(s)
	for i := 0; i < ss.NumField(); i++ {
		fmt.Println(ss.Field(i).Name)
	}
	fmt.Println("-----------Server--------------")
	sss := reflect.ValueOf(s)
	for i := 0; i < sss.NumField(); i++ {
		fmt.Println(ss.Field(i).Name)
	}
	// p.Elem()  实际上就是*p

	// set
	fmt.Println(v.Elem().FieldByName("Name").CanSet())
	if v.Elem().FieldByName("Name").CanSet() {
		v.Elem().FieldByName("Name").SetString("zch")
	}
	fmt.Println(v.Elem().FieldByName("Name").Type())

	// get
	fmt.Println(v.Elem().FieldByName("Name").String() + "-----")

	fmt.Println("----------t field-------------")
	// 属性
	for i := 0; i < t.Elem().NumField(); i++ {
		fmt.Println(t.Elem().Field(i))
		fmt.Println(t.Elem().Field(i).Type)
		fmt.Println(t.Elem().Field(i).Name)
	}

	fmt.Println("----------v field-------------")
	for i := 0; i < v.Elem().NumField(); i++ {
		fmt.Println(v.Elem().Field(i))
		fmt.Println(v.Elem().Field(i).Type())
		fmt.Println(v.Elem().Field(i).CanSet())
	}
	fmt.Println(v.Elem().FieldByName("Name"))
	fmt.Println("----------------------------")

	// 方法
	fmt.Println("---------t method-----------")
	for i := 0; i <= t.Elem().NumMethod(); i++  {
		fmt.Println(t.Method(i).Name)
	}
	fmt.Println("---------v method-----------")
	for i := 0; i <= t.Elem().NumMethod(); i++  {
		args := []reflect.Value{reflect.ValueOf("反射")}
		fmt.Println(v.Method(i).Call(args))
		fmt.Println(v.Method(i))
	}
	fmt.Println("-----------------------------")

	// 类型信息
	methodt, _ := t.MethodByName("Eat")
	fmt.Println(methodt.Type)
	fmt.Println(t)
	fmt.Println(v)

	// 通过反射修改信息
	//v2 := reflect.ValueOf(&man)
	//myMethod := v2.MethodByName("eat")
	//args := []reflect.Value{reflect.ValueOf("haha")}
	//fmt.Println(myMethod.Call(args))

	// 指针new
	var i *int
	i = new(int)
	*i = 10
	fmt.Println(*i)
}

func main() {
	go TestRflect()
	http.HandleFunc("/test", controller.Test)
	http.Handle("/", http.FileServer(http.Dir("static/")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("fffff")
	} else {
		fmt.Println("hello, ", input)
	}

	type mystring string

	ints := []int{1, 2, 3, 4, 5}
	for i, t := range ints  {
		fmt.Println(i, t)
	}

}
