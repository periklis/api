// Code generated by goyacc -p expr -o logql/v2/expr.y.go logql/v2/expr.y. DO NOT EDIT.

package v2

import __yyfmt__ "fmt"


import (
	"github.com/prometheus/prometheus/pkg/labels"
	"time"
)

type exprSymType struct {
	yys                  int
	Expr                 Expr
	LogBinaryOpExpr      LogBinaryOpExpr
	LogFilterExpr        LogFilterExpr
	LogFiltersExpr       LogFiltersExpr
	LogFormatExpr        *LogFormatExpr
	LogQueryExpr         LogSelectorExpr
	LogMetricExpr        LogMetricSampleExpr
	LogNumberExpr        LogNumberExpr
	LogPipelineExpr      LogPipelineExpr
	LogPipelineStageExpr LogPipelineStageExpr
	LogRangeQueryExpr    LogSelectorExpr
	LogOffsetExpr        *LogOffsetExpr
	Matcher              *labels.Matcher
	Matchers             []*labels.Matcher
	MetricOp             string
	BinaryOpOptions      BinaryOpOptions
	Range                string
	Filter               string
	Selector             []*labels.Matcher
	Grouping             *grouping
	Labels               []string
	str                  string
	binaryOp             string
	ComparisonOp         string
	duration             time.Duration
	ConvOp               string
}

const IDENTIFIER = 57346
const STRING = 57347
const RANGE = 57348
const NUMBER = 57349
const DURATION = 57350
const MATCHERS = 57351
const LABELS = 57352
const EQ = 57353
const RE = 57354
const NRE = 57355
const OPEN_BRACE = 57356
const CLOSE_BRACE = 57357
const OPEN_BRACKET = 57358
const CLOSE_BRACKET = 57359
const COMMA = 57360
const OPEN_PARENTHESIS = 57361
const CLOSE_PARENTHESIS = 57362
const COUNT_OVER_TIME = 57363
const RATE = 57364
const SUM = 57365
const AVG = 57366
const MAX = 57367
const MIN = 57368
const COUNT = 57369
const STDDEV = 57370
const STDVAR = 57371
const BOTTOMK = 57372
const TOPK = 57373
const BYTES_OVER_TIME = 57374
const BYTES_RATE = 57375
const BOOL = 57376
const JSON = 57377
const REGEXP = 57378
const LOGFMT = 57379
const PIPE_MATCH = 57380
const PIPE_EXACT = 57381
const PIPE = 57382
const LINE_FMT = 57383
const LABEL_FMT = 57384
const UNWRAP = 57385
const AVG_OVER_TIME = 57386
const SUM_OVER_TIME = 57387
const MIN_OVER_TIME = 57388
const MAX_OVER_TIME = 57389
const STDVAR_OVER_TIME = 57390
const STDDEV_OVER_TIME = 57391
const QUANTILE_OVER_TIME = 57392
const FIRST_OVER_TIME = 57393
const LAST_OVER_TIME = 57394
const ABSENT_OVER_TIME = 57395
const BY = 57396
const WITHOUT = 57397
const LABEL_REPLACE = 57398
const IP = 57399
const UNPACK = 57400
const PATTERN = 57401
const OFFSET = 57402
const BYTES_CONV = 57403
const DURATION_CONV = 57404
const DURATION_SECONDS_CONV = 57405
const OR = 57406
const AND = 57407
const UNLESS = 57408
const CMP_EQ = 57409
const NEQ = 57410
const LT = 57411
const LTE = 57412
const GT = 57413
const GTE = 57414
const ADD = 57415
const SUB = 57416
const MUL = 57417
const DIV = 57418
const MOD = 57419
const POW = 57420

var exprToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENTIFIER",
	"STRING",
	"RANGE",
	"NUMBER",
	"DURATION",
	"MATCHERS",
	"LABELS",
	"EQ",
	"RE",
	"NRE",
	"OPEN_BRACE",
	"CLOSE_BRACE",
	"OPEN_BRACKET",
	"CLOSE_BRACKET",
	"COMMA",
	"OPEN_PARENTHESIS",
	"CLOSE_PARENTHESIS",
	"COUNT_OVER_TIME",
	"RATE",
	"SUM",
	"AVG",
	"MAX",
	"MIN",
	"COUNT",
	"STDDEV",
	"STDVAR",
	"BOTTOMK",
	"TOPK",
	"BYTES_OVER_TIME",
	"BYTES_RATE",
	"BOOL",
	"JSON",
	"REGEXP",
	"LOGFMT",
	"PIPE_MATCH",
	"PIPE_EXACT",
	"PIPE",
	"LINE_FMT",
	"LABEL_FMT",
	"UNWRAP",
	"AVG_OVER_TIME",
	"SUM_OVER_TIME",
	"MIN_OVER_TIME",
	"MAX_OVER_TIME",
	"STDVAR_OVER_TIME",
	"STDDEV_OVER_TIME",
	"QUANTILE_OVER_TIME",
	"FIRST_OVER_TIME",
	"LAST_OVER_TIME",
	"ABSENT_OVER_TIME",
	"BY",
	"WITHOUT",
	"LABEL_REPLACE",
	"IP",
	"UNPACK",
	"PATTERN",
	"OFFSET",
	"BYTES_CONV",
	"DURATION_CONV",
	"DURATION_SECONDS_CONV",
	"OR",
	"AND",
	"UNLESS",
	"CMP_EQ",
	"NEQ",
	"LT",
	"LTE",
	"GT",
	"GTE",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"MOD",
	"POW",
}

