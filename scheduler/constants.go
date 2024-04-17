package scheduler

const DateFormat = "20060102"

var AllowedRepeatsMap = map[string]int{
	"d": 400,
	"w": 7,
	"m": 31,
	"y": 1,
}
