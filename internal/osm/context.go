package osm

import (
	"context"

	"github.com/hiendv/geojson/internal/shared"
	"github.com/paulmach/osm"
)

type ctxKey string

const DEFAULT_OUTDIR = "./geojson"

const ctxKeyRaw ctxKey = "raw"
const ctxKeySeparated ctxKey = "separated"
const ctxKeyOut ctxKey = "out"
const ctxKeyRoot ctxKey = "root"
const ctxKeyLog ctxKey = "log"

// NewContext is the utility to encapsulate pkg-scoped context values by preventing context key collision
func NewContext(ctx context.Context, log shared.Logger, raw bool, separated bool, out string) context.Context {
	ctxx := map[ctxKey]interface{}{
		ctxKeyLog:       log,
		ctxKeyRaw:       raw,
		ctxKeySeparated: separated,
		ctxKeyOut:       out,
	}

	if log != nil {
		log.Debugw("context", "values", ctxx)
	}

	for k, v := range ctxx {
		ctx = context.WithValue(ctx, k, v)
	}

	return ctx
}

func ctxShouldNormalize(ctx context.Context) bool {
	raw, ok := ctx.Value(ctxKeyRaw).(bool)
	return !(ok && raw)
}

func ctxShouldPrint(ctx context.Context) bool {
	out, ok := ctx.Value(ctxKeyOut).(string)
	return ok && out == ""
}

func ctxShouldCombine(ctx context.Context) bool {
	separated, ok := ctx.Value(ctxKeySeparated).(bool)
	return !(ok && separated)
}

func ctxOutDir(ctx context.Context) string {
	out, ok := ctx.Value(ctxKeyOut).(string)
	if !ok {
		return DEFAULT_OUTDIR
	}

	return out
}

func ctxRoot(ctx context.Context) (*osm.Relation, bool) {
	v, ok := ctx.Value(ctxKeyRoot).(*osm.Relation)
	return v, ok
}

func ctxLog(ctx context.Context) shared.Logger {
	v, ok := ctx.Value(ctxKeyLog).(shared.Logger)
	if !ok {
		return shared.LoggerNoop
	}

	return v
}

func ctxSetRoot(ctx context.Context, root *osm.Relation) context.Context {
	return context.WithValue(ctx, ctxKeyRoot, root)
}
