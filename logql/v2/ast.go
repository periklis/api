// nolint:exhaustivestruct
package v2

import (
	"fmt"
	"strings"

	"github.com/prometheus/prometheus/pkg/labels"
)

type Expr interface {
	logQLExpr()
	fmt.Stringer
}

type LogSelectorExpr interface {
	Matchers() []*labels.Matcher
	Expr
}

type LogMetricSampleExpr interface {
	Selector() LogSelectorExpr
	Expr
}

type defaultLogQLExpr struct{}

func (defaultLogQLExpr) logQLExpr() {}

type StreamMatcherExpr struct {
	defaultLogQLExpr
	matchers []*labels.Matcher
}

func newStreamMatcherExpr(matchers []*labels.Matcher) *StreamMatcherExpr {
	return &StreamMatcherExpr{matchers: matchers}
}

func (sme *StreamMatcherExpr) Matchers() []*labels.Matcher {
	return sme.matchers
}

func (sme *StreamMatcherExpr) AppendMatchers(m []*labels.Matcher) {
	sme.matchers = append(sme.matchers, m...)
}

func (sme *StreamMatcherExpr) String() string {
	var sb strings.Builder

	sb.WriteString("{")

	for i, m := range sme.matchers {
		sb.WriteString(m.String())

		if i+1 != len(sme.matchers) {
			sb.WriteString(", ")
		}
	}

	sb.WriteString("}")

	return sb.String()
}

func mustNewLabelMatcher(t labels.MatchType, n, v string) *labels.Matcher {
	m, err := labels.NewMatcher(t, n, v)
	if err != nil {
		panic(err.Error())
	}

	return m
}

type LogFiltersExpr []LogFilterExpr

