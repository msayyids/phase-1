package main

import "fmt"

type Hero struct {
	Name           int
	BaseAttack     int
	Defence        int
	CriticalDamage int
	HealtPoint     int
	weapon         Weapon
}

type Weapon struct {
	Attack int
}

func (h *Hero) CountDamage() int {
	DamageStandar := h.BaseAttack + h.weapon.Attack
	h.CriticalDamage = 7
	totalDamage := DamageStandar + h.CriticalDamage

	return totalDamage
}

func (h *Hero) IsAttackedBy(char int) {
	totalTakenDamage := h.CountDamage() - h.Defence

	if totalTakenDamage <= h.Defence {
		fmt.Println("no damage taken")
	} else {
		h.HealtPoint -= totalTakenDamage
	}

}

func battle(char1, char2 *Hero) {
	damageChar1 := char1.CountDamage()
	char2.IsAttackedBy(damageChar1)

	damageChar2 := char2.CountDamage()
	char1.IsAttackedBy(damageChar2)
}

func soal2() {
	hero := []Hero{
		{
			Name:           2,
			BaseAttack:     4,
			Defence:        4,
			CriticalDamage: 5,
			HealtPoint:     33,
			weapon:         Weapon{Attack: 10},
		},
		{
			Name:           3,
			BaseAttack:     4,
			Defence:        4,
			CriticalDamage: 5,
			HealtPoint:     40,
			weapon:         Weapon{Attack: 10},
		},
	}

	battle(&hero[0], &hero[1])

}
