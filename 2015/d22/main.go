package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type SpellId int

const (
	MagicMissile SpellId = iota
	Drain

	// effects
	Shield
	Poison
	Recharge
)

var Spells = []SpellId{MagicMissile, Drain, Shield, Poison, Recharge}

var Spell = map[SpellId]int{
	MagicMissile: 53,
	Drain:        73,
	Shield:       113,
	Poison:       173,
	Recharge:     229,
}

type Effect struct {
	id   SpellId
	left int
}

type Boss struct {
	hp  int
	dmg int
}

func (b Boss) clone() *Boss {
	return &Boss{
		hp:  b.hp,
		dmg: b.dmg,
	}
}

type Player struct {
	hp        int
	mana      int
	armor     int
	effects   []Effect
	manaSpent int
}

func (p Player) availableSpells() []SpellId {
	spells := []SpellId{}
	for _, id := range Spells {
		if p.hasActiveEffect(id) {
			continue
		}
		if Spell[id] > p.mana {
			continue
		}
		spells = append(spells, id)
	}
	return spells
}

func (p Player) clone() *Player {
	effects := make([]Effect, len(p.effects))
	for i, e := range p.effects {
		effects[i] = Effect{e.id, e.left}
	}
	return &Player{
		hp:        p.hp,
		mana:      p.mana,
		armor:     p.armor,
		manaSpent: p.manaSpent,
		effects:   effects,
	}
}

func (p Player) hasActiveEffect(id SpellId) bool {
	return slices.ContainsFunc(p.effects,
		func(e Effect) bool {
			return e.id == id && e.left > 0
		})
}

func (p *Player) addEffect(id SpellId) {
	switch id {
	case Shield:
		p.effects = append(p.effects, Effect{Shield, 6})
	case Poison:
		p.effects = append(p.effects, Effect{Poison, 6})
	case Recharge:
		p.effects = append(p.effects, Effect{Recharge, 5})
	}
}

func (p *Player) clearEmptyEffects() {
	filtered := []Effect{}
	for _, effect := range p.effects {
		if effect.left <= 0 {
			if effect.id == Shield {
				p.armor = 0
			}
			continue
		}
		filtered = append(filtered, effect)
	}
	p.effects = filtered
}

func (p *Player) applyEffects(boss *Boss) {
	p.clearEmptyEffects()
	for i, ef := range p.effects {
		switch ef.id {
		case Shield:
			p.armor = 7
		case Poison:
			boss.hp -= 3
		case Recharge:
			p.mana += 101
		}
		p.effects[i].left -= 1
	}
}

func (p *Player) castSpell(spellId SpellId, boss *Boss) {
	if p.mana < Spell[spellId] {
		fmt.Println("Not enough mana")
		return
	}
	p.mana -= Spell[spellId]
	p.manaSpent += Spell[spellId]

	switch spellId {
	case MagicMissile:
		boss.hp -= 4
	case Drain:
		boss.hp -= 2
		p.hp += 2
	default:
		p.addEffect(spellId)
	}
}

func parser(lines []string) Boss {
	boss := Boss{}
	for _, line := range lines {
		if strings.HasPrefix(line, "Hit Points: ") {
			boss.hp, _ = strconv.Atoi(strings.TrimPrefix(line, "Hit Points: "))
		}
		if strings.HasPrefix(line, "Damage: ") {
			boss.dmg, _ = strconv.Atoi(strings.TrimPrefix(line, "Damage: "))
		}
	}
	return boss
}

func game(player *Player, boss *Boss, bestValue int, isHard bool) int {
	if isHard {
		player.hp -= 1
		if player.hp <= 0 {
			return bestValue
		}
	}

	player.applyEffects(boss)
	spells := player.availableSpells()
	if len(spells) == 0 {
		return bestValue
	}
	for _, spell := range spells {
		playerClone, bossClone := player.clone(), boss.clone()
		playerClone.castSpell(spell, bossClone)
		playerClone.applyEffects(bossClone)
		if bossClone.hp <= 0 {
			// Win
			return min(bestValue, playerClone.manaSpent)
		}
		playerClone.hp -= max(1, bossClone.dmg-playerClone.armor)
		if playerClone.hp <= 0 {
			// Lose
			continue
		}
		// Continue
		if playerClone.manaSpent > bestValue {
			continue
		}
		bestValue = min(bestValue, game(playerClone, bossClone, bestValue, isHard))
	}

	return bestValue
}

func PartOne(lines []string, player Player) int {
	boss := parser(lines)
	best := game(&player, &boss, math.MaxInt, false)
	return best
}

func PartTwo(lines []string, player Player) int {
	boss := parser(lines)
	best := game(&player, &boss, math.MaxInt, true)
	return best
}

func main() {
	lines := []string{"Hit Points: 13", "Damage: 8"}
	player := Player{hp: 10, mana: 250}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
		player = Player{hp: 50, mana: 500}
	}
	fmt.Println("PartOne: ", PartOne(lines, player))
	fmt.Println("PartTwo: ", PartTwo(lines, player))
}
