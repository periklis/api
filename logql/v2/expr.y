%{
package v2

import (
  "github.com/prometheus/prometheus/pkg/labels"
)
%}

%union{
  Expr                 Expr
  LogFilterExpr        LogFilterExpr
  LogFiltersExpr       LogFiltersExpr
  LogFormatExpr        *LogFormatExpr
  LogQueryExpr         LogSelectorExpr
  LogMetricExpr        LogMetricSampleExpr
  LogPipelineExpr      LogPipelineExpr
  LogPipelineStageExpr LogPipelineStageExpr
  LogRangeQueryExpr    LogSelectorExpr
  Matcher              *labels.Matcher
  Matchers             []*labels.Matcher
  MetricOp             string
  Range                string
  Filter               string
  Selector             []*labels.Matcher
  Grouping             *grouping
  Labels               []string
  str                  string
  binOp                string
  cmpOp                string
}

%start root

%type <Expr>                 expr
%type <Filter>               filter
%type <LogFilterExpr>        logFilterExpr
%type <LogFiltersExpr>       logFiltersExpr
%type <LogFormatExpr>        logFormatExpr
%type <LogQueryExpr>         logQueryExpr
%type <LogMetricExpr>        logMetricExpr
%type <LogPipelineExpr>      logPipelineExpr
%type <LogPipelineStageExpr> logPipelineStageExpr
%type <LogRangeQueryExpr>    logRangeQueryExpr
%type <Matcher>              matcher
%type <Matchers>             matchers
%type <MetricOp>             metricOp
%type <Selector>             selector
%type <Labels>               labels
%type <Grouping>             grouping

%token  <str>      IDENTIFIER STRING RANGE NUMBER
%token  <val>      MATCHERS LABELS EQ RE NRE OPEN_BRACE CLOSE_BRACE OPEN_BRACKET CLOSE_BRACKET COMMA
                   OPEN_PARENTHESIS CLOSE_PARENTHESIS COUNT_OVER_TIME RATE SUM AVG MAX MIN COUNT STDDEV STDVAR BOTTOMK TOPK
                   BYTES_OVER_TIME BYTES_RATE BOOL JSON REGEXP LOGFMT PIPE_MATCH PIPE_EXACT PIPE LINE_FMT LABEL_FMT UNWRAP AVG_OVER_TIME SUM_OVER_TIME MIN_OVER_TIME
                   MAX_OVER_TIME STDVAR_OVER_TIME STDDEV_OVER_TIME QUANTILE_OVER_TIME FIRST_OVER_TIME LAST_OVER_TIME ABSENT_OVER_TIME
                   BY WITHOUT LABEL_REPLACE IP UNPACK PATTERN

/* Operators are listed with increasing precedence. */
%left <binOp> OR
%left <binOp> AND UNLESS
%left <binOp> CMP_EQ NEQ LT LTE GT GTE
%left <binOp> ADD SUB
%left <binOp> MUL DIV MOD
%right <binOp> POW
%%

root: expr { exprlex.(*parser).expr = $1 };

expr:
                logQueryExpr                                { $$ = $1 }
        |       logMetricExpr                               { $$ = $1 }
                ;

logQueryExpr:
                selector                                         { $$ = newStreamMatcherExpr($1) }
        |       selector logPipelineExpr                         { $$ = newLogQueryExpr(newStreamMatcherExpr($1), $2) }
        |       OPEN_PARENTHESIS logQueryExpr CLOSE_PARENTHESIS  { $$ = $2 }
                ;

logPipelineExpr:
                logPipelineStageExpr                 { $$ = LogPipelineExpr { $1 } }
        |       logPipelineExpr logPipelineStageExpr { $$ = append($1, $2)}
        ;

logPipelineStageExpr:
                logFiltersExpr               { $$ = newLogPipelineStageExpr("", nil, $1) }
        |       PIPE LOGFMT                  { $$ = newLogPipelineStageExpr("logfmt", nil, nil) }
        |       PIPE JSON                    { $$ = newLogPipelineStageExpr("json", nil, nil) }
        |       PIPE UNPACK                  { $$ = newLogPipelineStageExpr("unpack", nil, nil) }
        |       PIPE REGEXP STRING           { $$ = newLogPipelineStageExpr("regexp", newLogFormatExpr("", map[string]string{"": $3}), nil) }
        |       PIPE PATTERN STRING          { $$ = newLogPipelineStageExpr("pattern", newLogFormatExpr("", map[string]string{"": $3}), nil) }
        |       PIPE LINE_FMT STRING         { $$ = newLogPipelineStageExpr("line_format", newLogFormatExpr("", map[string]string{"": $3}), nil) }
        |       PIPE LABEL_FMT logFormatExpr { $$ = newLogPipelineStageExpr("label_format", $3, nil)}
        |       PIPE IDENTIFIER EQ IP OPEN_PARENTHESIS STRING CLOSE_PARENTHESIS { $$ = newLogPipelineStageExpr("", nil, LogFiltersExpr{newLogFilterExpr("|", $2, "ip", $6)}) }
                ;

