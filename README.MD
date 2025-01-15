# Git Configuration Tools

GitHubのIssue管理とGitの設定を簡単に行うためのツール集です。

## 概要
- GitHubにアサインされたIssueの一覧取得
- プロジェクト固有のGitユーザー設定
- 現在のGitユーザー情報の確認

## 使用方法

### GitHubアサインIssueの確認
```bash
# 環境変数を使用する場合
export GITHUB_TOKEN="your-github-token"
go run main.go

# 特定の組織のissueのみを表示
go run main.go -token="your-github-token" -org="organization-name"
```