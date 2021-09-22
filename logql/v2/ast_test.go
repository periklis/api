// nolint:testpackage,lll
package v2

import (
	"testing"

	"github.com/prometheus/prometheus/pkg/labels"
)

func Test_AstWalker_SimpleCountExpr(t *testing.T) {
	type tt struct {
		input string
		total int
	}

	tc := []tt{
		{
			input: "100 * 100",
			total: 3,
		},
	}
	for _, tc := range tc {
		tc := tc

		expr, err := ParseExpr(tc.input)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		total := 0

		expr.Walk(func(e interface{}) {
			total++
		})

		if total != tc.total {
			t.Fatalf("got: %d, want: %d", total, tc.total)
		}
	}
}

func Test_AstWalker_AppendMatcher(t *testing.T) {
	type tt struct {
		input  string
		output string
	}

	l := []*labels.Matcher{
		{
			Type:  labels.MatchEqual,
			Name:  "flip",
			Value: "flop",
		},
	}

	tc := []tt{
		// log number expressions
		{
			input:  "100 * 100",
			output: "(100.000000 * 100.000000)",
		},
		{
			input:  "100 * -100",
			output: "(100.000000 * -100.000000)",
		},
		// log selector expressions
		{
			input:  `{foo="bar"}`,
			output: `{foo="bar", flip="flop"}`,
		},
		{
			input:  `{foo="bar"} |= "baz" |= ip("123.123.123.123")`,
			output: `{foo="bar", flip="flop"} |= "baz" |= ip("123.123.123.123")`,
		},
		{
			input:  `{ foo = "bar" }|logfmt|addr>=ip("1.2.3.4")`,
			output: `{foo="bar", flip="flop"} | logfmt | addr>=ip("1.2.3.4")`,
		},
		// log metric expressions
		{
			input:  `sum(rate({foo="bar"}[5m]))`,
			output: `sum(rate({foo="bar", flip="flop"}[5m]))`,
		},
		{
			input:  `max without (bar) (count_over_time({foo="bar"}[5h]))`,
			output: `max(count_over_time({foo="bar", flip="flop"}[5h])) without(bar)`,
		},
		// log binary expressions
		{
			input:  `sum(rate({foo="bar"}[5m])) / sum(rate({foo="bar"}[5m]))`,
			output: `(sum(rate({foo="bar", flip="flop"}[5m])) / sum(rate({foo="bar", flip="flop"}[5m])))`,
		},
		{
			input: `sum by (job) (
							count_over_time({namespace="tns"} |= "level=error"[5m])
						/
							count_over_time({namespace="tns"}[5m])
						)  * 100`,
			output: `(sum((count_over_time({namespace="tns", flip="flop"} |= "level=error" [5m]) / count_over_time({namespace="tns", flip="flop"}[5m]))) by(job) * 100.000000)`,
		},
		// multiline expressions:
		{
			input: `avg(
					label_replace(
						count_over_time({foo="bar"}[5h]),
						"bar",
						"$1$2",
						"foo",
						"(.*).(.*)"
					)
				) by(bar,foo)`,
			output: `avg(label_replace(count_over_time({foo="bar", flip="flop"}[5h]),"bar","$1$2","foo","(.*).(.*)")) by(bar,foo)`,
		},
	}
	for _, tc := range tc {
		tc := tc

		expr, err := ParseExpr(tc.input)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		expr.Walk(func(e interface{}) {
			switch ex := e.(type) {
			case *StreamMatcherExpr:
				ex.AppendMatchers(l)
			}
		})

		got := expr.String()
		if got != tc.output {
			t.Fatalf("\ngot:  %s\nwant: %s", got, tc.output)
		}
	}
}
