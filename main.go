package main

import (
	"fmt"
	"learn/helper"
	"math/rand/v2"
	"strconv"
)

var (
	choice int
	input  string
	lines  []string
	dog    = `
	 	                  \
		                   \
			      	     \
					 ⣠⣴⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣮⣵⣄⠀⠀⠀ ⠀⠀
					⢾⣻⣿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⡀⠀ ⠀
					⠸⣽⣻⠃⣿⡿⠋⣉⠛⣿⣿⣿⣿⣿⣿⣿⣿⣏⡟⠉⡉⢻⣿⡌⣿⣳⡥⠀ ⠀
					⢜⣳⡟⢸⣿⣷⣄⣠⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣤⣠⣼⣿⣇⢸⢧⢣⠀ ⠀
					⠨⢳⠇⣸⣿⣿⢿⣿⣿⣿⣿⡿⠿⠿⠿⢿⣿⣿⣿⣿⣿⣿⣿⣿⠀⡟⢆⠀ ⠀⠀
					 ⠈⠀⣾⣿⣿⣼⣿⣿⣿⣿⡀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣽⣿⣿⠐⠈⠀⠀ ⠀
					 ⢀⣀⣼⣷⣭⣛⣯⡝⠿⢿⣛⣋⣤⣤⣀⣉⣛⣻⡿⢟⣵⣟⣯⣶⣿⣄⡀⠀
					⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣶⣶⣶⣾⣶⣶⣴⣾⣿⣿⣿⣿⣿⣿⢿⣿⣿⣧
					 ⣿⣿⣿⠿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⣿⡿
	`
)

func main() {
	path := "comTrangCaoBoCollection.txt"
	fmt.Print("Nhập số bài thơ bạn muốn xem (1-8) hoặc nhập 'r' để random: ")
	fmt.Scan(&input)

	if input == "r" {
		choice = rand.IntN(8) + 1

	} else {
		c, _ := strconv.Atoi(input)
		choice = c
	}
	lines = helper.ReadFile(path, choice)

	maxwidth := helper.Caculate_max_width(lines)
	messages := helper.NormalizeStringsLength(lines, maxwidth)
	balloon := helper.Build_ballon(messages, maxwidth)

	helper.PrintRPG(balloon)
	fmt.Println(dog)
	fmt.Println()

}
