package main

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

func generateJWT(AppID int, PrivateKeyPath string) (string, error) {
	privateKey, err := ioutil.ReadFile(PrivateKeyPath)
	if err != nil {
		return "", err
	}

	parsedKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 10).Unix(),
		"iss": AppID,
	})

	tokenString, err := token.SignedString(parsedKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getInstallationAccessToken(jwt string, InstallationID int64) (string, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: jwt},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	token, _, err := client.Apps.CreateInstallationToken(ctx, InstallationID, nil)
	if err != nil {
		return "", err
	}

	return token.GetToken(), nil
}

func githubClient() *github.Client {
	jwt, err := generateJWT(AppID, PrivateKeyPath)
	if err != nil {
		panic(err)
	}

	token, err := getInstallationAccessToken(jwt, InstallationID)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}
