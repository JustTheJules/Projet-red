package src

import (
	"fmt"
	"strings"
)

type Equipement struct {
	tete  string
	torse string
	pied  string
}

func accesssInventory(p1 *Character) {
	var input string
	fmt.Println("Voici votre inventaire:")
	fmt.Println(p1.inventory)
	fmt.Print("Voulez-vous utiliser une potion ? (oui/non) : ")
	fmt.Scan(&input)
	input = strings.TrimSpace(strings.ToLower(input))
	if input == "oui" {
		takePot(p1)
	} else {
		fmt.Println("Vous n'avez pas utilisé de potion")
	}
	if checkItem(p1, "Livre de sort : Boule de feu") {
		SpellBook(*p1)
		removeInventory(p1, "Livre de sort : Boule de feu")
	}
	fmt.Println("Voulez-vous équiper ou déséquiper un objet ? (équip/déséquip/exit) : ")
	fmt.Scan(&input)
	input = strings.TrimSpace(strings.ToLower(input))
	if input == "équip" {
		fmt.Print("Quel objet voulez-vous équiper ? : ")
		fmt.Println("[1] Chapeau de l'aventurier")
		fmt.Println("[2] Tunique de l'aventurier")
		fmt.Println("[3] Bottes de l'aventurier")
		fmt.Println("[4] Retour")
		var choix int
		fmt.Scan(&choix)
		switch choix {
		case 1:
			if checkItem(p1, "Chapeau de l'aventurier") {
				equipItem(p1, "Chapeau de l'aventurier")
			} else {
				fmt.Println("Vous ne possédez pas cet objet.")
			}
		case 2:
			if checkItem(p1, "Tunique de l'aventurier") {
				equipItem(p1, "Tunique de l'aventurier")
			} else {
				fmt.Println("Vous ne possédez pas cet objet.")
			}
		case 3:
			if checkItem(p1, "Bottes de l'aventurier") {
				equipItem(p1, "Bottes de l'aventurier")
			} else {
				fmt.Println("Vous ne possédez pas cet objet.")
			}
		case 4:
			fmt.Println("Retour à l'inventaire")
			accesssInventory(p1)
		default:
			fmt.Println("Choix invalide")
		}
	} else if input == "déséquip" {
		fmt.Print("Quel objet voulez-vous déséquiper ? : ")
		fmt.Println("[1] Chapeau de l'aventurier")
		fmt.Println("[2] Tunique de l'aventurier")
		fmt.Println("[3] Bottes de l'aventurier")
		fmt.Println("[4] Retour")
		var choix int
		fmt.Scan(&choix)
		switch choix {
		case 1:
			if checkEquip(p1, "Chapeau de l'aventurier", "tete") {
				unequipItem(p1, "Chapeau de l'aventurier")
			} else {
				fmt.Println("Vous n'avez pas cet objet équipé.")
			}
		case 2:
			if checkEquip(p1, "Tunique de l'aventurier", "torse") {
				unequipItem(p1, "Tunique de l'aventurier")
			} else {
				fmt.Println("Vous n'avez pas cet objet équipé.")
			}
		case 3:
			if checkEquip(p1, "Bottes de l'aventurier", "pied") {
				unequipItem(p1, "Bottes de l'aventurier")
			} else {
				fmt.Println("Vous n'avez pas cet objet équipé.")
			}
		case 4:
			fmt.Println("Retour à l'inventaire")
			accesssInventory(p1)
		default:
			fmt.Println("Choix invalide")
		}
	} else {
		fmt.Println("Retour au menu principal")
		AfficherMenu(*p1)
	}
}

func takePot(p1 *Character) {
	for i := 0; i < len(p1.inventory); i++ {
		if p1.inventory[i] == "potion" {
			if p1.act_pv == p1.max_pv {
				fmt.Println("Vos points de vie sont déjà au maximum.")
				return
			} else {
				p1.act_pv += 50
				removeInventory(p1, "potion")
			}

			if p1.act_pv > p1.max_pv {
				p1.act_pv = p1.max_pv
			}
			fmt.Println("Vous avez bu une potion, vos points de vie actuels sont de:", p1.act_pv)
			return
		}
	}
	fmt.Println("Vous n'avez plus de potion")
}

func addInventory(p1 *Character, item string) {
	if len(p1.inventory) >= p1.inventorySize {
		fmt.Println("Votre inventaire est plein, vous ne pouvez pas ajouter", item)
		return
	}
	p1.inventory = append(p1.inventory, item)
	fmt.Println("Vous avez ajouté", item, "à votre inventaire.")
}

func removeInventory(p1 *Character, item string) {
	for i := 0; i < len(p1.inventory); i++ {
		if p1.inventory[i] == item {
			p1.inventory[i] = ""
			fmt.Println("Vous avez retiré", item, "de votre inventaire.")
			return
		}
	}
	fmt.Println("L'objet", item, "n'est pas dans votre inventaire.")
}