func (lf *LogFiltersExpr) String() string {
	var sb strings.Builder

	for i, l := range *lf {
		sb.WriteString(l.String())

		if i+1 != len(*lf) {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}

type LogFilterExpr struct {
	defaultLogQLExpr // nolint:unused
	filter           string
	alias            string
	filterOp         string
	value            string
}

func (LogFilterExpr) logQLExpr() {}

func newLogFilterExpr(filter, alias, filterOp, value string) LogFilterExpr {
	return LogFilterExpr{filter: filter, alias: alias, filterOp: filterOp, value: value}
}

func (lf *LogFilterExpr) String() string {
	var sb strings.Builder

	sb.WriteString(lf.filter)
	sb.WriteString(" ")

	if lf.filterOp != "" {
		if lf.alias != "" {
			sb.WriteString(lf.alias)
			sb.WriteString("=")
		}

		sb.WriteString(lf.filterOp)
		sb.WriteString("(")
		sb.WriteString(`"`)
		sb.WriteString(lf.value)
		sb.WriteString(`"`)
		sb.WriteString(")")
	} else {
		sb.WriteString(`"`)
		sb.WriteString(lf.value)
		sb.WriteString(`"`)
	}

	return sb.String()
}

type LogFormatExpr struct {
	defaultLogQLExpr // nolint:unused
	kv               map[string]string
	sep              string
}

func newLogFormatExpr(sep string, kv map[string]string) *LogFormatExpr {
	return &LogFormatExpr{sep: sep, kv: kv}
}

func (LogFormatExpr) logQLExpr() {}

func (l *LogFormatExpr) String() string {
	if l == nil {
		return ""
	}

	var (
		sb strings.Builder
		i  int
	)

	for key, value := range l.kv {
		if key != "" {
			sb.WriteString(key)
			sb.WriteString("=")
		} else {
			sb.WriteString(" ")
		}

		sb.WriteString(`"`)
		sb.WriteString(value)
		sb.WriteString(`"`)
		i++

		if i+1 != len(l.kv) {
			sb.WriteString(l.sep)
		}
	}

	return sb.String()
}

func mergeFormatMaps(lhs, rhs map[string]string) map[string]string {
	for _, _ = range lhs {
		for rk, rv := range rhs {
			lhs[rk] = rv
		}
	}

	return lhs
}

type LogPipelineExpr []LogPipelineStageExpr

func (LogPipelineExpr) logQLExpr() {}

func (l LogPipelineExpr) String() string {
	var sb strings.Builder

	for i, p := range l {
		sb.WriteString(p.String())

		if i+1 != len(l) {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}

type LogPipelineStageExpr struct {
	parser  string
	matcher *LogFormatExpr
	stages  LogFiltersExpr
}

func newLogPipelineStageExpr(parser string, matcher *LogFormatExpr, stage LogFiltersExpr) LogPipelineStageExpr {
	return LogPipelineStageExpr{parser: parser, matcher: matcher, stages: stage}
}

func (LogPipelineStageExpr) logQLExpr() {}

func (l *LogPipelineStageExpr) String() string {
	var sb strings.Builder

	if l.parser != "" || l.matcher != nil {
		sb.WriteString("| ")
		sb.WriteString(l.parser)
		sb.WriteString(l.matcher.String())
	}

	for i, stage := range l.stages {
		sb.WriteString(stage.String())

		if i+1 != len(l.stages) {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}

type LogQueryExpr struct {
	defaultLogQLExpr // nolint:unused
	left             *StreamMatcherExpr
	filter           LogPipelineExpr
	Expr
}

func newLogQueryExpr(m *StreamMatcherExpr, filter LogPipelineExpr) LogSelectorExpr {
	return &LogQueryExpr{left: m, filter: filter}
}

func (LogQueryExpr) logQLExpr() {}

func (le *LogQueryExpr) Matchers() []*labels.Matcher {
	return le.left.matchers
}

func (le *LogQueryExpr) String() string {
	var sb strings.Builder

	sb.WriteString(le.left.String())

	if le.filter != nil {
		sb.WriteString(" ")
		sb.WriteString(le.filter.String())
	}

	return sb.String()
}

type LogRangeQueryExpr struct {
	defaultLogQLExpr // nolint:unused
	left             LogSelectorExpr
	rng              string
	grouping         *grouping
	rngLast          bool
	Expr
}

func newLogRangeQueryExpr(m LogSelectorExpr, rng string, grouping *grouping, rngLast bool) LogSelectorExpr {
	return &LogRangeQueryExpr{left: m, rng: rng, grouping: grouping, rngLast: rngLast}
}

func (LogRangeQueryExpr) logQLExpr() {}

func (l *LogRangeQueryExpr) Matchers() []*labels.Matcher {
	return l.left.Matchers()
}

func (l *LogRangeQueryExpr) String() string {
	var sb strings.Builder

	if l.grouping != nil {
		sb.WriteString("(")
	}

	if l.rngLast {
		sb.WriteString(l.left.String())
		sb.WriteString(" ")
		sb.WriteString(l.rng)
	} else {
		sl := strings.Replace(l.left.String(), "}", fmt.Sprintf("}%s", l.rng), 1)
		sb.WriteString(sl)
	}

	if l.grouping != nil {
		sb.WriteString(") ")
		sb.WriteString(l.grouping.String())
	}

	return sb.String()
}

type LogMetricExpr struct {
	defaultLogQLExpr // nolint:unused
	left             LogSelectorExpr
	metricOp         string
	preamble         string
	grouping         *grouping
	params           []string
	Expr
}

func newLogMetricExpr(
	e LogMetricSampleExpr,
	m LogSelectorExpr,
	op, preamble string,
	grouping *grouping,
	params []string,
) LogMetricSampleExpr {
	return &LogMetricExpr{
		Expr:     e,
		left:     m,
		metricOp: op,
		preamble: preamble,
		grouping: grouping,
		params:   params,
	}
}

func (LogMetricExpr) logQLExpr() {}

func (lme *LogMetricExpr) Selector() LogSelectorExpr {
	return lme.left
}

func (lme *LogMetricExpr) String() string {
	var sb strings.Builder

	sb.WriteString(lme.metricOp)
	sb.WriteString("(")

	if lme.preamble != "" {
		sb.WriteString(lme.preamble)
		sb.WriteString(",")
	}

	if lme.Expr != nil {
		sb.WriteString(lme.Expr.String())
	} else {
		sb.WriteString(lme.left.String())
	}

	if lme.metricOp == OpLabelReplace {
		sb.WriteString(",")

		for i, p := range lme.params {
			sb.WriteString(`"`)
			sb.WriteString(p)
			sb.WriteString(`"`)

			if i+1 != len(lme.params) {
				sb.WriteString(",")
			}
		}
	}

	sb.WriteString(")")

	if lme.grouping != nil {
		sb.WriteString(lme.grouping.String())
	}

	return sb.String()
}

type grouping struct {
	without bool
	groups  []string
}

func (g grouping) String() string {
	var sb strings.Builder
	if g.without {
		sb.WriteString(" without")
	} else if len(g.groups) > 0 {
		sb.WriteString(" by")
	}

	if len(g.groups) > 0 {
		sb.WriteString("(")
		sb.WriteString(strings.Join(g.groups, ","))
		sb.WriteString(")")
	}

	return sb.String()
}
