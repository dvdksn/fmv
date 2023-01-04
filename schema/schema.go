package schema

import (
	"fmt"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	cueyaml "cuelang.org/go/encoding/yaml"
	"gopkg.in/yaml.v3"
)

type Validator struct {
	schema cue.Value
}

func Init(schemafile string) *Validator {
	return &Validator{
		schema: createSchema(schemafile),
	}
}

func createSchema(schemafile string) cue.Value {
	cueCtx := cuecontext.New()
	schemaBytes, err := os.ReadFile(schemafile)
	if err != nil {
		fmt.Println(err)
	}
	cueSchema := cueCtx.CompileBytes(schemaBytes)
	if cueSchema.Err() != nil {
		fmt.Println(err)
	}
	cueSchema.Validate(cue.Schema())

	return cueSchema
}

func (v *Validator) Validate(metadata interface{}) bool {
	// Marshalling the metadata to YAML here, but would have
	// preferred using validate with the Go types directly
	//
	// Unfortunately, the Cue gocodec library still relies
	// on the old, deprecated cue.Runtime in a very entangled way.
	yamlMetadata, _ := yaml.Marshal(metadata)
	err := cueyaml.Validate(yamlMetadata, v.schema)
	return err == nil
}
