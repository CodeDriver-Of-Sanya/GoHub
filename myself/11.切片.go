package main

import "fmt"

func main() {
	/*
		//切片英文slice slicesliceslice

		//切片的基本语法
		//var 变量名[]
		//var a [] int
		var intArr [5]int = [...]int{1, 2, 3, 4, 5}
		//声明/定义一个切片
		//slice 就是切片名
		//slice := intArr [1:3]表示第一到第四个元素
		//intArr[1:3] 表示引用intArr这个数组
		slice := intArr[1:3]
		//注意：引用数组的起始下标为 1，最后的下标为3但是不包含3
		fmt.Println("intArr", intArr)
		fmt.Println("slice 的元素是 =", slice)
		fmt.Println("slice 元素个数=", len(slice)) //长度
		fmt.Println("slice 容量是 =", cap(slice)) //容量目前可以存放的最大个数

		fmt.Println()
		fmt.Println()
		//两者地址相同
		fmt.Printf("intArr[1]的地址=%p\n", &intArr[1])
		fmt.Printf("slice[0]的地址=%p slice[0]=%v\n", &slice[0], slice[0])
	*/
	/*切片的使用
	  1 定义一个切片，然后切片去引用一个几经创建好的数组，比如前面的案例就是这样
	  2 第二种方式：通过make来创建切片，基本语法 var 切片名 []type = make([],len,[cap])
	    type 数据类型  len 大小  cap 指定切片容量，可选择
	*/
	/*
		//y演示切片的使用 make
		var slice []float64 = make([]float64, 5, 10)
		fmt.Println(slice)

		slice[1] = 10
		slice[3] = 99
		slice[4] = 88
		fmt.Println(slice)
		fmt.Println("slice的大小=", len(slice))
		fmt.Println("slice的容量=", cap(slice))
		fmt.Println("slice的元素是=", slice)
	*/
	/*
		//常规的for循环遍历切片
		var arr [5]int = [...]int{10, 20, 30, 40, 50}
		slice := arr[1:4]
		for i := 0; i < len(slice); i++ {

			fmt.Printf("slice[%v]=%v\n", i, slice[i])

		}

		//使用for--rang 方式遍历

		for i, v := range slice {
			fmt.Printf("i=%v v=%v\n", i, v)

		}

		//切片append内置函数，可以对切片进行动态追加。
		var slice3 []int = []int{100, 200, 300}
		//通过append追加元素,slice3本质没有变化
		slice3 = append(slice3, 400, 500, 900)
		fmt.Println("slice3", slice3)
		//通过append将切片slice3追加给slice3，自我追加
		slice3 = append(slice3, slice3...) //三点固定写法
		fmt.Println("slice3", slice3)
	*/
	//切片的拷贝操作，数组不可以拷贝

	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println("slice4", slice4)
	fmt.Println("slice5", slice5)

}
