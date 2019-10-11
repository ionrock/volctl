package views

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type tmpl struct {
	templates map[string]*template.Template
}

func (t *tmpl) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		return errors.Errorf("Missing template %s", name)
	}
	return tmpl.Execute(w, data)
}

func (t *tmpl) Add(name, content string) {
	t.templates[name] = template.Must(template.New(name).Parse(content))
}

var Templates = &tmpl{
	templates: make(map[string]*template.Template),
}

func Renderer() echo.Renderer {
	Templates.Add("index", `<html>
  <head>
    <title>Volume Control</title>
  </head>
  <body>
    <div id="slider">
      <input type="range" min="0" max="100" value="{{ .CurrentVolume }}" id="vol-level">
    </div>
    <h3 id="vol-value">{{ .CurrentVolume }}</h3>
  </body>
  <script>
  var slider = document.getElementById("vol-level");

  slider.oninput = function () {
    document.getElementById("vol-value").innerHTML = this.value;
  }

  slider.onchange = function () {

    try {
      fetch("/update", {
        method: "POST",
        body: JSON.stringify({volume: this.value}),
        headers: {'Content-Type': 'application/json'}
      }).then(() => {
        document.getElementById("vol-value").innerHTML = this.value;
      })
    } catch (error) {
      console.error('Error:', error);
    }
  }
  </script>
</html>
`)

	return Templates
}
