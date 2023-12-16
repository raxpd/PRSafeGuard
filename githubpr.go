package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-github/v57/github"
)

func getPRDetails(client *github.Client, owner string, repo string, number int) string {
	pr, _, err := client.PullRequests.Get(context.Background(), owner, repo, number)
	if err != nil {
		panic(err)
	}

	files, _, err := client.PullRequests.ListFiles(context.Background(), owner, repo, number, nil)
	if err != nil {
		panic(err)
	}

	fileDetails := make([]map[string]interface{}, len(files))
	for i, file := range files {
		fileDetail := map[string]interface{}{
			"filename":  *file.Filename,
			"additions": *file.Additions,
			"deletions": *file.Deletions,
		}
		fileDetails[i] = fileDetail
	}

	compareCommitsRaw, _, err := client.Repositories.CompareCommitsRaw(context.Background(), owner, repo, *pr.Base.Ref, *pr.Head.Ref, github.RawOptions{Type: github.Diff})
	if err != nil {
		panic(err)
	}
	commitsDiff := string(compareCommitsRaw)

	prDescription := *pr.Body
	prTitle := *pr.Title
	prState := *pr.State
	prURL := *pr.HTMLURL
	prUser := *pr.User.Login
	prBranch := *pr.Head.Ref
	prBaseBranch := *pr.Base.Ref
	prDiffURL := *pr.DiffURL

	if prState == "closed" {
		return "skip"
	}
	if prDescription == "" {
		prDescription = "No description"
	}
	if prTitle == "" {
		prTitle = "No title"
	}

	prJSON := map[string]interface{}{
		"description": prDescription,
		"title":       prTitle,
		"state":       prState,
		"url":         prURL,
		"user":        prUser,
		"branch":      prBranch,
		"baseBranch":  prBaseBranch,
		"diffURL":     prDiffURL,
		"files":       fileDetails,
		"commitsDiff": commitsDiff,
	}

	jsonData, err := json.Marshal(prJSON)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(jsonData))

	return string(jsonData)
}

func AddLabelsToIssue(client *github.Client, owner string, repo string, number int, labels []string) error {
	_, _, err := client.Issues.AddLabelsToIssue(context.Background(), owner, repo, number, labels)
	if err != nil {
		return err
	}

	fmt.Println("Added labels to issue:", owner, repo, number)
	return nil
}

func RemoveLabelFromIssue(client *github.Client, owner string, repo string, number int, label string) error {
	_, err := client.Issues.RemoveLabelForIssue(context.Background(), owner, repo, number, label)
	if err != nil {
		return err
	}

	fmt.Println("Removed labels from issue:", owner, repo, number)
	return nil
}

func AddReviewersToPR(client *github.Client, owner string, repo string, number int, reviewers []string) error {
	_, _, err := client.PullRequests.RequestReviewers(context.Background(), owner, repo, number, github.ReviewersRequest{Reviewers: reviewers})
	if err != nil {
		return err
	}

	fmt.Println("Added reviewers to PR:", owner, repo, number)
	return nil
}

func AddCommntToPR(client *github.Client, owner string, repo string, number int, comment string) error {
	_, _, err := client.Issues.CreateComment(context.Background(), owner, repo, number, &github.IssueComment{
		Body: &comment,
	})
	if err != nil {
		return err
	}

	fmt.Println("Added comment to PR:", owner, repo, number)
	return nil
}

func handlePR(user string, repo string, prNumber int) (string, string, int, error) {
	fmt.Println("Webhook received for PR")
	fmt.Printf("Repo: %s, Owner: %s, PR Number: %d\n", repo, user, prNumber)

	return user, repo, prNumber, nil
}
