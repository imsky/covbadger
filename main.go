package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"strconv"
	"text/template"
)

//Badge represents a coverage badge
type Badge struct {
	Coverage int
	Color    string
}

var colors = map[string]string{
	"brightgreen": "#44cc11",
	"green":       "#97ca00",
	"yellow":      "#dfb317",
	"orange":      "#fe7d37",
	"red":         "#e05d44",
}

var _badgeTemplate string = `<svg xmlns="http://www.w3.org/2000/svg" width="96" height="20">
    <title>{{.Coverage}}</title>
    <desc>Generated with covbadger (https://github.com/imsky/covbadger)</desc>
    <linearGradient id="smooth" x2="0" y2="100%">
        <stop offset="0" stop-color="#bbb" stop-opacity=".1" />
        <stop offset="1" stop-opacity=".1" />
    </linearGradient>
    <rect rx="3" width="96" height="20" fill="#555" />
    <rect rx="3" x="60" width="36" height="20" fill="{{.Color}}" />
    <rect x="60" width="4" height="20" fill="{{.Color}}" />
    <rect rx="3" width="96" height="20" fill="url(#smooth)" />
    <g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,sans-serif" font-size="11">
        <text x="30" y="15" fill="#010101" fill-opacity=".3">coverage</text>
        <text x="30" y="14">coverage</text>
        <text x="78" y="15" fill="#010101" fill-opacity=".3">{{.Coverage}}%</text>
        <text x="78" y="14">{{.Coverage}}%</text>
    </g>
</svg>`

func RenderBadge(coverage int) (string, error) {
	if coverage < 0 || coverage > 100 {
		return "", errors.New("Invalid coverage: " + strconv.Itoa(coverage))
	}

	var buffer bytes.Buffer
	badgeTemplate, _ := template.New("badge").Parse(_badgeTemplate)

	color := colors["red"]

	if coverage > 95 {
		color = colors["brightgreen"]
	} else if coverage > 80 {
		color = colors["green"]
	} else if coverage > 60 {
		color = colors["yellow"]
	} else if coverage > 40 {
		color = colors["orange"]
	}

	_ = badgeTemplate.Execute(&buffer, &Badge{coverage, color})
	return buffer.String(), nil
}

func Run(args []string) {
	if len(args) != 1 {
		flag.Usage()
		return
	}

	coverage, _ := strconv.Atoi(args[0])
	badge, err := RenderBadge(coverage)

	if err != nil {
		panic(err)
	} else {
		fmt.Println(badge)
	}
}

func main() {
	flag.Usage = func() {
		fmt.Println(`Usage: covbadger [coverage]`)
	}

	flag.Parse()
	Run(flag.Args())
}
