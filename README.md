github_flavored_markdown
========================

[![Go Reference](https://pkg.go.dev/badge/github.com/shurcooL/github_flavored_markdown.svg)](https://pkg.go.dev/github.com/shurcooL/github_flavored_markdown)
[![Build Status](https://travis-ci.org/shurcooL/github_flavored_markdown.svg?branch=master)](https://travis-ci.org/shurcooL/github_flavored_markdown)

Package github_flavored_markdown provides a GitHub Flavored Markdown renderer
with fenced code block highlighting, clickable heading anchor links.

The functionality should be equivalent to the GitHub Markdown API endpoint specified at
https://developer.github.com/v3/markdown/#render-a-markdown-document-in-raw-mode, except
the rendering is performed locally.

See examples for how to generate a complete HTML page, including CSS styles.

Installation
------------

```bash
go get github.com/shurcooL/github_flavored_markdown
```

Directories
-----------

| Path                                                                                 | Synopsis                                                                     |
|--------------------------------------------------------------------------------------|------------------------------------------------------------------------------|
| [gfmstyle](https://pkg.go.dev/github.com/shurcooL/github_flavored_markdown/gfmstyle) | Package gfmstyle contains CSS styles for rendering GitHub Flavored Markdown. |

License
-------

-	[MIT License](LICENSE)
