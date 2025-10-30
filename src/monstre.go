package src

import (
	"fmt"
)

type monstre struct {
	nom             string
	maxhp           int
	act_pv          int
	pts_attaque     int
	pts_attaque_spe int
	status          bool
}

func InitDragonRouge(dr monstre) monstre {
	dr.nom = "dragon rouge"
	dr.maxhp = 40
	dr.act_pv = 40
	dr.pts_attaque = 5
	dr.pts_attaque_spe = 10
	return dr
}

func infoMonstres(mob monstre) string {
	fmt.Println("Nom :", mob.nom)
	fmt.Println("pv :", mob.act_pv, "/", mob.maxhp)
	fmt.Println("point d'attaque", mob.pts_attaque)
	fmt.Println("point d'attaque spe", mob.pts_attaque_spe)
	return fmt.Sprintf("Nom : %s\nPv actuel : %d\nPv max : %d\nPoint d'attaque : %d\npoint d'attaque spe : %d\n", mob.nom, mob.act_pv, mob.maxhp, mob.pts_attaque, mob.pts_attaque_spe)
}

func InitLoup_Garou(lg monstre) monstre {
	lg.maxhp = 55
	lg.act_pv = 55
	lg.pts_attaque = 5
	lg.pts_attaque_spe = 10
	return lg
}
