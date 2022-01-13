package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

const (
	clearScreen = "\033[H\033[2J"
)

type Player struct {
	startingHand []int
	keptDice     []int
	Score        int
	diceLeft     int
}

func main() {

	showRules()

	player := newPlayer()

	playHand(player)

}

func showRules() {

	fmt.Print(clearScreen)
	fmt.Println("Welcome to Farkle! \n")

	t := table.NewWriter()
	t.AppendHeader(table.Row{"Playable Set", "Point Value"})
	t.AppendRow(table.Row{"1 Five", 50})
	t.AppendRow(table.Row{"1 One", 100})
	t.AppendRow(table.Row{"3 Ones", 300})
	t.AppendRow(table.Row{"3 Twos", 200})
	t.AppendRow(table.Row{"3 Threes", 300})
	t.AppendRow(table.Row{"3 Fours", 400})
	t.AppendRow(table.Row{"3 Fives", 500})
	t.AppendRow(table.Row{"3 Sixes", 600})
	t.AppendRow(table.Row{"Four of a Kind", 1000})
	t.AppendRow(table.Row{"Five of a Kind", 2000})
	t.AppendRow(table.Row{"Six 0f a Kind", 3000})
	t.AppendRow(table.Row{"1-6 Straight", 1500})
	t.AppendRow(table.Row{"3 Pairs", 1500})
	t.AppendRow(table.Row{"2 Triplets", 2500})
	t.AppendRow(table.Row{"Four of a Kind + 1 Pair", 1500})
	fmt.Println(t.Render())

	fmt.Println("You must keep at least one set, then reroll the remaining dice until you want to keep.")
	fmt.Println("If a roll has no playable set, you've FARKLED and lose your turn.")
	fmt.Println("Score 10,000 points to win! \n")
	fmt.Println("Press Enter to start playing!")

	fmt.Scanln()

}

func playHand(player Player) Player {

	fmt.Print(clearScreen)
	fmt.Println("Here is your starting hand:")

	t := table.NewWriter()
	t.AppendHeader(table.Row{"Die 1", "Die 2", "Die 3", "Die 4", "Die 5", "Die 6"})
	t.AppendRow(table.Row{player.startingHand[0], player.startingHand[1], player.startingHand[2],
		player.startingHand[3], player.startingHand[4], player.startingHand[5]})
	fmt.Println(t.Render())

	for {
		var option string

		fmt.Println("Which dice will you keep?")

		fmt.Printf("Press 1-%v, then press enter.  Enter D when done choosing.", len(player.startingHand))
		fmt.Scanln(&option)

		if option != "d" {
			intOption, _ := strconv.Atoi(option)
			player.keptDice = append(player.keptDice, player.startingHand[(intOption-1)])
			player.diceLeft -= 1
			fmt.Println(player.keptDice)
		} else {
			scoreHand()
			break
		}
	}
	return player
}
func scoreHand() {

}
func rollDice() int {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	roll := int(r.Intn(6))
	interval := int(r.Intn(200))
	time.Sleep(time.Duration(interval) * time.Millisecond)

	if roll == 0 {
		roll = roll + 1
	}

	return roll
}

func newPlayer() Player {

	var player Player
	for i := 0; i < 6; i++ {

		player.startingHand = append(player.startingHand, rollDice())
	}
	player.Score = 0
	player.diceLeft = 6

	//hand.rolls = sort(hand.rolls)

	return player
}

func rollHand(player Player) Player {
	for i := 0, i < player.diceLeft; i++ {

		player.startingHand = append(player.startingHand, rollDice())
	}

	return player
}

func ProcessHand(hand Player) Player {

	last := hand.startingHand[0]
	isSame := make([]int, 6)

	for i := 1; i < len(hand.startingHand); i++ {
		if hand.startingHand[i] == last {

		}
	}

	for i, roll := range hand.startingHand {
		if roll == last {
			isSame[i]++
		} else {
			isSame[i] = 0
		}
	}

	return hand
}

// func findSets(hand Hand) Hand {

// 	var dupes []int
// 	visited := make(map[int]bool, 0)

// 	for i := 0; i < len(hand.rolls); i++ {
// 		if visited[hand.rolls[i]] == true {
// 			dupes = append(dupes, hand.rolls[i])
// 		} else {
// 			visited[hand.rolls[i]] = true
// 		}
// 	}

// 	if len(dupes) == 0 {
// 		hand.sixes = true
// 		return hand
// 	}

// 	if len(distinct(dupes)) == len(dupes) {
// 		hand.pairs = append(hand.pairs, dupes...)
// 		return hand
// 	}
// 	//count how many times each duplicate number appears
// 	secondPass := make(map[int]int, 0)
// 	for _, dupe := range dupes {
// 		secondPass[dupe]++
// 	}

// 	return hand
// }

// func findPairs(hand Hand) Hand {

// 	visited := make(map[int]bool, 0)

// 	for i := 0; i < len(hand.rolls); i++ {
// 		if visited[hand.rolls[i]] == true {
// 			hand.pairs = append(hand.pairs, hand.rolls[i])
// 		} else {
// 			visited[hand.rolls[i]] = true
// 		}
// 	}
// 	return hand
//}

// func findTrips(hand Hand) Hand {

// 	visited := make(map[int]bool, 0)

// 	for i := 0; i < len(hand.pairs); i++ {
// 		if visited[hand.pairs[i]] == true {
// 			if !contains(hand.trips, hand.pairs[i]) {
// 				hand.pairs = append(hand.trips, hand.pairs[i])
// 			}
// 		} else {
// 			visited[hand.pairs[i]] = true
// 		}
// 	}

// 	hand.pairs = distinct(hand.pairs)
// 	return hand
// }

func sort(n []int) []int {

	var isDone = false
	for !isDone {
		isDone = true
		for i := 0; i < len(n)-1; i++ {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
				isDone = false
			}
		}
	}
	return n
}

func distinct(intSlice []int) []int {

	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func contains(i []int, val int) bool {

	for _, v := range i {
		if v == val {
			return true
		}
	}
	return false
}
