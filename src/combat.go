package src

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/fatih/color"
)

func InitMonstre() monstre {
	var ennemi monstre
	var n int
	n = rand.IntN(2) // 0 ou 1
	if n == 0 {
		ennemi = InitDragonRouge(ennemi)
	} else {
		ennemi = InitLoup_Garou(ennemi)
	}
	return ennemi
}
func StartCombat(p1 Character) {
	ennemi := InitMonstre()
	combat(&p1, &ennemi)
}
func combat(p1 *Character, ennemi *monstre) {
	red := color.New(color.FgRed).SprintFunc()
	tour := 1
	dot := 1
	fmt.Println("Un", ennemi.nom, "sauvage apparaît !")
	fmt.Println(infoMonstres(*ennemi))
	for p1.act_pv > 0 && ennemi.act_pv > 0 {
		fmt.Println("----------")
		fmt.Println("Tour :", tour)
		fmt.Println("Vos PV :", p1.act_pv, "/", p1.max_pv)
		fmt.Println("PV du monstre :", ennemi.act_pv, "/", ennemi.maxhp)
		fmt.Println("----------")
		fmt.Print("\n")
		fmt.Println("Que voulez-vous faire ?")
		fmt.Println("[1] Attaques")
		fmt.Println("[2] Inventaire")
		fmt.Println("[3] Fuir (retour au menu principal)")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			SkillsUseMenu(p1, ennemi)
		case 2:
			accesssInventory(p1)
		case 3:
			fmt.Println(red("Vous fuyez le combat..."))
			AfficherMenu(*p1)
			break
		default:
			fmt.Println("Choix invalide, vous perdez votre tour...")
		}

		if dot == 4 {
			ennemi.status = false
			dot = 0
		}
		if ennemi.status {
			fmt.Println(red("Le monstre est empoisonné et perd 5 PV !"))
			ennemi.act_pv -= 5
		}
		if ennemi.act_pv == 0 {
			fmt.Println("Bravo ! Vous avez vaincu le monstre !")
			AfficherMenu(*p1)
			break
		}
		fmt.Println("Le monstre attaque et inflige", ennemi.pts_attaque, "dégâts !")
		p1.act_pv -= ennemi.pts_attaque
		fmt.Println("Vos PV :", p1.act_pv, "/", p1.max_pv)

		if isDead(*p1) {
			fmt.Print("\n")
			fmt.Println(red("############################################"))
			fmt.Println(red("#                                          #"))
			fmt.Println(red("#                                          #"))
			fmt.Println(red("#             T'ES MORT BOLOSS             #"))
			fmt.Println(red("#                                          #"))
			fmt.Println(red("#                                          #"))
			fmt.Println(red("############################################"))
			fmt.Print("\n")
			fmt.Println("Voulez-vous :")
			fmt.Println("[1] Recommencer le combat")
			fmt.Println("[2] Retour au menu principal")
			fmt.Print("Votre choix : ")
			var choix string
			fmt.Scan(&choix)
			choix = strings.ToLower(choix)

			if choix == "1" || choix == "r" || choix == "recommencer" {
				ennemi.act_pv = ennemi.maxhp
				p1.act_pv = p1.max_pv
				combat(p1, ennemi)
			} else if choix == "2" || choix == "m" || choix == "menu" {
				AfficherMenu(*p1)
			} else {
				fmt.Println("Choix invalide, retour au menu.")
				AfficherMenu(*p1)
			}
		}
		dot++
		tour++
	}
}
