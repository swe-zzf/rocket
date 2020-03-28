// 必须加上package
package first

import (
	"errors"
	"fmt"
	"github.com/swe-zzf/rocket/test/learn/second"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

/**
 * go入门-基础语法
 *
 * @author gavin.z, swe.zzf@gmail.com
 * @since  2020-02-23
 */
func init() {
	fmt.Println("first-init()")
}

// 变量定义
var userAge = 30

// 全局变量声明和赋值
var userName = "姓名"

// 一般类型声明
type userSex int

// 结构的声明-相当于java的模型对象类
type StructName struct {
}

// 接口的声明
type InterName interface {
}

// 函数的声明
type funcName func()

// 公开函数定义
func Main() {
	fmt.Println("first公用函数Main()")
	//func101()
	//func102()
	//func103()
	//func104()
	//func105()
	//func106()
	func107()

}

// 函数定义
func func101() {
	fmt.Println("内部函数func101()")
	fmt.Println(os.Getuid())
}

func func102() {
	// 单号注释
	/*
	  多号注释
	 */
	// 常量定义
	const USER_AGE int = 20
	fmt.Println("内部函数func102()")
	fmt.Println("你好")
	fmt.Println("var年龄=", userAge, ",const年龄=", USER_AGE)
	fmt.Println("userName=", userName)
}

func func103() {
	//var k uint8 = 255
	//var k int = 1
	//var k uint = 1
	//var k float64 = 1.5
	//var k bool = false
	//var k rune = 2
	//var k string
	// complex128复数
	var k complex128
	fmt.Println("k=", k)
	fmt.Println("k.size=", unsafe.Sizeof(k))
	//
	fmt.Println("------------- 类型别名")
	// 类型别名 bm
	type bm int32
	var i bm = 1
	var j bm = 2
	//var j int32
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("i.size=", unsafe.Sizeof(i))
	fmt.Println("j.size=", unsafe.Sizeof(j))
	//
	fmt.Println("i+j=", i+j)
	//
	fmt.Println("------------- 变量与常量")
	var _ string
	fmt.Println("下划线变量")
	var (
		a int32 = 2
		b       = "你好"
	)
	fmt.Println("多个变量声明=", a, b)
	c, d, _ := 10, 20, 30
	fmt.Println(":=赋值和声明变量=", c, d)
	//
	var x int32 = 12
	y := float64(x)
	fmt.Println("x,y类型转换=", x, y)
	// 大写字母开头变量为可调用变量
	fmt.Println("公用变量UserAge=", second.UserAge)
	//
	const age int32 = 20
	const name = "名称"
	const _bool bool = false
	fmt.Println("常量=", age, name, _bool)
	const (
		a1 int32  = 20
		a2 string = "中文用3字节"
		a3 bool   = false
	)
	const _len = len(a2)
	fmt.Println("分组常量=", a1, a2, a3, ",类型=", reflect.TypeOf(a3), "len()表达式=", _len)
	//
	const a4 = iota
	const a5 = iota + 2
	const a6 = iota
	fmt.Println("iota赋值常量=", a4, a5, a6)
	const (
		a7 = iota * 10
		_
		a8 = iota
		_
		a9  = iota
		ax  = 900
		a10 = iota
	)
	fmt.Println("iota分组赋值常量=", a7, a8, a9, ax, a10)
	const (
		a20 = iota * 2
		a21
		a22
	)
	fmt.Println("iota分组隐式赋值常量=", a20, a21, a22)
	const (
		a30, a31 = iota, iota + 9
		a32, a33
		a34 = iota
	)
	fmt.Println("iota分组单行赋值常量=", a30, a31, a32, a33, a34)
}

func func104() {
	fmt.Println("------------- 算术运算符")
	var a int32 = 10
	var b int32 = 20
	fmt.Println("a+b=", a+b)
	fmt.Println("a-b=", a-b)
	fmt.Println("a*b=", a*b)
	fmt.Println("b/a=", b/a)
	fmt.Println("求余b%a=", b%a)
	a++
	fmt.Println("a++=", a)
	b--
	fmt.Println("b--=", b)
	fmt.Println("------------- 关系运算符")
	fmt.Println("a>b:", a > b)
	fmt.Println("a<b:", a < b)
	fmt.Println("a==b:", a == b)
	fmt.Println("------------- 逻辑运算符")
	var b1 = true
	var b2 = false
	fmt.Println("b1 && b2=", b1 && b2)
	fmt.Println("b1 || b2=", b1 || b2)
	fmt.Println("!b1=", !b1)
	fmt.Println("------------- 按位运算符")
	var c1 int32 = 0
	var c2 int32 = 1
	var c3 int32 = 12
	fmt.Println("按位与c1&c2=", c1&c2)
	fmt.Println("按位或c1|c2=", c1|c2)
	fmt.Println("按位异或c1^c2=", c1^c2)
	fmt.Println("1左移3位=", c2<<3)
	fmt.Println("10右移2位=", c3>>2)
	c3 >>= 2
	fmt.Println("10右移2位后赋值=", c3)

}

func func105() {
	fmt.Println("------------- if else条件语句")
	x := 10
	if x > 1 {
		fmt.Println("x > 1")
		if x > 5 {
			fmt.Println("x > 5")
		}
	} else {
		fmt.Println("x < 1")
	}
	fmt.Println("")
	fmt.Println("--------- switch选择语句")
	var y interface{}
	y = 32
	switch y.(type) { // 类型判断
	case int:
		fmt.Println("整型int")
	case string:
		fmt.Println("字符串string")
	default:
		fmt.Println("以上不满足条件")
	}
	fmt.Println("")
	fmt.Println("--------- for循环语句")
	//for {
	//	fmt.Println("无限循环。。。")
	//}
	fmt.Println("")
	for i := 0; i < 10; i++ {
		fmt.Println("普通for-i=", i)
	}
	fmt.Println("")
	array := []string{"a", "b", "c", "d", "e"}
	for key, value := range (array) {
		fmt.Println("foreach-key=", key, ",value=", value)
	}
	rangeMap := map[int]string{1: "Golang", 2: "Java", 3: "Python", 4: "C"}
	for k, v := range (rangeMap) {
		fmt.Printf("rangeMap-%d: %s\n", k, v)
	}
	fmt.Println("")
	fmt.Println("--------- goto break控制语句")
	//	k := 1
	//L1:
	//	{
	//		fmt.Println("goto无限循环语句k=", k)
	//		k++
	//		time.Sleep(1 * time.Second)
	//		if (k > 5) {
	//			return
	//		}
	//	}
	//	goto L1
	fmt.Println("")
	for j := 0; j < 10; j++ {
		if j == 4 {
			fmt.Println("continue test...j=", j)
			continue
		}
		if j > 5 {
			fmt.Println("break test...j=", j)
			break
		}
		fmt.Println("for...j=", j)
	}
	n1 := -0x1000
	fmt.Printf("16进制数%X，表示的是%d\n", n1, n1)
	fmt.Println("")
	var n2 = 5.89E-4
	// 这里用到了字符串格式化函数
	// %E用于以带指数部分的表示法显示浮点数类型值，
	// %f用于以通常的方法显示浮点数类型值
	fmt.Printf("浮点数%E，表示的是%f\n", n2, n2)
	fmt.Println("")
	var n3 = 3.7E+1 + 5.98E-2i
	// 这里用到了字符串格式化函数
	// %E用于以带指数部分的表示法显示浮点数类型值，
	// %f用于以通常的方法显示浮点数类型值
	fmt.Printf("浮点数%E，表示的是%f\n", n3, n3)
	fmt.Println("")
	var char1 rune = '赞'
	// 这里用到了字符串格式化函数。其中，%c用于显示rune类型值代表的字符。
	fmt.Printf("字符'%c'的Unicode代码点是%s\n", char1, ("U+8D5E"))
	fmt.Println("")
	var str1 string = "\\\""
	// 这里用到了字符串格式化函数
	// %q用于显示字符串值的表象值并用双引号包裹
	fmt.Printf("用解释型字符串表示法表示的 %q ，所代表的是%s\n", str1, `"\\\""`)
	fmt.Println("")
	var numbers2 [5]int
	numbers2[0] = 2
	numbers2[3] = numbers2[0] - 3
	numbers2[1] = numbers2[2] + 5
	numbers2[4] = len(numbers2)
	sum := 11
	// “==”用于两个值的相等性判断
	fmt.Printf("%v\n",
		(sum == numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]))
	fmt.Println("------------ 切片类型1")
	var numbers3 = [5]int{1, 2, 3, 4, 5}
	slice3 := numbers3[2:len(numbers3)]
	fmt.Println("slice3=", slice3)
	length := len(slice3)
	capacity := cap(slice3)
	fmt.Println("capacity=", capacity)
	fmt.Printf("%v, %v\n", (length == len(slice3)), (capacity == cap(slice3)))
	fmt.Println("------------ 切片类型2")
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice5 := numbers4[4:6:8]
	length5 := len(slice5)
	capacity5 := cap(slice5)
	fmt.Printf("%v, %v\n", length5 == len(slice5), capacity5 == cap(slice5))
	slice5 = slice5[:cap(slice5)]
	slice5 = append(slice5, 11, 12, 13)
	length5 = len(slice5)
	fmt.Printf("%v\n", length5 == len(slice5))
	slice6 := []int{0, 0, 0}
	copy(slice5, slice6)
	e2 := slice5[2]
	e3 := slice5[3]
	e4 := slice5[4]
	fmt.Printf("%v, %v, %v\n", e2 == slice5[2], e3 == slice5[3], e4 == slice5[4])
	fmt.Println("------------ 字典（Map）类型")
	mm2 := map[string]int{"golang": 42, "java": 1, "python": 8}
	mm2["scala"] = 25
	mm2["erlang"] = 50
	mm2["python"] = 0
	fmt.Printf("%d, %d, %d \n", mm2["scala"], mm2["erlang"], mm2["python"])
	fmt.Println("")
	fmt.Println("------------ 通道（Channel）类型")
	ch2 := make(chan string, 1)
	// 下面就是传说中的通过启用一个Goroutine来并发的执行代码块的方法。
	// 关键字 go 后跟的就是需要被并发执行的代码块，它由一个匿名函数代表。
	// 对于 go 关键字以及函数编写方法，我们后面再做专门介绍。
	// 在这里，我们只要知道在花括号中的就是将要被并发执行的代码就可以了。
	go func() {
		ch2 <- "已到达！"
	}()
	var value string = "数据"
	value = value + <-ch2
	fmt.Println(value)
	//
	fmt.Println("")
	fmt.Println("------------ 非缓冲通道（Channel）类型")
	// 发送方在向通道值发送数据的时候会立即被阻塞，直到有某一个接收方已从该通道值中接收了这条数据。
	var myChannel = make(chan int, 0)
	var number = 6
	go func() {
		var sender Sender = myChannel
		sender <- number
		fmt.Println("Sent!")
	}()
	go func() {
		var receiver Receiver = myChannel
		fmt.Println("Received!", <-receiver)
	}()
	// 让执行结束的时间延迟1秒，
	// 以使上面两个代码块有机会被执行。
	time.Sleep(time.Second)

}

