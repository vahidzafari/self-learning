package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Pattern matching is a technique for searching a string for some set of characters
// based on a specific search pattern that is based on regular expressions and
// grammars.

// A regular expression is a sequence of characters that defines a search pattern. Every
// regular expression is compiled into a recognizer by building a generalized transition
// diagram called a finite automaton. A finite automaton can be either deterministic
// or nondeterministic. Nondeterministic means that more than one transition out of a
// state can be possible for the same input. A recognizer is a program that takes a string
// x as input and is able to tell whether x is a sentence of a given language.

// A grammar is a set of production rules for strings in a formal language—the
// production rules describe how to create strings from the alphabet of the language
// that are valid according to the syntax of the language. A grammar does not describe
// the meaning of a string or what can be done with it in whatever context—it only
// describes its form. What is important here is to realize that grammars are at the
// heart of regular expressions because without a grammar, you cannot define or use a
// regular expression.

// Expression      Description
// .	           Matches any character
// *               Means any number of times—cannot be used on its own
// ?               Zero or one time—cannot be used on its own
// +               Means one or more times—cannot be used on its own
// ^               This denotes the beginning of the line
// $               This denotes the end of the line
// []              [] is for grouping characters
// [A-Z]           This means all characters from capital A to capital Z
// \d              Any digit in 0-9
// \D              A non-digit
// \w              Any word character: [0-9A-Za-z_]
// \W              Any non-word character
// \s              A whitespace character
// \S              A non-whitespace character

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func matchRecord(s string) bool {
	fields := strings.Split(s, ",")
	if len(fields) != 3 {
		return false
	}
	if !matchNameSur(fields[0]) {
		return false
	}
	if !matchNameSur(fields[1]) {
		return false
	}
	return matchInt(fields[2])
}

func main() {
	mihhalis := "Mihalis"
	result := matchNameSur(mihhalis)
	fmt.Println("Result: ", result)

	positive := "+123.2"
	result = matchNameSur(positive)
	fmt.Println("Result: ", result)

	fileds := "Name,Surname,2109416471"
	result = matchRecord(fileds)
	fmt.Println("Result: ", result)
}
