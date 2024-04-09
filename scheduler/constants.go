package scheduler

const DATE_FORMAT = "20060102"

var ALLOWED_REPEATS_MAP = map[string]int{
	"d": 400,
	"w": 7,
	"m": 31,
	"y": 1,
}
