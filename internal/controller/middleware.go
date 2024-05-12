package controller

import (
	"context"
	"net/http"
)

type CapbValue string

const Pricing = "pricing"
const capbPricing = CapbValue(Pricing)
const CapbHeader = "Capability"

func WithCapability(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		capb := r.Header.Get(CapbHeader)

		ctx = context.WithValue(ctx, capbPricing, capb == Pricing)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func IsPriced(ctx context.Context) bool {
	// bl := ctx.Value(capbPricing)
	// fmt.Printf("wooot %s\n", bl)

	u, _ := ctx.Value(capbPricing).(bool)
	return u
}
