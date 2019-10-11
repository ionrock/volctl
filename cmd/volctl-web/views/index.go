package views

import (
	"html/template"
)

func Index() *template.Template {
	t := `<html>
  <head>
    <title>Volume Control</title>
  </head>
  <body>
    <h3>{{ CurrentVolume }}</h3>
  </body>
</html>
`
	return template.Must(template.New("index").Parse(t))
}
