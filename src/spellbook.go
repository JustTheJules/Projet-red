package src

import "fmt"

func SpellBook(p1 Character) {
	for _, m := range p1.skill {
		if m == "Boule de feu" {
			fmt.Println("Sort deja appris")
			return
		}
	}

	for i, m := range p1.skill {
		if m == "" {
			p1.skill[i] = "Boule de feu"
			fmt.Println("nouveau sort appris : Boule de feu")
			return
		}
	}
	fmt.Println("Impossible d'apprendre de nouveaux sorts (emplacements pleins)")
}
func UnlockedSkill(p1 Character, skill string) bool {
	for _, m := range p1.skill {
		if m == skill {
			return true
		}
	}
	return false
}

func AttaqueDeBase(p1 *Character, enne *monstre) {
	enne.act_pv -= 5
}
func CoupDePoing(p1 *Character, enne *monstre) {
	enne.act_pv -= 8
}
func BouleDeFeu(p1 *Character, enne *monstre) {
	if !UnlockedSkill(*p1, "Boule de feu") {
		fmt.Println("Vous n'avez pas appris ce sort.")
		SkillsUseMenu(p1, enne)
	} else {
		enne.act_pv -= 18
	}
}
