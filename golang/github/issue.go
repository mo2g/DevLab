package main

/*
go run issue.go -u user -r repo -t token -o out_dir
explain: https://github.com/mo2g/OpenWrt-Action
go run issue.go -u mo2g -r OpenWrt-Action -t abcdefg -o /tmp/OpenWrt-Action
https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#get-an-issue
https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#list-repository-issues
https://docs.github.com/en/rest/issues/comments?apiVersion=2022-11-28#list-issue-comments
*/
import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "regexp"

    "github.com/cheggaaa/pb/v3"
)

type Issue struct {
    Number    int    `json:"number"`
    Title     string `json:"title"`
    User      User   `json:"user"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    Labels    []Label `json:"labels"`
    Body      string `json:"body"`
    Comments  int    `json:"comments"`
    CommentsURL string `json:"comments_url"`
}

type User struct {
    Login string `json:"login"`
}

type Label struct {
    Name string `json:"name"`
}

type Comment struct {
    User      User   `json:"user"`
    CreatedAt string `json:"created_at"`
    Body      string `json:"body"`
}

var (
    username    = flag.String("u", "", "github account username")
    repo        = flag.String("r", "", "github account repo")
    issueNumber = flag.String("i", "", "github account repo issue number")
    token       = flag.String("t", "", "github personal access token")
    outDir      = flag.String("o", "", "output directory")
)

func getAPIURL() string {
    if *issueNumber != "" {
        return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", *username, *repo, *issueNumber)
    }
    return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", *username, *repo)
}


// get fetches data from a URL with authentication
func getInfo(url string, token string, params map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
    req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	if params != nil {
		q := req.URL.Query()
		for key, value := range params {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body,_ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode > 300 {
		return nil, fmt.Errorf("error: %s", string(body))
	}

	return body, nil
}

// fetchAllComments fetches all comments for an issue and handles pagination
func fetchAllComments(url, token string) ([]Comment, error) {
	var allComments []Comment
	page := 1
	perPage := 100

	for {
		params := map[string]string{"page": fmt.Sprintf("%d", page), "per_page": fmt.Sprintf("%d", perPage)}
		data, err := getInfo(url, token, params)
		if err != nil {
			return nil, err
		}

		if len(data) == 0 {
			break // No more comments
		}

		var comments []Comment
		err = json.Unmarshal(data, &comments)
		if err != nil {
			return nil, err
		}

		allComments = append(allComments, comments...)
		page++

		if len(comments) < perPage {
			break // Reached the last page
		}
	}

	return allComments, nil
}

func fetchIssues(token string, params map[string]string) ([]Issue, error) {
    var issues []Issue
    data, err := getInfo(getAPIURL(), token, params)
     if err != nil {
        return issues, err
    }

    err = json.Unmarshal(data, &issues)
    if err != nil {
        return issues, err
    }

    return issues, nil
}

func toMarkdown(issue Issue, comments []Comment) string {
    markdown := fmt.Sprintf("# %s\n\n", issue.Title)
    markdown += fmt.Sprintf("user: %s\n", issue.User.Login)
    markdown += fmt.Sprintf("created_at: %s\n", issue.CreatedAt)
    markdown += fmt.Sprintf("updated_at: %s\n", issue.UpdatedAt)

    labels := "label: "
    for idx, label := range issue.Labels {
        if idx == len(issue.Labels)-1 {
            labels += label.Name + "\n"
        } else {
            labels += label.Name + ","
        }
    }
    markdown += labels + "\n"

    if issue.Body != "" {
        markdown += issue.Body + "\n\n"
    } else {
        fmt.Printf("issue %d has no body\n", issue.Number)
    }

    for _, c := range comments {
        markdown += fmt.Sprintf("# %s commented on %s\n\n", c.User.Login, c.CreatedAt)
        markdown += c.Body + "\n\n"
    }

    return markdown
}

func saveIssue(issues []Issue, token string, dir string) {

    bar := pb.StartNew(len(issues))
    for _, issue := range issues {
        bar.Increment()
        filename := fmt.Sprintf("%d_%s.md", issue.Number, issue.Title)
        filename = regexp.MustCompile(`[/:*?"<>|]`).ReplaceAllString(filename, " ")
        filename = fmt.Sprintf("%s/%s", dir, filename)

        // Check if file exists
		if _, err := os.Stat(filename); err == nil {
			fmt.Printf("Skipping issue %d, file already exists: %s\n", issue.Number, filename)
			continue
		}

        var comments []Comment
        if issue.Comments > 0 {
            comments, _ = fetchAllComments(issue.CommentsURL, token)
        }

        mdContent := toMarkdown(issue, comments)
        ioutil.WriteFile(filename, []byte(mdContent), 0644)
    }
 bar.Finish()
}

func main() {
    flag.Parse()

    if *username == "" || *repo == "" || *token == "" || *outDir == "" {
        fmt.Println("Missing required arguments")
        return
    }

    if _, err := os.Stat(*outDir); os.IsNotExist(err) {
        os.MkdirAll(*outDir, os.ModePerm)
    }

    page := 1
	perPage := 100

    for {
        params := map[string]string{
            "page": fmt.Sprintf("%d", page),
            "per_page": fmt.Sprintf("%d", perPage),
        }
        if *issueNumber != "" {
            params["state"] = "all"
        }

        issues, err := fetchIssues(*token, params)
        if err != nil {
            fmt.Println(err)
            return
        }

        if len(issues) == 0 {
			break // No more issues
		}

        page++

        saveIssue(issues, *token, *outDir)
    }


}