var exprStatenames = [...]string{}

const exprEofCode = 1
const exprErrCode = 2
const exprInitialStackSize = 16


var exprExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const exprPrivate = 57344

const exprLast = 468

var exprAct = [...]int{
	66, 54, 53, 4, 164, 5, 74, 150, 7, 107,
	175, 3, 64, 113, 2, 43, 44, 45, 46, 11,
	63, 153, 155, 156, 46, 190, 14, 123, 125, 126,
	142, 8, 77, 15, 16, 29, 30, 32, 33, 31,
	34, 35, 36, 37, 17, 18, 41, 42, 43, 44,
	45, 46, 162, 265, 162, 92, 19, 20, 21, 22,
	23, 24, 25, 26, 27, 28, 224, 206, 10, 109,
	222, 110, 163, 117, 111, 118, 264, 191, 154, 159,
	160, 157, 158, 257, 124, 12, 13, 144, 145, 146,
	127, 251, 128, 129, 130, 131, 132, 133, 134, 135,
	136, 137, 138, 139, 140, 141, 165, 165, 217, 216,
	165, 103, 165, 209, 169, 211, 64, 209, 173, 210,
	174, 171, 67, 68, 63, 247, 237, 178, 182, 38,
	39, 40, 47, 48, 51, 52, 49, 50, 41, 42,
	43, 44, 45, 46, 39, 40, 47, 48, 51, 52,
	49, 50, 41, 42, 43, 44, 45, 46, 57, 204,
	246, 218, 65, 104, 193, 242, 61, 240, 199, 200,
	197, 92, 198, 202, 203, 111, 196, 236, 207, 47,
	48, 51, 52, 49, 50, 41, 42, 43, 44, 45,
	46, 59, 60, 56, 232, 221, 215, 67, 68, 177,
	177, 223, 225, 226, 92, 92, 220, 229, 208, 230,
	61, 231, 194, 168, 93, 179, 176, 245, 167, 106,
	105, 62, 219, 239, 187, 241, 120, 243, 161, 116,
	115, 92, 114, 11, 69, 59, 60, 56, 261, 119,
	14, 252, 121, 253, 122, 112, 254, 15, 16, 29,
	30, 32, 33, 31, 34, 35, 36, 37, 17, 18,
	201, 61, 260, 256, 255, 62, 189, 61, 228, 188,
	19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
	235, 234, 10, 205, 181, 180, 59, 60, 56, 172,
	166, 195, 59, 60, 56, 71, 14, 70, 244, 12,
	13, 8, 263, 15, 16, 29, 30, 32, 33, 31,
	34, 35, 36, 37, 17, 18, 62, 61, 262, 259,
	258, 61, 62, 250, 227, 249, 19, 20, 21, 22,
	23, 24, 25, 26, 27, 28, 248, 238, 10, 213,
	233, 170, 59, 60, 56, 108, 59, 60, 61, 212,
	192, 186, 14, 185, 61, 12, 13, 112, 184, 15,
	16, 29, 30, 32, 33, 31, 34, 35, 36, 37,
	17, 18, 62, 59, 60, 56, 62, 183, 149, 59,
	60, 56, 19, 20, 21, 22, 23, 24, 25, 26,
	27, 28, 148, 147, 10, 151, 214, 102, 73, 75,
	75, 143, 9, 62, 72, 6, 55, 58, 14, 62,
	152, 12, 13, 8, 76, 15, 16, 29, 30, 32,
	33, 31, 34, 35, 36, 37, 17, 18, 95, 98,
	94, 1, 0, 0, 100, 101, 97, 0, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 0, 0,
	10, 96, 99, 0, 78, 79, 80, 81, 82, 83,
	84, 85, 86, 87, 88, 89, 90, 91,
}

var exprPact = [...]int{
	12, -1000, 65, -1000, -1000, -1000, -1000, 341, 394, 143,
	215, -1000, 290, 288, 396, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -2, -2,
	-2, -2, -2, -2, -2, -2, -2, -2, -2, -2,
	-2, -2, -2, 341, -1000, 308, 393, -1000, 106, -1000,
	-1000, -1000, -1000, 200, 199, 338, 213, 211, 210, 12,
	-1000, -1000, 224, 229, -1000, 16, 12, -1000, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, -1000, -1000, -1000, -1000, -1000, 26, 388, 387,
	373, 391, 10, -1000, 209, -1000, -1000, 52, 272, 198,
	193, 335, 394, 65, 282, 196, 195, 267, 266, -1000,
	-1000, 395, -1000, 372, 353, 348, 346, 79, 112, 112,
	-60, -60, -54, -54, -54, -54, -27, -27, -27, -27,
	-27, -27, -1000, 205, -1000, -1000, -1000, -1000, -1000, -1000,
	251, 255, 20, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 345, -1000, 68, 192, 283, 226, 68, 68, 254,
	341, 153, 265, 47, 188, 99, -1000, -1000, 95, -1000,
	344, 334, -1000, -1000, -1000, -1000, -1000, 392, 391, 104,
	-1000, 203, 186, -1000, 68, -1000, 50, 46, 183, -1000,
	-1000, -1000, 341, 304, 248, 12, -1000, 174, -1000, 336,
	-1000, -1000, 263, 262, 157, 251, -1000, -1000, 107, 332,
	-1000, -1000, 68, 147, 68, 145, 68, 292, -1000, 197,
	140, 105, -1000, -1000, 331, 320, -1000, 318, 71, -1000,
	68, -1000, 68, -1000, -1000, 68, -1000, -1000, 246, 245,
	63, -1000, -1000, -1000, -1000, 315, 314, -1000, 244, 220,
	313, 297, 56, 33, -1000, -1000,
}

