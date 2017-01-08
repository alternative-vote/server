package utils

import (
	"context"

	"github.com/alternative-vote/server/authentication"
)

const (
	claimsKey contextKey = iota
)

type contextKey int

func SetClaims(ctx context.Context, claims authentication.CustomClaims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}

func GetClaims(ctx context.Context) authentication.CustomClaims {
	val, ok := ctx.Value(claimsKey).(authentication.CustomClaims)
	if !ok {
		panic("expected to be able to get claims from the context but it failed")
	}
	return val
}
