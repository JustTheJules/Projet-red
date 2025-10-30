package src

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type Character struct {
	nom           string
	classe        string
	niveau        int
	max_pv        int
	act_pv        int
	inventory     []string
	inventorySize int
	money         int
	equip         Equipement
	skill         [2]string
	perso_init    bool
	atk           int
}

func InitCharacterHumain(p1 Character) Character {
	p1.nom = ""
	p1.classe = "Humain"
	p1.niveau = 1
	p1.max_pv = 100
	p1.act_pv = 50
	p1.money = 50
	p1.skill = [2]string{"Coup de poing"}
	p1.inventory = make([]string, p1.inventorySize)
	p1.inventory = append(p1.inventory, "potion", "potion", "potion")
	p1.inventorySize = 10
	p1.atk = 5
	return p1
}

func InitCharacterElfe(p1 Character) Character {
	p1.nom = ""
	p1.classe = "Elfe"
	p1.niveau = 1
	p1.max_pv = 80
	p1.act_pv = 40
	p1.money = 50
	p1.skill = [2]string{"Coup de poing"}
	p1.inventory = make([]string, p1.inventorySize)
	p1.inventory = append(p1.inventory, "potion", "potion", "potion")
	p1.inventorySize = 10
	p1.atk = 5
	return p1
}

func InitCharacterNain(p1 Character) Character {
	p1.nom = ""
	p1.classe = "nain"
	p1.niveau = 1
	p1.max_pv = 120
	p1.act_pv = 60
	p1.money = 50
	p1.skill = [2]string{"Coup de poing"}
	p1.inventory = make([]string, p1.inventorySize)
	p1.inventory = append(p1.inventory, "potion", "potion", "potion")
	p1.inventorySize = 10
	p1.atk = 5
	return p1
}
func InitCharacterAdmin(p1 Character) Character {
	p1.nom = "----Admin----"
	p1.classe = "Admin"
	p1.niveau = 1
	p1.max_pv = 5
	p1.act_pv = 5
	p1.money = 1000
	p1.inventory = make([]string, p1.inventorySize)
	p1.inventory = append(p1.inventory, "potion", "potion", "potion", "Potion de Poison", "Livre de sort : Boule de feu")
	p1.inventorySize = 10
	p1.atk = -1
	return p1
}
func displayInfo(p1 Character) string {
	hgre := color.New(color.FgHiCyan).SprintFunc()
	fmt.Println(hgre("===Stats==="))
	fmt.Println("Nom: " + p1.nom)
	fmt.Println("Classe: " + p1.classe)
	fmt.Println("Niveau: ", p1.niveau)
	fmt.Println("Pv : ", p1.act_pv, "/", p1.max_pv)
	fmt.Println("Inventory: ", p1.inventory)
	fmt.Println("Taille de l'inventaire: ", p1.inventorySize)
	fmt.Println("Argent: ", p1.money)
	fmt.Println("Sort: ", p1.skill)
	return fmt.Sprintf("Nom: %s\nClasse: %s\nNiveau: %d\nPv max: %d\nPv actuel: %d\nInvantaire: %v\nArgent: %d\n", p1.nom, p1.classe, p1.niveau, p1.max_pv, p1.act_pv, p1.inventory, p1.money)
}

func attribName(p1 Character) {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Choose your name")
	for scanner.Scan() {
		param := strings.Split(scanner.Text(), ",")
		if param[0] == "" {
			fmt.Println(red("tu sais meme plus ecrire ton nom ??"))
			continue
		}
		if len(param) >= 1 {
			if IsAlpha(param[0]) {
				p1.nom = Capitalize(param[0])
			} else {
				fmt.Println(red("Merci de ne rentrer que des lettres"))
				attribName(p1)
			}
			if param[0] == "Admin" || param[0] == "A" || param[0] == "admin" || param[0] == "a" {
				p1 = InitCharacterAdmin(p1)
			}
			fmt.Print("\n")
			fmt.Println(green("Personnage créé"))
			AfficherMenu(p1)
			break
		}
	}
}

func InitCharacter(p1 Character) Character {
	//green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Choose you calsse : Humain, Elfe, Nain")
	if !p1.perso_init {
		for scanner.Scan() {
			param := strings.Split(scanner.Text(), ",")
			if len(param) >= 1 {
				p1.classe = param[0]
				if param[0] == "Humain" || param[0] == "H" || param[0] == "humain" || param[0] == "h" {
					fmt.Println("classe choisie : Humain")
					p1 = InitCharacterHumain(p1)
				} else if param[0] == "Elfe" || param[0] == "E" || param[0] == "elfe" || param[0] == "e" {
					fmt.Println("classe choisie : Elfe")
					p1 = InitCharacterElfe(p1)
				} else if param[0] == "Nain" || param[0] == "N" || param[0] == "nain" || param[0] == "n" {
					fmt.Println("classe choisie : Nain")
					p1 = InitCharacterNain(p1)
				} else if param[0] == "Admin" || param[0] == "A" || param[0] == "admin" || param[0] == "a" {
					fmt.Println("classe choisie : Admin")
					p1 = InitCharacterAdmin(p1)
				} else {
					fmt.Println(red("Classe invalide,"), "veuillez choisir entre : Humain, Elfe, Nain")
					continue
				}
				p1.perso_init = true
				break
			}
		}
		attribName(p1)
	} else {
		fmt.Println(red("Tu as cru que tu pouvais changer de vie comme ça ?"))
		AfficherMenu(p1)
	}
	return p1
}
