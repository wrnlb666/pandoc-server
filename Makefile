all: templ go

templ: html.templ
	templ generate

go: *
	go mod tidy
	go build
