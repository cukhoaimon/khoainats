package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
)

func processFile(filename, pkgName string) error {
	src, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse Go source: %w", err)
	}

	var typeName string
	var constants []string
	var constDecls []string

	// Extract type, constants, and existing methods
	for _, decl := range node.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch ts := spec.(type) {
				case *ast.TypeSpec:
					if ident, ok := ts.Type.(*ast.Ident); ok && ident.Name == "string" {
						typeName = ts.Name.Name
					}
				case *ast.ValueSpec:
					if ts.Type == nil {
						continue
					}
					if ident, ok := ts.Type.(*ast.Ident); ok && ident.Name == typeName {
						for i, name := range ts.Names {
							val := ""
							if len(ts.Values) > i {
								valBuf := new(bytes.Buffer)
								_ = format.Node(valBuf, fset, ts.Values[i])
								val = valBuf.String()
							}
							constDecls = append(constDecls, fmt.Sprintf("%s %s = %s", name.Name, typeName, val))
							constants = append(constants, name.Name)
						}
					}
				}
			}
		}
	}

	if typeName == "" {
		return errors.New("no string-based custom type found")
	}
	if len(constants) == 0 {
		return fmt.Errorf("no constants found for type: %s", typeName)
	}

	// Begin rewriting full file
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("package %s\n\n", pkgName))
	buf.WriteString(`import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)
`)

	// Re-declare type
	buf.WriteString(fmt.Sprintf("\ntype %s string\n\n", typeName))

	// Re-declare constants
	buf.WriteString("const (\n")
	for _, decl := range constDecls {
		buf.WriteString("\t" + decl + "\n")
	}
	buf.WriteString(")\n\n")

	// Generate is valid method
	buf.WriteString(fmt.Sprintf("func (p *%s) IsValid() bool {\n", typeName))
	buf.WriteString("\tswitch *p {\n")
	for _, c := range constants {
		buf.WriteString(fmt.Sprintf("\tcase %s:\n", c))
	}
	buf.WriteString("\t\treturn true\n\t}\n\treturn false\n}\n\n")

	// Generate MarshalJSON method
	buf.WriteString(fmt.Sprintf(`
// MarshalJSON Implement json.Marshaler
func (p *%[1]s) MarshalJSON() ([]byte, error) {
	if !p.IsValid() {
		return nil, fmt.Errorf("invalid %[1]s: %%s", *p)
	}
	return json.Marshal(string(*p))
}
`, typeName))

	// Generate UnmarshalJSON method
	buf.WriteString(fmt.Sprintf(`
// UnmarshalJSON Implement json.Unmarshaler
func (p *%[1]s) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	pt := %[1]s(s)
	if !pt.IsValid() {
		return fmt.Errorf("invalid %[1]s: %%s", s)
	}
	*p = pt
	return nil
}
`, typeName))

	// Generate Value method
	buf.WriteString(fmt.Sprintf(`
// Value Implement driver.Valuer for SQL
func (p *%[1]s) Value() (driver.Value, error) {
	if !p.IsValid() {
		return nil, fmt.Errorf("invalid %[1]s: %%s", *p)
	}
	return string(*p), nil
}
`, typeName))

	// Scan Value method
	buf.WriteString(fmt.Sprintf(`
// Scan Implement sql.Scanner for SQL
func (p *%[1]s) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("%[1]s should be a string")
	}
	*p = %[1]s(str)
	if !p.IsValid() {
		return fmt.Errorf("invalid %[1]s: %%s", str)
	}
	return nil
}
`, typeName))

	// Write final buffer to file (overwrite)
	if err = os.WriteFile(filename, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write generated file: %w", err)
	}

	fmt.Printf("✅ Rewrote file with generated methods for %s: %s\n", typeName, filename)
	return nil
}
