package roulette

import (
	"fmt"
	"math/rand"

	"sajjadgozal/gameserver/internal/models"
)

type Game struct {
	UserID   int `json:"user_id"`  // User ID
	Result   int `json:"result"`   // Result of the game
	Amount   int `json:"amount"`   // Amount of the bet
	Winnings int `json:"winnings"` // Winnings of the user
	Balance  int `json:"balance"`  // Balance of the user
}

type Bet struct {
	Type   string `json:"type"`
	Fields []int  `json:"fields"`
	Amount int    `json:"amount"`
}

const (
	Red   = iota
	Black = iota
)

func PlayRoulette(a models.Account, bets []Bet) (game Game, err error) {

	// Check if the user has enough balance to place the bets
	total := 0
	for _, bet := range bets {
		if bet.Amount <= 0 {
			return game, fmt.Errorf("invalid bet amount")
		}
		total += bet.Amount
	}
	if int64(total) > a.Balance {
		return game, fmt.Errorf("insufficient balance to place the bets")
	}

	// Generate a random number from 0 to 38
	spinResult := generateRandomNumber() // Random number from 0 to 38

	// Calculate the winnings
	winnings := 0
	for _, bet := range bets {
		if win(getWinnerNumbers(bet), spinResult) {
			switch bet.Type {
			case "straight":
				winnings += bet.Amount * 35
			case "split":
				winnings += bet.Amount * 17
			case "street":
				winnings += bet.Amount * 11
			case "corner":
				winnings += bet.Amount * 8
			case "sixline":
				winnings += bet.Amount * 5
			case "dozen":
				winnings += bet.Amount * 2
			case "column":
				winnings += bet.Amount * 2
			case "redblack":
				winnings += bet.Amount * 2
			case "evenodd":
				winnings += bet.Amount * 2
			case "lowhigh":
				winnings += bet.Amount * 2
			case "snake":
				winnings += bet.Amount * 1
			case "fivenumber":
				winnings += bet.Amount * 6
			default:
				fmt.Println("Invalid Bet Type")
			}
		}
	}

	// Update the user's balance
	a.Balance += int64(winnings) - int64(total)

	game = Game{
		UserID:   int(a.ID),
		Result:   spinResult,
		Amount:   total,
		Winnings: winnings,
		Balance:  int(a.Balance),
	}

	return game, nil
}

func generateRandomNumber() int {
	// Generate a random number from 0 to 38
	return rand.Intn(39)
}

// Straight-Up Bet: Betting on a single number, including 0 or 00 (if playing American Roulette). Pays 35 to 1. 39 == 00 in this case
// Split Bet: Betting on two adjacent numbers by placing chips on the line between them. Pays 17 to 1.
// Street Bet (or Line Bet): Betting on a row of three numbers by placing chips on the outer edge of the row. Pays 11 to 1.
// Corner Bet (or Quad Bet): Betting on four numbers that meet at one corner by placing chips at the intersection of those numbers. Pays 8 to 1.
// Six Line Bet: Betting on two adjacent rows of numbers by placing chips on the outer corner shared by both rows. Pays 5 to 1.
// Dozen Bet: Betting on one of the three groups of 12 numbers: 1-12, 13-24, or 25-36. Pays 2 to 1.
// Column Bet: Betting on one of the three vertical columns of numbers. Pays 2 to 1.
// Red/Black Bet: Betting on whether the winning number will be red or black. Pays 1 to 1.
// Even/Odd Bet: Betting on whether the winning number will be even or odd. Pays 1 to 1.
// Low/High Bet: Betting on whether the winning number will be in the range of 1-18 (low) or 19-36 (high). Pays 1 to 1.
// Snake Bet: A special bet that covers the numbers 1, 5, 9, 12, 14, 16, 19, 23, 27, 30, 32, and 34. This bet looks like a snake winding through the numbers. It's a combination of various bets, and each bet is a one-unit bet.
// Five Number Bet (Only in American Roulette): Betting on the numbers 0, 00, 1, 2, and 3 by placing chips at the intersection of 0 and 1. Pays 6 to 1.

func getWinnerNumbers(bet Bet) []int {

	switch bet.Type {
	case "straight":
		return bet.Fields
	case "split":
		return splitBet(bet)
	case "street":
		return streetBet(bet)
	case "corner":
		return cornerBet(bet)
	case "sixline":
		return sixLineBet(bet)
	case "dozen":
		return dozenBet(bet)
	case "column":
		return columnBet(bet)
	case "redblack":
		return redBlackBet(bet)
	case "evenodd":
		return evenOddBet(bet)
	case "lowhigh":
		return lowHighBet(bet)
	case "snake":
		return snakeBet(bet)
	case "fivenumber":
		return fiveNumberBet(bet)
	default:
		fmt.Println("Invalid Bet Type")
	}
	return []int{}
	// This is a placeholder for the roulette game logic

}

// Split Bet: Betting on two adjacent numbers by placing chips on the line between them. Pays 17 to 1.
// TODO: correct the splitBet function
func splitBet(bet Bet) []int {
	var numbers []int
	for _, num := range bet.Fields {
		if num > 0 && num < 37 { // Ensure the number is within the valid range for a Split bet
			if num%3 != 0 && num != 36 { // Ensure the number is not at the right edge of the table and not at the bottom-right corner
				numbers = append(numbers, num, num+1)
			}
		}
	}
	return numbers
}