var exprPgo = [...]int{
	0, 431, 414, 410, 13, 407, 10, 5, 158, 406,
	7, 11, 3, 405, 2, 1, 9, 4, 6, 404,
	402, 8, 0, 401,
}

var exprR1 = [...]int{
	0, 1, 4, 4, 4, 4, 11, 11, 11, 14,
	14, 23, 23, 23, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 9, 9, 8, 8,
	10, 10, 10, 10, 17, 16, 16, 16, 16, 16,
	16, 16, 16, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 13, 13, 13, 2, 2, 21, 21,
	21, 19, 19, 18, 18, 18, 18, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	5, 5, 5, 5, 3, 3, 3, 3, 3, 3,
	3, 3, 6, 6, 22, 22, 22, 22,
}

var exprR2 = [...]int{
	0, 1, 1, 1, 1, 1, 1, 2, 3, 1,
	2, 1, 1, 1, 1, 2, 2, 2, 3, 6,
	3, 3, 3, 3, 4, 7, 1, 2, 2, 5,
	3, 3, 6, 3, 2, 2, 3, 3, 4, 5,
	5, 6, 2, 4, 5, 6, 7, 7, 8, 5,
	6, 5, 4, 6, 7, 8, 7, 5, 6, 12,
	5, 4, 6, 7, 7, 5, 12, 3, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 1, 2, 2, 0, 1, 3, 3,
	3, 1, 3, 3, 3, 3, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 4, 4, 3, 3,
}

var exprChk = [...]int{
	-1000, -1, -4, -11, -12, -7, -13, -21, 19, -20,
	56, 7, 73, 74, 14, 21, 22, 32, 33, 44,
	45, 46, 47, 48, 49, 50, 51, 52, 53, 23,
	24, 27, 25, 26, 28, 29, 30, 31, 64, 65,
	66, 73, 74, 75, 76, 77, 78, 67, 68, 71,
	72, 69, 70, -14, -15, -9, 40, -8, -5, 38,
	39, 13, 68, -11, -12, 19, -22, 54, 55, 19,
	7, 7, -19, 2, -18, 4, -2, 34, -2, -2,
	-2, -2, -2, -2, -2, -2, -2, -2, -2, -2,
	-2, -2, -15, -8, 37, 35, 58, 43, 36, 59,
	41, 42, 4, 5, 57, 20, 20, -16, 7, -12,
	-7, -21, 19, -4, 19, 19, 19, -12, -7, 15,
	2, 18, 15, 11, 68, 12, 13, -4, -4, -4,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4,
	-4, -4, 4, -23, 61, 62, 63, 5, 5, 5,
	-10, 4, -3, 11, 68, 12, 13, 71, 72, 69,
	70, 19, 2, 20, -17, 60, 18, 20, 20, -14,
	6, -21, 7, -12, -7, -6, 20, 4, -6, 20,
	18, 18, -18, 5, 5, 5, 5, 19, 18, 11,
	5, 57, 5, -22, 20, 8, -16, -12, -7, -22,
	-22, 6, -14, -14, 6, 18, 20, -17, 20, 18,
	20, 20, 5, 5, 4, -10, 5, 4, 57, 19,
	20, -22, 20, -17, 20, -17, 20, 20, 20, -14,
	-12, -7, 20, 4, 18, 18, 20, 19, 5, -22,
	20, -22, 20, -22, 6, 20, 20, 20, 5, 5,
	5, 20, -22, -22, -22, 18, 18, 20, 5, 5,
	18, 18, 5, 5, 20, 20,
}

