package core

import (
	"bufio"
	"finance-definitions-quiz/vars"
	"fmt"
	"os"
	"strconv"
)

func Core(language string) {
	var number int
	for _, elem := range vars.CoreQuestionsAnswers[language] {
		number = -1
		println(elem.Question)
		fmt.Printf("1. %s\n2. %s\n3. %s\n4. %s\n",
			elem.PossibleAnswers[0],
			elem.PossibleAnswers[1],
			elem.PossibleAnswers[2],
			elem.PossibleAnswers[3])
		for number < 0 || number > 4 {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			responseNumber, _ := strconv.Atoi(scanner.Text())
			number = responseNumber - 1
		}
		fmt.Printf("%s\n %s\n",
			vars.YourAnswerIsText[language],
			elem.PossibleAnswers[number])
		if elem.PossibleAnswers[number] == elem.CorrectAnswer {
			println(vars.CorrectAnswerText[language])
		} else {
			fmt.Printf("%s %s\n\n", vars.WrongAnswerText[language], elem.CorrectAnswer)
		}
	}
	println("Thanks for playing ! 😎")
}