// Street Bet (or Line Bet): Betting on a row of three numbers by placing chips on the outer edge of the row. Pays 11 to 1.
func streetBet(bet Bet) []int {
	var numbers []int
	for _, num := range bet.Fields {
		startNum := (num-1)*3 + 1 // Calculate the starting number for the street
		for i := 0; i < 3; i++ {
			numbers = append(numbers, startNum+i)
		}
	}
	return numbers
}

func cornerBet(bet Bet) []int {
	var numbers []int
	for _, num := range bet.Fields {
		if num > 0 && num < 33 { // Ensure the number is within the valid range for a corner bet
			if num%3 != 0 && num != 31 { // Ensure the number is not at the right edge of the table and not at the bottom-right corner
				numbers = append(numbers, num, num+1, num+3, num+4)
			}
		}
	}
	return numbers
}

// Six Line Bet: Betting on two adjacent rows of numbers by placing chips on the outer corner shared by both rows. Pays 5 to 1.
func sixLineBet(bet Bet) []int {
	var numbers []int
	for _, num := range bet.Fields {
		if num > 0 && num < 31 { // Ensure the number is within the valid range for a Six Line bet
			if num%3 != 0 { // Ensure the number is not at the right edge of the table
				numbers = append(numbers, num, num+1, num+2, num+3, num+4, num+5)
			}
		}
	}
	return numbers
}

// Dozen Bet: Betting on one of the three groups of 12 numbers: 1-12, 13-24, or 25-36. Pays 2 to 1. field 0 = 1-12, 1 = 13-24, 2 = 25-36
func dozenBet(bet Bet) []int {
	var numbers []int
	for i := 1; i <= 36; i++ {
		switch bet.Fields[0] {
		case 0:
			if i <= 12 {
				numbers = append(numbers, i)
			}
		case 1:
			if i > 12 && i <= 24 {
				numbers = append(numbers, i)
			}
		case 2:
			if i > 24 {
				numbers = append(numbers, i)
			}
		}
	}
	return numbers
}

// Column Bet: Betting on one of the three vertical columns of numbers. Pays 2 to 1. field 0 = 1st column, 1 = 2nd column, 2 = 3rd column
func columnBet(bet Bet) []int {
	var numbers []int
	for i := 1; i <= 36; i++ {
		switch bet.Fields[0] {
		case 0:
			if i%3 == 1 {
				numbers = append(numbers, i)
			}
		case 1:
			if i%3 == 2 {
				numbers = append(numbers, i)
			}
		case 2:
			if i%3 == 0 {
				numbers = append(numbers, i)
			}
		}
	}
	return numbers
}

// Red/Black Bet: Betting on whether the winning number will be red or black. Pays 1 to 1. field 0 = red, 1 = black
func redBlackBet(bet Bet) []int {
	var numbers []int
	for i := 1; i <= 36; i++ {
		switch bet.Fields[0] {
		case 0: // Red
			if (i <= 10 || (i >= 19 && i <= 28)) && i%2 != 0 {
				numbers = append(numbers, i)
			}
		case 1: // Black
			if ((i >= 11 && i <= 18) || (i >= 29 && i <= 36)) && i%2 == 0 {
				numbers = append(numbers, i)
			}
		}
	}
	return numbers
}

// Even/Odd Bet: Betting on whether the winning number will be even or odd. Pays 1 to 1. field 0 = even, 1 = odd
func evenOddBet(bet Bet) []int {
	var numbers []int
	for i := 1; i <= 36; i++ {
		switch bet.Fields[0] {
		case 0: // Even
			if i%2 == 0 {
				numbers = append(numbers, i)
			}
		case 1: // Odd
			if i%2 != 0 {
				numbers = append(numbers, i)
			}
		}
	}
	return numbers
}

// Low/High Bet: Betting on whether the winning number will be in the range of 1-18 (low) or 19-36 (high). Pays 1 to 1. field 0 = low, 1 = high
func lowHighBet(bet Bet) []int {
	var numbers []int
	for i := 1; i <= 36; i++ {
		switch bet.Fields[0] {
		case 0: // Low
			if i <= 18 {
				numbers = append(numbers, i)
			}
		case 1: // High
			if i > 18 {
				numbers = append(numbers, i)
			}
		}
	}
	return numbers
}

// Snake Bet: A special bet that covers the numbers 1, 5, 9, 12, 14, 16, 19, 23, 27, 30, 32, and 34. This bet looks like a snake winding through the numbers. It's a combination of various bets, and each bet is a one-unit bet.
func snakeBet(bet Bet) []int {
	return []int{1, 5, 9, 12, 14, 16, 19, 23, 27, 30, 32, 34}
}

// Five Number Bet (Only in American Roulette): Betting on the numbers 0, 00, 1, 2, and 3 by placing chips at the intersection of 0 and 1. Pays 6 to 1.
// 39 == 00 in this case
func fiveNumberBet(bet Bet) []int {
	return []int{0, 39, 1, 2, 3}
}

// Determine the winning numbers based on the bets
func win(numbers []int, nid int) bool {
	for _, num := range numbers {
		if num == nid {
			return true
		}
	}
	return false
}
