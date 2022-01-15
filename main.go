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
	rolls       []int
	keptDice    []int
	score       int
	diceLeft    int
	donePlaying bool
}

func main() {

	showRules()

	player := newPlayer()

	//player = playStartingHand(player)

	for {
		player = playSubsequentHand(player)

		if player.donePlaying {
			return
		}
	}
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

func playStartingHand(player Player) Player {

	fmt.Print(clearScreen)
	fmt.Println("Here is your starting hand:")

	t := generateDiceTable(player.rolls)
	fmt.Println(t.Render())

	for {
		player = handleRollingOptions(player)

		if !player.donePlaying {
			player = playSubsequentHand(player)
		}

		return player
	}
}

func playSubsequentHand(player Player) Player {

	fmt.Print(clearScreen)
	fmt.Println("You have rolled...")

	player = rollHand(player)

	rollTable := generateDiceTable(player.rolls)
	fmt.Println(rollTable.Render())

	for {
		player = handleRollingOptions(player)

		if player.donePlaying {
			return player
		}
	}
}

func handleRollingOptions(player Player) Player {

	for {
		var option string

		fmt.Println("Which dice will you keep?")
		fmt.Printf("Enter 1-%v to keep a die.\n", player.diceLeft)
		fmt.Println("Enter R to reroll remaining dice, or E to end turn.")
		fmt.Scanln(&option)

		if option == "" {
			fmt.Printf("You must enter a number 1-%v, R to reroll or E to end turn.\n", player.diceLeft)
			return player
		}

		if option != "r" && option != "e" {
			intOption, _ := strconv.Atoi(option)

			if intOption > 0 {
				player.keptDice = append(player.keptDice, player.rolls[(intOption-1)])
			} else {
				player.keptDice = append(player.keptDice, player.rolls[0])
			}

			player.rolls = append(player.rolls[:intOption-1], player.rolls[intOption:]...)

			player.diceLeft -= 1

			remainingTable := generateDiceTable(player.rolls)
			fmt.Println(remainingTable.Render())

			if len(player.keptDice) > 0 {
				fmt.Println("\nYou're holding:")

				holdingTable := generateDiceTable(player.keptDice)

				if player.score > 0 {
					holdingTable.SetCaption("Current Score: %v", player.score)
				}

				fmt.Println(holdingTable.Render())
			}

			//fmt.Println(player.keptDice)

			// if player.diceLeft == 0 {
			// 	player.donePlaying = true
			// }
		} else if option == "r" {
			player.score += calculateScore(player.keptDice)
			return player
		} else if option == "e" {
			player.score += calculateScore(player.keptDice)
			player.donePlaying = true
			return player
		}
	}
}

func calculateScore(keptDice []int) int {

	score := 0
	keptDice = sort(keptDice)

	if len(keptDice) == 1 {

		switch keptDice[0] {
		case 1:
			score += 100
		case 5:
			score += 50
		}
	}

	if len(keptDice) == 3 {

		switch sum(keptDice) {
		case 3:
			score += 300
		case 6:
			score += 200
		case 9:
			score += 300
		case 12:
			score += 400
		case 15:
			score += 500
		case 18:
			score += 600
		}
	}

	if len(keptDice) == 4 {
		if findOfAKinds(keptDice) {
			score += 1000
		}
	}

	if len(keptDice) == 5 {
		if findOfAKinds(keptDice) {
			score += 2000
		}
	}

	if len(keptDice) == 6 {
		if findOfAKinds(keptDice) {
			score += 3000
		} else if findThreePairs(keptDice) {
			score += 1500
		} else if findTwoTriplets(keptDice) {
			score += 2500
		}
	}

	return score
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

func generateDiceTable(values []int) table.Writer {

	t := table.NewWriter()
	header := table.Row{}
	row := table.Row{}

	for i := 0; i < len(values); i++ {
		dieNo := "Die " + strconv.Itoa(i+1)
		header = append(header, dieNo)
		row = append(row, values[i])
	}

	t.AppendHeader(header)
	t.AppendRow(row)

	return t

}

func newPlayer() Player {

	var player Player
	player.score = 0
	player.diceLeft = 6
	player = rollHand(player)
	player.donePlaying = false

	//hand.rolls = sort(hand.rolls)

	return player
}

func rollHand(player Player) Player {

	player.rolls = nil

	for i := 0; i < player.diceLeft; i++ {
		player.rolls = append(player.rolls, rollDice())
	}

	return player
}

func ProcessHand(hand Player) Player {

	last := hand.rolls[0]
	isSame := make([]int, 6)

	for i := 1; i < len(hand.rolls); i++ {
		if hand.rolls[i] == last {

		}
	}

	for i, roll := range hand.rolls {
		if roll == last {
			isSame[i]++
		} else {
			isSame[i] = 0
		}
	}

	return hand
}

func findOfAKinds(keptDice []int) bool {

	last := 0
	score := 0

	for i := 0; i < len(keptDice); i++ {
		if last == keptDice[i] {
			score++
		} else {
			last = keptDice[i]
		}
	}

	if score == (len(keptDice) - 1) {
		return true
	}
	return false

	// if len(distinct(dupes)) == len(dupes) {
	// 	player.pairs = append(player.pairs, dupes...)
	// 	return player
	// }
	// //count how many times each duplicate number appears
	// secondPass := make(map[int]int, 0)
	// for _, dupe := range dupes {
	// 	secondPass[dupe]++
	// }

	// return player
}

func findThreePairs(keptDice []int) bool {

	visited := make(map[int]bool, 0)
	var pairs []int

	for i := 0; i < len(keptDice); i++ {
		if visited[keptDice[i]] == true {
			pairs = append(pairs, keptDice[i])
		} else {
			visited[keptDice[i]] = true
		}
	}

	if len(distinct(pairs)) == 3 {
		return true
	}
	return false
}

func findTwoTriplets(keptDice []int) bool {

	visited := make(map[int]bool, 0)
	var dupes []int
	var trips []int

	for i := 0; i < len(keptDice); i++ {
		if visited[keptDice[i]] == true {
			if contains(dupes, keptDice[i]) {
				trips = append(trips, keptDice[i])
			} else {
				dupes = append(dupes, keptDice[i])
			}
		} else {
			visited[keptDice[i]] = true
		}
	}

	if len(distinct(trips)) == 2 {
		return true
	}
	return false
}

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

func sum(array []int) int {
	result := 0

	for _, v := range array {
		result += v
	}

	return result
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
