package src

import (
	"fmt"

	"github.com/fatih/color"
)

func SeeYouNextTime() {
	colors := []func(a ...interface{}) string{
		color.New(color.FgRed).SprintFunc(),
		color.New(color.FgGreen).SprintFunc(),
		color.New(color.FgYellow).SprintFunc(),
		color.New(color.FgBlue).SprintFunc(),
		color.New(color.FgMagenta).SprintFunc(),
		color.New(color.FgCyan).SprintFunc(),
		color.New(color.FgHiRed).SprintFunc(),
		color.New(color.FgHiGreen).SprintFunc(),
		color.New(color.FgHiYellow).SprintFunc(),
		color.New(color.FgHiBlue).SprintFunc(),
		color.New(color.FgHiMagenta).SprintFunc(),
		color.New(color.FgHiCyan).SprintFunc(),
	}

	text := "See you next time"
	result := ""

	for i, char := range text {
		colorFunc := colors[i%len(colors)]
		result += colorFunc(string(char))
	}

	fmt.Println(result)
}

func AfficherMenu(p1 Character) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Print("\n")
	fmt.Println(green("===Menu Principale==="))
	fmt.Println("[1] Création de personnage")
	fmt.Println("[2] Infos du Personnage")
	fmt.Println("[3] Marché")
	fmt.Println("[4] Forgeron")
	fmt.Println("[5] Combat")
	fmt.Println("[6] Accès à l'inventaire")
	fmt.Println("[0] Quitter")
	var choix int
	fmt.Print("Votre choix : ")

	_, err := fmt.Scan(&choix)
	if err != nil {
		fmt.Println(red("⚠️  Entrée invalide, veuillez entrer un chiffre."))
		var trash string
		fmt.Scanln(&trash)
		AfficherMenu(p1)
		return
	}
	switch choix {
	case 1:
		p1 = InitCharacter(p1)
	case 2:
		if p1.classe == "" {
			fmt.Print("\n")
			fmt.Println(red("créez d'abord un personnage"))
		} else {
			fmt.Print("\n")
			displayInfo(p1)
		}
		AfficherMenu(p1)
	case 3:
		afficheMarché(p1)
	case 4:
		afficheForgeron(p1)
	case 5:
		StartCombat(p1)
	case 6:
		accesssInventory(&p1)
	case 0:
		SeeYouNextTime()
		return
	default:
		fmt.Println(red("Choix invalide"))
		AfficherMenu(p1)
	}

	return // Default return value if no valid choice is made
}

func afficheMarché(p1 Character) {
	blu := color.New(color.FgBlue).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	fmt.Print("\n")
	fmt.Println(blu("===Marché==="))
	fmt.Println("[1] Potion de vie -- 3G")
	fmt.Println("[2] Épée -- 20G")
	fmt.Println("[3] Chapeau -- 50G")
	fmt.Println("[4] Composants")
	fmt.Println("[5] Augmentation de l'inventaire -- 30G")
	fmt.Println("[6] Livre de sort : Boule de feu -- 15G")
	fmt.Println("[7] Potion de Poison -- 6G")
	fmt.Println("[9] Retour")
	fmt.Println("[0] Quitter")
	var choix int
	fmt.Print("Votre choix : ")
	_, err := fmt.Scan(&choix)
	if err != nil {
		fmt.Print("\n")
		fmt.Println(red("⚠️  Entrée invalide, veuillez entrer un chiffre."))
		fmt.Print("\n")
		var trash string
		fmt.Scanln(&trash)
		afficheMarché(p1)
		return
	}

	switch choix {
	case 1:
		buyStuff(p1, "potion", 3)
	case 2:
		buyStuff(p1, "épée", 20)
	case 3:
		buyStuff(p1, "chapeau", 50)
	case 4:
		afficheComposants(p1)
	case 5:
		upgradeInventory(&p1, 30)
	case 6:
		buyStuff(p1, "Livre de sort : Boule de feu", 15)
	case 7:
		buyStuff(p1, "potion de poison", 6)
	case 9:
		AfficherMenu(p1)
	case 0:
		SeeYouNextTime()
		return
	default:
		fmt.Println("Choix invalide")
	}
}

func afficheComposants(p1 Character) {
	cy := color.New(color.FgHiCyan).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	fmt.Print("\n")
	fmt.Println(cy("===Composants==="))
	fmt.Println("[1] Plume de Corbeau -- 5G")
	fmt.Println("[2] Cuir de Sanglier -- 15G")
	fmt.Println("[3] Fourrure de Loup -- 10G")
	fmt.Println("[4] Peau de Troll -- 25G")
	fmt.Println("[5] Retour")
	fmt.Println("[0] Quitter")
	var choix int
	fmt.Print("Votre choix : ")
	_, err := fmt.Scan(&choix)
	if err != nil {
		fmt.Print("\n")
		fmt.Println(red("⚠️  Entrée invalide, veuillez entrer un chiffre."))
		fmt.Print("\n")
		var trash string
		fmt.Scanln(&trash)
		afficheComposants(p1)
		return
	}

	switch choix {
	case 1:
		buyStuff(p1, "Plume de Corbeau", 5)
	case 2:
		buyStuff(p1, "Cuir de Sanglier", 15)
	case 3:
		buyStuff(p1, "Fourrure de Loup", 10)
	case 4:
		buyStuff(p1, "Peau de Troll", 25)
	case 5:
		afficheMarché(p1)
	case 0:
		SeeYouNextTime()
		return
	default:
		fmt.Println("Choix invalide")
	}
}

func SkillsUseMenu(p1 *Character, cible *monstre) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Print("\n")
	fmt.Println(green("===Menu des compétences==="))
	fmt.Println("Vos compétences : ", p1.skill)
	fmt.Println("[1] Attaque de base (5 dégâts)")
	fmt.Println("[2] Coup de poing (8 dégâts)")
	fmt.Println("[3] Boule de feu (18 dégâts)")
	fmt.Println("[4] Potion de poison (15 dégâts sur 3 tours)")
	fmt.Println("[0] Retour")
	var choix int
	fmt.Print("Votre choix : ")
	_, err := fmt.Scan(&choix)
	if err != nil {
		fmt.Println(red("⚠️  Entrée invalide, veuillez entrer un chiffre."))
	}

	switch choix {
	case 1:
		AttaqueDeBase(p1, cible)
	case 2:
		CoupDePoing(p1, cible)
	case 3:
		BouleDeFeu(p1, cible)
	case 4:
		PoisonPot(p1, cible)
	case 0:
		AfficherMenu(*p1)
	default:
		fmt.Println(red("Choix invalide"))
		SkillsUseMenu(p1, cible)
	}
}
