package main

import(
//  "database/sql"
//  _"github.com/go-sql-driver/mysql"
//  "log"
  "fmt"
  "context"
  "github.com/google/go-github/github"
)
// Function to fetch the repositories of the given user
func FetchRepos(username string)([]*github.Repository, error){
  client := github.NewClient(nil)
  repos, _, err := client.Repositories.List(context.Background(), username, nil)
  return repos, err
}

// Function to fetch the organi of the given user
func FetchOrganisations(username string)([]*github.Organization, error){
  client := github.NewClient(nil)
  orgs, _, err := client.Organizations.List(context.Background(), username, nil)
  return orgs, err
  }

func main() {
  var username string
  fmt.Print("Please enter your GitHub username: ")
  fmt.Scanf("%s", &username)

  // Access extractdb from MySql, the database where we are going to store the extracted information
  /*db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/extractdb")
  if err != nil {
    log.Fatal(err)
  }*/

  repos, err := FetchRepos(username)
  organisations, err := FetchOrganisations(username)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  // REPOSITORIES
  fmt.Printf("REPOSITORIES\n\n")
  for i, repo := range repos {
    fmt.Printf("ID: %v\n", repo.GetID()) // ID int64
    fmt.Printf("Name: %v\n", repo.GetName()) // Name string
    fmt.Printf("Description: %v\n", repo.GetDescription()) // Description string
    fmt.Printf("Private: %v\n", repo.GetPrivate()) // Private boolean
    fmt.Printf("Size: %v\n", repo.GetSize()) // Size int
    fmt.Printf("Language: %v\n\n", repo.GetLanguage()) // Language string
    i++
  }
  // ORGANISATIONS
  fmt.Printf("ORGANISATIONS\n\n")
  for i, organisation := range organisations {
    fmt.Printf("ID: %v\n", organisation.GetID()) // ID int64
    fmt.Printf("Name: %v\n", organisation.GetLogin()) // Name string
    fmt.Printf("Description: %v\n", organisation.GetDescription()) // Description string
    fmt.Printf("URL: %v\n\n", organisation.GetURL()) // URL string
    i++
  }
  //defer db.Close() // Close database

}
