package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/adrg/frontmatter"
	"github.com/bek-shoyatbek/beksblog/interfaces"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

func PostHandler(sl interfaces.SlugReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post interfaces.Post
		post.Slug = r.PathValue("slug")
		postMarkdown, err := sl.Read(post.Slug)
		if err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}
		rest, err := frontmatter.Parse(strings.NewReader(postMarkdown), &post)
		if err != nil {
			http.Error(w, "Error parsing frontmatter", http.StatusInternalServerError)
			return
		}

		mdRenderer := goldmark.New(
			goldmark.WithExtensions(
				highlighting.NewHighlighting(
					highlighting.WithStyle("dracula"),
				),
			),
		)

		var buf bytes.Buffer
		err = mdRenderer.Convert(rest, &buf)
		if err != nil {
			http.Error(w, "Error converting markdown", http.StatusInternalServerError)
			return
		}

		tpl, err := template.ParseFiles("./public/post.gohtml")
		if err != nil {
			fmt.Printf("err %s", err)
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}
		post.Content = template.HTMLEscapeString(buf.String())
		err = tpl.Execute(w, post)
	}
}
