package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

/**
"login": "rajatparida86",
"id": 11991018,
"node_id": "MDQ6VXNlcjExOTkxMDE4",
"avatar_url": "https://avatars0.githubusercontent.com/u/11991018?v=4",
"gravatar_id": "",
"url": "https://api.github.com/users/rajatparida86",
"html_url": "https://github.com/rajatparida86",
"followers_url": "https://api.github.com/users/rajatparida86/followers",
"following_url": "https://api.github.com/users/rajatparida86/following{/other_user}",
"gists_url": "https://api.github.com/users/rajatparida86/gists{/gist_id}",
"starred_url": "https://api.github.com/users/rajatparida86/starred{/owner}{/repo}",
"subscriptions_url": "https://api.github.com/users/rajatparida86/subscriptions",
"organizations_url": "https://api.github.com/users/rajatparida86/orgs",
"repos_url": "https://api.github.com/users/rajatparida86/repos",
"events_url": "https://api.github.com/users/rajatparida86/events{/privacy}",
"received_events_url": "https://api.github.com/users/rajatparida86/received_events",
"type": "User",
"site_admin": false,
"name": "Rajat Parida",
"company": null,
"blog": "",
"location": null,
"email": null,
"hireable": null,
"bio": null,
"public_repos": 14,
"public_gists": 1,
"followers": 0,
"following": 0,
"created_at": "2015-04-17T08:20:39Z",
"updated_at": "2019-01-01T21:51:22Z"
**/

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

	if resp, err := http.Get(url); err == nil {
		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			log.Fatalf("Unable to decode response: %s", err.Error())
		}
		//Check for status code
		if resp.StatusCode != 200 {
			log.Fatalf("Failed to fetch from Github. Github returned with Status code: %v and Message: %s", resp.StatusCode, user.Message)
		}
	} else {
		log.Fatalf("Unable to fetch from GIT: %s", err.Error())
	}
	return user
}
