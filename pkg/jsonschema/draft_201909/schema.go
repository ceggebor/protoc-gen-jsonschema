package draft_201909

import (
	"github.com/ceggebor/protoc-gen-jsonschema/pkg/jsonschema"
	"github.com/ceggebor/protoc-gen-jsonschema/pkg/utils"
	"github.com/iancoleman/orderedmap"
)

type Schema struct {
	Version         string                 `json:"$schema,omitempty"`
	ID              string                 `json:"$id,omitempty"`
	Anchor          string                 `json:"$anchor,omitempty"`
	RecursiveAnchor string                 `json:"$recursiveAnchor,omitempty"`
	Ref             string                 `json:"$ref,omitempty"`
	RecursiveRef    string                 `json:"$recursiveRef,omitempty"`
	Definitions     *orderedmap.OrderedMap `json:"$defs,omitempty"`
	Comments        string                 `json:"$comment,omitempty"`

	AllOf []*Schema `json:"allOf,omitempty"`
	AnyOf []*Schema `json:"anyOf,omitempty"`
	OneOf []*Schema `json:"oneOf,omitempty"`
	Not   *Schema   `json:"not,omitempty"`

	If               *Schema                `json:"if,omitempty"`
	Then             *Schema                `json:"then,omitempty"`
	Else             *Schema                `json:"else,omitempty"`
	DependentSchemas *orderedmap.OrderedMap `json:"dependentSchemas,omitempty"`

	PrefixItems []*Schema `json:"items,omitempty"`
	Items       *Schema   `json:"additionalItems,omitempty"`
	Contains    *Schema   `json:"contains,omitempty"`

	Properties           *orderedmap.OrderedMap `json:"properties,omitempty"`
	PatternProperties    *orderedmap.OrderedMap `json:"patternProperties,omitempty"`
	AdditionalProperties *Schema                `json:"additionalProperties,omitempty"`
	PropertyNames        *Schema                `json:"propertyNames,omitempty"`

	Type              string              `json:"type,omitempty"`
	Enum              []any               `json:"enum,omitempty"`
	Const             *any                `json:"const,omitempty"`
	MultipleOf        *int                `json:"multipleOf,omitempty"`
	Maximum           *float64            `json:"maximum,omitempty"`
	ExclusiveMaximum  *float64            `json:"exclusiveMaximum,omitempty"`
	Minimum           *float64            `json:"minimum,omitempty"`
	ExclusiveMinimum  *float64            `json:"exclusiveMinimum,omitempty"`
	MaxLength         *int                `json:"maxLength,omitempty"`
	MinLength         *int                `json:"minLength,omitempty"`
	Pattern           string              `json:"pattern,omitempty"`
	MaxItems          *int                `json:"maxItems,omitempty"`
	MinItems          *int                `json:"minItems,omitempty"`
	UniqueItems       *bool               `json:"uniqueItems,omitempty"`
	MaxContains       *int                `json:"maxContains,omitempty"`
	MinContains       *int                `json:"minContains,omitempty"`
	MaxProperties     *int                `json:"maxProperties,omitempty"`
	MinProperties     *int                `json:"minProperties,omitempty"`
	Required          []string            `json:"required,omitempty"`
	DependentRequired map[string][]string `json:"dependentRequired,omitempty"`

	Format string `json:"format,omitempty"`

	ContentEncoding  string  `json:"contentEncoding,omitempty"`
	ContentMediaType string  `json:"contentMediaType,omitempty"`
	ContentSchema    *Schema `json:"contentSchema,omitempty"`

	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Default     *any   `json:"default,omitempty"`
	Deprecated  *bool  `json:"deprecated,omitempty"`
	ReadOnly    *bool  `json:"readOnly,omitempty"`
	WriteOnly   *bool  `json:"writeOnly,omitempty"`
	Examples    []any  `json:"examples,omitempty"`

	Extras map[string]any `json:"-"`
}

func New(schema *jsonschema.Schema) *Schema {
	return deepCopy(schema)
}