func removeMoney(p1 *Character, amount int) {
	p1.money -= amount
	if p1.money < 0 {
		p1.money = 0
	}
}

func checkItem(p1 *Character, item string) bool {
	for i := 0; i < len(p1.inventory); i++ {
		if p1.inventory[i] == item {
			return true
		}
	}
	return false
}

func equipItem(p1 *Character, item string) {
	switch item {
	case "Chapeau de l'aventurier":
		p1.equip.tete = item
		fmt.Println("Vous avez équipé le", item)
	case "Tunique de l'aventurier":
		p1.equip.torse = item
		fmt.Println("Vous avez équipé la", item)
	case "Bottes de l'aventurier":
		p1.equip.pied = item
		fmt.Println("Vous avez équipé les", item)
	default:
		fmt.Println("Cet objet ne peut pas être équipé.")
	}
	modifyHpWithArmor(p1)
}

func unequipItem(p1 *Character, item string) {
	switch item {
	case "Chapeau de l'aventurier":
		p1.equip.tete = ""
		fmt.Println("Vous avez retiré le", item)
	case "Tunique de l'aventurier":
		p1.equip.torse = ""
		fmt.Println("Vous avez retiré la", item)
	case "Bottes de l'aventurier":
		p1.equip.pied = ""
		fmt.Println("Vous avez retiré les", item)
	default:
		fmt.Println("Cet objet n'est pas équipé.")
	}
	modifyHpWithArmor(p1)
}

func checkEquip(p1 *Character, item string, part string) bool {
	switch part {
	case "tete":
		return p1.equip.tete == item
	case "torse":
		return p1.equip.torse == item
	case "pied":
		return p1.equip.pied == item
	default:
		return false
	}
}

func modifyHpWithArmor(p1 *Character) {
	if !checkEquip(p1, "Chapeau de l'aventurier", "tete") && !checkEquip(p1, "Tunique de l'aventurier", "torse") && !checkEquip(p1, "Bottes de l'aventurier", "pied") {
		p1.max_pv = 100
		p1.act_pv = 100
	} //rien d'équipé
	if !checkEquip(p1, "Chapeau de l'aventurier", "tete") && checkEquip(p1, "Tunique de l'aventurier", "torse") && !checkEquip(p1, "Bottes de l'aventurier", "pied") {
		p1.max_pv = 125
		p1.act_pv = 125
	} //juste la tunique
	if checkEquip(p1, "Chapeau de l'aventurier", "tete") && !checkEquip(p1, "Tunique de l'aventurier", "torse") && !checkEquip(p1, "Bottes de l'aventurier", "pied") {
		p1.max_pv = 110
		p1.act_pv = 110
	} //juste le chapeau
	if !checkEquip(p1, "Chapeau de l'aventurier", "tete") && !checkEquip(p1, "Tunique de l'aventurier", "torse") && checkEquip(p1, "Bottes de l'aventurier", "pied") {
		p1.max_pv = 115
		p1.act_pv = 115
	} //juste les bottes
	if checkEquip(p1, "Chapeau de l'aventurier", "tete") && checkEquip(p1, "Tunique de l'aventurier", "torse") && !checkEquip(p1, "Bottes de l'aventurier", "pied") {
		p1.max_pv = 150
		p1.act_pv = 150
	} //chapeau et tunique
	if checkEquip(p1, "Chapeau de l'aventurier", "tete") && !checkEquip(p1, "Tunique de l'aventurier", "torse") && checkEquip(p1, "Bottes de l'aventurier", "pied") {
		p1.max_pv = 135
		p1.act_pv = 135
	} //chapeau et bottes
	if !checkEquip(p1, "Chapeau de l'aventurier", "tete") && checkEquip(p1, "Tunique de l'aventurier", "torse") && checkEquip(p1, "Bottes de l'aventurier", "pied") {
		p1.max_pv = 140
		p1.act_pv = 140
	} //tunique et bottes
	if !checkEquip(p1, "Chapeau de l'aventurier", "tete") && checkEquip(p1, "Tunique de l'aventurier", "torse") && checkEquip(p1, "Bottes de l'aventurier", "pied") {
		p1.max_pv = 160
		p1.act_pv = 160
	} //tout est équipé
}

func upgradeInventory(p1 *Character, price int) {
	if p1.money < price {
		fmt.Println("Vous n'avez pas assez d'or pour augmenter votre inventaire.")
		return
	} else if p1.inventorySize == 40 {
		fmt.Println("Votre inventaire est déjà au maximum.")
		afficheMarché(*p1)
	} else {
		p1.inventorySize += 10
		p1.money -= price
		fmt.Println("Vous avez augmenté la taille de votre inventaire de 10 emplacements. Nouvelle taille:", p1.inventorySize)
	}
	afficheMarché(*p1)
}
