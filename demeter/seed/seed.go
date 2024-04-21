package seed

import (
	"reflect"
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
	ReadOne() any

	// Method to return an specified amount of seeds
	ReadMany(int) []any

	// Method to return all seeds, only 1 if no limit is configured
	ReadAll() []any

	// This method gets the size of the seed source, -1 if unlimited
	GetSize() int

    // This method gets the metadata of the seed, forces SeedMeta usage
    GetMeta() *SeedMeta
}

// Seeds netadata to perform operations on data and keep state of it
type SeedMeta struct {
    
    // Seed name
    Name string

    // Seed props
    Props []prop
}

type prop struct {
    Name string 
    Kind reflect.Kind
}

// An utility function to get the props of a struct of type any,
// allows seed implementations to follow the same guideline when
// setting up the initialization of the inner struct
func ObtainProps(base any, name string) SeedMeta {

    meta := SeedMeta {
        Name: name,
        Props: []prop{},
    }

	// check if base is a struct otherwise return no headers
	t := reflect.TypeOf(base)

	if t.Kind() != reflect.Struct {
		return meta
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		
        switch kind := field.Type.Kind(); kind {
            case reflect.Int, reflect.String:
                meta.Props = append(meta.Props, prop{Name: field.Name, Kind: kind,})
        }
	}

	return meta
}

type CountSeed struct {
	counter int
    meta SeedMeta
}

func NewCountSeed(base any, start int) *CountSeed {
	return &CountSeed{
		counter: start,
		meta: ObtainProps(base, "CountSeed"),
	}
}

func (s *CountSeed) GetProps() []string {
    
    props := make([]string, len(s.meta.Props))
    for i, prop := range s.meta.Props {
        props[i] = prop.Name
    }

	return props
}

func (s *CountSeed) GetSize() int {
	return -1
}
