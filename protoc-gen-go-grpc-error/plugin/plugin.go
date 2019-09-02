package plugin

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/zoncoen/grpc-error-annotation/gen/annotation/api/grpc"
	"google.golang.org/grpc/codes"
)

type plugin struct {
	*generator.Generator
	generator.PluginImports
}

// NewPlugin implements generator.Plugin interface.
func NewPlugin() generator.Plugin {
	return &plugin{}
}

// Name implements generator.Plugin interface.
func (p *plugin) Name() string {
	return "grpc-error"
}

// Init implements generator.Plugin interface.
func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

// Generate implements generator.Plugin interface.
func (p *plugin) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)

	p.AddImport(generator.GoImportPath("context"))
	p.AddImport(generator.GoImportPath("google.golang.org/grpc"))
	p.AddImport(generator.GoImportPath("google.golang.org/grpc/codes"))
	p.AddImport(generator.GoImportPath("google.golang.org/grpc/status"))

	for _, s := range file.GetService() {
		p.generateMap(file.GetPackage(), s)
		p.generateInterceptor(s.GetName())
	}
}

func (p *plugin) generateMap(packageName string, svc *descriptor.ServiceDescriptorProto) {
	p.block(fmt.Sprintf(`var map%s = map[string]struct{}`, svc.GetName()), func(p *plugin) {
		for _, m := range svc.GetMethod() {
			v, err := proto.GetExtension(m.GetOptions(), grpc.E_Error)
			if err != nil {
				continue
			}
			opt, ok := v.(*grpc.Error)
			if !ok {
				continue
			}
			if opt == nil {
				continue
			}
			methodName := fmt.Sprintf("/%s.%s/%s", packageName, svc.GetName(), m.GetName())
			if opt.GetCancelled() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.Canceled))
			}
			if opt.GetUnknown() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.Unknown))
			}
			if opt.GetInvalidArgument() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.InvalidArgument))
			}
			if opt.GetDeadlineExceeded() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.DeadlineExceeded))
			}
			if opt.GetNotFound() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.NotFound))
			}
			if opt.GetAlreadyExists() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.AlreadyExists))
			}
			if opt.GetPermissionDenied() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.PermissionDenied))
			}
			if opt.GetResourceExhausted() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.ResourceExhausted))
			}
			if opt.GetFailedPrecondition() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.FailedPrecondition))
			}
			if opt.GetAborted() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.Aborted))
			}
			if opt.GetOutOfRange() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.OutOfRange))
			}
			if opt.GetUnimplemented() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.Unimplemented))
			}
			if opt.GetInternal() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.Internal))
			}
			if opt.GetUnavailable() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.Unavailable))
			}
			if opt.GetDataLoss() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.DataLoss))
			}
			if opt.GetUnauthenticated() != nil {
				p.P(fmt.Sprintf(`"%s/%d": struct{}{},`, methodName, codes.Unauthenticated))
			}
		}
	})
}

func (p *plugin) block(leading string, inner func(p *plugin)) {
	p.P(leading, `{`)
	p.In()
	inner(p)
	p.Out()
	p.P(`}`)
}

func (p *plugin) generateInterceptor(serviceName string) {
	p.P(fmt.Sprintf(`// New%sCheckErrorUnaryServerInterceptor returns an interceptor to check status code.`, serviceName))
	p.P(`// The interceptor calls function f if handler returns the status code which is not included in zoncoen.api.grpc.error annotation.`)
	p.block(fmt.Sprintf(`func New%sCheckErrorUnaryServerInterceptor(f func(context.Context, *grpc.UnaryServerInfo, *status.Status)) grpc.UnaryServerInterceptor`, serviceName), func(p *plugin) {
		p.block(`return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)`, func(p *plugin) {
			p.P(`resp, err := handler(ctx, req)`)
			p.block(`if err == nil`, func(p *plugin) {
				p.P(`return resp, nil`)
			})
			p.P(`s := status.Convert(err)`)
			p.block(`if s.Code() != codes.OK`, func(p *plugin) {
				p.P(`key := fmt.Sprintf("%s/%d", info.FullMethod, s.Code())`)
				p.block(fmt.Sprintf(`if _, ok := map%s[key]; !ok`, serviceName), func(p *plugin) {
					p.P(`f(ctx, info, s)`)
				})
			})
			p.P(`return resp, err`)
		})
	})
}
