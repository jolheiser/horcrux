package router

import "html/template"

func init() {
	infoTmpl = template.Must(template.New("info").Parse(`<html lang="en">
<head>
    <title>{{.Name}}</title>
</head>
<body>
<p><strong>Name:</strong> {{.Name}}</p>
<dl>
    {{if .Gitea}}
        <dt>Gitea</dt>
        {{range .Gitea}}
            <dd><a href="{{.RepoURL}}">{{.HumanURL}}</a></dd>
        {{end}}
    {{end}}
    {{if .GitHub}}
        <dt>GitHub</dt>
        {{range .GitHub}}
            <dd><a href="{{.RepoURL}}">{{.HumanURL}}</a></dd>
        {{end}}
    {{end}}
    {{if .GitLab}}
        <dt>GitLab</dt>
        {{range .GitLab}}
            <dd><a href="{{.RepoURL}}">{{.HumanURL}}</a></dd>
        {{end}}
    {{end}}
</dl>
</body>
</html>`))
}
