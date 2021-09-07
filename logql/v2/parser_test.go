package v2

import (
	"strings"
	"testing"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/stretchr/testify/require"
)

func TestParseExpr(t *testing.T) { //nolint:paralleltest
	type tt struct {
		input            string
		expr             Expr
		err              error
		doNotcheckString bool
	}

	tc := []tt{
		// log selector expressions
		{
			input: `{foo="bar"}`,
			expr: &StreamMatcherExpr{
				matchers: []*labels.Matcher{
					{
						Type:  labels.MatchEqual,
						Name:  "foo",
						Value: "bar",
					},
				},
			},
		},
		{
			input: `{foo="bar", bar!="baz"}`,
			expr: &StreamMatcherExpr{
				matchers: []*labels.Matcher{
					{
						Type:  labels.MatchEqual,
						Name:  "foo",
						Value: "bar",
					},
					{
						Type:  labels.MatchNotEqual,
						Name:  "bar",
						Value: "baz",
					},
				},
			},
		},
		// log query expressions with filter
		{
			input: `{foo="bar"} |= "baz"`,
			expr: &LogQueryExpr{
				filter: LogPipelineExpr{
					{
						stages: LogFiltersExpr{
							{
								filter: "|=",
								value:  "baz",
							},
						},
					},
				},
				left: &StreamMatcherExpr{
					matchers: []*labels.Matcher{
						{
							Type:  labels.MatchEqual,
							Name:  "foo",
							Value: "bar",
						},
					},
				},
			},
		},
		{
			input: `{foo="bar"} |= "baz" |= ip("123.123.123.123")`,
			expr: &LogQueryExpr{
				filter: LogPipelineExpr{
					{
						stages: LogFiltersExpr{
							{
								filter: "|=",
								value:  "baz",
							},
							{
								filter:   "|=",
								filterOp: "ip",
								value:    `123.123.123.123`,
							},
						},
					},
				},
				left: &StreamMatcherExpr{
					matchers: []*labels.Matcher{
						{
							Type:  labels.MatchEqual,
							Name:  "foo",
							Value: "bar",
						},
					},
				},
			},
		},
		{
			input: `{foo="bar"} |= "baz" |~ "blip" != "flip" !~ "flap"`,
			expr: &LogQueryExpr{
				filter: LogPipelineExpr{
					{
						stages: LogFiltersExpr{
							{
								filter: "|=",
								value:  "baz",
							},
							{
								filter: "|~",
								value:  "blip",
							},
							{
								filter: "!=",
								value:  "flip",
							},
							{
								filter: "!~",
								value:  "flap",
							},
						},
					},
				},
				left: &StreamMatcherExpr{
					matchers: []*labels.Matcher{
						{
							Type:  labels.MatchEqual,
							Name:  "foo",
							Value: "bar",
						},
					},
				},
			},
		},
		// log query expressions with parsers
		{
			input: `{foo="bar"} | logfmt | addr=ip("1.2.3.4")`,
			expr: &LogQueryExpr{
				filter: LogPipelineExpr{
					{
						parser: "logfmt",
					},
					{
						stages: LogFiltersExpr{
							{
								filter:   "|",
								alias:    "addr",
								filterOp: "ip",
								value:    "1.2.3.4",
							},
						},
					},
				},
				left: &StreamMatcherExpr{
					matchers: []*labels.Matcher{
						{
							Type:  labels.MatchEqual,
							Name:  "foo",
							Value: "bar",
						},
					},
				},
			},
		},
		{
			input: `{foo="bar"} | json | addr=ip("1.2.3.4")`,
			expr: &LogQueryExpr{
				filter: LogPipelineExpr{
					{
						parser: "json",
					},
					{
						stages: LogFiltersExpr{
							{
								filter:   "|",
								alias:    "addr",
								filterOp: "ip",
								value:    "1.2.3.4",
							},
						},
					},
				},
				left: &StreamMatcherExpr{
					matchers: []*labels.Matcher{
						{
							Type:  labels.MatchEqual,
							Name:  "foo",
							Value: "bar",
						},
					},
				},
			},
		},
		{
			input: `{foo="bar"} | unpack | addr=ip("1.2.3.4")`,
			expr: &LogQueryExpr{
				filter: LogPipelineExpr{
					{
						parser: "unpack",
					},
					{
						stages: LogFiltersExpr{
							{
								filter:   "|",
								alias:    "addr",
								filterOp: "ip",
								value:    "1.2.3.4",
							},
						},
					},
				},
				left: &StreamMatcherExpr{
					matchers: []*labels.Matcher{
						{
							Type:  labels.MatchEqual,
							Name:  "foo",
							Value: "bar",
						},
					},
				},
			},
		},
		{
			input: `{foo="bar"} | regexp "(.)*" | addr=ip("1.2.3.4")`,
			expr: &LogQueryExpr{
				filter: LogPipelineExpr{
					{
						parser: "regexp",
						matcher: &LogFormatExpr{
							sep: "",
							kv:  map[string]string{"": "(.)*"},
						},
					},
					{
						stages: LogFiltersExpr{
							{
								filter:   "|",
								alias:    "addr",
								filterOp: "ip",
								value:    "1.2.3.4",
							},
						},
					},
				},
				left: &StreamMatcherExpr{
					matchers: []*labels.Matcher{
						{
							Type:  labels.MatchEqual,
							Name:  "foo",
							Value: "bar",
						},
					},
				},
			},
		},
		{
			input: `{foo="bar"} | pattern "(.)*" | addr=ip("1.2.3.4")`,
			expr: &LogQueryExpr{
				filter: LogPipelineExpr{
					{
						parser: "pattern",
						matcher: &LogFormatExpr{
							sep: "",
							kv:  map[string]string{"": "(.)*"},
						},
					},
					{
						stages: LogFiltersExpr{
							{
								filter:   "|",
								alias:    "addr",
								filterOp: "ip",
								value:    "1.2.3.4",
							},
						},
					},
				},
				left: &StreamMatcherExpr{
					matchers: []*labels.Matcher{
						{
							Type:  labels.MatchEqual,
							Name:  "foo",
							Value: "bar",
						},
					},
				},
			},
		},
		// log query expressions with format expressions
		{
			input: `{app="foo"} |= "bar" | json | line_format "blip{{ .foo }}blop {{.status_code}}" | label_format foo=bar,status_code="buzz{{.bar}}"`, //nolint:lll
			expr: &LogQueryExpr{
				filter: LogPipelineExpr{
					{
						stages: LogFiltersExpr{
							{
								filter: "|=",
								value:  "bar",
							},
						},
					},
					{
						parser: "json",
					},
					{
						parser: "line_format",
						matcher: &LogFormatExpr{
							sep: "",
							kv:  map[string]string{"": "blip{{ .foo }}blop {{.status_code}}"},
						},
					},
					{
						parser: "label_format",
					},
				},
				left: &StreamMatcherExpr{
					matchers: []*labels.Matcher{
						{
							Type:  labels.MatchEqual,
							Name:  "app",
							Value: "foo",
						},
					},
				},
			},
		},
		// log metric expressions
		{
			input: `rate({foo="bar"}[1m])`,
			expr: &LogMetricExpr{
				metricOp: "rate",
				left: &LogRangeQueryExpr{
					rng: `[1m]`,
					left: &LogQueryExpr{
						left: &StreamMatcherExpr{
							matchers: []*labels.Matcher{
								{
									Type:  labels.MatchEqual,
									Name:  "foo",
									Value: "bar",
								},
							},
						},
					},
				},
			},
		},
		{
			input: `sum(rate({foo="bar"}[1m]))`,
			expr: &LogMetricExpr{
				metricOp: "sum",
				Expr: &LogMetricExpr{
					metricOp: "rate",
					left: &LogRangeQueryExpr{
						rng: `[1m]`,
						left: &LogQueryExpr{

							left: &StreamMatcherExpr{
								matchers: []*labels.Matcher{
									{
										Type:  labels.MatchEqual,
										Name:  "foo",
										Value: "bar",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			input: `count_over_time({foo="bar"}[12h] |= "error")`,
			expr: &LogMetricExpr{
				metricOp: `count_over_time`,
				left: &LogRangeQueryExpr{
					rng: `[12h]`,
					left: &LogQueryExpr{
						filter: LogPipelineExpr{
							{
								stages: LogFiltersExpr{
									{
										filter: "|=",
										value:  "error",
									},
								},
							},
						},
						left: &StreamMatcherExpr{
							matchers: []*labels.Matcher{
								{
									Type:  labels.MatchEqual,
									Name:  "foo",
									Value: "bar",
								},
							},
						},
					},
				},
			},
		},
		{
			input: `count_over_time({foo="bar"} |= "error" [12h])`,
			expr: &LogMetricExpr{
				metricOp: `count_over_time`,
				left: &LogRangeQueryExpr{
					rng:     `[12h]`,
					rngLast: true,
					left: &LogQueryExpr{
						filter: LogPipelineExpr{
							{
								stages: LogFiltersExpr{
									{
										filter: "|=",
										value:  "error",
									},
								},
							},
						},
						left: &StreamMatcherExpr{
							matchers: []*labels.Matcher{
								{
									Type:  labels.MatchEqual,
									Name:  "foo",
									Value: "bar",
								},
							},
						},
					},
				},
			},
		},
		// metric expressions with groupings
		{
			input: `avg(count_over_time({foo="bar"}[5h])) by ()`,
			expr: &LogMetricExpr{
				metricOp: "avg",
				grouping: &grouping{without: false, groups: nil},
				Expr: &LogMetricExpr{
					metricOp: "count_over_time",
					left: &LogRangeQueryExpr{
						rng: "[5h]",
						left: &LogQueryExpr{
							left: &StreamMatcherExpr{
								matchers: []*labels.Matcher{
									{
										Type:  labels.MatchEqual,
										Name:  "foo",
										Value: "bar",
									},
								},
							},
						},
					},
				},
			},
		},
		// metric expressions with preamble
		{
			input: `topk(10,count_over_time({foo="bar"}[5h])) without(bar)`,
			expr: &LogMetricExpr{
				metricOp: "topk",
				preamble: "10",
				grouping: &grouping{without: true, groups: []string{"bar"}},
				Expr: &LogMetricExpr{
					metricOp: "count_over_time",
					left: &LogRangeQueryExpr{
						rng: "[5h]",
						left: &LogQueryExpr{
							left: &StreamMatcherExpr{
								matchers: []*labels.Matcher{
									{
										Type:  labels.MatchEqual,
										Name:  "foo",
										Value: "bar",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			input: `max without (bar) (count_over_time({foo="bar"}[5h]))`,
			expr: &LogMetricExpr{
				metricOp: "max",
				grouping: &grouping{without: true, groups: []string{"bar"}},
				Expr: &LogMetricExpr{
					metricOp: "count_over_time",
					left: &LogRangeQueryExpr{
						rng: "[5h]",
						left: &LogQueryExpr{
							left: &StreamMatcherExpr{
								matchers: []*labels.Matcher{
									{
										Type:  labels.MatchEqual,
										Name:  "foo",
										Value: "bar",
									},
								},
							},
						},
					},
				},
			},
			doNotcheckString: true,
		},
		// multi-line expressions
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
			expr: &LogMetricExpr{
				metricOp: "avg",
				grouping: &grouping{groups: []string{"bar", "foo"}},
				Expr: &LogMetricExpr{
					metricOp: "label_replace",
					params:   []string{"bar", "$1$2", "foo", "(.*).(.*)"},
					Expr: &LogMetricExpr{
						metricOp: "count_over_time",
						left: &LogRangeQueryExpr{
							rng: "[5h]",
							left: &LogQueryExpr{
								left: &StreamMatcherExpr{
									matchers: []*labels.Matcher{
										{
											Type:  labels.MatchEqual,
											Name:  "foo",
											Value: "bar",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tc := range tc { //nolint:paralleltest
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			expr, err := ParseExpr(tc.input)
			require.Equal(t, tc.err, err)
			require.Equal(t, tc.expr, expr)
			if !tc.doNotcheckString {
				require.Equal(t, trimInput(tc.input), expr.String())
			}
		})
	}
}

func trimInput(s string) string {
	if s == "" {
		return s
	}

	s = strings.ReplaceAll(s, "by ()", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\t", "")
	return strings.TrimSpace(s)
}
