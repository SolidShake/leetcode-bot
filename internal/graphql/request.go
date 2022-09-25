package graphql

const profileStats = `
query getUserProfile($username: String!) {
	allQuestionsCount {
		difficulty
		count
	}
	matchedUser(username: $username) {
		submitStats {
			acSubmissionNum {
				difficulty
				count
				submissions
			}
		}
	}
}
`

const randomQuestion = `
query randomQuestion($categorySlug: String, $filters: QuestionListFilterInput) {
	randomQuestion(categorySlug: $categorySlug, filters: $filters) {
		titleSlug
	}
}
`
