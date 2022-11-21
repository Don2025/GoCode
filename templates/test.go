package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

type Package struct {
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}

func main() {
	ctx := context.Background()
	token := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	oauthClient := oauth2.NewClient(ctx, token)
	// Create a new GitHub client using the OAuth2 client
	client := github.NewClient(oauthClient)
	username := "Don2025"
	/*
		// Get the list of repositories owned by a user
		repos, _, err := client.Repositories.List(ctx, "Don2025", nil)
		if err != nil {
			fmt.Println(err)
		}
		// Print the name of each repository
		for _, repo := range repos {
			fmt.Println(*repo.Name)
		}
		// Get specific repository Info
		repo, _, err := client.Repositories.Get(ctx, "Don2025", "Love")
		pack := &Package{
			FullName:    *repo.FullName,
			Description: *repo.Description,
			ForksCount:  *repo.ForksCount,
			StarsCount:  *repo.StargazersCount,
		}
		fmt.Printf("%+v\n", pack)
		// Get Readme.md content
		readme, _, err := client.Repositories.GetReadme(ctx, "Don2025", "Love", nil)
		if err != nil {
			fmt.Printf("Problem in getting readme information %v\n", err)
		}
		fmt.Println(readme.GetContent())
		// List followers of a user
		followers, _, err := client.Users.ListFollowers(ctx, "Don2025", nil)
		if err != nil {
			fmt.Println(err)
		}
		for _, follower := range followers {
			// fmt.Printf("%+v\n", follower)
			fmt.Printf("Name:%v, URL:%v\n", *follower.Login, *follower.HTMLURL)
		}
	*/
	/*
		// List followers of a user. Lists the people following the specified user.
		numFollower := 0
		var FollowersList []*github.User
		for page := 1; ; page++ {
			option := &github.ListOptions{Page: page, PerPage: 100}
			followers, _, err := client.Users.ListFollowers(ctx, username, option)
			if err != nil {
				fmt.Println(err)
			}
			if len(followers) == 0 {
				break
			}
			FollowersList = append(FollowersList, followers...)
			numFollower += len(followers)
		}

		fmt.Printf("%v is followed by %v users.\n", username, numFollower)
		for _, user := range FollowersList {
			fmt.Printf("Name:%v, URL:%v\n", *user.Login, *user.HTMLURL)
		}
	*/
	// List the people a user follows
	numFollowing := 0
	var FollowingList []*github.User
	for page := 1; ; page++ {
		option := &github.ListOptions{Page: page, PerPage: 100}
		following, _, err := client.Users.ListFollowing(ctx, username, option)
		if err != nil {
			fmt.Println(err)
		}
		if len(following) == 0 {
			break
		}
		// fmt.Println(page, len(following))
		FollowingList = append(FollowingList, following...)
		numFollowing += len(following)
	}
	fmt.Printf("%v follows %v users.\n", username, numFollowing)
	// Lists users who don't follow the specified user.
	var unfriendlyList []*github.User
	for _, user := range FollowingList {
		if flag, _, _ := client.Users.IsFollowing(ctx, *user.Login, username); !flag {
			fmt.Printf("Name:%v, URL:%v\n", *user.Login, *user.HTMLURL)
			unfriendlyList = append(unfriendlyList, user)
		}
	}
	fmt.Printf("Number of unfriendly users: %d", len(unfriendlyList))
}
