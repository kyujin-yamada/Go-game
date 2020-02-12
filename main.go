package main

import (
	"fmt"
	"math/rand"
	"time"
)

// map 縦横
const MAP_height = 5
const MAP_width  = 5 


// player 構造体
type Player struct{
	x, y, attack, defence, speed, hp int 
	name string
}
var person = Player{x: 0, y: 0, attack:5, defence:3, hp:10}

// enemy 構造体
type Enemy struct{
	x, y, attack, defence, speed, hp, life int 
}

// randome 関数
func random(ran int) int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(ran)
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

// バトルターン
func BattleFaise(enemy *Enemy) int {
	var damage int
	var defenceTmp float32
	var move int

	for{
		fmt.Println("your turn. Do you select move")
		fmt.Print("fight = 1 or defence = 2\n >> ")
		fmt.Scan(&move)
		
		// プレイヤーターン
		if move == 1{
			damage = (person.attack - enemy.defence) * ((random(5) + 8) / 10)
		
			if damage < 0 {
				damage = random(3) 
			}
			enemy.hp -= damage
			
			fmt.Printf("%d damage\n", damage)
			fmt.Println()
			
			if enemy.hp < 0{
				fmt.Println("you win")
				enemy.hp = 0
				enemy.life = 0
				return 1
			}
		
		} else if move == 2{
			defenceTmp = float32(person.defence)
			person.defence = int(defenceTmp * 1.5)
		}
		
		// エネミーターン
		fmt.Println("enemy turn")

		damage = (enemy.attack - person.defence) * ((random(5) + 8) / 10) 
		if damage < 0 {
			damage = random(3)
		}
		person.hp -= damage
		if move == 2{
			person.defence = int(defenceTmp)
		}
		fmt.Printf("%d damage\n", damage)
		fmt.Println()
		if person.hp < 0{
			return 0
		}
	}

	return 0
}

// RPGモード main
func rpgMode() int{
	var move string
	var battle int 
	var enemy1 = Enemy{x:MAP_width-1, y: MAP_height-1, attack:10, defence:1, hp:3, life:1}
	var enemy2 [100]Enemy
	var count int

	// フィールドマップ
	areaMap := [MAP_width][MAP_height]int{} 
	for i := 0 ; i < MAP_width ; i++{
		for j := 0 ; j < MAP_height ; j++{
			areaMap[i][j] = 0
		}
	}

	
	fmt.Println("you are welcome to gread iland")
	person.name = recodeName()
	fmt.Println(person.name)

	// マップ
	for{
		areaMap[person.x][person.y] = 1
		if enemy1.hp != 0 {
			areaMap[enemy1.x][enemy1.y] = 2
		}

		if random(2) == 1 && count <= 5{
			fmt.Println("hello")
			//enemy2 = new(Enemy)
			enemy2[count].x = random(MAP_width)
			enemy2[count].y = random(MAP_height)
			enemy2[count].hp = random(20)
			enemy2[count].life = 1
			enemy2[count].attack = 10
			enemy2[count].defence = random(3)
			areaMap[enemy2[count].x][enemy2[count].y] = 2
			count++
		}
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

		// バトルフェイズ
	//	if person.x == enemy2.x && person.y == enemy2.y && enemy2.life == 1{
		if person.x == enemy2[].x && person.y == enemy2[].y && enemy2[].life
			fmt.Println("Change of battle faise -----------------------")
				battle = BattleFaise(&enemy2)
				if battle == 0{
					return finish()
				}else if enemy2.hp < 0{
					areaMap[enemy2.x][enemy2.y] = 0
				}
		}
		fmt.Printf("%s HP:%d\n",person.name, person.hp)

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