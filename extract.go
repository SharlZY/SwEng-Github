package main

import(
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "log"
  "fmt"
  "context"
  "github.com/google/go-github/github"
)

// Function to fetch the followers of the given user
func FetchFollowers(username string)([]*github.User, error){
  client := github.NewClient(nil)
  followers, _, err := client.Users.ListFollowers(context.Background(), username, nil)
  return followers, err
}
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
  db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/extractdb")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

// REPOSITORIES
  stmtIns, err := db.Prepare("INSERT INTO repo (id, name, description, private, size, language) VALUES (?, ?, ?, ?, ?, ?)")
  if err != nil{
    log.Fatal(err)
  }
  defer stmtIns.Close()

  repos, err := FetchRepos(username)
  fmt.Printf("REPOSITORIES\n\n")
  for j, repo := range repos {
    fmt.Printf("ID: %v\n", repo.GetID()) // ID int64
    fmt.Printf("Name: %v\n", repo.GetName()) // Name string
    fmt.Printf("Description: %v\n", repo.GetDescription()) // Description string
    fmt.Printf("Private: %v\n", repo.GetPrivate()) // Private boolean
    fmt.Printf("Size: %v\n", repo.GetSize()) // Size int
    fmt.Printf("Language: %v\n\n", repo.GetLanguage()) // Language string
      _, err = stmtIns.Exec(repo.GetID(), repo.GetName(), repo.GetDescription(), repo.GetPrivate(), repo.GetSize(), repo.GetLanguage())
    j++
  }

// ORGANISATIONS
  organisations, err := FetchOrganisations(username)

  fmt.Printf("ORGANISATIONS\n\n")
  for k, organisation := range organisations {
    fmt.Printf("ID: %v\n", organisation.GetID()) // ID int64
    fmt.Printf("Name: %v\n", organisation.GetLogin()) // Name string
    fmt.Printf("Description: %v\n", organisation.GetDescription()) // Description string
    fmt.Printf("URL: %v\n\n", organisation.GetURL()) // URL string
    k++
  }


// FOLLOWERS
  followers, err := FetchFollowers(username)




  fmt.Printf("FOLLOWERS\n\n")
  for i, follower := range followers {
    fmt.Printf("ID: %v\n", follower.GetID()) // ID int64
    fmt.Printf("Name: %v\n", follower.GetLogin()) // Name string
    fmt.Printf("Repo URL: %v\n", follower.GetReposURL()) // URL string
    fmt.Printf("Type: %v\n", follower.GetType()) // Type string
    fmt.Printf("Site Admin: %v\n\n", follower.GetSiteAdmin()) // Site Admin boolean
    i++
  }

  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }

}
