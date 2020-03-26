// 必须加上package
package second

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/hashicorp/golang-lru"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

/**
 * go入门-应用
 *
 * @author gavin.z, swe.zzf@gmail.com
 * @since  2020-02-23
 */
func init() {
	fmt.Println("second-init()")
}

// 公用变量
var UserAge int32 = 20

// 私有变量
var userAge int32 = 30

// 公用函数定义
func Main() {
	fmt.Println("second公用函数Main()")
	//func201()
	//func202()
	//func203()
	func204()

}

type Color interface {
	Type() string
}

type Blue struct {
	This string
}

type Red struct {
	This string
}

func varArgs(args ...string) string {
	for k, v := range args {
		fmt.Println("k=", k, ",v=", v)
	}
	return ""
}

// 实现Color接口的方法Type()
func (b *Blue) Type() string {
	return b.This + " type"
}

// 实现Color接口的方法Type()
func (r *Red) Type() string {
	return r.This + " type"
}

func Action(c Color) string {
	return c.Type()
}

type Student struct {
	Person      // 继承Person
	grade int64 // 年级
}

func (s *Student) getGrade() int64 {
	return s.grade
}

type Person struct {
	Id   int64  `xml:"id,attr"` // 设置为id属性
	Name string `xml:"name"`
	Age  int64  `json:"年龄"` // Age命名为'年龄'
}

func (p Person) String() string {
	p.Age = p.Age + 1
	return fmt.Sprintf("[ id:%v, name:%v, age:%v ]", p.Id, p.Name, p.Age)
}

func readFrom(reader io.Reader, num int64) ([]byte, error) {
	p := make([]byte, num)
	// 调用Read读取IO
	n, err := reader.Read(p)
	if n > 0 {
		// 返回部分数据和无异常信息
		return p[:n], nil
	}
	// 全部返回
	return p, err
}

func writeFrom(writer io.Writer, num int64) ([]byte, error) {
	p := make([]byte, num)
	// 调用Read读取IO
	n, err := writer.Write(p)
	if n > 0 {
		// 返回部分数据和无异常信息
		return p[:n], nil
	}
	// 全部返回
	return p, err
}

func countFileLine(filename string) {
	fi, err := os.Open(filename)
	fmt.Println("fi=", fi, ",err=", err)
	defer func() {
		if xErr := fi.Close(); xErr != nil {
			panic(xErr)
		}
	}()
	if err != nil {
		panic(err)
	}
	// count line
	reader := bufio.NewReader(fi)
	line := 0
	for {
		data, isPrefix, err := reader.ReadLine()
		//fmt.Println("line=", line, ",isPrefix=", isPrefix, ",error=", err)
		if err != nil {
			fmt.Println("error=", err)
			break
		}
		// _len > 0 去掉空行
		_len := len(string(data))
		if !isPrefix && _len > 0 {
			line++
		}
	}
	fmt.Println("文件名:", filename, " 去掉空行行数:", line)
}

type ImgHeader struct {
	Size           uint32
	Width          int32
	Height         int32
	Places         uint16
	BigCount       uint16
	Compression    uint32
	SizeImage      uint32
	XperlsPerMeter int32
	YperlsPerMeter int32
	ClsrUsed       uint32
	ClrImportant   uint32
}

// 重写String函数
func (ih ImgHeader) String() string {
	return fmt.Sprintf("[Size:%v, Width:%v, Height:%v, Places:%v, \nBigCount:%v, Compression:%v,"+
		" SizeImage:%v, XperlsPerMeter:%v, \nYperlsPerMeter:%v, ClsrUsed:%v, ClrImportant:%v ]",
		ih.Size, ih.Width, ih.Height, ih.Places,
		ih.BigCount, ih.Compression, ih.SizeImage, ih.XperlsPerMeter,
		ih.YperlsPerMeter, ih.ClsrUsed, ih.ClrImportant)
}

func (p *Person) PersonM1() *Person {
	return p
}

func (p *Person) PersonM2() string {
	return p.Name
}

var MsgCh = make(chan int, 10)
var TimeoutCh = make(chan bool)

func SendMsg() {
	TimeoutCh <- false
	for i := 1; i <= 5; i++ {
		MsgCh <- i
	}
	TimeoutCh <- true
}

func ReceiveMsg() {
	for {
		select {
		case num := <-MsgCh:
			fmt.Println("receive-num=", num)
		case <-TimeoutCh:
			fmt.Println("receive-timeout...")
		}
	}
}

func Loop() {
	for i := 1; i <= 10; i++ {
		fmt.Println("loop-i=", i)
		time.Sleep(time.Microsecond * 2)
	}
}

var RwSync sync.WaitGroup

