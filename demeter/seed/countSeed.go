package seed

// A CountSeed allows the drill of seeds that keep an inner counter
// at time of writing only the basic prop++ counter is defined
// **Props type accepted:  int
//   any other type is ignored
type CountSeed struct {
	counter int
    meta SeedMeta
    base any
}

func NewCountSeed(base any, start int) *CountSeed {
	return &CountSeed{
		counter: start,
		meta: ObtainProps(base, "CountSeed", Int),
        base: base,
	}
}

func (s *CountSeed) GetProps() []string {
    
    props := make([]string, len(s.meta.Props))
    for i, prop := range s.meta.Props {
        props[i] = prop.Name
    }

	return props
}

func (s *CountSeed) ReadOne() []prop {

    for _, prop := range s.meta.Props {
        (&s.meta).GetField(prop.Name).SetInt(10)
    }

    return s.meta.Props
}

// GetSize must return -1 for infinite seeds
func (s *CountSeed) GetSize() int {
	return -1
}

func (s *CountSeed) GetMeta() *SeedMeta {
    return &s.meta
}
