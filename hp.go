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

const usage = `usage: hp -h xp \
[-a xp] [-s xp] [-d xp] [-r xp] \
[-A level] [-S level] [-D level] [-R level]

hp estimates the Hitpoints level achieved from obtaining other combat levels in
Old School RuneScape. -h, -a, -s, -d, and -r specify the Hitpoints, Attack,
Strength, Defence, and Ranged experience points to begin with, respectively.
-A, -S, -D, and -R specify the target Attack, Strength, Defence, and Ranged
levels, respectively.`

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
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
	}
	flag.Parse()
	if *hp == 0 {
		flag.Usage()
		os.Exit(1)
	}
	xpTotal := 0
	if xp := XP(*atkLvl) - *atk; xp > 0 {
		xpTotal += xp
	}
	if xp := XP(*strLvl) - *str; xp > 0 {
		xpTotal += xp
	}
	if xp := XP(*defLvl) - *def; xp > 0 {
		xpTotal += xp
	}
	if xp := XP(*rngLvl) - *rng; xp > 0 {
		xpTotal += xp
	}
	*hp += xpTotal * 3 / 4
	hpLvl := 0
	for lvl := 10; XP(lvl) <= *hp; lvl++ {
		hpLvl = lvl
	}
	fmt.Println(hpLvl)
}
