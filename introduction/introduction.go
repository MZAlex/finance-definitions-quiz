package introduction

import (
	"bufio"
	"os"
	"strings"

	"finance-definitions-quiz/vars"
)

func Introduction() string {
	language := selectLanguage()
	displayGreetings(language)
	return language
}

func selectLanguage() string {
	println("Please type FR for french or EN for english")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	language := strings.ToLower(scanner.Text())
	if isLanguageCorrect(language) {
		return language
	}
	return selectLanguage()
}

func displayGreetings(language string) {
	println(vars.Greetings[language])
}

func isLanguageCorrect(language string) bool {
	switch language {
	case vars.FR:
		println("Langage sélectionné : français. Bienvenue !")
	case vars.EN:
		println("Selected language: English. Welcome !")
	default:
		return false
	}
	return true
}
