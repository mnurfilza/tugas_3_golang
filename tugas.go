package tugas3

import "fmt"

func Tugas3(arr []string, aditional string) {
	arr = append(arr, aditional)
	fmt.Println("Jumlah Element:", len(arr))
	fmt.Println("isi Element :", arr)
	for index, item := range arr {
		fmt.Println("Element ke - :", index, item)
	}
}
