// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.19

package main

import (
	"context"
	"fmt"
	"github.com/ducesoft/ulsp/log"
	"strings"
)

var typeNames = make(map[*Type]string)
var genTypes []*newType

func findTypeNames(ctx context.Context, model Model) {
	for _, s := range model.Structures {
		for _, e := range s.Extends {
			nameType(ctx, e, nil) // all references
		}
		for _, m := range s.Mixins {
			nameType(ctx, m, nil) // all references
		}
		for _, p := range s.Properties {
			nameType(ctx, p.Type, []string{s.Name, p.Name})
		}
	}
	for _, t := range model.Enumerations {
		nameType(ctx, t.Type, []string{t.Name})
	}
	for _, t := range model.TypeAliases {
		nameType(ctx, t.Type, []string{t.Name})
	}
	for _, r := range model.Requests {
		nameType(ctx, r.Params, []string{"Param", r.Method})
		nameType(ctx, r.Result, []string{"Result", r.Method})
		nameType(ctx, r.RegistrationOptions, []string{"RegOpt", r.Method})
	}
	for _, n := range model.Notifications {
		nameType(ctx, n.Params, []string{"Param", n.Method})
		nameType(ctx, n.RegistrationOptions, []string{"RegOpt", n.Method})
	}
}

// nameType populates typeNames[t] with the computed name of the type.
// path is the list of enclosing constructs in the JSON model.
func nameType(ctx context.Context, t *Type, path []string) string {
	if t == nil || typeNames[t] != "" {
		return ""
	}
	switch t.Kind {
	case "base":
		typeNames[t] = t.Name
		return t.Name
	case "reference":
		typeNames[t] = t.Name
		return t.Name
	case "array":
		nm := "[]" + nameType(ctx, t.Element, append(path, "Elem"))
		typeNames[t] = nm
		return nm
	case "map":
		key := nameType(ctx, t.Key, nil) // never a generated type
		value := nameType(ctx, t.Value.(*Type), append(path, "Value"))
		nm := "map[" + key + "]" + value
		typeNames[t] = nm
		return nm
	// generated types
	case "and":
		nm := nameFromPath("And", path)
		typeNames[t] = nm
		for _, it := range t.Items {
			nameType(ctx, it, append(path, "Item"))
		}
		genTypes = append(genTypes, &newType{
			name:  nm,
			typ:   t,
			kind:  "and",
			items: t.Items,
			line:  t.Line,
		})
		return nm
	case "literal":
		nm := nameFromPath("Lit", path)
		typeNames[t] = nm
		for _, p := range t.Value.(ParseLiteral).Properties {
			nameType(ctx, p.Type, append(path, p.Name))
		}
		genTypes = append(genTypes, &newType{
			name:       nm,
			typ:        t,
			kind:       "literal",
			properties: t.Value.(ParseLiteral).Properties,
			line:       t.Line,
		})
		return nm
	case "tuple":
		nm := nameFromPath("Tuple", path)
		typeNames[t] = nm
		for _, it := range t.Items {
			nameType(ctx, it, append(path, "Item"))
		}
		genTypes = append(genTypes, &newType{
			name:  nm,
			typ:   t,
			kind:  "tuple",
			items: t.Items,
			line:  t.Line,
		})
		return nm
	case "or":
		nm := nameFromPath("Or", path)
		typeNames[t] = nm
		for i, it := range t.Items {
			// these names depend on the ordering within the "or" type
			nameType(ctx, it, append(path, fmt.Sprintf("Item%d", i)))
		}
		// this code handles an "or" of stringLiterals (_InitializeParams.trace)
		names := make(map[string]int)
		msg := ""
		for _, it := range t.Items {
			if line, ok := names[typeNames[it]]; ok {
				// duplicate component names are bad
				msg += fmt.Sprintf("lines %d %d dup, %s for %s\n", line, it.Line, typeNames[it], nm)
			}
			names[typeNames[it]] = t.Line
		}
		// this code handles an "or" of stringLiterals (_InitializeParams.trace)
		if len(names) == 1 {
			var solekey string
			for k := range names {
				solekey = k // the sole name
			}
			if solekey == "string" { // _InitializeParams.trace
				typeNames[t] = "string"
				return "string"
			}
			// otherwise unexpected
			log.Info(ctx, "unexpected: single-case 'or' type has non-string key %s: %s", nm, solekey)
			log.Fatal(ctx, msg)
		} else if len(names) == 2 {
			// if one of the names is null, just use the other, rather than generating an "or".
			// This removes about 40 types from the generated code. An entry in goplsStar
			// could be added to handle the null case, if necessary.
			newNm := ""
			sawNull := false
			for k := range names {
				if k == "null" {
					sawNull = true
				} else {
					newNm = k
				}
			}
			if sawNull {
				typeNames[t] = newNm
				return newNm
			}
		}
		genTypes = append(genTypes, &newType{
			name:  nm,
			typ:   t,
			kind:  "or",
			items: t.Items,
			line:  t.Line,
		})
		return nm
	case "stringLiteral": // a single type, like 'kind' or 'rename'
		typeNames[t] = "string"
		return "string"
	default:
		log.Fatal(ctx, "nameType: %T unexpected, line:%d path:%v", t, t.Line, path)
		panic("unreachable in nameType")
	}
}

func nameFromPath(prefix string, path []string) string {
	nm := prefix + "_" + strings.Join(path, "_")
	// methods have slashes
	nm = strings.ReplaceAll(nm, "/", "_")
	return nm
}
