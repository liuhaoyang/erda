package sqllint

import (
	"bytes"
	"strings"

	"github.com/pingcap/parser/ast"
)

type IndexNameLinter struct {
	script Script
	err    error
	text   string
}

func NewIndexNameLinter(script Script) Rule {
	return &IndexNameLinter{script: script}
}

func (l *IndexNameLinter) Enter(in ast.Node) (ast.Node, bool) {
	if l.text == "" || in.Text() != "" {
		l.text = in.Text()
	}

	constraint, ok := in.(*ast.Constraint)
	if !ok {
		return in, false
	}

	switch constraint.Tp {
	case ast.ConstraintIndex:
		if !strings.HasPrefix(constraint.Name, "idx_") {
			l.err = NewLintError(l.script, l.text, "普通索引名没有以 idx_ 开头",
				func(line []byte) bool {
					return bytes.Contains(bytes.ToLower(line), []byte("index")) &&
						bytes.Contains(bytes.ToLower(line), bytes.ToLower([]byte(constraint.Name)))
				})
			return in, true
		}
	case ast.ConstraintUniq, ast.ConstraintUniqKey, ast.ConstraintUniqIndex:
		if !strings.HasPrefix(constraint.Name, "uk_") {
			l.err = NewLintError(l.script, l.text, "唯一索引名没有以 uk_ 开头",
				func(line []byte) bool {
					return bytes.Contains(bytes.ToLower(line), []byte("unique")) &&
						bytes.Contains(bytes.ToLower(line), bytes.ToLower([]byte(constraint.Name)))
				})
			return in, true
		}
	}

	return in, true
}

func (l *IndexNameLinter) Leave(in ast.Node) (ast.Node, bool) {
	return in, l.err == nil
}

func (l *IndexNameLinter) Error() error {
	return l.err
}
