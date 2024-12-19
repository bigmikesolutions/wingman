//go:build tools

// Tools pkg is a dummy package to enable graphql code generation using go directives.
package tools

import (
	_ "github.com/99designs/gqlgen"
)
