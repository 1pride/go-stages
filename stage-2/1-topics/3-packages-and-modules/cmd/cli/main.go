package main

import (
	"random-number/internal/random"
)

/*
It's important to explain that code in a package can access and use all types, constants,
variables and functions within that package — even if they are declared in a different .go file.
*/

/*
Generally don't export things unless you actually have a reason to (i.e.,
don't capitalize a name just because it looks nicer!). Additionally,
a main package should never normally be imported by anything, so it probably shouldn't have any exported things in it.

ps: one package == one directory. That is, all .go files for a package should be contained in the same directory, and a directory should contain the .go files for one package only. You shouldn't ever have .go files with different package names in the same directory.

important: For all non-main packages, the directory name that the code lives in should be the same as the package name.
*/

// A module is... a tree of Go source files with a go.mod file in the tree's root directory.
func main() {
	random.Guess()
}
