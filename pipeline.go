package pipeline

type PipeValue interface{}
type PipeNext func(PipeValue) PipeValue

type PipeInterface interface {
	Handle(PipeValue, PipeNext) PipeValue
}

type Pipeline struct {
	passable PipeValue
	pipes    []PipeInterface
}

func Send(passable interface{}) *Pipeline {
	return &Pipeline{passable: passable}
}

func (p *Pipeline) Through(pipes []PipeInterface) *Pipeline {
	p.pipes = pipes
	return p
}

func (p *Pipeline) Then(destination PipeNext) PipeValue {
	carry := p.carry()
	stack := destination
	for i := len(p.pipes) - 1; i >= 0; i-- {
		stack = carry(stack, p.pipes[i])
	}
	return stack(p.passable)
}

func (p *Pipeline) ThenReturn() PipeValue {
	return p.Then(func(passable PipeValue) PipeValue {
		return passable
	})
}

func (p *Pipeline) carry() func(PipeNext, PipeInterface) PipeNext {
	return func(stack PipeNext, pipe PipeInterface) PipeNext {
		return func(passable PipeValue) PipeValue {
			return pipe.Handle(passable, stack)
		}
	}
}
