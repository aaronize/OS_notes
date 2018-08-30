package main

import "fmt"

func sum(values []int, resultChan chan int) {
    sum := 0
    for _, value := range values {
        sum += value
    }
    // 将计算结果发送到channel中
    resultChan <- sum
}
func main() {
    values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    resultChan := make(chan int, 3)
    go sum(values[:len(values)/2], resultChan) //15
    go sum(values[len(values)/2:], resultChan) //45
    go sum(values[len(values)/3:], resultChan) //
    sum1, sum2, sum3 := <-resultChan, <-resultChan, <-resultChan
    fmt.Println("Result:", sum1, sum2, sum3)
}
