package main

import (
	"fmt"
	"log"
	"os"

	gql "github.com/SolidShake/leetcode-bot/internal/graphql"
	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("username")
	if username == "" {
		log.Fatal("Empty username env variable")
	}

	client := gql.NewClient(graphql.NewClient("https://leetcode.com/graphql"), username)

	resp, err := client.GetStats()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	// random, err := client.GetRandomProblem()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(random)
}