func deepCopy(origin *jsonschema.Schema) *Schema {
	if origin == nil {
		return nil
	}

	dst := &Schema{}
	dst.Version = origin.Version
	dst.ID = origin.ID
	dst.Anchor = origin.Anchor
	dst.RecursiveAnchor = origin.DynamicAnchor
	if origin.Ref.String() != "" {
		dst.Ref = "#/$defs/" + origin.Ref.String()
	}
	dst.RecursiveRef = origin.DynamicRef
	dst.Definitions = deepCopyMap(origin.Definitions)
	dst.Comments = origin.Comments

	dst.AllOf = deepCopyArray(origin.AllOf)
	dst.AnyOf = deepCopyArray(origin.AnyOf)
	dst.OneOf = deepCopyArray(origin.OneOf)
	dst.Not = deepCopy(origin.Not)

	dst.If = deepCopy(origin.If)
	dst.Then = deepCopy(origin.Then)
	dst.Else = deepCopy(origin.Else)
	dst.DependentSchemas = deepCopyMap(origin.DependentSchemas)

	dst.PrefixItems = deepCopyArray(origin.PrefixItems)
	dst.Items = deepCopy(origin.Items)
	dst.Contains = deepCopy(origin.Contains)

	dst.Properties = deepCopyMap(origin.Properties)
	dst.PatternProperties = deepCopyMap(origin.PatternProperties)
	dst.AdditionalProperties = deepCopy(origin.AdditionalProperties)
	dst.PropertyNames = deepCopy(origin.PropertyNames)

	dst.Type = origin.Type
	dst.Enum = utils.CopyAnyArray(origin.Enum)
	dst.Const = utils.CopyAnyP(origin.Const)
	dst.MultipleOf = utils.CopyIntP(origin.MultipleOf)
	dst.Maximum = utils.CopyFloat64P(origin.Maximum)
	dst.ExclusiveMaximum = utils.CopyFloat64P(origin.ExclusiveMaximum)
	dst.Minimum = utils.CopyFloat64P(origin.Minimum)
	dst.ExclusiveMinimum = utils.CopyFloat64P(origin.ExclusiveMinimum)
	dst.MaxLength = utils.CopyIntP(origin.MaxLength)
	dst.MinLength = utils.CopyIntP(origin.MinLength)
	dst.Pattern = origin.Pattern
	dst.MaxItems = utils.CopyIntP(origin.MaxItems)
	dst.MinItems = utils.CopyIntP(origin.MinItems)
	dst.UniqueItems = utils.CopyBoolP(origin.UniqueItems)
	dst.MaxContains = utils.CopyIntP(origin.MaxContains)
	dst.MinContains = utils.CopyIntP(origin.MinContains)
	dst.MaxProperties = utils.CopyIntP(origin.MaxProperties)
	dst.MinProperties = utils.CopyIntP(origin.MinProperties)
	dst.Required = utils.CopyStringArray(origin.Required)
	dst.DependentRequired = utils.CopyMapStringArray(origin.DependentRequired)

	dst.Format = origin.Format

	dst.ContentEncoding = origin.ContentEncoding
	dst.ContentMediaType = origin.ContentMediaType
	dst.ContentSchema = deepCopy(origin.ContentSchema)

	dst.Title = origin.Title
	dst.Description = origin.Description
	dst.Default = utils.CopyAnyP(origin.Default)
	dst.Deprecated = utils.CopyBoolP(origin.Deprecated)
	dst.ReadOnly = utils.CopyBoolP(origin.ReadOnly)
	dst.WriteOnly = utils.CopyBoolP(origin.WriteOnly)

	dst.Extras = utils.CopyMapAny(origin.Extras)
	return dst
}

func deepCopyArray(arr []*jsonschema.Schema) []*Schema {
	dst := make([]*Schema, len(arr))
	for i, schema := range arr {
		dst[i] = deepCopy(schema)
	}
	return dst
}

func deepCopyMap(schemaMap jsonschema.SchemaMap) *orderedmap.OrderedMap {
	if schemaMap == nil {
		return nil
	}
	dst := orderedmap.New()
	for _, key := range schemaMap.Keys() {
		schema, _ := schemaMap.Get(key)
		dst.Set(key, deepCopy(schema))
	}
	return dst
}