func readData(flag string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second * 1)
		RwSync.Done()
		fmt.Println(flag, ":read data done,i=", i)
	}
}

func writeData() {
	for i := 1; i <= 10; i++ {
		// 添加计数器个数
		RwSync.Add(1)
	}
	fmt.Println("write data done.")
}

func func201() {
	fmt.Println("内部函数func201()")
	fmt.Println("--------------- IO操作fmt模块")
	// %v自动匹配
	a, b := 10, "中文"
	fmt.Printf("百分v自动匹配：%v - %v \n", a, b)
	// 格式化输出
	str := fmt.Sprintf("float %f", 3.14158977897978)
	fmt.Println("Sprintf格式化:", str)
	// 文件流输出
	fmt.Fprintln(os.Stdout, "文件流输出到控制台")
	// 输出对象
	fmt.Println("输出对象=", Person{1, "张三", 30})
	//
	x := make([]string, 2)
	x[0] = "a"
	x[1] = "b"
	fmt.Println(x)
	fmt.Println("")
	fmt.Println("--------------- IO操作Reader模块")
	// 控制台输入方式
	fmt.Println("请输入字符串：")
	//rs, err := readFrom(os.Stdin, 5)
	//fmt.Println("Stdin控制台方式5字符串:", string(rs), ",error:", err)
	// 字符串方式输出
	rs2, err2 := readFrom(strings.NewReader("123456789"), 5)
	if err2 != nil {
		panic(err2) // 抛出异常
	}
	fmt.Println("字符串方式5个=", string(rs2), ",error2=", err2)
	// 文件方式输出
	outFile, err := os.Open("/app/tmp/test.txt")
	fmt.Println("outFile=", outFile, ",error=", err)
	defer outFile.Close() // 关闭文件流
	if err != nil {
		panic(err) // 抛出异常
	}
	rs3, err3 := readFrom(outFile, 3)
	if err3 != nil {
		panic(err3) // 抛出异常
	}
	fmt.Println("文件方式3个=", string(rs3), ",error3=", err3)
	bufReader := bufio.NewReader(strings.NewReader("123456789"))
	// Peek只读方式提取n个字节，不改变缓存大小。
	rs4, err4 := bufReader.Peek(3)
	fmt.Println("缓存流error4=", err4)
	if err4 != nil {
		panic(err4) // 抛出异常
	}
	fmt.Println("缓存流3个=", string(rs4), ",error4=", err4, "Buffered=", bufReader.Buffered())
	fmt.Println("")
	fmt.Println("--------------- IO操作Writer模块")
	// ioutil方式
	// read the whole file at once
	data, err := ioutil.ReadFile("/app/tmp/test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("ioutil读文件=", string(data))
	newData := string(data) + "111\n"
	// write the whole body at once
	err = ioutil.WriteFile("/app/tmp/test.txt", []byte(newData), os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println("ioutil写文件err=", err)
	// 计算文件行数
	//args := os.Args
	//if len(args) != 2 {
	//	fmt.Println("控制台中输入格式有误！")
	//	return
	//}
	//countFileLine(args[1])
	dir, _ := os.Getwd()
	countFileLine(dir + "/src/learn/learn.go")
	fmt.Println("")
	//
	fmt.Println("--------------- IO操作binary模块")
	// read
	fi98, err98 := os.Open("/app/tmp/test-pic.jpg")
	defer fi98.Close()
	if err98 != nil {
		fmt.Println("err98=", err98)
		return
	}
	imgHeader := new(ImgHeader)
	err88 := binary.Read(fi98, binary.LittleEndian, imgHeader)
	fmt.Println("binary.read=", imgHeader, ",\nerr88=", err88)
	// write
	buf := new(bytes.Buffer)
	var pi float64 = math.Pi
	err99 := binary.Write(buf, binary.LittleEndian, pi)
	if err2 != nil {
		fmt.Println("binary.write failed=", err99)
	}
	// 18 2d 44 54 fb 21 09 40
	fmt.Println("binary.write=", buf.Bytes())

}

func func202() {
	// 判断前缀、合并、包含、连接、分割、取索引、消除空格
	fmt.Println("")
	fmt.Println("--------------- 文本编码处理strings模块")
	str := "hello world"
	// 包含字符串
	fmt.Println("包含Contains:", strings.Contains(str, "w"))

	// 分割字符串
	str2 := "a$b$c$d"
	fmt.Println("分割Split:", strings.Split(str2, "$"))

	// 获取索引
	fmt.Println("索引Index:", strings.Index(str, "w"))

	// 合并字符串
	str3 := strings.Split(str2, "$")
	fmt.Println("合并Join:", strings.Join(str3, "2"))

	// 判断前缀和后缀
	fmt.Println("前缀和后缀HasPrefix:", strings.HasPrefix(str, "w"))

	fmt.Println("")
	fmt.Println("--------------- 文本编码处理strconv模块")
	// string类型转换
	fmt.Println("strconv.Itoa=", strconv.Itoa(10))
	x, _ := strconv.Atoi("10")
	fmt.Println("strconv.Atoi=", x)

	// 类型解析
	x2, _ := strconv.ParseBool("false")
	fmt.Println("strconv.ParseBool=", x2)

	// 类型格式化
	fmt.Println("strconv.FormatInt=", strconv.FormatInt(int64(99), 10))
	fmt.Println("")
	fmt.Println("--------------- 文本编码处理struct和xml转换")
	var data []byte
	var err error
	person := Person{Id: 1, Name: "张三", Age: 30}
	if data, err = xml.MarshalIndent(person, "", " "); err != nil {
		fmt.Println("to xml,error=", err)
		return
	}
	fmt.Println("对象转换xml:\n", string(data))
	//
	person2 := new(Person)
	if err2 := xml.Unmarshal(data, person2); err2 != nil {
		fmt.Println("to object,error=", err2)
		return
	}
	fmt.Println("xml转换对象:", person2)
	//
	fmt.Println("")
	fmt.Println("--------------- 文本编码处理os.Args命令行解析")
	// xxx.exe a b c  返回string[a,b,c]
	//args := os.Args
	//fmt.Println("命令行参数:", args)
	// xxx.exe -m get
	//strFg := flag.String("-m", "默认值", "备注")
	//i := flag.Int("-m", -1, "备注")
	//flag.Parse() // 解析
	//fmt.Println("flag-str=", *strFg, ",int=", *i)
	fmt.Println("")
	fmt.Println("--------------- 文本编码处理获取xml文件信息")
	content, err := ioutil.ReadFile("/app/tmp/person.xml")
	if err != nil {
		fmt.Println("error=", err)
		return
	}
	var t xml.Token
	decoder := xml.NewDecoder(strings.NewReader(string(content)))
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch v := t.(type) {
		case xml.StartElement:
			fmt.Println("name=", v.Name.Local)
		}
	}
	fmt.Println("")
	fmt.Println("--------------- 文本编码处理模拟命令行cp")
	//io.Copy("writer", "reader")

}

func func203() {
	fmt.Println("")
	fmt.Println("--------------- 内建函数")
	// make返回引用的类型
	strs := make([]string, 4)
	strs[0] = "a"
	strs[1] = "b"
	strs[2] = "c"
	strs[3] = "d"
	fmt.Println("make=", strs)
	// new返回输入类型的指针地址
	strs2 := new([]string)
	fmt.Println("new-TypeOf=", reflect.TypeOf(strs2))
	// append
	strs = append(strs, "e")
	fmt.Println("append-len长度=", len(strs))
	fmt.Println("append-cap容量=", cap(strs))
	// copy
	dest := make([]string, 6)
	dest[5] = "f"
	copy(dest, strs)
	//dest[0] = "1"
	fmt.Println("copy=", dest)
	// delete
	delMap := make(map[int]string)
	delMap[0] = "a"
	delMap[1] = "b"
	delMap[2] = "c"
	fmt.Println("delete11=", delMap)
	delete(delMap, 0)
	fmt.Println("delete22=", delMap)
	// close
	ch := make(chan int, 2)
	ch <- 1
	defer close(ch) // 程序执行完成后（defer）再关闭通道
	ch <- 2
	fmt.Println("关闭通道:", ch)
	// 异常处理：抛出panic 捕获recover
	// 捕获异常recover
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("error=", e)
		}
	}()
	// 抛出异常panic
	//panic("this is panic")
	//
	fmt.Println("")
	fmt.Println("--------------- struct结构体-面向对象")
	// 初始化方式
	p := Person{Name: "名称"}
	fmt.Println("初始化方式1", p)
	p2 := new(Person)
	p2.Name = "名称2"
	fmt.Println("初始化方式2", p2)
	// 结构体方法
	p3 := new(Person)
	p3.Name = "名称3"
	fmt.Println("结构体方法1:", p3.PersonM1())
	fmt.Println("结构体方法2:", p3.PersonM2())
	// 结构体继承
	st := new(Student)
	st.grade = 2
	st.Name = "名称4"
	fmt.Println("结构体继承:", st.getGrade(), ",name=", st.Name)
	fmt.Println("")
	fmt.Println("--------------- interface接口")
	// 接口的实现和多态
	blue := new(Blue)
	blue.This = "blue"
	fmt.Println("接口的多态1:", Action(blue))
	red := new(Red)
	red.This = "red"
	fmt.Println("接口的多态2:", Action(red))
	//
	fmt.Println("")
	fmt.Println("--------------- 并发操作")
	// 使用方法
	//fmt.Println("cpu num=", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU() - 1) // 设置协程使用的最大cpu数
	//go Loop()
	//go Loop()
	// 协程通信
	fmt.Println("")
	fmt.Println("----协程通信")
	//go SendMsg()
	//go SendMsg()
	//go ReceiveMsg()
	//go ReceiveMsg()
	//time.Sleep(time.Second * 3)
	// 协程同步sync.WaitGroup：添加记录Add 移除记录Done 等待协程Wait
	writeData()       // 主线程同步写
	go readData("t1") // 子协程异步读
	go readData("t2") // 子协程异步读
	RwSync.Wait()
	fmt.Println("协程同步end！！！")

	//
	fmt.Println("")
	fmt.Println("--------------- 指针操作")
	var count int64 = 70
	var countPt *int64
	var countPt2 *int64 // 未分配内存
	// go不支持指针类型变量运算，如：countPt2++
	//countPt2++
	countPt = &count
	fmt.Println("count变量值:", count)
	fmt.Println("count变量地址:", &count)
	fmt.Println("countPt指针地址:", countPt)
	fmt.Println("countPt指针地址的值:", *countPt)
	fmt.Println("countPt2指针地址:", countPt2)
	// 指针数组
	a, b, c := 1, 2, 3
	pointArr := []*int{&a, &b, &c}
	fmt.Println("指针数组:", pointArr)
	// 数组指针
	arr := []int{1, 2, 3}
	arrPoint := &arr
	fmt.Println("数组指针:", arrPoint)

	//
	fmt.Println("")
	fmt.Println("--------------- JSON序列和反序列化和Tag")
	// 序列化
	var jdata []byte
	var err error
	if jdata, err = json.MarshalIndent(Person{Id: 2, Name: "李四", Age: 40}, "", " "); err != nil {
		fmt.Println("json-error=", err)
	}
	fmt.Println("JSON序列化:\n", string(jdata))
	// 反序列化
	jp := new(Person)
	if err := json.Unmarshal(jdata, jp); err != nil {
		fmt.Println("json-error2=", err)
	}
	fmt.Println("JSON反序列化:", jp)

	//
	fmt.Println("")
	fmt.Println("--------------- 可变参数")
	fmt.Println("可变参数:")
	// 可变参数
	varArgs("a", "b", "c")

	//
	fmt.Println("")
	fmt.Println("--------------- module包依赖管理")
	kv, _ := lru.New(128)
	for i := 0; i < 256; i++ {
		kv.Add(i, math.Round(100))
	}
	if kv.Len() != 128 {
		panic(fmt.Sprintf("bad len: %v", kv.Len()))
	}
	fmt.Println("len=", kv.Len())
	log.Println("日志输出。。。")

}

