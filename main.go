package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"flag"
	"io/ioutil"
	"math"
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

func RenderBadge(reports []CoverageReport) string {
	var buffer bytes.Buffer
	var coverageSum float64 = 0
	badgeTemplate, err := template.New("badge").Parse(_badgeTemplate)

	if err != nil {
		panic(err)
	}

	for _, report := range reports {
		coverageSum += report.LineRate
	}

	averageCoverage := coverageSum / float64(len(reports))
	aggregateReport := &CoverageReport{LineRate: math.Floor(averageCoverage * 100)}

	err = badgeTemplate.Execute(&buffer, aggregateReport)

	if err != nil {
		panic(err)
	}

	return buffer.String()
}

func ParseFilesToReports(files []string) []CoverageReport {
	reports := make([]CoverageReport, 0, len(files))
	i := 0

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

		reports = append(reports, report)
		i += 1
	}

	if i == 0 {
		panic(errors.New("No valid coverage reports provided"))
	}

	return reports
}

func main() {
	flag.Parse()
	files := flag.Args()
	RenderBadge(ParseFilesToReports(files))
}
