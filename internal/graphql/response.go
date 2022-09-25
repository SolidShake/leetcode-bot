package graphql

type SubmissionData struct {
	AllQuestionsCount []struct {
		Difficulty string `json:"difficulty"`
		Count      int    `json:"count"`
	} `json:"allQuestionsCount"`
	MatchedUser struct {
		SubmitStats struct {
			AcSubmissionNum []struct {
				Difficulty  string `json:"difficulty"`
				Count       int    `json:"count"`
				Submissions int    `json:"submissions"`
			} `json:"acSubmissionNum"`
		} `json:"submitStats"`
	} `json:"matchedUser"`
}

type RandomQuestionData struct {
	RandomQuestion struct {
		TitleSlug string `json:"titleSlug"`
	} `json:"randomQuestion"`
}
