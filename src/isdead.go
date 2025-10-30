package src

func isDead(p Character) bool {
	if p.act_pv <= 0 {
		return true
	}
	return false
}
