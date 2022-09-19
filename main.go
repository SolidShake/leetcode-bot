package main

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

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

func main() {
	client := newGraphqlClient(graphql.NewClient("https://leetcode.com/graphql"), "SolidShake")

	resp, err := client.getStats()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	random, err := client.getRandomProblem()
	if err != nil {
		panic(err)
	}
	fmt.Println(random)
}

type graphqlClient struct {
	client   *graphql.Client
	username string
}

func newGraphqlClient(client *graphql.Client, username string) *graphqlClient {
	return &graphqlClient{
		client:   client,
		username: username,
	}
}

func (c *graphqlClient) getStats() (*SubmissionData, error) {
	request := graphql.NewRequest(`
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
	`)
	request.Var("username", c.username)

	var respData *SubmissionData
	if err := c.client.Run(context.Background(), request, &respData); err != nil {
		return nil, err
	}

	return respData, nil
}

func (c *graphqlClient) getRandomProblem() (*RandomQuestionData, error) {
	request := graphql.NewRequest(`
	query randomQuestion($categorySlug: String, $filters: QuestionListFilterInput) {
		randomQuestion(categorySlug: $categorySlug, filters: $filters) {
			titleSlug
		}
	}
	`)
	request.Var("categorySlug", "")
	request.Var("filters", struct{}{})

	var respData *RandomQuestionData
	if err := c.client.Run(context.Background(), request, &respData); err != nil {
		return nil, err
	}

	return respData, nil
}
