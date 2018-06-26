package interceptor

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/open-policy-agent/opa/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage"
	"github.com/open-policy-agent/opa/storage/inmem"
)

const policy = `package poc.opasvc.Hello

default Say = false

allowed_names = ["leonardo", "donatello", "raphael", "michaelangelo"]

is_welcome(name) {
	lower(allowed_names[_], allowed_lower)
	lower(input.name, name_lower)
	allowed_lower = name_lower
}

Say {
	is_welcome(input.name)
}
`

// PolicyInterceptor ...
func PolicyInterceptor() grpc.UnaryServerInterceptor {

	store := inmem.New()
	compiler := ast.NewCompiler()
	module := ast.MustParseModule(policy)

	compiler.Compile(map[string]*ast.Module{"": module})
	if compiler.Failed() {
		grpclog.Fatalln("Could not compile policy:", compiler.Errors)
	}

	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		json, err := json.Marshal(req)
		if err != nil {
			grpclog.Fatalln("error could not marshal request:", err)
		}
		input := util.MustUnmarshalJSON(json)
		txn := storage.NewTransactionOrDie(ctx, store)

		query := fmt.Sprintf("data.%s", strings.Replace(strings.TrimPrefix(info.FullMethod, "/"), "/", ".", -1))

		r := rego.New(
			rego.Compiler(compiler),
			rego.Store(store),
			rego.Transaction(txn),
			rego.Input(input),
			rego.Query(query),
		)
		start := time.Now()
		rs, err := r.Eval(ctx)
		if err != nil {
			grpclog.Fatalln("Unexpected error(s):", err)
		}
		grpclog.Infoln("Policy evaluation took:", time.Since(start))
		// grpclog.Infoln("Result:", rs[0].Expressions[0].Value)
		// for i, exp := range rs[0].Expressions {
		// 	grpclog.Infof("Expression(%v): %v (%v) = %v", i, exp.Text, exp.Location, exp.Value)
		// }
		allow := rs[0].Expressions[0].Value.(bool)
		if !allow {
			return nil, status.Error(codes.PermissionDenied, "Access denied")
		}

		h, err := handler(ctx, req)
		return h, err
	}

	return interceptor
}
