package github_flavored_markdown_test

import (
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/shurcooL/github_flavored_markdown"
	"github.com/shurcooL/github_flavored_markdown/gfmstyle"
)

func ExampleMarkdown() {
	text := []byte("Hello world github/linguist#1 **cool**, and #1!")

	os.Stdout.Write(github_flavored_markdown.Markdown(text))

	// Output:
	// <p>Hello world github/linguist#1 <strong>cool</strong>, and #1!</p>
}

// An example of how to generate a complete HTML page, including CSS styles.
func ExampleMarkdown_completeHTMLPage() {
	// Serve the "/assets/gfm.css" file.
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(gfmstyle.Assets)))

	var w io.Writer = os.Stdout // It can be an http.ResponseWriter.
	markdown := []byte("# GitHub Flavored Markdown\n\nHello.")

	io.WriteString(w, `<html><head><meta charset="utf-8"><link href="/assets/gfm.css" media="all" rel="stylesheet" type="text/css" /><link href="//cdnjs.cloudflare.com/ajax/libs/octicons/2.1.2/octicons.css" media="all" rel="stylesheet" type="text/css" /></head><body><article class="markdown-body entry-content" style="padding: 30px;">`)
	w.Write(github_flavored_markdown.Markdown(markdown))
	io.WriteString(w, `</article></body></html>`)

	// Output:
	// <html><head><meta charset="utf-8"><link href="/assets/gfm.css" media="all" rel="stylesheet" type="text/css" /><link href="//cdnjs.cloudflare.com/ajax/libs/octicons/2.1.2/octicons.css" media="all" rel="stylesheet" type="text/css" /></head><body><article class="markdown-body entry-content" style="padding: 30px;"><h1><a name="github-flavored-markdown" class="anchor" href="#github-flavored-markdown" rel="nofollow" aria-hidden="true"><span class="octicon octicon-link"></span></a>GitHub Flavored Markdown</h1>
	//
	// <p>Hello.</p>
	// </article></body></html>
}

func TestComponents(t *testing.T) {
	tests := []struct {
		text string
		want string
	}{
		{
			// Header.
			text: "## git diff",
			want: `<h2><a name="git-diff" class="anchor" href="#git-diff" rel="nofollow" aria-hidden="true"><span class="octicon octicon-link"></span></a>git diff</h2>` + "\n",
		},
		{
			// Header Link.
			text: "### [Some **bold** _italic_ link](http://www.example.com)",
			want: `<h3><a name="some-bold-italic-link" class="anchor" href="#some-bold-italic-link" rel="nofollow" aria-hidden="true"><span class="octicon octicon-link"></span></a><a href="http://www.example.com" rel="nofollow">Some <strong>bold</strong> <em>italic</em> link</a></h3>` + "\n",
		},
		{
			// Task List.
			text: `- [ ] This is an incomplete task.
- [x] This is done.
`,
			want: `<ul>
<li><input type="checkbox" disabled=""> This is an incomplete task.</li>
<li><input type="checkbox" checked="" disabled=""> This is done.</li>
</ul>
`,
		},
	}

	for _, test := range tests {
		if got := string(github_flavored_markdown.Markdown([]byte(test.text))); got != test.want {
			t.Errorf("\ngot %q\nwant %q", got, test.want)
		}
	}
}
