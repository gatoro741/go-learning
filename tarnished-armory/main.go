package main

import (
	"fmt"
	"math/rand"
)

// Training interfaces , structs , slices , maps and using them all together

type Weapon interface {
	Attack() int
	Name() string
}

type Player struct {
	name      string
	inventory []Weapon
}
type Sword struct {
	baseDamage int
	name       string
}

type Katana struct {
	baseDamage int
	bleed      int
	name       string
}

type Staff struct {
	spellDamage int
	manacost    int
	name        string
}

func (s *Sword) Attack() int {
	return s.baseDamage

}

func (k *Katana) Attack() int {
	random := rand.Intn(2)
	if random == 1 {
		return k.baseDamage + k.bleed
	} else {
		return k.baseDamage
	}

}

func (stf *Staff) Attack() int {
	return stf.spellDamage

}

func (s *Sword) Name() string {
	return s.name

}

func (k *Katana) Name() string {
	return k.name

}

func (stf *Staff) Name() string {
	return stf.name

}

func (p *Player) Equip(weaponName string, arsenal map[string]Weapon) {
	if val, ok := arsenal[weaponName]; ok {
		p.inventory = append(p.inventory, val)
		fmt.Println("\nWeapon equiped successfully:\n", weaponName)
	} else {
		fmt.Println("\nWeapon isn't found\n")
	}
}

func (p *Player) ShowInventory() {
	var counter int
	fmt.Printf("\nInventory: %s\n", p.name)
	for _, item := range p.inventory {
		if item != nil {
			counter += 1
			fmt.Printf("%d. %s (Damage:%d)\n\n", counter, item.Name(), item.Attack())
		}
	}
}

func (p *Player) FullAtack(inventory []Weapon) {
	var counter int
	var counterDamage int
	var attackDamage int
	fmt.Printf("Full attack:\n")
	for _, item := range p.inventory {
		if item != nil {
			counter += 1
			attackDamage = item.Attack()
			fmt.Printf("%d. %s did %d damage.\n", counter, item.Name(), attackDamage)
			counterDamage += attackDamage

		}
	}
	fmt.Printf("Overall damage : %d\n\n", counterDamage)
}

func (p *Player) DeleteItem(weaponName string) {

	for i, value := range p.inventory {
		if value.Name() == weaponName {
			p.inventory = append(p.inventory[:i], p.inventory[i+1:]...)
		}
	}

}

func main() {
	mainArsenal := map[string]Weapon{
		"Claymore":        &Sword{baseDamage: 100, name: "Claymore"},
		"Uchigatana":      &Katana{baseDamage: 80, bleed: 35, name: "Uchigatana"},
		"Meteorite Staff": &Staff{spellDamage: 50, manacost: 20, name: "Meteorite Staff"},
	}
	gatoro := Player{
		name:      "gatoro",
		inventory: make([]Weapon, 0, 3),
	}

	gatoro.Equip("Claymore", mainArsenal)
	gatoro.Equip("Uchigatana", mainArsenal)
	gatoro.Equip("Meteorite Staff", mainArsenal)
	gatoro.ShowInventory()
	gatoro.FullAtack(gatoro.inventory)
	gatoro.DeleteItem("Claymore")
	gatoro.ShowInventory()
}
