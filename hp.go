package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type skill int

const (
	attack skill = iota + 1
	defence
	strength
	hitpoints
	ranged
)

// XP returns the experience points necessary to achieve level.
func XP(lvl int) int {
	sum := 0
	for i := 1; i < lvl; i++ {
		sum += int(float64(i) + (300.0 * math.Pow(2.0, float64(i)/7.0)))
	}
	return sum / 4
}

const usage = `usage: hp [-n name] \
[-h xp] [-a xp] [-s xp] [-d xp] [-r xp] \
[-A level] [-S level] [-D level] [-R level]

hp estimates the Hitpoints experience achieved from obtaining other combat
levels in Old School RuneScape. -n specifies a player name to look up in the
Hiscores; otherwise, -h, -a, -s, -d, and -r specify the Hitpoints, Attack,
Strength, Defence, and Ranged experience points to begin with, respectively. -A,
-S, -D, and -R specify the target Attack, Strength, Defence, and Ranged levels,
respectively.`

func main() {
	name := flag.String("n", "", "player name")
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
	if *name == "" && *hp == 0 {
		flag.Usage()
		os.Exit(1)
	}
	xp := map[skill]int{
		attack:    *atk,
		strength:  *str,
		defence:   *def,
		hitpoints: *hp,
		ranged:    *rng,
	}
	if *name != "" {
		u, err := url.Parse("https://secure.runescape.com/m=hiscore_oldschool/index_lite.ws")
		if err != nil {
			log.Fatal(err)
		}
		v := url.Values{}
		v.Set("player", *name)
		u.RawQuery = v.Encode()
		res, err := http.Get(u.String())
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		r := csv.NewReader(res.Body)
		if _, err := r.Read(); err != nil {
			log.Fatal(err)
		}
		for skill := attack; skill <= ranged; skill++ {
			record, err := r.Read()
			if err != nil {
				log.Fatal(err)
			}
			xp[skill], err = strconv.Atoi(record[2])
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	levels := map[skill]int{
		attack:   *atkLvl,
		defence:  *defLvl,
		strength: *strLvl,
		ranged:   *rngLvl,
	}
	combatXP := 0
	for skill := attack; skill <= ranged; skill++ {
		if skill == hitpoints {
			continue
		}
		lvl := levels[skill]
		if lvl == 0 {
			continue
		}
		if xp := XP(lvl) - xp[skill]; xp > 0 {
			combatXP += xp
		}
	}
	*hp = xp[hitpoints] + (combatXP * 3 / 4)
	fmt.Println(*hp)
}
