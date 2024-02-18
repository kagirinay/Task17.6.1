package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// RandInt Генерирует случайные числа
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	rand.Seed(time.Now().Unix())
	c1 := make(chan int)
	c2 := make(chan int)
	countOfc1 := 0.
	countOfc2 := 0.
	var wg sync.WaitGroup
	// Отправляет сообщения
	sender := func() {
		for {
			// С помощью этой переменной будем определять
			// в какой канал отправлять сообщение
			chance := RandInt(1, 100)
			// Если число в chance меньше либо равно 50, отправляем сообщение в канал c1,
			// если больше 50, то в канал c2
			if chance <= 50 {
				c1 <- RandInt(1, 100)
			} else {
				chance = RandInt(1, 100)
				if chance <= 50 {
					c2 <- RandInt(1, 100)
				} else {
					continue
				}
			}
		}
	}

	// Принимает сообщения
	receiving := func() {
		for {
			select {
			case num := <-c1:
				{
					countOfc1++
					//Вычисляет процентное соотношение полученных сообщений канала с1 по отношению к каналу с2
					var ration = (countOfc1 / (countOfc1 + countOfc2)) * 100
					str := fmt.Sprintf("Канал с1 принимает %.2f%% сообщений, сообщение из канала: %d", ration, num)
					fmt.Println(str)
				}
			case num := <-c2:
				{
					countOfc2++
					//Вычисляет процентное соотношение полученных сообщений канала с2 по отношению к каналу с1
					var ration = (countOfc2 / (countOfc1 + countOfc2)) * 100
					str := fmt.Sprintf("Канал с2 принимает %.2f%% сообщений, сообщение из канала: %d", ration, num)
					fmt.Println(str)
				}
			default:
				fmt.Println("Сообщений не поступило")
			}
		}
	}

	go sender()
	go receiving()
	wg.Add(2)
	wg.Wait()
}
