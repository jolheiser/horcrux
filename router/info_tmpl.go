package router

import "html/template"

func init() {
	infoTmpl = template.Must(template.New("info").Parse(`<html lang="en">
<head>
    <title>{{.Name}}</title>
</head>
<body>
<p><a href="../">back</a></p>
<p><strong>Name:</strong> {{.Name}}</p>
<dl>
    {{if .Gitea}}
        <dt><strong>Gitea</strong></dt>
        {{range .Gitea}}
            <dd><a href="{{.RepoURL}}">{{.HumanURL}}</a></dd>
        {{end}}
    {{end}}
    {{if .GitHub}}
        <dt><strong>GitHub</strong></dt>
        {{range .GitHub}}
            <dd><a href="{{.RepoURL}}">{{.HumanURL}}</a></dd>
        {{end}}
    {{end}}
    {{if .GitLab}}
        <dt><strong>GitLab</strong></dt>
        {{range .GitLab}}
            <dd><a href="{{.RepoURL}}">{{.HumanURL}}</a></dd>
        {{end}}
    {{end}}
</dl>
</body>
</html>`))
}
