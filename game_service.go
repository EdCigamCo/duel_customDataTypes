package main

import "bufio"

type GameService struct {
	Enemies []*Character
	Scanner *bufio.Scanner
}
