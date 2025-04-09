package graphqlctx

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

const (
	queryEnv = "environment"
	varEnv   = "env"
)

var (
	ErrNotQuery            = errors.New("not query")
	ErrNotEnvironmentQuery = errors.New("not an environment related query")
	ErrEnvironmentMissing  = errors.New("environment value missing")
)

type Query struct {
	ctx *graphql.OperationContext
}

func NewQuery(ctx context.Context) *Query {
	return &Query{
		ctx: graphql.GetOperationContext(ctx),
	}
}

func (q *Query) IsQuery() bool {
	if q.ctx == nil || q.ctx.Operation == nil {
		return false
	}

	return q.ctx.Operation.Operation == ast.Query
}

func (q *Query) QueryVars() map[string]any {
	if !q.IsQuery() {
		return nil
	}

	return q.ctx.Variables
}

func (q *Query) QueryName() (string, bool) {
	if !q.IsQuery() {
		return "", false
	}

	f, ok := q.ctx.Operation.SelectionSet[0].(*ast.Field)

	if !ok {
		return "", false
	}

	return f.Name, true
}

func (q *Query) Environment() (string, error) {
	queryName, ok := q.QueryName()
	if !ok {
		return "", ErrNotQuery
	}

	if queryName != queryEnv {
		return "", ErrNotEnvironmentQuery
	}

	v, ok := q.QueryVars()[varEnv].(string)
	if !ok || v == "" {
		return "", ErrEnvironmentMissing
	}

	return v, nil
}

func Environment(ctx context.Context) (string, error) {
	q := NewQuery(ctx)
	return q.Environment()
}