logFiltersExpr:
                logFilterExpr                      { $$= LogFiltersExpr{ $1 } }
        |       logFiltersExpr logFilterExpr       { $$= append($1, $2)}
        ;

logFilterExpr:
                filter STRING                                                   { $$ = newLogFilterExpr($1, "", "", $2) }
        |       filter IP OPEN_PARENTHESIS STRING CLOSE_PARENTHESIS             { $$ = newLogFilterExpr($1, "", "ip", $4) }
                ;

logFormatExpr:
                IDENTIFIER EQ STRING                                       { $$ = newLogFormatExpr("", map[string]string{$1: $3}) }
        |       IDENTIFIER EQ IDENTIFIER                                   { $$ = newLogFormatExpr("", map[string]string{$1: $3}) }
        |       IDENTIFIER EQ IP OPEN_PARENTHESIS STRING CLOSE_PARENTHESIS { $$ = newLogFormatExpr("", map[string]string{$1: "ip("+$5+")"}) }
        |       logFormatExpr COMMA logFormatExpr                          { $$ = newLogFormatExpr(",", mergeFormatMaps($1.kv, $3.kv))}
        ;

logRangeQueryExpr:
                selector RANGE                                                             { $$ = newLogRangeQueryExpr(newLogQueryExpr(newStreamMatcherExpr($1), nil), $2, nil, false) }
        |       selector RANGE logPipelineExpr                                             { $$ = newLogRangeQueryExpr(newLogQueryExpr(newStreamMatcherExpr($1), $3), $2, nil, false) }
        |       selector logPipelineExpr RANGE                                             { $$ = newLogRangeQueryExpr(newLogQueryExpr(newStreamMatcherExpr($1), $2), $3, nil, true) }
        |       OPEN_PARENTHESIS selector RANGE CLOSE_PARENTHESIS                          { $$ = newLogRangeQueryExpr(newLogQueryExpr(newStreamMatcherExpr($2), nil), $3, nil, false) }
        |       OPEN_PARENTHESIS selector RANGE logPipelineExpr CLOSE_PARENTHESIS          { $$ = newLogRangeQueryExpr(newLogQueryExpr(newStreamMatcherExpr($2), $4), $3, nil, false) }
        |       OPEN_PARENTHESIS selector RANGE logPipelineExpr CLOSE_PARENTHESIS grouping { $$ = newLogRangeQueryExpr(newLogQueryExpr(newStreamMatcherExpr($2), $4), $3, $6, false) }
        |       logRangeQueryExpr error
        ;

logMetricExpr:
                metricOp OPEN_PARENTHESIS logRangeQueryExpr CLOSE_PARENTHESIS                                                      { $$ = newLogMetricExpr(nil, $3, $1, "", nil, nil)  }
        |       metricOp OPEN_PARENTHESIS NUMBER COMMA logRangeQueryExpr CLOSE_PARENTHESIS                                         { $$ = newLogMetricExpr(nil, $5, $1, $3, nil, nil) }
        |       metricOp OPEN_PARENTHESIS logRangeQueryExpr CLOSE_PARENTHESIS grouping                                             { $$ = newLogMetricExpr(nil, $3, "", "", $5, nil)  }
        |       metricOp OPEN_PARENTHESIS logMetricExpr CLOSE_PARENTHESIS grouping                                                 { $$ = newLogMetricExpr($3, nil, $1, "", $5, nil)  }
        |       metricOp OPEN_PARENTHESIS logMetricExpr CLOSE_PARENTHESIS                                                          { $$ = newLogMetricExpr($3, nil, $1, "", nil, nil)  }
        |       metricOp OPEN_PARENTHESIS NUMBER COMMA logMetricExpr CLOSE_PARENTHESIS                                             { $$ = newLogMetricExpr($5, nil, $1, $3, nil, nil) }
        |       metricOp OPEN_PARENTHESIS NUMBER COMMA logMetricExpr CLOSE_PARENTHESIS grouping                                    { $$ = newLogMetricExpr($5, nil, $1, $3, $7, nil) }
        |       metricOp grouping OPEN_PARENTHESIS NUMBER COMMA logMetricExpr CLOSE_PARENTHESIS                                    { $$ = newLogMetricExpr($6, nil, $1, $4, $2, nil) }
        |       metricOp grouping OPEN_PARENTHESIS logMetricExpr CLOSE_PARENTHESIS                                                 { $$ = newLogMetricExpr($4, nil, $1, "", $2, nil)  }
        |       LABEL_REPLACE OPEN_PARENTHESIS logMetricExpr COMMA STRING COMMA STRING COMMA STRING COMMA STRING CLOSE_PARENTHESIS { $$ = newLogMetricExpr($3, nil, OpLabelReplace, "", nil, []string{$5,$7,$9,$11}) }
        |       OPEN_PARENTHESIS logMetricExpr CLOSE_PARENTHESIS                                                                   { $$ = $2 }
                ;

