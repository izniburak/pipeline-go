package pipeline

type PipeValue interface{}
type PipeNext func(PipeValue) PipeValue

type PipeInterface interface {
	Handle(PipeValue, PipeNext) PipeValue
}

type Pipeline struct {
	value PipeValue
	pipes []PipeInterface
}

func Send(value interface{}) *Pipeline {
	return &Pipeline{value: value}
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
	return stack(p.value)
}

func (p *Pipeline) ThenReturn() PipeValue {
	return p.Then(func(value PipeValue) PipeValue {
		return value
	})
}

func (p *Pipeline) carry() func(PipeNext, PipeInterface) PipeNext {
	return func(stack PipeNext, pipe PipeInterface) PipeNext {
		return func(value PipeValue) PipeValue {
			return pipe.Handle(value, stack)
		}
	}
}
