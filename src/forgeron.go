package src

import (
	"fmt"

	"github.com/fatih/color"
)

func afficheForgeron(p1 Character) {
	ma := color.New(color.FgMagenta).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	fmt.Print("\n")
	fmt.Println(ma("===Forge==="))
	fmt.Println("[1] Chapeau de l'aventurier")
	fmt.Println("[2] Tunique de l'aventurier")
	fmt.Println("[3] Bottes de l'aventurier")
	fmt.Println("[4] Retour")
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
		fmt.Println("materiaux requis : \n")
		fmt.Println("- 1 Plume de Corbeau")
		fmt.Println("- 1 Cuir de Sanglier \n")
		if checkItem(&p1, "Plume de Corbeau") && checkItem(&p1, "Cuir de Sanglier") {
			removeInventory(&p1, "Plume de Corbeau")
			removeInventory(&p1, "Cuir de Sanglier")
			fmt.Println("Vous avez fabriqué un Chapeau de l'aventurier ! \n")
			addInventory(&p1, "Chapeau de l'aventurier")
		} else {
			fmt.Println(red("Vous n'avez pas les matériaux requis.\n"))
		}
		afficheForgeron(p1)
	case 2:
		fmt.Println("materiaux requis : \n")
		fmt.Println("- 2 Fourrures de Loup")
		fmt.Println("- 1 Peau de Troll \n")
		if checkItem(&p1, "Fourrure de Loup") && checkItem(&p1, "Peau de Troll") {
			removeInventory(&p1, "Fourrure de Loup")
			removeInventory(&p1, "Fourrure de Loup")
			removeInventory(&p1, "Peau de Troll")
			fmt.Println("Vous avez fabriqué une Tunique de l'aventurier ! \n")
			addInventory(&p1, "Tunique de l'aventurier")
		} else {
			fmt.Println(red("Vous n'avez pas les matériaux requis.\n"))
		}
		afficheForgeron(p1)
	case 3:
		fmt.Println("materiaux requis : \n")
		fmt.Println("- 1 Fourrure de Loup")
		fmt.Println("- 1 Cuir de Sanglier \n")
		if checkItem(&p1, "Fourrure de Loup") && checkItem(&p1, "Cuir de Sanglier") {
			removeInventory(&p1, "Fourrure de Loup")
			removeInventory(&p1, "Cuir de Sanglier")
			fmt.Println("Vous avez fabriqué des Bottes de l'aventurier ! \n")
			addInventory(&p1, "Bottes de l'aventurier")
		} else {
			fmt.Println(red("Vous n'avez pas les matériaux requis.\n"))
		}
		afficheForgeron(p1)
	case 4:
		AfficherMenu(p1)
	case 0:
		SeeYouNextTime()
		return
	default:
		fmt.Println(red("Choix invalide"))
	}
}
