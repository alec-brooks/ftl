// Code generated by FTL. DO NOT EDIT.
package main

import (
	"context"
	"reflect"

	"github.com/TBD54566975/ftl/backend/protos/xyz/block/ftl/v1/ftlv1connect"
	"github.com/TBD54566975/ftl/backend/schema"
	"github.com/TBD54566975/ftl/common/plugin"
	"github.com/TBD54566975/ftl/go-runtime/ftl/typeregistry"
	"github.com/TBD54566975/ftl/go-runtime/server"

	"ftl/another"
	"ftl/other"
)

func main() {
	verbConstructor := server.NewUserVerbServer("other",
		server.HandleCall(other.Echo),
	)
	ctx := context.Background()

	goTypeRegistry := typeregistry.NewTypeRegistry()
	schemaTypeRegistry := schema.NewTypeRegistry()
	goTypeRegistry.RegisterSumType(reflect.TypeFor[another.TypeEnum](), map[string]reflect.Type{
		"A": reflect.TypeFor[another.A](),
		"B": reflect.TypeFor[another.B](),
	})
	schemaTypeRegistry.RegisterSumType("another.TypeEnum", map[string]schema.Type{
		"A": &schema.Int{},
		"B": &schema.String{},
	})
	goTypeRegistry.RegisterSumType(reflect.TypeFor[other.SecondTypeEnum](), map[string]reflect.Type{
		"A": reflect.TypeFor[other.A](),
		"B": reflect.TypeFor[other.B](),
	})
	schemaTypeRegistry.RegisterSumType("other.SecondTypeEnum", map[string]schema.Type{
		"A": &schema.String{},
		"B": &schema.Ref{Module: "other", Name: "B"},
	})
	goTypeRegistry.RegisterSumType(reflect.TypeFor[other.TypeEnum](), map[string]reflect.Type{
		"Bool": reflect.TypeFor[other.Bool](),
		"Bytes": reflect.TypeFor[other.Bytes](),
		"Float": reflect.TypeFor[other.Float](),
		"Int": reflect.TypeFor[other.Int](),
		"Time": reflect.TypeFor[other.Time](),
		"List": reflect.TypeFor[other.List](),
		"Map": reflect.TypeFor[other.Map](),
		"String": reflect.TypeFor[other.String](),
		"Struct": reflect.TypeFor[other.Struct](),
		"Option": reflect.TypeFor[other.Option](),
		"Unit": reflect.TypeFor[other.Unit](),
	})
	schemaTypeRegistry.RegisterSumType("other.TypeEnum", map[string]schema.Type{
		"Bool": &schema.Bool{},
		"Bytes": &schema.Bytes{},
		"Float": &schema.Float{},
		"Int": &schema.Int{},
		"Time": &schema.Time{},
		"List": &schema.Array{Element: &schema.String{}},
		"Map": &schema.Map{Key: &schema.String{}, Value: &schema.String{}},
		"String": &schema.String{},
		"Struct": &schema.Ref{Module: "other", Name: "Struct"},
		"Option": &schema.Optional{Type: &schema.String{}},
		"Unit": &schema.Unit{},
	})
	ctx = typeregistry.ContextWithTypeRegistry(ctx, goTypeRegistry)
	ctx = schema.ContextWithTypeRegistry(ctx, schemaTypeRegistry.ToProto())

	plugin.Start(ctx, "other", verbConstructor, ftlv1connect.VerbServiceName, ftlv1connect.NewVerbServiceHandler)
}