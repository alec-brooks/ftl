package reflection

import (
	"reflect"

	"github.com/alecthomas/types/optional"
)

// singletonTypeRegistry is the global type registry that all public functions in this
// package interface with. It is not truly threadsafe. However, everything is initialized
// in init() calls, which are safe, and the type registry is never mutated afterwards.
var singletonTypeRegistry = newTypeRegistry()

// TypeRegistry is used for dynamic type resolution at runtime. It stores associations between sum type discriminators
// and their variants, for use in encoding and decoding.
//
// FTL manages the type registry for you, so you don't need to create one yourself.
type TypeRegistry struct {
	sumTypes                 map[reflect.Type][]sumTypeVariant
	variantsToDiscriminators map[reflect.Type]reflect.Type
}

type sumTypeVariant struct {
	name   string
	goType reflect.Type
}

// WithSumType adds a sum type and its variants to the type registry.
func WithSumType[Discriminator any](variants ...Discriminator) func(t *TypeRegistry) {
	return func(t *TypeRegistry) {
		variantMap := map[string]reflect.Type{}
		for _, v := range variants {
			ref := TypeRefFromValue(v)
			variantMap[ref.Name] = reflect.TypeOf(v)
		}
		t.registerSumType(reflect.TypeFor[Discriminator](), variantMap)
	}
}

// newTypeRegistry creates a new [TypeRegistry] for instantiating types by their qualified
// name at runtime.
func newTypeRegistry(options ...func(t *TypeRegistry)) *TypeRegistry {
	t := &TypeRegistry{
		sumTypes:                 map[reflect.Type][]sumTypeVariant{},
		variantsToDiscriminators: map[reflect.Type]reflect.Type{},
	}
	for _, o := range options {
		o(t)
	}
	return t
}

// Register applies all the provided options to the singleton TypeRegistry
func Register(options ...func(t *TypeRegistry)) {
	for _, o := range options {
		o(singletonTypeRegistry)
	}
}

// registerSumType registers a Go sum type with the type registry.
//
// Sum types are represented as enums in the FTL schema.
func (t *TypeRegistry) registerSumType(discriminator reflect.Type, variants map[string]reflect.Type) {
	var values []sumTypeVariant
	for name, v := range variants {
		t.variantsToDiscriminators[v] = discriminator
		values = append(values, sumTypeVariant{
			name:   name,
			goType: v,
		})
	}
	t.sumTypes[discriminator] = values
}

// ResetTypeRegistry clears the contents of the singleton type registry for tests to
// guarantee determinism.
func ResetTypeRegistry() {
	singletonTypeRegistry = newTypeRegistry()
}

// IsSumTypeDiscriminator returns true if the given type is a sum type discriminator.
func IsSumTypeDiscriminator(discriminator reflect.Type) bool {
	return singletonTypeRegistry.isSumTypeDiscriminator(discriminator)
}

func (t *TypeRegistry) isSumTypeDiscriminator(discriminator reflect.Type) bool {
	return t.getSumTypeVariants(discriminator).Ok()
}

// GetDiscriminatorByVariant returns the discriminator type for the given variant type.
func GetDiscriminatorByVariant(variant reflect.Type) optional.Option[reflect.Type] {
	return singletonTypeRegistry.getDiscriminatorByVariant(variant)
}

func (t *TypeRegistry) getDiscriminatorByVariant(variant reflect.Type) optional.Option[reflect.Type] {
	return optional.Zero(t.variantsToDiscriminators[variant])
}

// GetVariantByName returns the variant type for the given discriminator and variant name.
func GetVariantByName(discriminator reflect.Type, name string) optional.Option[reflect.Type] {
	return singletonTypeRegistry.getVariantByName(discriminator, name)
}

func (t *TypeRegistry) getVariantByName(discriminator reflect.Type, name string) optional.Option[reflect.Type] {
	variants, ok := t.getSumTypeVariants(discriminator).Get()
	if !ok {
		return optional.None[reflect.Type]()
	}
	for _, v := range variants {
		if v.name == name {
			return optional.Some(v.goType)
		}
	}
	return optional.None[reflect.Type]()
}

// GetVariantByType returns the variant name for the given discriminator and variant type.
func GetVariantByType(discriminator reflect.Type, variantType reflect.Type) optional.Option[string] {
	return singletonTypeRegistry.getVariantByType(discriminator, variantType)
}

func (t *TypeRegistry) getVariantByType(discriminator reflect.Type, variantType reflect.Type) optional.Option[string] {
	variants, ok := t.getSumTypeVariants(discriminator).Get()
	if !ok {
		return optional.None[string]()
	}
	for _, v := range variants {
		if v.goType == variantType {
			return optional.Some(v.name)
		}
	}
	return optional.None[string]()
}

func (t *TypeRegistry) getSumTypeVariants(discriminator reflect.Type) optional.Option[[]sumTypeVariant] {
	variants, ok := t.sumTypes[discriminator]
	if !ok {
		return optional.None[[]sumTypeVariant]()
	}

	return optional.Some(variants)
}
