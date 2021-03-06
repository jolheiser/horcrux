package config

func init() {
	defaultConfig = `port: "4000"
log_level: "info"

repositories:
  - name: "horcrux"
    gitea:
      repo_url: "https://gitea.com/jolheiser/horcrux.git"
      secret: "GiteaSecret"
      username: ""
      access_token: ""
    github:
      repo_url: "https://github.com/jolheiser/horcrux.git"
      secret: "GitHubSecret"
      username: ""
      access_token: ""
    gitlab:
      repo_url: "https://gitlab.com/jolheiser/horcrux.git"
      secret: "GitLabSecret"
      username: ""
      access_token: ""`
}
