package main

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"   // Library used for accessing mysql database
	"github.com/google/go-github/github" // Library used for accessing github api
)

// Commented out Database code, as data has already been entered into the database

// Function to fetch users that the given user is following
// func FetchFollowing(username string) ([]*github.User, error) {
// 	ctx := context.Background()
// 	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "insert access token here"})
// 	tc := oauth2.NewClient(ctx, ts)
// 	client := github.NewClient(tc)
// 	following, _, err := client.Users.ListFollowing(context.Background(), username, nil)
// 	return following, err
// }

// Function to fetch the followers of the given user
func FetchFollowers(username string) ([]*github.User, error) {
	client := github.NewClient(nil)
	followers, _, err := client.Users.ListFollowers(context.Background(), username, nil)
	return followers, err
}

// Function to fetch the repositories of the given user
func FetchRepos(username string) ([]*github.Repository, error) {
	client := github.NewClient(nil)
	repos, _, err := client.Repositories.List(context.Background(), username, nil)
	return repos, err
}

// Function to fetch the collaborators of certain repo
func FetchReposStats1(username, repo string) ([]*github.ContributorStats, error) {
	client := github.NewClient(nil)
	collabs, _, err := client.Repositories.ListContributorsStats(context.Background(), username, repo)
	return collabs, err
}

// Function to fetch the stats of certain repo
func FetchReposStats2(username, repo string) ([]*github.WeeklyStats, error) {
	client := github.NewClient(nil)
	act, _, err := client.Repositories.ListCodeFrequency(context.Background(), username, repo)
	return act, err
}

// Function to fetch the organisation of the given user
func FetchOrganisations(username string) ([]*github.Organization, error) {
	client := github.NewClient(nil)
	orgs, _, err := client.Organizations.List(context.Background(), username, nil)
	return orgs, err
}

func main() {
	var username string
	var reponame string
	fmt.Print("Please enter your GitHub username: ")
	fmt.Scanf("%s", &username)
	fmt.Print("Please enter a repo name: ")
	fmt.Scanf("%s", &reponame)

	// Access extractdb from MySql, the database where we are going to store the extracted information
	// For privacy reasons, I cannot leave my username and password for the database I used in the code
	// Replace user with username of database and password with the password to the database
	/*db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/extractdb")
		if err != nil {
			log.Fatal(err)
		}
	  defer db.Close()*/

	// FOLLOWING
	// followings, err := FetchFollowing(username)
	// fmt.Printf("FOLLOWING\n\n")
	// for m, following := range followings {
	// 	fmt.Printf("ID: %v\n", following.GetID())                  // ID int64
	// 	fmt.Printf("Name: %v\n", following.GetLogin())             // Name string
	// 	fmt.Printf("Repo URL: %v\n", following.GetReposURL())      // URL string
	// 	fmt.Printf("Type: %v\n", following.GetType())              // Type string
	// 	fmt.Printf("Site Admin: %v\n\n", following.GetSiteAdmin()) // Site Admin boolean
	// 	m++
	// }

	// REPOSITORIES
	/*  stmtIns, err := db.Prepare("INSERT INTO repo (id, name, description, private, size, language) VALUES (?, ?, ?, ?, ?, ?)")
	    if err != nil{
	      log.Fatal(err)
	    }
	    defer stmtIns.Close()*/

	repos, err := FetchRepos(username)
	fmt.Printf("REPOSITORIES\n\n")
	for j, repo := range repos {
		fmt.Printf("ID: %v\n", repo.GetID())                   // ID int64
		fmt.Printf("Name: %v\n", repo.GetName())               // Name string
		fmt.Printf("Description: %v\n", repo.GetDescription()) // Description string
		fmt.Printf("Private: %v\n", repo.GetPrivate())         // Private boolean
		fmt.Printf("Size: %v\n", repo.GetSize())               // Size int
		fmt.Printf("Language: %v\n\n", repo.GetLanguage())     // Language string
		//  _, err = stmtIns.Exec(repo.GetID(), repo.GetName(), repo.GetDescription(), repo.GetPrivate(), repo.GetSize(), repo.GetLanguage())
		/*    if err != nil {
		      log.Fatal(err)
		    }*/
		j++
	}

	// ORGANISATIONS
	/*  stmtIns1, err := db.Prepare("INSERT INTO org (id, login, description, url) VALUES (?, ?, ?, ?)")
	    if err != nil{
	    log.Fatal(err)
	    }
	    defer stmtIns1.Close()*/
	organisations, err := FetchOrganisations(username)

	fmt.Printf("ORGANISATIONS\n\n")
	for k, organisation := range organisations {
		fmt.Printf("ID: %v\n", organisation.GetID())                   // ID int64
		fmt.Printf("Name: %v\n", organisation.GetLogin())              // Name string
		fmt.Printf("Description: %v\n", organisation.GetDescription()) // Description string
		fmt.Printf("URL: %v\n\n", organisation.GetURL())               // URL string
		//  _, err = stmtIns1.Exec(organisation.GetID(), organisation.GetLogin(), organisation.GetDescription(), organisation.GetURL())
		/*    if err != nil {
		      log.Fatal(err)
		    }*/
		k++
	}

	collabs, err := FetchReposStats1(username, reponame)

	fmt.Printf("COLLABORATORS\n\n")
	for i, collab := range collabs {
		fmt.Printf("Author: %v\n", collab.GetAuthor()) //
		fmt.Printf("Total: %v\n", collab.GetTotal())   //
		/*      if err != nil {
		        log.Fatal(err)
		      }*/
		i++
	}

	acts, err := FetchReposStats2(username, reponame)

	fmt.Printf("ACTIVITY\n\n")
	for i, act := range acts {
		fmt.Printf("Additions: %v\n", act.GetAdditions()) //
		fmt.Printf("Commits: %v\n", act.GetCommits())
		fmt.Printf("Deletions: %v\n\n", act.GetDeletions()) // /
		/*      if err != nil {
		        log.Fatal(err)
		      }*/
		i++
	}
	// FOLLOWERS
	/*  stmtIns2, err := db.Prepare("INSERT INTO follower (id, login, repo, type, site_admin) VALUES (?, ?, ?, ?, ?)")
	    if err != nil{
	      log.Fatal(err)
	    }
	    defer stmtIns2.Close()*/
	followers, err := FetchFollowers(username)

	fmt.Printf("FOLLOWERS\n\n")
	for i, follower := range followers {
		fmt.Printf("ID: %v\n", follower.GetID())                  // ID int64
		fmt.Printf("Name: %v\n", follower.GetLogin())             // Name string
		fmt.Printf("Repo URL: %v\n", follower.GetReposURL())      // URL string
		fmt.Printf("Type: %v\n", follower.GetType())              // Type string
		fmt.Printf("Site Admin: %v\n\n", follower.GetSiteAdmin()) // Site Admin boolean
		//  _, err = stmtIns2.Exec(follower.GetID(), follower.GetLogin(), follower.GetReposURL(), follower.GetType(), follower.GetSiteAdmin())
		/*      if err != nil {
		        log.Fatal(err)
		      }*/
		i++
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	// Retrieving information from the database
	/*stmtOut, err := db.Prepare("SELECT name FROM repo WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmtOut.Close()
	var result string
	err = stmtOut.QueryRow(141727516).Scan(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Query Result: %s\n", result)*/
}
