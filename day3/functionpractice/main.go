package main

import (
	"fmt"
)

// 要求
// 50硬币分给users切片里的人
// 名字里每有一个'e'或'E'分一个硬币
// 'i','I'两个
// 'o','O'三个
// 'u','U'四个
// 每个人获得几个硬币,返回值为剩下的硬币数,实现函数 dispatchCoin

var ( // 硬币,存放人名的切片,人名和硬币数对应的map放在全局变量里
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func nameParse(name string) (coins int) { // 求这个名字值几个硬币
	for _, value := range name {
		switch {
		case value == 'e' || value == 'E':
			coins++
		case value == 'i' || value == 'I':
			coins += 2
		case value == 'o' || value == 'O':
			coins += 3
		case value == 'u' || value == 'U':
			coins += 4
		}
	}
	return
}

func dispatchCoin() int { // 给每个人分硬币，分完后把剩下的硬币数量返回出去
	var currentUserCoins int
	var nowCoins = coins

	for _, user := range users {
		currentUserCoins = nameParse(user)
		count, ok := distribution[user]
		if !ok {
			distribution[user] = currentUserCoins
		}
		distribution[user] = count + currentUserCoins
		nowCoins -= currentUserCoins
	}

	return nowCoins
}

func main() {
	var left = dispatchCoin()
	fmt.Println(distribution)
	fmt.Println("剩下: ", left)
}
