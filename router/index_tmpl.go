package router

import "html/template"

func init() {
	indexTmpl = template.Must(template.New("index").Parse(`<html lang="en">
<head>
    <title>Horcrux</title>
</head>
<body>
<h1>Welcome to Horcrux</h1>
<p>Configured repositories:</p>
<ul>
    {{range .}}
        <li><a href="{{.Name}}">{{.Name}}</a></li>
    {{end}}
</ul>
</body>
</html>`))
}