selector:
                OPEN_BRACE matchers CLOSE_BRACE  { $$ = $2 }
        |       OPEN_BRACE matchers error        { $$ = $2 }
        |       OPEN_BRACE error CLOSE_BRACE     { }
                ;

matchers:
                matcher                          { $$ = []*labels.Matcher{ $1 } }
        |       matchers COMMA matcher           { $$ = append($1, $3) }
                ;

matcher:
                IDENTIFIER EQ STRING             { $$ = mustNewLabelMatcher(labels.MatchEqual, $1, $3) }
        |       IDENTIFIER NEQ STRING            { $$ = mustNewLabelMatcher(labels.MatchNotEqual, $1, $3) }
        |       IDENTIFIER RE STRING             { $$ = mustNewLabelMatcher(labels.MatchRegexp, $1, $3) }
        |       IDENTIFIER NRE STRING            { $$ = mustNewLabelMatcher(labels.MatchNotRegexp, $1, $3) }
                ;

metricOp:
                COUNT_OVER_TIME    { $$ = RangeOpTypeCount }
        |       RATE               { $$ = RangeOpTypeRate }
        |       BYTES_OVER_TIME    { $$ = RangeOpTypeBytes }
        |       BYTES_RATE         { $$ = RangeOpTypeBytesRate }
        |       AVG_OVER_TIME      { $$ = RangeOpTypeAvg }
        |       SUM_OVER_TIME      { $$ = RangeOpTypeSum }
        |       MIN_OVER_TIME      { $$ = RangeOpTypeMin }
        |       MAX_OVER_TIME      { $$ = RangeOpTypeMax }
        |       STDVAR_OVER_TIME   { $$ = RangeOpTypeStdvar }
        |       STDDEV_OVER_TIME   { $$ = RangeOpTypeStddev }
        |       QUANTILE_OVER_TIME { $$ = RangeOpTypeQuantile }
        |       FIRST_OVER_TIME    { $$ = RangeOpTypeFirst }
        |       LAST_OVER_TIME     { $$ = RangeOpTypeLast }
        |       ABSENT_OVER_TIME   { $$ = RangeOpTypeAbsent }
        |       SUM                { $$ = VectorOpTypeSum }
        |       AVG                { $$ = VectorOpTypeAvg }
        |       COUNT              { $$ = VectorOpTypeCount }
        |       MAX                { $$ = VectorOpTypeMax }
        |       MIN                { $$ = VectorOpTypeMin }
        |       STDDEV             { $$ = VectorOpTypeStddev }
        |       STDVAR             { $$ = VectorOpTypeStdvar }
        |       BOTTOMK            { $$ = VectorOpTypeBottomK }
        |       TOPK               { $$ = VectorOpTypeTopK }
                ;

filter:
                PIPE_MATCH                       { $$ = "|~" }
        |       PIPE_EXACT                       { $$ = "|=" }
        |       NRE                              { $$ = "!~" }
        |       NEQ                              { $$ = "!=" }
                ;

labels:
                IDENTIFIER                 { $$ = []string{ $1 } }
        |       labels COMMA IDENTIFIER    { $$ = append($1, $3) }
                ;

grouping:
                BY OPEN_PARENTHESIS labels CLOSE_PARENTHESIS        { $$ = &grouping{ without: false , groups: $3 } }
        |       WITHOUT OPEN_PARENTHESIS labels CLOSE_PARENTHESIS   { $$ = &grouping{ without: true , groups: $3 } }
        |       BY OPEN_PARENTHESIS CLOSE_PARENTHESIS               { $$ = &grouping{ without: false , groups: nil } }
        |       WITHOUT OPEN_PARENTHESIS CLOSE_PARENTHESIS          { $$ = &grouping{ without: true , groups: nil } }
                ;
%%