type Sender chan<- int
type Receiver <-chan int

func func106() {
	fmt.Println("------------ 函数")
	fmt.Println("匿名函数rs=", anonymousRs)
	fmt.Println("")
	fmt.Println("string类型转换:", strConv(10, 20))
	fmt.Println("")
	fmt.Println("------------ 结构体struct")
	person := struct {
		Name   string
		Gender string
		Age    uint8
	}{"Robert", "Male", 33}
	fmt.Println("person=", person)
	fmt.Println("")
	fmt.Println("------------ 指针方法")
	mi := MyInt{}
	fmt.Println("n0=", mi.n)
	mi.Increase()
	fmt.Println("n1=", mi.n)
	mi.Increase()
	fmt.Println("n2=", mi.n)
	mi.Decrease()
	fmt.Println("nx0=", mi.n)
	mi.Decrease()
	fmt.Println("nx1=", mi.n)
	mi.Increase()
	fmt.Println("n3=", mi.n)
	fmt.Printf("mi.n == 1: %v\n", mi.n == 1)
	fmt.Println("")
	fmt.Println("------------ select选择语句")
	ch3 := make(chan int, 1)
	// 省略若干条语句
	select {
	case ch3 <- 1:
		fmt.Printf("Sent %d\n", 1)
	case ch3 <- 2:
		fmt.Printf("Sent %d\n", 2)
	default:
		fmt.Println("Full channel")
	}
	fmt.Println("")
	ch4 := make(chan int, 1)
	for i := 0; i < 4; i++ {
		select {
		case e, ok := <-ch4:
			if !ok {
				fmt.Println("End")
				return
			}
			fmt.Println(e)
			close(ch4) // 关闭通道ch4
			//return
		default:
			fmt.Println("No Data")
			ch4 <- 100
		}
	}
}

