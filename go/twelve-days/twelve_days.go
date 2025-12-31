package twelve

import (
	"fmt"
	"strings"
)

func Verse(i int) string {
	numberToVerse := []struct {
		word  string
		verse string
	}{
		{"first", "a Partridge in a Pear Tree"},
		{"second", "two Turtle Doves"},
		{"third", "three French Hens"},
		{"fourth", "four Calling Birds"},
		{"fifth", "five Gold Rings"},
		{"sixth", "six Geese-a-Laying"},
		{"seventh", "seven Swans-a-Swimming"},
		{"eighth", "eight Maids-a-Milking"},
		{"ninth", "nine Ladies Dancing"},
		{"tenth", "ten Lords-a-Leaping"},
		{"eleventh", "eleven Pipers Piping"},
		{"twelfth", "twelve Drummers Drumming"},
	}
	verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me:", numberToVerse[i-1].word)
	for index := i - 1; index >= 0; index-- {
		if index != 0 {
			verse += fmt.Sprintf(" %s,", numberToVerse[index].verse)
		} else if i != 1 {
			verse += fmt.Sprintf(" and %s.", numberToVerse[index].verse)
		} else {
			verse += fmt.Sprintf(" %s.", numberToVerse[index].verse)
		}
	}
	return verse
}

func Song() string {
	song := ""
	for i := range 12 {
		song += Verse(i+1) + "\n"
	}

	return strings.TrimSuffix(song, "\n")
}
