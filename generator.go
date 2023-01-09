package iter

import "github.com/rodaine/iter/generator"

// FromGenerator creates an Iterator from a generator.Generator.
func FromGenerator[E any](gen *generator.Generator[E]) Iterator[E] {
	return FromCore[E](generatorCore[E]{gen: gen})
}

// FromGeneratorFunc creates a generator-based Iterator. This function is
// equivalent to calling generator.New, then FromGenerator. If the generator
// will not be exhausted or if it's infinite, FromGenerator should be used
// instead so that resources can be released via Generator.Close.
func FromGeneratorFunc[E any](genFn func(yield func(E))) Iterator[E] {
	gen := generator.New(genFn)
	return FromGenerator(gen)
}

type generatorCore[E any] struct {
	gen *generator.Generator[E]
}

func (g generatorCore[E]) Next() (next E, ok bool) { return g.gen.Next() }
