package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/google/go-github/v57/github"
)

func main() {
	http.HandleFunc("/", handleWebhook)
	http.ListenAndServe(":8000", nil)
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}
	payload, err := github.ValidatePayload(r, []byte(""))
	if err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		http.Error(w, "Could not parse webhook", http.StatusBadRequest)
		return
	}

	switch e := event.(type) {

	case *github.PullRequestEvent:
		fmt.Println("Unhandled event type:", reflect.TypeOf(e))
		user := e.GetRepo().GetOwner().GetLogin()
		repo := e.GetRepo().GetName()
		prNumber := e.GetPullRequest().GetNumber()

		handlePR(user, repo, prNumber)

		input := getPRDetails(githubClient(), user, repo, prNumber)
		switch investigation(input) {
		case true && input != "skip":
			AddLabelsToIssue(githubClient(), user, repo, prNumber, []string{"threataware review pass"})
		case false && input != "skip":
			AddLabelsToIssue(githubClient(), user, repo, prNumber, []string{"threataware review fail"})
			AddReviewersToPR(githubClient(), user, repo, prNumber, []string{securityResearchers})
			AddCommntToPR(githubClient(), user, repo, prNumber, botComment)
		}
	default:
		fmt.Println("Unhandled event type:", reflect.TypeOf(e))
	}
}
