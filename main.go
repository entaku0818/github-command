package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

func main() {
	// コマンドライン引数の設定
	token := flag.String("token", "", "GitHub personal access token")
	org := flag.String("org", "", "GitHub organization name (optional)")
	flag.Parse()

	// 環境変数からトークンを取得（コマンドライン引数が空の場合）
	if *token == "" {
		*token = os.Getenv("GITHUB_TOKEN")
		if *token == "" {
			log.Fatal("GitHub token is required. Set it via -token flag or GITHUB_TOKEN environment variable")
		}
	}

	// GitHubクライアントの初期化
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// 検索オプションの設定
	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	// 検索クエリの構築
	query := "is:open assignee:@me"
	if *org != "" {
		query += fmt.Sprintf(" org:%s", *org)
	}

	// issueの検索
	result, _, err := client.Search.Issues(ctx, query, opts)
	if err != nil {
		log.Fatalf("Error searching issues: %v", err)
	}

	// 結果の表示
	if len(result.Issues) == 0 {
		fmt.Println("No assigned issues found.")
		return
	}

	fmt.Printf("Found %d assigned issues:\n\n", result.GetTotal())
	for _, issue := range result.Issues {
		repo := *issue.RepositoryURL
		repoName := repo[strings.LastIndex(repo, "/")+1:]

		fmt.Printf("Repository: %s\n", repoName)
		fmt.Printf("Title: %s\n", *issue.Title)
		fmt.Printf("URL: %s\n", *issue.HTMLURL)
		fmt.Printf("State: %s\n", *issue.State)
		if issue.Labels != nil && len(issue.Labels) > 0 {
			labels := make([]string, len(issue.Labels))
			for i, label := range issue.Labels {
				labels[i] = *label.Name
			}
			fmt.Printf("Labels: %s\n", strings.Join(labels, ", "))
		}
		fmt.Println("---")
	}
}
