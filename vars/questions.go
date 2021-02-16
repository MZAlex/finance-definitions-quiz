package vars

type QuestionAnswersStruct struct {
	Question        string
	PossibleAnswers [4]string
	CorrectAnswer   string
}

var (
	QuestionsAnswersEnglish = []QuestionAnswersStruct{
		{
			Question:        "What is an acquisition ?",
			PossibleAnswers: [4]string{"When the stocks's prices are raising", "When a company acquires new equipment", "When one company decides to take over another one", "When a government nationalize a company"},
			CorrectAnswer:   "When one company decides to take over another one",
		},
		{
			Question:        "What is an American Depositary Receipt, or ADR ?",
			PossibleAnswers: [4]string{"An American Depositary Receipt (or ADR, for short) is a way in which US investors can trade shares of non-US companies without using their local exchanges.", "Possible answer 3", "Possible answer 4", "Right answer"},
			CorrectAnswer:   "An American Depositary Receipt (or ADR, for short) is a way in which US investors can trade shares of non-US companies without using their local exchanges.",
		},
	}
	QuestionsAnswersFrench = []QuestionAnswersStruct{
		{
			Question:        "Qu'est-ce qu'une acquisition ?",
			PossibleAnswers: [4]string{"C'est quand le prix des actions monte", "C'est lorsqu'une entreprise investit dans de l'équipement", "C'est lorsqu'une entreprise en rachète une autre", "C'est lorsque l'État nationalise une entreprise"},
			CorrectAnswer:   "C'est lorsqu'une entreprise en rachète une autre",
		},
		{
			Question:        "Voilà une question 2 ?",
			PossibleAnswers: [4]string{"Réponse possible 1", "Réponse possible 1", "Réponse correcte", "Réponse possible 1"},
			CorrectAnswer:   "Réponse correcte",
		},
	}
	CoreQuestionsAnswers = map[string][]QuestionAnswersStruct{
		FR: QuestionsAnswersFrench,
		EN: QuestionsAnswersEnglish,
	}
)