var exprDef = [...]int{
	0, -2, 1, 2, 3, 4, 5, 6, 0, 0,
	0, 83, 0, 0, 0, 97, 98, 99, 100, 101,
	102, 103, 104, 105, 106, 107, 108, 109, 110, 111,
	112, 113, 114, 115, 116, 117, 118, 119, 86, 86,
	86, 86, 86, 86, 86, 86, 86, 86, 86, 86,
	86, 86, 86, 7, 9, 14, 0, 26, 0, 120,
	121, 122, 123, 0, 0, 0, 0, 0, 0, 0,
	84, 85, 0, 0, 91, 0, 0, 87, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 10, 27, 15, 16, 17, 0, 0, 0,
	0, 0, 0, 28, 0, 8, 67, 0, 83, 3,
	4, 6, 0, 0, 0, 0, 0, 3, 4, 88,
	89, 0, 90, 0, 0, 0, 0, 68, 69, 70,
	71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
	81, 82, 18, 0, 11, 12, 13, 20, 21, 22,
	23, 0, 0, 124, 125, 126, 127, 128, 129, 130,
	131, 0, 42, 43, 0, 0, 0, 52, 61, 7,
	35, 6, 83, 3, 4, 0, 136, 132, 0, 137,
	0, 0, 92, 93, 94, 95, 96, 0, 0, 0,
	24, 0, 0, 49, 44, 34, 0, 3, 4, 51,
	60, 37, 36, 0, 0, 0, 57, 0, 65, 0,
	134, 135, 0, 0, 0, 33, 30, 31, 0, 0,
	29, 50, 45, 0, 53, 0, 62, 0, 38, 0,
	3, 4, 58, 133, 0, 0, 19, 0, 0, 46,
	47, 54, 0, 63, 40, 39, 56, 64, 0, 0,
	0, 25, 48, 55, 41, 0, 0, 32, 0, 0,
	0, 0, 0, 0, 59, 66,
}

var exprTok1 = [...]int{
	1,
}

var exprTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78,
}

var exprTok3 = [...]int{
	0,
}

var exprErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}


/*	parser for yacc output	*/

var (
	exprDebug        = 0
	exprErrorVerbose = false
)

type exprLexer interface {
	Lex(lval *exprSymType) int
	Error(s string)
}

type exprParser interface {
	Parse(exprLexer) int
	Lookahead() int
}

type exprParserImpl struct {
	lval  exprSymType
	stack [exprInitialStackSize]exprSymType
	char  int
}

func (p *exprParserImpl) Lookahead() int {
	return p.char
}

func exprNewParser() exprParser {
	return &exprParserImpl{}
}

const exprFlag = -1000

