{
  Forge(name, username, tokenFile, repo, apiURL=''):: {
    name: name,
    username: username,
    tokenFile: tokenFile,
    repoName: repo,
    apiURL: apiURL,
  },
  GitHub(username, tokenFile, repo, apiURL='https://api.github.com'):: self.Forge('github', username, tokenFile, repo, apiURL),
  GitLab(username, tokenFile, repo, apiURL='https://gitlab.com/api/v4'):: self.Forge('gitlab', username, tokenFile, repo, apiURL),
  Gitea(username, tokenFile, repo, apiURL='https://gitea.com/api/v1'):: self.Forge('gitea', username, tokenFile, repo, apiURL),
  SourceHut(username, tokenFile, repo, apiURL='https://git.sr.ht/api'):: self.Forge('sourcehut', username, tokenFile, repo, apiURL),
  Codeberg(username, tokenFile, repo, apiURL='https://codeberg.org/api/v1'):: self.Forge('gitea', username, tokenFile, repo, apiURL),
}
