package graphql

import (
	"context"

	"github.com/machinebox/graphql"
)

type Client struct {
	client   *graphql.Client
	username string
}

func NewClient(client *graphql.Client, username string) *Client {
	return &Client{
		client:   client,
		username: username,
	}
}

func (c *Client) GetStats() (*SubmissionData, error) {
	request := graphql.NewRequest(profileStats)
	request.Var("username", c.username)

	var respData *SubmissionData
	if err := c.client.Run(context.Background(), request, &respData); err != nil {
		return nil, err
	}

	return respData, nil
}

func (c *Client) GetRandomProblem() (*RandomQuestionData, error) {
	request := graphql.NewRequest(randomQuestion)
	request.Var("categorySlug", "")
	request.Var("filters", struct{}{})

	var respData *RandomQuestionData
	if err := c.client.Run(context.Background(), request, &respData); err != nil {
		return nil, err
	}

	return respData, nil
}
