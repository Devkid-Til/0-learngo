package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sumTo100() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile() {
	file, err := os.Open("abc.txt")
	if err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}

}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Printf("%q\n",
		convertToBin(0),
	)
	sumTo100()
	printFile()
	forever()
}
