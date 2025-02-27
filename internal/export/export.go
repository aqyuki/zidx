package export

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"path"
	"text/template"

	"github.com/aqyuki/zidx/internal/article"
)

const tocTemplate = `# Table of Contents
{{ range . }}
- [{{ .Emoji }} {{ .Title }}]({{ .URL }}){{ end }}
`

type Content struct {
	URL   string
	Title string
	Emoji string
}

func ConvertAndTrim(metas []article.Meta) []Content {
	var contents []Content
	for _, meta := range metas {
		if !meta.IsPublished {
			continue
		}

		u := url.URL{
			Scheme: "https",
			Host:   "zenn.dev",
			Path:   path.Join(meta.Username, "articles", meta.Slug),
		}

		contents = append(contents, Content{
			URL:   u.String(),
			Title: meta.Title,
			Emoji: meta.Emoji,
		})
	}
	return contents
}

func Generate(metas []Content) (io.Reader, error) {
	tmpl, err := template.New("toc").Parse(tocTemplate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	w := new(bytes.Buffer)
	if err := tmpl.Execute(w, metas); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}
	return w, nil
}
