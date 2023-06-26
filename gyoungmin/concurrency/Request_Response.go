package concurrency

type Response struct {
	Num      int
	WorkerID int
}

type Request struct {
	Num  int
	Resp chan Response
}

func PlusOnService(reqs <-chan Request, WorkerID int) {

	for req := range reqs {
		go func(req Request) {
			defer close(req.Resp)
			req.Resp <- Response{req.Num, WorkerID}
		}(req)
	}
}