func func107() {
	fmt.Println("------------ defer语句")
	fmt.Println("多个defer执行顺序是逆序的")
	multDefer()
	fmt.Println("")
	// 在panic语句后面的defer语句不被执行
	//panic("panic")
	//defer fmt.Println("defer after panic")
	finally()
	fmt.Println("")
	fmt.Println("------------ error异常")
	ret, ok := readFile("/app/tmp/test.jpg")
	fmt.Println("返回结果=", ret, ",error=", ok)
	fmt.Println("")
	fmt.Println("------------ panic运行时异常")
	panicMain()
	fmt.Println("")
	fmt.Println("------------ go语句")
	go fmt.Println("Go!")
	time.Sleep(100 * time.Millisecond)
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	go func() {
		fmt.Println("1")
		ch1 <- 1
	}()
	go func() {
		<-ch1
		fmt.Println("2")
		ch2 <- 2
	}()
	go func() {
		<-ch2
		fmt.Println("3")
		ch3 <- 3
	}()
	<-ch3
	fmt.Println("")
	// 循环
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("")
	exit := make(chan int, 1)
	fmt.Println("task start")
	go func(str string) {
		fmt.Println(str)
		close(exit) // 关闭通道，发出信号
	}("go task")
	<-exit // 关闭通道，立即解除阻塞
	fmt.Println("task stop")

}

func multDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func finally() {
	defer func() {
		fmt.Println("最后执行...")
	}()
	fmt.Println("先执行...")
}

func errInnerFunc() {
	fmt.Println("Enter innerFunc")
	panic(errors.New("Occur a panic!"))
	fmt.Println("Quit innerFunc")
}

func errOuterFunc() {
	fmt.Println("Enter outerFunc")
	errInnerFunc()
	fmt.Println("Quit outerFunc")
}

func panicMain() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Fatal error: %s\n", p)
		}
	}()
	fmt.Println("Enter main")
	errOuterFunc()
	fmt.Println("Quit main")
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

func strConv(args0 int, args1 int) string {
	rs := "args0=" + strconv.Itoa(args0) + ",args1=" + strconv.Itoa(args1) + ",return=str"
	return rs
}

// 匿名函数
var anonymousRs = func(part1 string, part2 string) string {
	return part1 + part2
}("abc,", "123")

type MyInt struct {
	n int
}

func (myInt *MyInt) Increase() {
	myInt.n++
}

func (myInt *MyInt) Decrease() {
	myInt.n--
	//fmt.Println("=",myInt.n)
}
