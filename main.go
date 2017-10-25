package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"text/template"
)

//CoverageReport represents an individual coverage report
type CoverageReport struct {
	LineRate float64 `xml:"line-rate,attr"`
}

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
		return "", errors.New("Invalid coverage: " + string(coverage))
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

func GetCoverageFromReports(files []string) int {
	var sum float64 = 0
	reports := make([]CoverageReport, 0, len(files))

	for _, fileName := range files {
		var report CoverageReport

		in, err := ioutil.ReadFile(fileName)

		if err != nil {
			panic(err)
		}

		xml.Unmarshal(in, &report)
		sum += report.LineRate
		reports = append(reports, report)
	}

	return int(math.Floor(sum / float64(len(reports)) * 100))
}

func Run(files []string, coverage int) {
	var badge string

	if len(files) == 0 && coverage == 0 {
		flag.Usage()
		return
	} else if len(files) > 0 {
		badge, _ = RenderBadge(GetCoverageFromReports(files))
	} else if coverage > 0 {
		badge, _ = RenderBadge(coverage)
	}

	fmt.Println(badge)
}

func main() {
	var coverageFlag int

	flag.IntVar(&coverageFlag, "coverage", 0, "custom coverage value")
	flag.Parse()
	flag.Usage = func() {
		fmt.Println(`Usage: covbadger [files]`)
		flag.PrintDefaults()
	}

	Run(flag.Args(), coverageFlag)
}
