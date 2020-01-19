package main

import (
	"fmt"
)

// player 構造体
type Player struct{
	x, y int 
}

// 終了関数
func finish() int{
	return 1
}

// ゲーム選択
func gameSelect() int {
	var gameSelect int
	fmt.Println("you are welcome golang hoge game");
	fmt.Print("play : 1\nexit : 2\n>> ")
	fmt.Scanf("%d", &gameSelect);

	switch;
	gameSelect{
	case 1:
		return 1
	case 2:
		return 2
	default:
		return 3
	}
}

// RPGモード
func rpgMode() int{
	var name string
	var person = Player{x: 1, y: 1}

	// フィールドマップ
	areaMap := [9] int{0, 0, 0, 
					   0, 0, 0, 
					   0, 0, 0}
	
	fmt.Println("you are welcome to gread iland")
	name = recodeName()
	fmt.Println(name)
	areaMap[person.x * 3 + person.y] = 1
	for i := 0 ; i < 3 ; i++ {
		for j := 0 ; j < 3 ; j++ {
			if areaMap[j * 3 + i] == 0 {
				fmt.Print("#")
			} else if areaMap[j * 3 + i] == 1{
				fmt.Print("P")
			}
		}
		fmt.Println()
	}
	return 1
}

// 名前登録
func recodeName() string{
	var userName , yes string

	fmt.Print("you recode of something name\n>> ")
	fmt.Scan(&userName)

	for{
		fmt.Println("Is you name " + userName + " OK ? yes or no")
		fmt.Scan(&yes)
		if yes == "yes" || yes == "ye" || yes == "y" {
			return userName
		} else{
			fmt.Println("you make name")
		}
	}
}

func main(){
	var exit , startf int;
	exit = 0
	startf = 0
	for{
		startf = gameSelect()
		if startf == 1 {
			fmt.Println("game start")
			exit = rpgMode()
		}else if startf == 2{
			fmt.Println("exit")
			exit = finish()
		} else{
			fmt.Println("Error restart select")
		}

		if exit == 1 {break}
	}
	fmt.Println("game over")
}