package tsp

type out struct {
	results <-chan Result
	errors  <-chan error
	finish  chan<- bool
}

type sync struct {
	out
}

type async struct {
	hasFinished bool

	out
}

type in struct {
	results chan<- Result
	errors  chan<- error
	finish  <-chan bool
}

func init(cfg TSPConfig) (TSPIn, TSPOut) {
	results := make(chan Result, 50)
	errors := make(chan error, 10)
	finish := make(chan bool)
	o := out{results, errors, finish}
	var out TSPOut
	if cfg.Async {
		out = &sync{o}
	} else {
		out = &async{false, o}
	}
	in := &in{results, errors, finish}
	return in, out
}

func (s *async) GetResult() (Result, error) {
	select {
	case r := <-s.results:
		return r, nil
	case err := <-s.errors:
		return Result{}, err
	default:
		return Result{}, ErrorWouldBlock{}
	}
}

func (s *in) SendResult(r Result) {
	//verify it is TSP
	//verify it is better than current best
	//make copy of result
	//send
}

func (s *in) SendError(e error) {
}

func (s *async) Finish() {
	if !s.hasFinished {
		s.finish <- true
	}
}

func (s *sync) GetResult() (Result, error) {
	select {
	case r := <-s.results:
		return r, nil
	case err := <-s.errors:
		return Result{}, err
	}
}

func (s *sync) Finish() {
	s.finish <- true
}
