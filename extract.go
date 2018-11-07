package main

import(
  "fmt"
  "context"
  "github.com/google/go-github/github"
)
func FetchRepos(username string)([]*github.Repository, error){
  client := github.NewClient(nil)
  repos, _, err := client.Repositories.List(context.Background(), username, nil)
  return repos, err
}

func FetchOrganisations(username string)([]*github.Organization, error){
  client := github.NewClient(nil)
  orgs, _, err := client.Organizations.List(context.Background(), username, nil)
  return orgs, err
  }

func main() {
  var username string
  fmt.Print("Please enter your GitHub username: ")
  fmt.Scanf("%s", &username)

  repos, err := FetchRepos(username)
  organisations, err := FetchOrganisations(username)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }
  for i, repo := range repos {
    fmt.Printf("Repository Name: %v. %v\n", i+1, repo.GetName())
  }
  for i, organisation := range organisations {
    fmt.Printf("Organisation Name: %v. %v\n", i+1, organisation.GetLogin())
  }
//  fmt.Println("Hello")


}
