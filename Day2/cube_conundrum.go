package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Hand struct {
	red   int
	green int
	blue  int
}

func main() {
	filename := "input.txt"

	file, err := os.Open(filename)
	if err != nil {
		print(err.Error())
		return
	}

	config := Hand{
		red:   12,
		green: 13,
		blue:  14,
	}
	inputReader := bufio.NewReader(file)

	sum := 0
	for game, err := inputReader.ReadString('\n'); err == nil; game, err = inputReader.ReadString('\n') {
		hands := getGameHands(game)
		isValidGame := true
		for _, hand := range hands {
			if !isValidHand(hand, config) {
				isValidGame = false
				break
			}
		}
		if isValidGame {
			sum += getGameId(game)
		}
	}
	println(sum)
}

func getGameId(game string) int {
	pattern := "^Game\\s(\\d+):.*"
	regex, _ := regexp.Compile(pattern)
	match := regex.FindAllStringSubmatch(game, -1)
	if match == nil {
		return 0
	}

	num, _ := strconv.ParseInt(match[0][1], 10, 32)
	return int(num)
}

func getGameHands(game string) []Hand {
	hands := []Hand{}

	handPattern := "Game \\d+:|;?([^;]+)"
	regex, _ := regexp.Compile(handPattern)
	match := regex.FindAllStringSubmatch(game, -1)
	if match == nil {
		return hands
	}

	for _, hand := range match[1:] {
		hands = append(hands, Hand{
			red:   getColorCount(hand[1], "red"),
			green: getColorCount(hand[1], "green"),
			blue:  getColorCount(hand[1], "blue"),
		})
	}

	return hands
}

func getColorCount(hand string, color string) int {
	pattern := "\\s*(\\d+)\\s*" + color
	regex, _ := regexp.Compile(pattern)
	match := regex.FindStringSubmatch(hand)
	if len(match) != 2 {
		return 0
	}
	num, _ := strconv.ParseInt(match[1], 10, 32)
	return int(num)
}

func isValidHand(hand Hand, configuration Hand) bool {
	return hand.blue <= configuration.blue && hand.red <= configuration.red && hand.green <= configuration.green
}
