package router

import "html/template"

func init() {
	indexTmpl = template.Must(template.New("index").Funcs(funcMap()).Parse(`<html lang="en">
<head>
    <title>Horcrux</title>
</head>
<body>
<h1>Index</h1>
<p>Configured repositories:</p>
<ul>
    {{range .}}
        <li><a href="{{.Name}}">{{.Name}}</a></li>
    {{end}}
</ul>
<footer>Horcrux Version: <pre>{{Version}}</pre></footer>
</body>
</html>`))
}
