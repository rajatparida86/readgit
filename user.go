package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	gitApiURL    = "https://api.github.com"
	userEndPoint = "/users/"
)

type User struct {
	Login             string    `json:"login"`
	ID                int       `json:"id"`
	NodeID            string    `json:"node_id"`
	AvatarURL         string    `json:"avatar_url"`
	GravatarID        string    `json:"gravatar_id"`
	URL               string    `json:"url"`
	HTMLURL           string    `json:"html_url"`
	FollowersURL      string    `json:"followers_url"`
	FollowingURL      string    `json:"following_url"`
	GistsURL          string    `json:"gists_url"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
	OrganizationsURL  string    `json:"organizations_url"`
	ReposURL          string    `json:"repos_url"`
	EventsURL         string    `json:"events_url"`
	ReceivedEventsURL string    `json:"received_events_url"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"site_admin"`
	Name              string    `json:"name"`
	Company           string    `json:"company"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             string    `json:"email"`
	Hireable          bool      `json:"hireable"`
	Bio               string    `json:"bio"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Message           string    `json:"message"`
}

func getUserFromGit(userName string) User {
	user := User{}
	url := gitApiURL + userEndPoint + userName

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Unable to fetch from GIT: %s", err.Error())
	}

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Fatalf("Unable to decode response: %s", err.Error())
	}
	//Check for status code
	if resp.StatusCode != 200 {
		log.Fatalf("Failed to fetch from Github. Github returned with Status code: %v and Message: %s", resp.StatusCode, user.Message)
	}
	return user
}
