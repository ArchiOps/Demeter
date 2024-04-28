package seed

import (
    "slices"
	"reflect"
)

const (
    Int     reflect.Kind = reflect.Int
    String  reflect.Kind = reflect.String
)

// The Seed interface allows the implementation of Seeds
// a Seed is the concept of a crop source, the seed must provide
// the availity to load data, keep internal state and refresh state.
// ...
// Each seed configures a inner struct of type Any, this struct is used
// to configure the data that is retrieved, since the struct fields are
// not known in advance, the seed must include logic to handle a proper
// usage of the struct
type Seed interface {

	// Get seed props
	GetProps() []string

	// Filter seed props, use in conjunction with GetProps to update them
	WithProps([]string) int

	// These methods returns seeds ready to drill
	// Method to return a single seed
	ReadOne() []prop

	// Method to return an specified amount of seeds
	ReadMany(int) []any

	// Method to return all seeds, only 1 if no limit is configured
	ReadAll() []any

	// This method gets the size of the seed source, -1 if unlimited
	GetSize() int

    // This method gets the metadata of the seed, forces SeedMeta usage
    GetMeta() *SeedMeta
}

// Seeds metadata to perform operations on data and keep state of it
type SeedMeta struct {
    
    // Seed name
    Name string

    // Seed props
    Props []prop

    // Seed names
    Names map[string]NameIndex

    // Seed struct
    Struct reflect.Value
}

func (s *SeedMeta) GetField(name string) reflect.Value {
    
    field, ok := s.Names[name]

    // If no name found then return the struct
    if !ok {
        return s.Struct
    }

    return s.Struct.Field(field.field)
}

// Tuple for prop and field indexes, used by names map to get field or prop directly
type NameIndex struct {
    prop  int
    field int
}

// Property that can be associated with a seed, seeds can have zero to N properties
type prop struct {
    Name    string 
    Kind    reflect.Kind
}

// An utility function to get the props of a struct of type any,
// allows seed implementations to follow the same guideline when
// setting up the initialization of the inner struct
func ObtainProps(base any, name string, types ...reflect.Kind) SeedMeta {

    // TODO: make sure base is a pointer to struct

    meta := SeedMeta {
        Name: name,
        Props: []prop{},
        Names: make(map[string]NameIndex),
        Struct: reflect.ValueOf(base).Elem(),
    }

	// check if base is a struct otherwise return no headers
	t := reflect.ValueOf(base).Elem().Type()
	if t.Kind() != reflect.Struct {
		return meta
	}

    // iterate over each field to create props
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

        switch kind := field.Type.Kind(); kind {
            case reflect.Int, reflect.String:
                if len(types) == 0 || slices.Contains(types, kind){
                    meta.Props = append(meta.Props, prop{Name: field.Name, Kind: kind,})
                    meta.Names[field.Name] = NameIndex{prop: len(meta.Props) - 1, field: i,}
                }
        }
	}

	return meta
}
