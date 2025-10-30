package src

import (
	"fmt"

	"github.com/fatih/color"
)

func PoisonPot(p1 *Character, enne *monstre) {
	red := color.New(color.FgRed).SprintFunc()
	if !checkItem(p1, "Potion de Poison") {
		fmt.Println(red("Tu ne possede pas un tel item !"))
		SkillsUseMenu(p1, enne)
	} else {
		removeInventory(p1, "Potion de Poison")
		fmt.Println("Vous utilisez la Potion de Poison")
		enne.status = true
	}
}
