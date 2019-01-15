hp
==

Estimates Hitpoints XP from combat XP in Old School RuneScape

Building
--------

    go build

Usage
-----

    hp [-n name] \
    [-h xp] [-a xp] [-s xp] [-d xp] [-r xp] \
    [-A level] [-S level] [-D level] [-R level]

hp estimates the Hitpoints experience achieved from obtaining other combat
levels in Old School RuneScape. -n specifies a player name to look up in the
Hiscores; otherwise, -h, -a, -s, -d, and -r specify the Hitpoints, Attack,
Strength, Defence, and Ranged experience points to begin with, respectively. -A,
-S, -D, and -R specify the target Attack, Strength, Defence, and Ranged levels,
respectively.
