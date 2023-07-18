package controller

import (
	"context"

	"github.com/alecthomas/errors"
	"github.com/bufbuild/connect-go"

	"github.com/TBD54566975/ftl/controller/internal/dal"
	"github.com/TBD54566975/ftl/internal/slices"
	ftlv1 "github.com/TBD54566975/ftl/protos/xyz/block/ftl/v1"
	pbconsole "github.com/TBD54566975/ftl/protos/xyz/block/ftl/v1/console"
	"github.com/TBD54566975/ftl/protos/xyz/block/ftl/v1/console/pbconsoleconnect"
	pschema "github.com/TBD54566975/ftl/protos/xyz/block/ftl/v1/schema"
	"github.com/TBD54566975/ftl/schema"
)

type ConsoleService struct {
	dal *dal.DAL
}

var _ pbconsoleconnect.ConsoleServiceHandler = (*ConsoleService)(nil)

func NewConsoleService(dal *dal.DAL) *ConsoleService {
	return &ConsoleService{
		dal: dal,
	}
}

func (*ConsoleService) Ping(context.Context, *connect.Request[ftlv1.PingRequest]) (*connect.Response[ftlv1.PingResponse], error) {
	return connect.NewResponse(&ftlv1.PingResponse{}), nil
}

func (c *ConsoleService) GetModules(ctx context.Context, req *connect.Request[pbconsole.GetModulesRequest]) (*connect.Response[pbconsole.GetModulesResponse], error) {
	deployments, err := c.dal.GetActiveDeployments(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	moduleNames, err := slices.MapErr(deployments, func(in dal.Deployment) (string, error) {
		return in.Module, nil
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	moduleCalls, err := c.dal.GetModuleCalls(ctx, moduleNames)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var modules []*pbconsole.Module
	for _, deployment := range deployments {
		var verbs []*pbconsole.Verb
		var data []*pschema.Data

		for _, decl := range deployment.Schema.Decls {
			switch decl := decl.(type) {
			case *schema.Verb:
				//nolint:forcetypeassert
				verbs = append(verbs, &pbconsole.Verb{
					Verb: decl.ToProto().(*pschema.Verb),
					Calls: getModuleCalls(moduleCalls[dal.ModuleCallKey{
						Module: deployment.Module,
						Verb:   decl.Name,
					}]),
				})
			case *schema.Data:
				//nolint:forcetypeassert
				data = append(data, decl.ToProto().(*pschema.Data))
			}
		}

		modules = append(modules, &pbconsole.Module{
			Name:     deployment.Module,
			Language: deployment.Language,
			Verbs:    verbs,
			Data:     data,
		})
	}

	return connect.NewResponse(&pbconsole.GetModulesResponse{
		Modules: modules,
	}), nil
}

func getModuleCalls(calls []dal.Call) []*pbconsole.Call {
	return slices.Map(calls, func(call dal.Call) *pbconsole.Call {
		var errorMessage string
		if call.Error != nil {
			errorMessage = call.Error.Error()
		}
		return &pbconsole.Call{
			Id:            call.ID,
			RunnerKey:     call.RunnerKey.String(),
			RequestId:     call.RequestID,
			ControllerKey: call.ControllerKey.String(),
			TimeStamp:     call.Time.Unix(),
			SourceModule:  call.SourceVerb.Module,
			SourceVerb:    call.SourceVerb.Name,
			DestModule:    call.DestVerb.Module,
			DestVerb:      call.DestVerb.Name,
			DurationMs:    call.Duration.Milliseconds(),
			Request:       call.Request,
			Response:      call.Response,
			Error:         errorMessage,
		}
	})
}