func func204() {
	fmt.Println("")
	fmt.Println("--------------- 函数参数类型-引用类型")
	a := 3
	b := 4
	c := "abc"
	d := [3]int{1, 2, 3}
	fmt.Printf("func204函数1：a=%v, b=%v, c=%v, d=%v \n", a, b, c, d)
	changeArgs(a, b, c, d)
	fmt.Printf("func204函数2：a=%v, b=%v, c=%v, d=%v \n", a, b, c, d)
	fmt.Println("")
	fmt.Println("--------------- 函数参数类型-指针类型")
	a2 := 3
	b2 := 4
	p1 := &a2 //获取变量a的内存地址，并将其赋值给变量p1
	p2 := &b2 //获取变量b的内存地址，并将其赋值给变量p2
	fmt.Printf("func204函数1：a2的值=%v, a2的指针=%v ，p1指向的变量的值=%v\n", a2, p1, *p1)
	fmt.Printf("func204函数1：b2的值=%v, b2的指针=%v ，p2指向的变量的值=%v\n", b2, p2, *p2)
	fmt.Println("changeArgs2函数:", changeArgs2(p1, p2))
	fmt.Printf("func204函数2：a2的值=%v, a2的指针=%v ，p1指向的变量的值=%v\n", a2, p1, *p1)
	fmt.Printf("func204函数2：b2的值=%v, b2的指针=%v ，p2指向的变量的值=%v\n", b2, p2, *p2)

}

func changeArgs(a, b int, c string, d [3]int) {
	a = 5
	b = 6
	c = "efg"
	d[0] = 0
	fmt.Printf("changeArgs函数：a=%v, b=%v, c=%v, d=%v \n", a, b, c, d)
}

func changeArgs2(a, b *int) int {
	*a = 5
	*b = 6
	//这里出现连续的两个*，Go编译器会根据上下文自动识别乘法与两个引用
	return *a * *b
}
