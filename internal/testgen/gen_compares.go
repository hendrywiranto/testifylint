package main

import (
	"strings"
	"text/template"

	"github.com/Antonboom/testifylint/internal/checkers"
)

type ComparesTestsGenerator struct{}

func (ComparesTestsGenerator) Checker() checkers.Checker {
	return checkers.NewCompares()
}

func (g ComparesTestsGenerator) TemplateData() any {
	var (
		checker = g.Checker().Name()
		report  = checker + ": use %s.%s"
	)

	return struct {
		CheckerName       CheckerName
		InvalidAssertions []Assertion
		ValidAssertions   []Assertion
	}{
		CheckerName: CheckerName(checker),
		InvalidAssertions: []Assertion{
			{Fn: "True", Argsf: "a == b", ReportMsgf: report, ProposedFn: "Equal", ProposedArgsf: "a, b"},
			{Fn: "True", Argsf: "a != b", ReportMsgf: report, ProposedFn: "NotEqual", ProposedArgsf: "a, b"},
			{Fn: "True", Argsf: "a > b", ReportMsgf: report, ProposedFn: "Greater", ProposedArgsf: "a, b"},
			{Fn: "True", Argsf: "a >= b", ReportMsgf: report, ProposedFn: "GreaterOrEqual", ProposedArgsf: "a, b"},
			{Fn: "True", Argsf: "a < b", ReportMsgf: report, ProposedFn: "Less", ProposedArgsf: "a, b"},
			{Fn: "True", Argsf: "a <= b", ReportMsgf: report, ProposedFn: "LessOrEqual", ProposedArgsf: "a, b"},

			{Fn: "False", Argsf: "a == b", ReportMsgf: report, ProposedFn: "NotEqual", ProposedArgsf: "a, b"},
			{Fn: "False", Argsf: "a != b", ReportMsgf: report, ProposedFn: "Equal", ProposedArgsf: "a, b"},
			{Fn: "False", Argsf: "a > b", ReportMsgf: report, ProposedFn: "LessOrEqual", ProposedArgsf: "a, b"},
			{Fn: "False", Argsf: "a >= b", ReportMsgf: report, ProposedFn: "Less", ProposedArgsf: "a, b"},
			{Fn: "False", Argsf: "a < b", ReportMsgf: report, ProposedFn: "GreaterOrEqual", ProposedArgsf: "a, b"},
			{Fn: "False", Argsf: "a <= b", ReportMsgf: report, ProposedFn: "Greater", ProposedArgsf: "a, b"},
		},
		ValidAssertions: []Assertion{
			{Fn: "Equal", Argsf: "a, b"},
			{Fn: "NotEqual", Argsf: "a, b"},
			{Fn: "Greater", Argsf: "a, b"},
			{Fn: "GreaterOrEqual", Argsf: "a, b"},
			{Fn: "Less", Argsf: "a, b"},
			{Fn: "LessOrEqual", Argsf: "a, b"},
		},
	}
}

func (ComparesTestsGenerator) ErroredTemplate() Executor {
	return template.Must(template.New("ComparesTestsGenerator.ErroredTemplate").
		Funcs(fm).
		Parse(comparesTestTmpl))
}

func (ComparesTestsGenerator) GoldenTemplate() Executor {
	return template.Must(template.New("ComparesTestsGenerator.GoldenTemplate").
		Funcs(fm).
		Parse(strings.ReplaceAll(comparesTestTmpl, "NewAssertionExpander", "NewAssertionExpander.AsGolden")))
}

const comparesTestTmpl = header + `

package {{ .CheckerName.AsPkgName }}

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func {{ .CheckerName.AsTestName }}(t *testing.T) {
	var a, b int

	// Invalid.
	{
		{{- range $ai, $assrn := $.InvalidAssertions }}
			{{ NewAssertionExpander.Expand $assrn "assert" "t" nil }}
		{{- end }}
	}

	// Valid.
	{
		{{- range $ai, $assrn := $.ValidAssertions }}
			{{ NewAssertionExpander.Expand $assrn "assert" "t" nil }}
		{{- end }}
	}
}
`