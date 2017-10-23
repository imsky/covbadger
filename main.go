package main

import (
	"encoding/xml"
	"errors"
	"flag"
	"io/ioutil"
	"math"
	"os"
	"text/template"
)

//CoverageReport represents an individual coverage report
type CoverageReport struct {
	LineRate float64 `xml:"line-rate,attr"`
}

var _badgeTemplate string = `<svg xmlns="http://www.w3.org/2000/svg" width="96" height="20">
    <title>{{.LineRate}}%</title>
    <desc>Generated with covbadger (https://github.com/imsky/covbadger)</desc>
    <linearGradient id="smooth" x2="0" y2="100%">
        <stop offset="0" stop-color="#bbb" stop-opacity=".1" />
        <stop offset="1" stop-opacity=".1" />
    </linearGradient>
    <rect rx="3" width="96" height="20" fill="#555" />
    <rect rx="3" x="60" width="36" height="20" fill="#6ccb08" />
    <rect x="60" width="4" height="20" fill="#6ccb08" />
    <rect rx="3" width="96" height="20" fill="url(#smooth)" />
    <g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,sans-serif" font-size="11">
        <text x="30" y="15" fill="#010101" fill-opacity=".3">coverage</text>
        <text x="30" y="14">coverage</text>
        <text x="78" y="15" fill="#010101" fill-opacity=".3">{{.LineRate}}%</text>
        <text x="78" y="14">{{.LineRate}}%</text>
    </g>
</svg>`

func main() {
	flag.Parse()
	files := flag.Args()

	var coverageSum float64 = 0
	fileCount := 0
	badgeTemplate, err := template.New("badge").Parse(_badgeTemplate)

	if err != nil {
		panic(err)
	}

	for _, fileName := range files {
		var report CoverageReport

		in, err := ioutil.ReadFile(fileName)

		if err != nil {
			panic(err)
		}

		err = xml.Unmarshal(in, &report)

		if err != nil {
			panic(err)
		}

		coverageSum += report.LineRate
		fileCount += 1
	}

	if fileCount == 0 {
		panic(errors.New("No valid coverage reports provided"))
	}

	averageCoverage := coverageSum / float64(fileCount)
	aggregateReport := &CoverageReport{LineRate: math.Floor(averageCoverage * 100)}

	err = badgeTemplate.Execute(os.Stdout, aggregateReport)

	if err != nil {
		panic(err)
	}
}
