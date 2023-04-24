package schema

import (
	"strings"

	"github.com/alecthomas/errors"
)

var (
	// Identifiers that can't be used as data or verb names.
	reservedIdentNames = map[string]bool{
		"Int": true, "Float": true, "String": true, "Bool": true, "Time": true,
	}
)

// Validate performs semantic validation of a schema.
func Validate(schema *Schema) error {
	modules := map[string]bool{}
	verbs := map[string]bool{}
	data := map[string]bool{}
	verbRefs := []*VerbRef{}
	dataRefs := []*DataRef{}
	merr := []error{}
	for _, module := range schema.Modules {
		if _, seen := modules[module.Name]; seen {
			merr = append(merr, errors.Errorf("%s: duplicate module %q", module.Pos, module.Name))
		}
		modules[module.Name] = true
		if err := ValidateModule(module); err != nil {
			merr = append(merr, err)
		}
		err := Visit(module, func(n Node, next func() error) error {
			switch n := n.(type) {
			case *VerbRef:
				verbRefs = append(verbRefs, n)
			case *DataRef:
				dataRefs = append(dataRefs, n)
			case *Verb:
				if n.Name == "" {
					merr = append(merr, errors.Errorf("%s: verb name is required", n.Pos))
					break
				}
				ref := makeRef(module.Name, n.Name)
				verbs[ref] = true
				verbs[n.Name] = true
			case *Data:
				if n.Name == "" {
					merr = append(merr, errors.Errorf("%s: data structure name is required", n.Pos))
					break
				}
				ref := makeRef(module.Name, n.Name)
				data[ref] = true
				data[n.Name] = true
			default:
			}
			return next()
		})
		if err != nil {
			merr = append(merr, err)
		}
	}
	for _, ref := range verbRefs {
		if !verbs[ref.String()] {
			merr = append(merr, errors.Errorf("%s: reference to unknown Verb %q", ref.Pos, ref))
		}
	}
	for _, ref := range dataRefs {
		if !data[ref.String()] {
			merr = append(merr, errors.Errorf("%s: reference to unknown data structure %q", ref.Pos, ref))
		}
	}
	return errors.Join(merr...)
}

// ValidateModule performs the subset of semantic validation possible on a single module.
func ValidateModule(module *Module) error {
	verbs := map[string]bool{}
	data := map[string]bool{}
	merr := []error{}
	if module.Name == "" {
		merr = append(merr, errors.Errorf("%s: module name is required", module.Pos))
	}
	err := Visit(module, func(n Node, next func() error) error {
		switch n := n.(type) {
		case *Verb:
			if _, ok := reservedIdentNames[n.Name]; ok {
				merr = append(merr, errors.Errorf("%s: Verb name %q is a reserved word", n.Pos, n.Name))
			}
			if _, ok := verbs[n.Name]; ok {
				merr = append(merr, errors.Errorf("%s: duplicate Verb %q", n.Pos, n.Name))
			}
			verbs[n.Name] = true

		case *Data:
			if _, ok := reservedIdentNames[n.Name]; ok {
				merr = append(merr, errors.Errorf("%s: data structure name %q is a reserved word", n.Pos, n.Name))
			}
			if _, ok := data[n.Name]; ok {
				merr = append(merr, errors.Errorf("%s: duplicate data structure %q", n.Pos, n.Name))
			}
			for _, md := range n.Metadata {
				if md, ok := md.(*MetadataCalls); ok {
					merr = append(merr, errors.Errorf("%s: metadata %q is not valid on data structures", md.Pos, strings.TrimSpace(md.String())))
				}
			}

		case *Array, *Bool, *DataRef, *Field, *Float, *Int, *Time, *Map, *MetadataCalls, *Module, *Schema, *String, *VerbRef:
		case Type, Metadata, Decl: // Union types.
		}
		return next()
	})
	if err != nil {
		merr = append(merr, err)
	}
	return errors.Join(merr...)
}