package vars

var (
	FR         = "fr"
	EN         = "en"
	NullString = ""

	Greetings = map[string]string{
		FR: "\nCe petit jeu est réalisé afin de rendre plus fun l'apprentissage du jargon de la finance et du trading. Amusez-vous bien !\n\nPour répondre à une question, entrez le numéro à gauche de la question (1, 2, 3, 4).",
		EN: "\nThis little game aims to make funnier the learning of the technical terminology of finance and trading. Have fun !\nTo answer a question, enter the number on the left of the question (1, 2, 3, 4).",
	}

	YourAnswerIsText = map[string]string{
		FR: "Votre réponse était :",
		EN: "Your answer is:",
	}

	PrefixAnswerText = map[string]string{
		FR: "La bonne réponse est :",
		EN: "The correct answer is:",
	}

	CorrectAnswerText = map[string]string{
		FR: "Bonne réponse ! ✅",
		EN: "Good answer ! ✅",
	}

	WrongAnswerText = map[string]string{
		FR: "Mauvaise réponse ! ❌\nLa bonne réponse était :",
		EN: "Wrong answer ! ❌\nThe correct answer was:",
	}
)
