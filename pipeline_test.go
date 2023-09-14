package pipeline

import (
	"strings"
	"testing"
)

//
// StringPipes
//

type UpperCasePipe struct{}

func (u *UpperCasePipe) Handle(value PipeValue, next PipeNext) PipeValue {
	// get value
	text := value.(string)

	capitalized := strings.ToUpper(text)
	return next(capitalized)
}

type TrimSpacePipe struct{}

func (t *TrimSpacePipe) Handle(value PipeValue, next PipeNext) PipeValue {
	// get value
	text := value.(string)

	trimmed := strings.Trim(text, " ")
	return next(trimmed)
}

type ReplaceTextPipe struct{}

func (r *ReplaceTextPipe) Handle(value PipeValue, next PipeNext) PipeValue {
	// get value
	text := value.(string)

	replaced := strings.ReplaceAll(text, "BURAKDEMIRTAS.ORG", "BUKI.DEV")
	return next(replaced)
}

func TestStringPipelines(t *testing.T) {
	pipes := []PipeInterface{new(UpperCasePipe), new(TrimSpacePipe)}
	data := Send("   burakdemirtas.org   ").Through(pipes).ThenReturn()

	if data != "BURAKDEMIRTAS.ORG" {
		t.Errorf("got %s, want BURAKDEMIRTAS.ORG", data)
	}

	pipes = append(pipes, new(ReplaceTextPipe))
	data = Send("   burakdemirtas.org   ").Through(pipes).ThenReturn()

	if data != "BUKI.DEV" {
		t.Errorf("got %s, want BUKI.DEV", data)
	}

	data = Send("   burakdemirtas.org   ").Through(pipes).Then(func(value PipeValue) PipeValue {
		return strings.ReplaceAll(value.(string), ".DEV", "")
	})

	if data != "BUKI" {
		t.Errorf("got %s, want BUKI", data)
	}
}

//
// ArrayPipes
//

type DoublePipe struct{}

func (s *DoublePipe) Handle(value PipeValue, next PipeNext) PipeValue {
	// get value
	numbers := value.([]int)

	result := make([]int, len(numbers))
	for i, v := range numbers {
		result[i] = v * v
	}

	return next(result)
}

type SumPipe struct{}

func (s *SumPipe) Handle(value PipeValue, next PipeNext) PipeValue {
	// get value
	numbers := value.([]int)

	sum := 0
	for _, v := range numbers {
		sum += v
	}

	return next(sum)
}

func TestArrayPipelines(t *testing.T) {
	pipes := []PipeInterface{new(DoublePipe), new(SumPipe)}
	result := Send([]int{1, 2, 3}).Through(pipes).ThenReturn()

	if result != 14 {
		t.Errorf("got %d, want 14", result)
	}
}

//
// MapPipes
//

func TestMapPipelines(t *testing.T) {

}
