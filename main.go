package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cycle = 2
var maxBatting = 0 // 最大バッティング金額
var minBenefit = 0 // 最低利益
var benefit = 0    // 利益

func game(m int) bool {
	rand.Seed(time.Now().UnixNano()) // go playgroundではrand数が固定される問題発生
	time.Sleep(1 * time.Second)

	// win
	if rand.Intn(100) > 50 {
		fmt.Println("ゲーム結果：Win")
		return true
	}

	// loose
	fmt.Println("ゲーム結果：Loose")
	return false
}

func delete(s []int, i int) []int {
	s = append(s[:i], s[i+1:]...)
	n := make([]int, len(s))
	copy(n, s)
	return n
}

func main() {
	for i := 0; i < cycle; i++ {
		fmt.Printf("=============サイクル%v回目=============\n", i+1)

		monteSlice := []int{1, 2, 3}

		for len(monteSlice) > 0 {
			fmt.Printf("配列：%v\n", monteSlice)
			fmt.Printf("現在残金：%v\n", benefit+4)

			money := monteSlice[0] + monteSlice[len(monteSlice)-1]
			fmt.Printf("バッティング：%v\n", money)

			if money > maxBatting {
				maxBatting = money
			}

			if game(money) == false { // loose
				monteSlice = append(monteSlice, money)
				benefit -= money

				if benefit < minBenefit {
					minBenefit = benefit
				}

			} else { // win
				monteSlice = delete(monteSlice, len(monteSlice)-1)
				benefit -= money     // バッティング金を引く
				benefit += money * 2 //勝ったら2倍の金をもらう

				if len(monteSlice) <= 0 {
					break
				}

				monteSlice = delete(monteSlice, 0)
			}
		}
		
		fmt.Println("Sliceが空になったので終了します。")
		fmt.Println("====================================")
		fmt.Printf("============サイクル%v回目結果===========\n", i+1)
		fmt.Printf("最大バッティング金額：%v\n", maxBatting)
		fmt.Printf("最低利得：%v\n", minBenefit)
		fmt.Printf("最終利得：%v\n", benefit)
		fmt.Printf("残金：%v\n", benefit+4)
		fmt.Println("====================================")
	}
	
	fmt.Println("Game Over")
}
