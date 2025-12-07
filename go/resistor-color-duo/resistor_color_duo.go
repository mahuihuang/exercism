package resistorcolorduo

var bandcolors = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	sum := 0
	for factor, i := 1, 0; i < len(colors); i++ {
		if colors[i] == "-" {
			continue
		}
		sum = factor*sum + bandcolors[colors[i]]
		factor *= 10
	}
	for sum > 99 {
		sum /= 10
	}
	return sum
}
