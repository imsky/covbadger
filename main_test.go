package main

import (
	"strings"
	"testing"
)

var expected string = `<svg xmlns="http://www.w3.org/2000/svg" width="96" height="20">
    <title>90%</title>
    <desc>Generated with covbadger (https://github.com/imsky/covbadger)</desc>
    <linearGradient id="smooth" x2="0" y2="100%">
        <stop offset="0" stop-color="#bbb" stop-opacity=".1" />
        <stop offset="1" stop-opacity=".1" />
    </linearGradient>
    <rect rx="3" width="96" height="20" fill="#555" />
    <rect rx="3" x="60" width="36" height="20" fill="#97ca00" />
    <rect x="60" width="4" height="20" fill="#97ca00" />
    <rect rx="3" width="96" height="20" fill="url(#smooth)" />
    <g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,sans-serif" font-size="11">
        <text x="30" y="15" fill="#010101" fill-opacity=".3">coverage</text>
        <text x="30" y="14">coverage</text>
        <text x="78" y="15" fill="#010101" fill-opacity=".3">90%</text>
        <text x="78" y="14">90%</text>
    </g>
</svg>`

func TestParseFilesToReports(t *testing.T) {
	files := []string{"test-report.xml"}
	reports := ParseFilesToReports(files)

	if len(reports) != len(files) {
		t.Errorf("Mismatch between parsed reports and input files")
	}

	if reports[0].LineRate != 0.9085 {
		t.Errorf("Expected line rate: 0.9085, actual line rate: %v", reports[0].LineRate)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Bad reports did not cause errors")
		}
	}()

	badFiles := []string{"xxx.go"}
	ParseFilesToReports(badFiles)
}

func TestRenderBadge(t *testing.T) {
	reports := []CoverageReport{CoverageReport{0.9}}
	badge := RenderBadge(reports)

	if badge != expected {
		t.Errorf("RenderBadge output is incorrect")
	}

	reports = []CoverageReport{CoverageReport{0.7}}
	badge = RenderBadge(reports)

	if strings.Contains(badge, "#dfb317") != true {
		t.Errorf("Incorrect color for coverage badge, expected yellow")
	}

	reports = []CoverageReport{CoverageReport{0.4}}
	badge = RenderBadge(reports)

	if strings.Contains(badge, "#e05d44") != true {
		t.Errorf("Incorrect color for coverage badge, expected red")
		t.Errorf(badge)
	}
}

func TestRun(t *testing.T) {
	run([]string{"test-report.xml"})
	run([]string{})
}
