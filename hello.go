package main

import (
	"bufio"
	"fmt"
	"github.com/ivahaev/russian-time"
	"log"
	"math/rand"
	"os"
	"rumod/gl"
	"time"
)

var name = "Роман"

func main() {
	rand.Seed(time.Now().UnixNano())
	file, err := os.Open("good.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	dobro := make([]string, 0)
	var i = 0
	for fileScanner.Scan() {
		dobro = append(dobro, fileScanner.Text())
		//	fmt.Println(dobro[i])
		i++
	}

	var max = len(dobro)
	//	max = 6
	var r1 = rand.Intn(max)
	var r2 = rand.Intn(max)
	var r3 = rand.Intn(max)
	var r4 = rand.Intn(max)

	// 	убираем дубль
	if r1 == r2 {
		r2 = r2 + 1
		if r2 >= max {
			r2 = 0
		}
	}

	for r3 == r2 || r3 == r1 {
		r3 = rand.Intn(max)
	}

	for r4 == r3 || r4 == r2 || r4 == r1 {
		r4 = rand.Intn(max)
	}

	var d1 = dobro[r1]
	var d2 = dobro[r2]
	var d3 = dobro[r3]
	var d4 = dobro[r4]

	var dop = rand.Intn(7)
	var mslov = 2
	if dop < 3 {
		mslov = 3
	}

	if dop == 0 {
		mslov = 4
	}

	var txt = d1 + " и " + d2
	if mslov == 3 {
		txt = d1 + " и " + d2 + " и " + d3
	}
	if mslov == 4 {
		txt = d1 + " и " + d2 + " и " + d3 + " и " + d4
	}

	fmt.Println("Доброе утро " + name + ". Вы самый " + txt + " человек во вселенной.")

	t := rtime.Now()
	// Or if you are using time.Time object: // 	standardTime := time.Now() 	t = rtime.Time(standardTime)

	fmt.Print("  Сегодня. ")
	fmt.Print(t.DayString() + " . ")
	fmt.Print(t.Month().StringInCase() + " . ")
	fmt.Println(t.Weekday().String() + " . ")
	//	fmt.Println()
	fmt.Println(" Задачи на сегодня.")
	fmt.Println(" изучить го...")
	fmt.Println(" Задачи на неделю...")
	fmt.Println(" Устроиться в. Яндекс .Алису.")

	fmt.Printf(gl.Read())
	//`date '+ %A, %B %d, %Y.'`
	//	currentTime := time.Now()
	//	fmt.Printf("%d %s %d ",
	//	currentTime.Day(),
	//	currentTime.Month(),
	//	currentTime.Year())
}
