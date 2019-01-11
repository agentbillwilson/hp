hp
==

Estimates Hitpoints level from combat XP in Old School RuneScape

Building
--------

    go build

Usage
-----

    hp -h xp \
    [-a xp] [-s xp] [-d xp] [-r xp] \
    [-A level] [-S level] [-D level] [-R level]

hp estimates the Hitpoints level achieved from obtaining other combat levels in
Old School RuneScape. -a, -s, -d, and -r specify the Attack, Strength, Defence,
and Ranged experience points to begin with, respectively. -A, -S, -D, and -R
specify the target Attack, Strength, Defence, and Ranged levels, respectively.
exit status 1
