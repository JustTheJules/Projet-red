package src

import (
	"fmt"
)

func buyStuff(p1 Character, item string, price int) {
	if item == "potion" {
		addInventory(&p1, "potion")
		fmt.Println("Vous avez acheté une potion pour", price, "or.")
		afficheMarché(p1)
	} else if item == "épée" {
		if p1.money < price {
			fmt.Println("Vous n'avez pas assez d'or pour acheter une épée.")
			afficheMarché(p1)
			return
		} else {
			addInventory(&p1, "épée")
			fmt.Println("Vous avez acheté une épée pour", price, "or.")
			removeMoney(&p1, price)
			afficheMarché(p1)
		}
	} else if item == "chapeau" {
		if p1.money < price {
			fmt.Println("Vous n'avez pas assez d'or pour acheter un chapeau.")
			afficheMarché(p1)
			return
		} else {
			addInventory(&p1, "chapeau")
			fmt.Println("Vous avez acheté un chapeau pour", price, "or.")
			removeMoney(&p1, price)
			afficheMarché(p1)
		}
	} else if item == "Livre de sort : Boule de feu" {
		if p1.money < price {
			fmt.Println("Vous n'avez pas assez d'or pour acheter un livre de sort.")
			afficheMarché(p1)
			return
		} else {
			addInventory(&p1, "Livre de sort : Boule de feu")
			fmt.Println("Vous avez acheté un livre de sort : Boule de feu pour", price, "or.")
			removeMoney(&p1, price)
			afficheMarché(p1)
		}
	} else if item == "Plume de Corbeau" || item == "Cuir de Sanglier" || item == "Fourrure de Loup" || item == "Peau de Troll" {
		if p1.money < price {
			fmt.Println("Vous n'avez pas assez d'or pour acheter", item)
			afficheComposants(p1)
			return
		} else {
			addInventory(&p1, item)
			fmt.Println("Vous avez acheté", item, "pour", price, "or.")
			removeMoney(&p1, price)
			afficheComposants(p1)
		}
	}
}
