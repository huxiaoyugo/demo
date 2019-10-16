package main

import "fmt"

func BubbleSort(s []int) {
	for i:=1; i< len(s); i++ {
		for j:=0; j < len(s) - i; j++ {
			if s[j] > s[j+1] {
				s[j],s[j+1] = s[j+1],s[j]
			}
		}
	}
}



// 选择排序
func SelectSort(s []int) {

	for i:=0; i < len(s) -1; i++ {
		minIndex := i
		for j := i; j< len(s); j++ {
			if s[j] < s[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			s[minIndex],s[i] = s[i],s[minIndex]
		}
	}
}


// 基数排序
func NumSort(s []int) {

	maxBit := maxBit(s)

	//fmt.Println(maxBit)
	t := 1
	for i:=1; i<= maxBit; i++ {
		b := [10][]int{}
		for _, num := range s {
			c := (num / t) % 10
			b[c] = append(b[c], num)
		}
		s = s[0:0]
		for i:=0; i<10; i++ {
			if len(b[i]) > 0 {
				s =append(s, b[i]...)
			}
		}
		t *= 10
	}
	return
}

func maxBit(s []int) int {
	maxBit := 1
	maxT := 10

	for _, item := range s {
		x := maxT
		for i:=maxBit;;i++{
			if item < x {
				maxBit = i
				maxT = x
				break
			}
			x = x*10
		}
	}
	return maxBit
}


// 归并排序

func MergeSort(s []int) {

	if len(s) == 1 {
		return
	}
	mid := len(s) / 2
	MergeSort(s[:mid])
	MergeSort(s[mid:])
	// 合并

	i:=0
	j:=mid

	t := make([]int, 0, len(s))
	for;i<mid&&j<len(s); {
		if s[i] < s[j] {
			t = append(t, s[i])
			i++
		} else {
			t = append(t, s[j])
			j++
		}
	}

	if i<mid {
		t = append(t, s[i:mid]...)
	}

	if j<len(s) {
		t = append(t, s[j:]...)
	}

	for i:=0; i< len(t);i++ {
		s[i] =  t[i]
	}
}


// 快速排序
func QuickSort(s []int) {

	if len(s) <= 1 {
		return
	}

	t := s[0]

	l := 0
	r := len(s) -1
	p := 0
	o: for {

		for {
			if l == r {
				p = l
				break o
			}
			if s[r] < t {
				s[l],s[r] = s[r],s[l]
				break
			}
			r--
		}

		for {
			if l== r {
				p=l
				break o
			}
			if s[l] > t {
				s[l],s[r] = s[r],s[l]
				break
			}
			l++
		}
	}

	s[p] = t

	QuickSort(s[0:p])
	if p+1 < len(s)-1 {
		QuickSort(s[p+1:])
	}
}

func main() {
	s := []int{-1,0,11,7,6,5,4,3,2,1,0}
	HeapSort(s)
	fmt.Println(s)
}

// 堆排序
func HeapSort(s []int) {

	for i:= 0; i < len(s)-1; i++ {
		buildHeap(s[0:len(s)-i])
		lastIndex := len(s) - i - 1
		s[0],s[lastIndex] = s[lastIndex],s[0]
	}

}


func buildHeap(s []int) {

	l := len(s)-1
	for i:= (l-1) / 2; i>=0; i-- {
		//左孩子
		left := 2 * i + 1
		lv := s[left]
		if lv > s[i] {
			s[left],s[i] = s[i], s[left]
		}
		// 右孩子
		right := left+1
		if right < len(s) && s[right] > s[i]{
			s[right],s[i] = s[i], s[right]
		}
	}
}