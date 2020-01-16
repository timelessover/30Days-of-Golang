package main

import "fmt"

// import "fmt"

// 类与实例
// 函数体外定义
// var arr = []int{1, 2, 3}

// type Person struct {
// 	name string
// 	age  int
// }

// type Phone interface {
// 	call()
// }

// type NokiaPhone struct {
// }

// func (nokiaPhone NokiaPhone) call() {
// 	fmt.Println("I am Nokia, I can call you!")
// }

// type IPhone struct {
// }

// func (iPhone IPhone) call() {
// 	fmt.Println("I am iPhone, I can call you!")
// }

func main() {
	var T int
	fmt.Scan(&T)
	fmt.Println(T)
	// println(arr)
	// 函数体内定义
	// // 字符串,变量的缩写
	// str  := "chris"
	// fmt.Println(str);

	// // 数字，同缩写
	// b,c  := 1,2
	// fmt.Println(b,c)

	// flag := true
	// fmt.Println(flag)

	// // 变量重新定义不用加冒号了
	// flag = false
	// fmt.Println(flag)

	// 常量,可用作枚举
	// const (
	// 	a = "abc"
	// 	b = len(a)
	// 	c = unsafe.Sizeof(a)
	// )
	// d := "ccc"
	// // if括号可以省略
	// if b>3 {
	// 	f = d
	// }

	// println(a, b, c,f)

	// 指针变量 * 和地址值 & 的区别：
	// 指针变量保存的是一个地址值，会分配独立的内存来
	// 存储一个整型数字。当变量前面有 * 标识时，才等同于 & 的用法，
	// 否则会直接输出一个整型数字。
	// var a int = 4
	// var ptr *int
	// ptr = &a
	// println("a的值为", a);    // 4
	// println("*ptr为", *ptr);  // 4
	// println("ptr为", ptr);    // 824633794744

	// for 循环

	// 	sum := 0
	// 	for i := 0; i <= 10; i++ {
	// 			sum += i
	// 			println(sum)
	// 	}

	// // foreach循环

	// list := []string{"google", "runoob"}

	// for key,value := range list {
	// 	println(key)
	// 	println(value)
	// }

	// string方法
	// scanner := bufio.NewScanner(os.Stdin)
	// println(strings.Contains("seafood", "foo"))
	// println(scanner)

	// 结构体 类似创建对象,需要在顶层定义

	// var person1 Person = createPerson("chris", 20)
	// sayPerson(person1)

	//数组方法, 关键字：slice,len,append,,make

	// var numbers []int
	// if numbers == nil {
	// 	fmt.Printf("切片是空的")
	// }
	// make创建动态数组
	// 切片是可索引的，并且可以由 len() 方法获取长度。

	// 切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少
	// sliceArr := make([]int, 2, 5)
	// fmt.Println(sliceArr)
	// printSlice(sliceArr)

	// 切片截取 numbers[:3]

	// numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// sliceNum1 := numbers[:5]
	// fmt.Println(sliceNum1) // [0,1,2,3,4]

	// sliceNum2 := numbers[5:]
	// fmt.Println(sliceNum2) // [5 6 7 8]

	// append() 向切片追加新元素,类似与push方法

	// var numbers []int

	// numbers = append(numbers, 2, 3)
	// fmt.Println(numbers)

	// /* 创建切片 numbers1 是之前切片的两倍容量*/
	// numbers1 := make([]int, len(numbers), (cap(numbers))*3)
	// fmt.Println(numbers1)
	// /* 拷贝 numbers 的内容到 numbers1 */
	// copy(numbers1, numbers)
	// printSlice(numbers1)

	// //range也可以用在map的键值对上。
	// kvs := map[string]string{"a": "apple", "b": "banana"}
	// for k, v := range kvs {
	// 	fmt.Printf("%s -> %s\n", k, v)
	// }
	// //range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	// for i, c := range "go" {
	// 	fmt.Println(i, c)
	// }

	// map =》 {k:v}
	// Map 是一种无序的键值对的集合。
	// Map 最重要的一点是通过 key 来快速检索数据,
	// key 类似于索引，指向数据的值。

	// var testMap map[string]string

	// testMap = make(map[string]string)

	// testMap["name"] = "chris"
	// testMap["age"] = "22"

	// fmt.Println(testMap)

	// 递归，
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(fibonacci(i))
	// }

	// 接口和实例

	// var phone Phone

	// // interface 拥有 new 方法,在接口中定义的方法，在对应的结构体中实现
	// phone = new(NokiaPhone)
	// phone.call()

	// phone = new(IPhone)
	// phone.call()

	// 简单练习：冒泡排序
	// arr := []int{2, 5, 3, 1, 6, 0}
	// lens := len(arr)
	// for i := 0; i < lens; i++ {
	// 	for j := 1; j < lens-i; j++ {
	// 		if arr[j] < arr[j-1] {
	// 			temp := arr[j]
	// 			arr[j] = arr[j-1]
	// 			arr[j-1] = temp
	// 		}
	// 	}
	// }
	// fmt.Println(arr)

}

// func fibonacci(n int) int {
// 	if n < 2 {
// 		return n
// 	}
// 	return fibonacci(n-2) + fibonacci(n-1)
// }

// func printSlice(x []int) {
// 	// len() 为数组长度，cap() 可以测量切片最长可以达到多少
// 	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
// }

// func createPerson(name string, age int) Person {
// 	var person Person
// 	person.name = name
// 	person.age = age

// 	return person
// }

// func sayPerson(person Person) {
// 	println("person.name : %s\n", person.name)
// 	println("person.age：%s\n", person.age)
// }
