package dndcharacter

import "math/rand"

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return (score+10)/2 - 10
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	minimum := rand.Intn(6) + 1
	score := minimum
	for i := 1; i < 4; i++ {
		v := rand.Intn(6) + 1
		score += v
		if v < minimum {
			minimum = v
		}
	}
	return score - minimum
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	c.Hitpoints = 10 + Modifier(c.Constitution)
	return c
}
