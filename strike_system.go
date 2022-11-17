package main

func CountDamage(damage int, penetration int, hp int, armor int) int {
	if penetration > armor {
		return hp - damage
	} else if penetration > (armor / 2) {
		return hp - (damage / 3)
	} else if penetration > (armor / 4) {
		return hp - (damage / 5)
	} else {
		return hp - (damage / 6)
	}
}
