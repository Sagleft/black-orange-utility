package main

type solution struct{}

func newSolution() solution {
	return solution{}
}

func (sol *solution) run() error {
	machine := newWeirdMachine()
	machine.work()
	// TODO
	return nil
}
