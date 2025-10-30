package main

import (
	"fmt"
	"github.com/fatih/color"
	"projetred/src"
)

func main() {
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Println(yellow("############################################"))
	fmt.Println(yellow("#                                          #"))
	fmt.Println(yellow("#                                          #"))
	fmt.Println(yellow("#    Très Obscur : Expédition GoLang       #"))
	fmt.Println(yellow("#                                          #"))
	fmt.Println(yellow("#                                          #"))
	fmt.Println(yellow("############################################"))
	var p1 src.Character
	src.AfficherMenu(p1)
}
