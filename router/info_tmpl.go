package router

import "html/template"

func init() {
	infoTmpl = template.Must(template.New("info").Funcs(funcMap()).Parse(`<html lang="en">
<head>
    <title>{{.Name}}</title>
</head>
<body>
<p><a href="../">back</a></p>
<p><strong>Name:</strong> {{.Name}}</p>
<dl>
    {{if .Gitea}}
        <dt><strong>Gitea</strong></dt>
        <dd><a href="{{.RepoURL}}">{{.HumanURL}}</a></dd>
    {{end}}
    {{if .GitHub}}
        <dt><strong>GitHub</strong></dt>
        <dd><a href="{{.RepoURL}}">{{.HumanURL}}</a></dd>
    {{end}}
    {{if .GitLab}}
        <dt><strong>GitLab</strong></dt>
        <dd><a href="{{.RepoURL}}">{{.HumanURL}}</a></dd>
    {{end}}
</dl>
</body>
<footer>Horcrux Version: <pre>{{Version}}</pre></footer>
</html>`))
}
