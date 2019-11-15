package main

import "fmt"

func modifyArray1(a [2]int) {
	a[0] = 100
}

func modifyArray2(a [3][2]int) {
	a[0][0] = 100
}

//求数组[1, 3, 5, 7, 8]所有元素的和
func countArray() {

	numArray := []int{1, 3, 5, 7, 8}
	total := 0
	for _, value := range numArray {
		total += value
	}
	fmt.Printf("数组的和 = %d\n", total)

}

//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
func Test1(numArray []int) {
	for i := 0; i < len(numArray); i++ {
		for j := i + 1; j < len(numArray); j++ {
			if numArray[i]+numArray[j] == 8 {
				fmt.Printf("(%d,%d)", i, j)
			}
		}
	}
}

func main() {

	testArray := [3]int{}
	numArray := [...]int{1, 2, 3}
	cityArray := [...]string{"广州", "上海", "北京"}
	fmt.Printf("%v,%T\n", testArray, testArray)
	fmt.Printf("%v,%T\n", numArray, numArray)
	fmt.Printf("%v,%T\n", cityArray, cityArray)

	a := [3][2]string{
		{"广州", "深圳"},
		{"长沙", "株洲"},
		{"海口", "三亚"},
	}

	fmt.Printf("%v,%T\n", a, a)

	numArray2 := [2]int{1, 2}
	modifyArray1(numArray2)
	fmt.Printf("%v,%T\n", numArray2, numArray2)

	numArray3 := [3][2]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	modifyArray2(numArray3)
	fmt.Printf("%v,%T\n", numArray3, numArray3)

	countArray()

	numArray4 := []int{0, 1, 3, 5, 7, 8}
	Test1(numArray4)

}
