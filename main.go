package main

import (
	"fmt"
)

// map 縦横
const MAP_height = 5
const MAP_width  = 5 


// player 構造体
type Player struct{
	x, y int 
}
// enemy 構造体
type Enemy struct{
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

// Player move
func pmove(move string, person Player) Player{
	if move == "w"{
		person.y -= 1
		if person.y < 0 {
			person.y = 0
		}
	}else if move == "s"{
		person.y += 1
		if person.y > MAP_height - 1 {
			person.y = MAP_height - 1
		}
	}else if move == "a" {
		person.x -= 1
		if person.x < 0 {
			person.x = 0
		}
	}else if move == "d" {
		person.x += 1
		if person.x > MAP_width - 1{
			person.x = MAP_width - 1
		}
	}else{
		person.x += 0
		person.y += 0
	}
	return person
}

// RPGモード
func rpgMode() int{
	var name , move string
	var person = Player{x: 0, y: 0}
	var enemy1 = Enemy{x:MAP_width-1, y: MAP_height-1}


	// フィールドマップ
	areaMap := [MAP_width][MAP_height]int{} 
	for i := 0 ; i < MAP_width ; i++{
		for j := 0 ; j < MAP_height ; j++{
			areaMap[i][j] = 0
		}
	}

	
	fmt.Println("you are welcome to gread iland")
	name = recodeName()
	fmt.Println(name)

	// マップ
	for{
		areaMap[person.x][person.y] = 1
		areaMap[enemy1.x][enemy1.y] = 2
		for i := 0 ; i < MAP_width ; i++ {
			for j := 0 ; j < MAP_height ; j++ {
				if areaMap[j][i] == 0 {
					fmt.Print(".")
				} else if areaMap[j][i] == 1{
					fmt.Print("P")
				} else if areaMap[j][i] == 2{
					fmt.Print("E")
				}
			}
			fmt.Println()
		}
		fmt.Scan(&move)
		areaMap[person.x][person.y] = 0
		person = pmove(move, person)
		areaMap[person.x][person.y] = 1

		// 移動
		// if move == "w"{
		// 	areaMap[person.x][person.y] = 0
		// 	person.y -= 1
		// 	if person.y < 0 {
		// 		person.y = 0
		// 		areaMap[person.x][person.y] = 1
		// 	}
		// }else if move == "s"{
		// 	areaMap[person.x][person.y] = 0
		// 	person.y += 1
		// 	if person.y > MAP_height - 1 {
		// 		person.y = MAP_height - 1
		// 		areaMap[person.x][person.y] = 1
		// 	}
		// }else if move == "a" {
		// 	areaMap[person.x][person.y] = 0
		// 	person.x -= 1
		// 	if person.x < 0 {
		// 		person.x = 0
		// 		areaMap[person.x][person.y] = 1
		// 	}
		// }else if move == "d" {
		// 	areaMap[person.x][person.y] = 0
		// 	person.x += 1
		// 	if person.x > MAP_width - 1{
		// 		person.x = MAP_width - 1
		// 		areaMap[person.x][person.y] = 1
		// 	}
		// }else{
		// 	person.x += 0
		// 	person.y += 0
		// }

	}
	fmt.Println(person)
	
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