package usecase

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"fmt"
)

func createfile(filename string) {
	err := ioutil.WriteFile(filename, []byte{}, 0666)
	if err != nil {
		panic(err)
	}
}

func GetList(user string, dir string, write, all bool) error {
	var list []Repo
	filename := dir + user + ".md"
	if write {
		createfile(filename)
	}
	b, err := request(user)
	if err != nil {
		panic(err)
	}
	if !all {
		err = json.Unmarshal(b, &list)
	}
	var f *os.File

	if write {
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
		_, err = fmt.Fprintln(f, "# "+user+" github starred list")
		_, err = fmt.Fprintln(f, "")
		if err != nil {
			panic(err)
		}
	}
	defer f.Close()

	if all {
		if write {
			_, err = fmt.Fprintln(f, "```json")
			_, err = fmt.Fprintln(f, string(b))
			_, err = fmt.Fprintln(f, "```")
			if err != nil {
				panic(err)
			}
		}
		fmt.Println(string(b))
	} else {
		for _, repo := range list {
			str := "- [" + repo.FullName + "](" + repo.HtmlUrl + "): " + repo.Description
			if write {
				_, err = fmt.Fprintln(f, str)
				if err != nil {
					panic(err)
				}
			}
			fmt.Println(str)
		}
	}

	return err
}

type Repo struct {
	ID               int     `json:"id"`
	NodeId           string  `json:"node_id"`
	Name             string  `json:"name"`
	FullName         string  `json:"full_name"`
	Private          bool    `json:"private"`
	Owner            Owner   `json:"owner"`
	HtmlUrl          string  `json:"html_url"`
	Description      string  `json:"description"`
	Fork             bool    `json:"fork"`
	Url              string  `json:"url"`
	ForksUrl         string  `json:"forks_url"`
	KeysUrl          string  `json:"keys_url"`
	CollaboratorsUrl string  `json:"collaborators_url"`
	TeamsUrl         string  `json:"teams_url"`
	HooksUrl         string  `json:"hooks_url"`
	IssueEventsUrl   string  `json:"issue_events_url"`
	EventsUrl        string  `json:"events_url"`
	AssigneesUrl     string  `json:"assignees_url"`
	BranchesUrl      string  `json:"branches_url"`
	TagsUrl          string  `json:"tags_url"`
	BlobsUrl         string  `json:"blobs_url"`
	GitTagsUrl       string  `json:"git_tags_url"`
	GitRefsUrl       string  `json:"git_refs_url"`
	TreesUrl         string  `json:"trees_url"`
	StatusesUrl      string  `json:"statuses_url"`
	LanguagesUrl     string  `json:"languages_url"`
	StargazersUrl    string  `json:"stargazers_url"`
	ContributorsUrl  string  `json:"contributors_url"`
	SubscribersUrl   string  `json:"subscribers_url"`
	SubscriptionUrl  string  `json:"subscription_url"`
	CommitsUrl       string  `json:"commits_url"`
	GitCommitsUrl    string  `json:"git_commits_url"`
	CommentsUrl      string  `json:"comments_url"`
	IssueCommentUrl  string  `json:"issue_comment_url"`
	ContentsUrl      string  `json:"contents_url"`
	CompareUrl       string  `json:"compare_url"`
	MergesUrl        string  `json:"merges_url"`
	ArchiveUrl       string  `json:"archive_url"`
	DownloadsUrl     string  `json:"downloads_url"`
	IssuesUrl        string  `json:"issues_url"`
	PullsUrl         string  `json:"pulls_url"`
	MilestonesUrl    string  `json:"milestones_url"`
	NotificationsUrl string  `json:"notifications_url"`
	LabelsUrl        string  `json:"labels_url"`
	ReleasesUrl      string  `json:"releases_url"`
	DeploymentsUrl   string  `json:"deployments_url"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	PushedAt         string  `json:"pushed_at"`
	GitUrl           string  `json:"git_url"`
	SshUrl           string  `json:"ssh_url"`
	CloneUrl         string  `json:"clone_url"`
	SvnUrl           string  `json:"svn_url"`
	Homepage         string  `json:"homepage"`
	Size             int     `json:"size"`
	StargazersCount  int     `json:"stargazers_count"`
	WatchersCount    int     `json:"watchers_count"`
	Language         string  `json:"language"`
	HasIssues        bool    `json:"has_issues"`
	HasProjects      bool    `json:"has_projects"`
	HasDownloads     bool    `json:"has_downloads"`
	HasWiki          bool    `json:"has_wiki"`
	HasPages         bool    `json:"has_pages"`
	ForksCount       int     `json:"forks_count"`
	MirrorUrl        string  `json:"mirror_url"`
	Archived         bool    `json:"archived"`
	Disabled         bool    `json:"disabled"`
	OpenIssuesCount  int     `json:"open_issues_count"`
	License          License `json:"license"`
	Forks            int     `json:"forks"`
	OpenIssues       int     `json:"open_issues"`
	Watchers         int     `json:"watchers"`
	DefaultBranch    string  `json:"default_branch"`
}

type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxId string `json:"spdx_id"`
	Url    string `json:"url"`
	NodeId string `json:"node_id"`
}

type Owner struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

//request
func request(user string) ([]byte, error) {
	var body []byte
	response, err := http.Get("https://api.github.com/users/" + user + "/starred?per_page=99999")
	if err != nil {
		return body, err
	}
	//程序在使用完回复后必须关闭回复的主体。
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
