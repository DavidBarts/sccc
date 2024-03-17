# Where things live. If defined, must end with a slash.

PANDOC_DIR =
GO_DIR = /usr/local/go/bin/

all: README.md sccc

# Because I dislike Markdown (syntictically significant end-of-line
# whitespace, really?), I maintain the README for this project in
# reStructuredText.

%.md: %.rst
	$(PANDOC_DIR)pandoc -f rst -t markdown_strict -o $@ $<

sccc: main.go runenames.go sccc.go
	$(GO_DIR)go build
