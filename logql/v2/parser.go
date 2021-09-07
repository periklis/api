package v2

import (
	"strings"
	"sync"
	"text/scanner"
)

var parserPool = sync.Pool{
	New: func() interface{} {
		p := &parser{
			p:      &exprParserImpl{},
			Reader: strings.NewReader(""),
			lexer:  &lexer{},
		}
		return p
	},
}

func init() {
	exprErrorVerbose = true
	exprDebug = 0
}

type parser struct {
	p *exprParserImpl
	*lexer
	expr Expr
	*strings.Reader
}

func ParseExpr(input string) (Expr, error) {
	p := parserPool.Get().(*parser)
	defer parserPool.Put(p)

	p.Reader.Reset(input)
	p.lexer.Init(p.Reader)
	p.lexer.errs = p.lexer.errs[:0]
	p.lexer.Scanner.Error = func(_ *scanner.Scanner, msg string) {
		p.lexer.Error(msg)
	}

	e := p.p.Parse(p)
	if e != 0 || len(p.lexer.errs) > 0 {
		return nil, p.lexer.errs[0]
	}

	return p.expr, nil
}
