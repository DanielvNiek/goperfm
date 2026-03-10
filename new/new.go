// Package new benchmarks New vs Init
package new

type Perfm struct {
	name string
}

func (p *Perfm) Name() string {
	return p.name
}

//go:noinline
func New(name string) *Perfm {
	return &Perfm{
		name: name,
	}
}

//go:noinline
func (p *Perfm) Init(name string) *Perfm {
	p.name = name
	return p
}
