package main

import (
	"fmt"
	"math"
)

type StockSpanner struct {
	stack [][2]int
	idx   int
}

func Constructor() StockSpanner {
	//初始化一个值，防止溢出
	return StockSpanner{[][2]int{{math.MaxInt32, -1}}, -1}
}

func (this *StockSpanner) Next(price int) int {
	this.idx++
	for price > this.stack[len(this.stack)-1][0] {
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, [2]int{price, this.idx})
	return this.idx - this.stack[len(this.stack)-2][1]
}
func main() {
	stockspanner := Constructor()
	fmt.Println(stockspanner.Next(28))
	fmt.Println(stockspanner.Next(14))
	fmt.Println(stockspanner.Next(28))
	fmt.Println(stockspanner.Next(35))
	fmt.Println(stockspanner.Next(46))
	fmt.Println(stockspanner.Next(53))
	fmt.Println(stockspanner.Next(66))
}
