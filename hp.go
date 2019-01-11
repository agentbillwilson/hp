package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

// XP returns the experience points necessary to achieve level.
func XP(lvl int) int {
	sum := 0
	for i := 1; i < lvl; i++ {
		sum += int(float64(i) + (300.0 * math.Pow(2.0, float64(i)/7.0)))
	}
	return sum / 4
}

var lvls = make([]int, 100)

func init() {
	for lvl := 1; lvl < 100; lvl++ {
		lvls[lvl] = XP(lvl)
	}
}

func main() {
	hp := flag.Int("h", 0, "Hitpoints XP")
	atk := flag.Int("a", 0, "Attack XP")
	str := flag.Int("s", 0, "Strength XP")
	def := flag.Int("d", 0, "Defence XP")
	rng := flag.Int("r", 0, "Ranged XP")
	atkLvl := flag.Int("A", 0, "target Attack level")
	strLvl := flag.Int("S", 0, "target Strength level")
	defLvl := flag.Int("D", 0, "target Defence level")
	rngLvl := flag.Int("R", 0, "target Ranged level")
	flag.Parse()
	if *hp == 0 {
		flag.Usage()
		os.Exit(1)
	}
	xpTotal := 0
	if *atkLvl != 0 {
		xpTotal += lvls[*atkLvl] - *atk
	}
	if *strLvl != 0 {
		xpTotal += lvls[*strLvl] - *str
	}
	if *defLvl != 0 {
		xpTotal += lvls[*defLvl] - *def
	}
	if *rngLvl != 0 {
		xpTotal += lvls[*rngLvl] - *rng
	}
	*hp += xpTotal * 3 / 4
	hpLvl := 0
	for lvl := 10; lvls[lvl] <= *hp; lvl++ {
		hpLvl = lvl
	}
	fmt.Println(hpLvl)
}