func exprTokname(c int) string {
	if c >= 1 && c-1 < len(exprToknames) {
		if exprToknames[c-1] != "" {
			return exprToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func exprStatname(s int) string {
	if s >= 0 && s < len(exprStatenames) {
		if exprStatenames[s] != "" {
			return exprStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func exprErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !exprErrorVerbose {
		return "syntax error"
	}

	for _, e := range exprErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + exprTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := exprPact[state]
	for tok := TOKSTART; tok-1 < len(exprToknames); tok++ {
		if n := base + tok; n >= 0 && n < exprLast && exprChk[exprAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if exprDef[state] == -2 {
		i := 0
		for exprExca[i] != -1 || exprExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; exprExca[i] >= 0; i += 2 {
			tok := exprExca[i]
			if tok < TOKSTART || exprExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if exprExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += exprTokname(tok)
	}
	return res
}

func exprlex1(lex exprLexer, lval *exprSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = exprTok1[0]
		goto out
	}
	if char < len(exprTok1) {
		token = exprTok1[char]
		goto out
	}
	if char >= exprPrivate {
		if char < exprPrivate+len(exprTok2) {
			token = exprTok2[char-exprPrivate]
			goto out
		}
	}
	for i := 0; i < len(exprTok3); i += 2 {
		token = exprTok3[i+0]
		if token == char {
			token = exprTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = exprTok2[1] /* unknown char */
	}
	if exprDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", exprTokname(token), uint(char))
	}
	return char, token
}

func exprParse(exprlex exprLexer) int {
	return exprNewParser().Parse(exprlex)
}

func (exprrcvr *exprParserImpl) Parse(exprlex exprLexer) int {
	var exprn int
	var exprVAL exprSymType
	var exprDollar []exprSymType
	_ = exprDollar // silence set and not used
	exprS := exprrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	exprstate := 0
	exprrcvr.char = -1
	exprtoken := -1 // exprrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		exprstate = -1
		exprrcvr.char = -1
		exprtoken = -1
	}()
	exprp := -1
	goto exprstack

ret0:
	return 0

ret1:
	return 1

exprstack:
	/* put a state and value onto the stack */
	if exprDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", exprTokname(exprtoken), exprStatname(exprstate))
	}

	exprp++
	if exprp >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprS[exprp] = exprVAL
	exprS[exprp].yys = exprstate

exprnewstate:
	exprn = exprPact[exprstate]
	if exprn <= exprFlag {
		goto exprdefault /* simple state */
	}
	if exprrcvr.char < 0 {
		exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
	}
	exprn += exprtoken
	if exprn < 0 || exprn >= exprLast {
		goto exprdefault
	}
	exprn = exprAct[exprn]
	if exprChk[exprn] == exprtoken { /* valid shift */
		exprrcvr.char = -1
		exprtoken = -1
		exprVAL = exprrcvr.lval
		exprstate = exprn
		if Errflag > 0 {
			Errflag--
		}
		goto exprstack
	}

exprdefault:
	/* default state action */
	exprn = exprDef[exprstate]
	if exprn == -2 {
		if exprrcvr.char < 0 {
			exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if exprExca[xi+0] == -1 && exprExca[xi+1] == exprstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			exprn = exprExca[xi+0]
			if exprn < 0 || exprn == exprtoken {
				break
			}
		}
		exprn = exprExca[xi+1]
		if exprn < 0 {
			goto ret0
		}
	}
	if exprn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			exprlex.Error(exprErrorMessage(exprstate, exprtoken))
			Nerrs++
			if exprDebug >= 1 {
				__yyfmt__.Printf("%s", exprStatname(exprstate))
				__yyfmt__.Printf(" saw %s\n", exprTokname(exprtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for exprp >= 0 {
				exprn = exprPact[exprS[exprp].yys] + exprErrCode
				if exprn >= 0 && exprn < exprLast {
					exprstate = exprAct[exprn] /* simulate a shift of "error" */
					if exprChk[exprstate] == exprErrCode {
						goto exprstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if exprDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", exprS[exprp].yys)
				}
				exprp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if exprDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", exprTokname(exprtoken))
			}
			if exprtoken == exprEofCode {
				goto ret1
			}
			exprrcvr.char = -1
			exprtoken = -1
			goto exprnewstate /* try again in the same state */
		}
	}

	/* reduction by production exprn */
	if exprDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", exprn, exprStatname(exprstate))
	}

	exprnt := exprn
	exprpt := exprp
	_ = exprpt // guard against "declared and not used"

	exprp -= exprR2[exprn]
	// exprp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if exprp+1 >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprVAL = exprS[exprp+1]

	/* consult goto table to find next state */
	exprn = exprR1[exprn]
	exprg := exprPgo[exprn]
	exprj := exprg + exprS[exprp].yys + 1

	if exprj >= exprLast {
		exprstate = exprAct[exprg]
	} else {
		exprstate = exprAct[exprj]
		if exprChk[exprstate] != -exprn {
			exprstate = exprAct[exprg]
		}
	}
	// dummy call; replaced with literal code
	switch exprnt {

	case 1:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprlex.(*parser).expr = exprDollar[1].Expr
		}
	case 2:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Expr = exprDollar[1].LogQueryExpr
		}
	case 3:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Expr = exprDollar[1].LogMetricExpr
		}
	case 4:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Expr = exprDollar[1].LogBinaryOpExpr
		}
	case 5:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Expr = exprDollar[1].LogNumberExpr
		}
	case 6:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.LogQueryExpr = newStreamMatcherExpr(exprDollar[1].Selector)
		}
	case 7:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LogQueryExpr = newLogQueryExpr(newStreamMatcherExpr(exprDollar[1].Selector), exprDollar[2].LogPipelineExpr)
		}
	case 8:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogQueryExpr = exprDollar[2].LogQueryExpr
		}
	case 9:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.LogPipelineExpr = LogPipelineExpr{exprDollar[1].LogPipelineStageExpr}
		}
	case 10:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LogPipelineExpr = append(exprDollar[1].LogPipelineExpr, exprDollar[2].LogPipelineStageExpr)
		}
	case 11:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.ConvOp = OpConvBytes
		}
	case 12:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.ConvOp = OpConvDuration
		}
	case 13:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.ConvOp = OpConvDurationSeconds
		}
	case 14:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.LogPipelineStageExpr = newLogPipelineStageExpr("", nil, exprDollar[1].LogFiltersExpr)
		}
	case 15:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LogPipelineStageExpr = newLogPipelineStageExpr("logfmt", nil, nil)
		}
	case 16:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LogPipelineStageExpr = newLogPipelineStageExpr("json", nil, nil)
		}
	case 17:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LogPipelineStageExpr = newLogPipelineStageExpr("unpack", nil, nil)
		}
	case 18:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogPipelineStageExpr = newLogPipelineStageExpr("unwrap", newLogFormatExpr("", LogFormatValues{"": newLogFormatValue(exprDollar[3].str, true)}, ""), nil)
		}
	case 19:
		exprDollar = exprS[exprpt-6 : exprpt+1]
		{
			exprVAL.LogPipelineStageExpr = newLogPipelineStageExpr("unwrap", newLogFormatExpr("", LogFormatValues{"": newLogFormatValue(exprDollar[5].str, true)}, exprDollar[3].ConvOp), nil)
		}
	case 20:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogPipelineStageExpr = newLogPipelineStageExpr("regexp", newLogFormatExpr("", LogFormatValues{"": newLogFormatValue(exprDollar[3].str, false)}, ""), nil)
		}
	case 21:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogPipelineStageExpr = newLogPipelineStageExpr("pattern", newLogFormatExpr("", LogFormatValues{"": newLogFormatValue(exprDollar[3].str, false)}, ""), nil)
		}
	case 22:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogPipelineStageExpr = newLogPipelineStageExpr("line_format", newLogFormatExpr("", LogFormatValues{"": newLogFormatValue(exprDollar[3].str, false)}, ""), nil)
		}
	case 23:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogPipelineStageExpr = newLogPipelineStageExpr("label_format", exprDollar[3].LogFormatExpr, nil)
		}
	case 24:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogPipelineStageExpr = newLogPipelineStageExpr("", nil, LogFiltersExpr{newLogFilterExpr("|", exprDollar[2].str, exprDollar[3].ComparisonOp, "", exprDollar[4].str)})
		}
	case 25:
		exprDollar = exprS[exprpt-7 : exprpt+1]
		{
			exprVAL.LogPipelineStageExpr = newLogPipelineStageExpr("", nil, LogFiltersExpr{newLogFilterExpr("|", exprDollar[2].str, exprDollar[3].ComparisonOp, OpIP, exprDollar[6].str)})
		}
	case 26:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.LogFiltersExpr = LogFiltersExpr{exprDollar[1].LogFilterExpr}
		}
	case 27:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LogFiltersExpr = append(exprDollar[1].LogFiltersExpr, exprDollar[2].LogFilterExpr)
		}
	case 28:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LogFilterExpr = newLogFilterExpr(exprDollar[1].Filter, "", "", "", exprDollar[2].str)
		}
	case 29:
		exprDollar = exprS[exprpt-5 : exprpt+1]
		{
			exprVAL.LogFilterExpr = newLogFilterExpr(exprDollar[1].Filter, "", "", OpIP, exprDollar[4].str)
		}
	case 30:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogFormatExpr = newLogFormatExpr("", LogFormatValues{exprDollar[1].str: newLogFormatValue(exprDollar[3].str, false)}, "")
		}
	case 31:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogFormatExpr = newLogFormatExpr("", LogFormatValues{exprDollar[1].str: newLogFormatValue(exprDollar[3].str, true)}, "")
		}
	case 32:
		exprDollar = exprS[exprpt-6 : exprpt+1]
		{
			exprVAL.LogFormatExpr = newLogFormatExpr("", LogFormatValues{exprDollar[1].str: newLogFormatValue(OpIP+"("+exprDollar[5].str+")", false)}, "")
		}
	case 33:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogFormatExpr = newLogFormatExpr(",", mergeLogFormatValues(exprDollar[1].LogFormatExpr.kv, exprDollar[3].LogFormatExpr.kv), "")
		}
	case 34:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			{
				exprVAL.LogOffsetExpr = newLogOffsetExpr(exprDollar[2].duration)
			}
		}
	case 35:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LogRangeQueryExpr = newLogRangeQueryExpr(newLogQueryExpr(newStreamMatcherExpr(exprDollar[1].Selector), nil), exprDollar[2].str, nil, false)
		}
	case 36:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogRangeQueryExpr = newLogRangeQueryExpr(newLogQueryExpr(newStreamMatcherExpr(exprDollar[1].Selector), exprDollar[3].LogPipelineExpr), exprDollar[2].str, nil, false)
		}
	case 37:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogRangeQueryExpr = newLogRangeQueryExpr(newLogQueryExpr(newStreamMatcherExpr(exprDollar[1].Selector), exprDollar[2].LogPipelineExpr), exprDollar[3].str, nil, true)
		}
	case 38:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogRangeQueryExpr = newLogRangeQueryExpr(newLogQueryExpr(newStreamMatcherExpr(exprDollar[2].Selector), nil), exprDollar[3].str, nil, false)
		}
	case 39:
		exprDollar = exprS[exprpt-5 : exprpt+1]
		{
			exprVAL.LogRangeQueryExpr = newLogRangeQueryExpr(newLogQueryExpr(newStreamMatcherExpr(exprDollar[2].Selector), exprDollar[4].LogPipelineExpr), exprDollar[3].str, nil, false)
		}
	case 40:
		exprDollar = exprS[exprpt-5 : exprpt+1]
		{
			exprVAL.LogRangeQueryExpr = newLogRangeQueryExpr(newLogQueryExpr(newStreamMatcherExpr(exprDollar[2].Selector), exprDollar[3].LogPipelineExpr), exprDollar[5].str, nil, true)
		}
	case 41:
		exprDollar = exprS[exprpt-6 : exprpt+1]
		{
			exprVAL.LogRangeQueryExpr = newLogRangeQueryExpr(newLogQueryExpr(newStreamMatcherExpr(exprDollar[2].Selector), exprDollar[4].LogPipelineExpr), exprDollar[3].str, exprDollar[6].Grouping, false)
		}
	case 43:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(nil, exprDollar[3].LogRangeQueryExpr, exprDollar[1].MetricOp, "", nil, nil, nil)
		}
	case 44:
		exprDollar = exprS[exprpt-5 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(nil, exprDollar[3].LogRangeQueryExpr, exprDollar[1].MetricOp, "", nil, nil, exprDollar[4].LogOffsetExpr)
		}
	case 45:
		exprDollar = exprS[exprpt-6 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(nil, exprDollar[5].LogRangeQueryExpr, exprDollar[1].MetricOp, exprDollar[3].str, nil, nil, nil)
		}
	case 46:
		exprDollar = exprS[exprpt-7 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(nil, exprDollar[5].LogRangeQueryExpr, exprDollar[1].MetricOp, exprDollar[3].str, exprDollar[7].Grouping, nil, nil)
		}
	case 47:
		exprDollar = exprS[exprpt-7 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(nil, exprDollar[5].LogRangeQueryExpr, exprDollar[1].MetricOp, exprDollar[3].str, nil, nil, exprDollar[6].LogOffsetExpr)
		}
	case 48:
		exprDollar = exprS[exprpt-8 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(nil, exprDollar[5].LogRangeQueryExpr, exprDollar[1].MetricOp, exprDollar[3].str, exprDollar[8].Grouping, nil, exprDollar[6].LogOffsetExpr)
		}
	case 49:
		exprDollar = exprS[exprpt-5 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(nil, exprDollar[3].LogRangeQueryExpr, "", "", exprDollar[5].Grouping, nil, nil)
		}
	case 50:
		exprDollar = exprS[exprpt-6 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(nil, exprDollar[3].LogRangeQueryExpr, "", "", exprDollar[6].Grouping, nil, exprDollar[4].LogOffsetExpr)
		}
	case 51:
		exprDollar = exprS[exprpt-5 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[3].LogMetricExpr, nil, exprDollar[1].MetricOp, "", exprDollar[5].Grouping, nil, nil)
		}
	case 52:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[3].LogMetricExpr, nil, exprDollar[1].MetricOp, "", nil, nil, nil)
		}
	case 53:
		exprDollar = exprS[exprpt-6 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[5].LogMetricExpr, nil, exprDollar[1].MetricOp, exprDollar[3].str, nil, nil, nil)
		}
	case 54:
		exprDollar = exprS[exprpt-7 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[5].LogMetricExpr, nil, exprDollar[1].MetricOp, exprDollar[3].str, exprDollar[7].Grouping, nil, nil)
		}
	case 55:
		exprDollar = exprS[exprpt-8 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[5].LogMetricExpr, nil, exprDollar[1].MetricOp, exprDollar[3].str, exprDollar[8].Grouping, nil, exprDollar[6].LogOffsetExpr)
		}
	case 56:
		exprDollar = exprS[exprpt-7 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[6].LogMetricExpr, nil, exprDollar[1].MetricOp, exprDollar[4].str, exprDollar[2].Grouping, nil, nil)
		}
	case 57:
		exprDollar = exprS[exprpt-5 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[4].LogMetricExpr, nil, exprDollar[1].MetricOp, "", exprDollar[2].Grouping, nil, nil)
		}
	case 58:
		exprDollar = exprS[exprpt-6 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[4].LogMetricExpr, nil, exprDollar[1].MetricOp, "", exprDollar[2].Grouping, nil, exprDollar[5].LogOffsetExpr)
		}
	case 59:
		exprDollar = exprS[exprpt-12 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[3].LogMetricExpr, nil, OpLabelReplace, "", nil, []string{exprDollar[5].str, exprDollar[7].str, exprDollar[9].str, exprDollar[11].str}, nil)
		}
	case 60:
		exprDollar = exprS[exprpt-5 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[3].LogBinaryOpExpr, nil, exprDollar[1].MetricOp, "", exprDollar[5].Grouping, nil, nil)
		}
	case 61:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[3].LogBinaryOpExpr, nil, exprDollar[1].MetricOp, "", nil, nil, nil)
		}
	case 62:
		exprDollar = exprS[exprpt-6 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[5].LogBinaryOpExpr, nil, exprDollar[1].MetricOp, exprDollar[3].str, nil, nil, nil)
		}
	case 63:
		exprDollar = exprS[exprpt-7 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[5].LogBinaryOpExpr, nil, exprDollar[1].MetricOp, exprDollar[3].str, exprDollar[7].Grouping, nil, nil)
		}
	case 64:
		exprDollar = exprS[exprpt-7 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[6].LogBinaryOpExpr, nil, exprDollar[1].MetricOp, exprDollar[4].str, exprDollar[2].Grouping, nil, nil)
		}
	case 65:
		exprDollar = exprS[exprpt-5 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[4].LogBinaryOpExpr, nil, exprDollar[1].MetricOp, "", exprDollar[2].Grouping, nil, nil)
		}
	case 66:
		exprDollar = exprS[exprpt-12 : exprpt+1]
		{
			exprVAL.LogMetricExpr = newLogMetricExpr(exprDollar[3].LogBinaryOpExpr, nil, OpLabelReplace, "", nil, []string{exprDollar[5].str, exprDollar[7].str, exprDollar[9].str, exprDollar[11].str}, nil)
		}
	case 67:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogMetricExpr = exprDollar[2].LogMetricExpr
		}
	case 68:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogBinaryOpExpr = newLogBinaryOpExpr("or", exprDollar[3].BinaryOpOptions, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 69:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogBinaryOpExpr = newLogBinaryOpExpr("and", exprDollar[3].BinaryOpOptions, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 70:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogBinaryOpExpr = newLogBinaryOpExpr("unless", exprDollar[3].BinaryOpOptions, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 71:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogBinaryOpExpr = newLogBinaryOpExpr("+", exprDollar[3].BinaryOpOptions, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 72:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogBinaryOpExpr = newLogBinaryOpExpr("-", exprDollar[3].BinaryOpOptions, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 73:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogBinaryOpExpr = newLogBinaryOpExpr("*", exprDollar[3].BinaryOpOptions, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 74:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogBinaryOpExpr = newLogBinaryOpExpr("/", exprDollar[3].BinaryOpOptions, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 75:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogBinaryOpExpr = newLogBinaryOpExpr("%", exprDollar[3].BinaryOpOptions, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 76:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogBinaryOpExpr = newLogBinaryOpExpr("^", exprDollar[3].BinaryOpOptions, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 77:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogBinaryOpExpr = newLogBinaryOpExpr("==", exprDollar[3].BinaryOpOptions, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 78:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogBinaryOpExpr = newLogBinaryOpExpr("!=", exprDollar[3].BinaryOpOptions, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 79:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogBinaryOpExpr = newLogBinaryOpExpr(">", exprDollar[3].BinaryOpOptions, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 80:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogBinaryOpExpr = newLogBinaryOpExpr(">=", exprDollar[3].BinaryOpOptions, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 81:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogBinaryOpExpr = newLogBinaryOpExpr("<", exprDollar[3].BinaryOpOptions, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 82:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogBinaryOpExpr = newLogBinaryOpExpr("<=", exprDollar[3].BinaryOpOptions, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 83:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.LogNumberExpr = newLogNumberExpr(exprDollar[1].str, false)
		}
	case 84:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LogNumberExpr = newLogNumberExpr(exprDollar[2].str, false)
		}
	case 85:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LogNumberExpr = newLogNumberExpr(exprDollar[2].str, true)
		}
	case 86:
		exprDollar = exprS[exprpt-0 : exprpt+1]
		{
			exprVAL.BinaryOpOptions = BinaryOpOptions{}
		}
	case 87:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.BinaryOpOptions = BinaryOpOptions{ReturnBool: true}
		}
	case 88:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Selector = exprDollar[2].Matchers
		}
	case 89:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Selector = exprDollar[2].Matchers
		}
	case 90:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
		}
	case 91:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Matchers = []*labels.Matcher{exprDollar[1].Matcher}
		}
	case 92:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Matchers = append(exprDollar[1].Matchers, exprDollar[3].Matcher)
		}
	case 93:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Matcher = newLabelMatcher(labels.MatchEqual, exprDollar[1].str, exprDollar[3].str)
		}
	case 94:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Matcher = newLabelMatcher(labels.MatchNotEqual, exprDollar[1].str, exprDollar[3].str)
		}
	case 95:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Matcher = newLabelMatcher(labels.MatchRegexp, exprDollar[1].str, exprDollar[3].str)
		}
	case 96:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Matcher = newLabelMatcher(labels.MatchNotRegexp, exprDollar[1].str, exprDollar[3].str)
		}
	case 97:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = RangeOpTypeCount
		}
	case 98:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = RangeOpTypeRate
		}
	case 99:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = RangeOpTypeBytes
		}
	case 100:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = RangeOpTypeBytesRate
		}
	case 101:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = RangeOpTypeAvg
		}
	case 102:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = RangeOpTypeSum
		}
	case 103:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = RangeOpTypeMin
		}
	case 104:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = RangeOpTypeMax
		}
	case 105:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = RangeOpTypeStdvar
		}
	case 106:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = RangeOpTypeStddev
		}
	case 107:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = RangeOpTypeQuantile
		}
	case 108:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = RangeOpTypeFirst
		}
	case 109:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = RangeOpTypeLast
		}
	case 110:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = RangeOpTypeAbsent
		}
	case 111:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = VectorOpTypeSum
		}
	case 112:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = VectorOpTypeAvg
		}
	case 113:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = VectorOpTypeCount
		}
	case 114:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = VectorOpTypeMax
		}
	case 115:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = VectorOpTypeMin
		}
	case 116:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = VectorOpTypeStddev
		}
	case 117:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = VectorOpTypeStdvar
		}
	case 118:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = VectorOpTypeBottomK
		}
	case 119:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricOp = VectorOpTypeTopK
		}
	case 120:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Filter = "|~"
		}
	case 121:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Filter = "|="
		}
	case 122:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Filter = "!~"
		}
	case 123:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Filter = "!="
		}
	case 124:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.ComparisonOp = "="
		}
	case 125:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.ComparisonOp = "!="
		}
	case 126:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.ComparisonOp = "=~"
		}
	case 127:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.ComparisonOp = "!~"
		}
	case 128:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.ComparisonOp = ">"
		}
	case 129:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.ComparisonOp = ">="
		}
	case 130:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.ComparisonOp = "<"
		}
	case 131:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.ComparisonOp = "<="
		}
	case 132:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Labels = []string{exprDollar[1].str}
		}
	case 133:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Labels = append(exprDollar[1].Labels, exprDollar[3].str)
		}
	case 134:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.Grouping = &grouping{without: false, groups: exprDollar[3].Labels}
		}
	case 135:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.Grouping = &grouping{without: true, groups: exprDollar[3].Labels}
		}
	case 136:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Grouping = &grouping{without: false, groups: nil}
		}
	case 137:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Grouping = &grouping{without: true, groups: nil}
		}
	}
	goto exprstack /* stack new state and value */
}
