// A super simple worker based concurrent job system that just transforms text

package maint

import (
	"time"
	"math/rand"
)

type WorkerFn func(data string) string

type Worker struct {
	jobs chan Job
	finished chan Response
	fn workerFn
}

// Brokers all the request of jobs to various workers
type Broker struct {
	Workers []Worker
	jobReqs chan Job
}

type Job struct {
	data string
}

type Response struct {
	job Job
	response interface{}
}

// Create a new broker with numWorkers Workers instanciated
func NewBroker(numWorkers int, finishedCh chan Response, fn WorkerFn) *Broker {
	jobCh := make(chan Job)

	broker = Broker{
		Workers: make([]Worker, 0)
	}
	for i := 0; i < numWorkers; i++ {
		broker.AddWorker(jobCh, finishedCh, fn)
	}
}

func (b *Broker) run() {
	for worker := range b.Workers {
		go worker.run()
	}
}

func (b *Broker) AddWorker(jobCh chan Job, finshedCh chan Response, fn WorkerFn) {
	b.Workers = append(b.Workers, NewWorker(jobCh, finshedCh, fn))
}

// A helper func to quickly submit a bunch of jobs to the broker
func (b *Broker) SubmitJob(jobData ...string) {
	for data := range jobData {
		b.jobReqs <- Job{data}
	}
}

func NewWorker(jobCh chan Job, finshedCh chan Response) Worker {
	return Worker{
		jobs: jobCh,
		response: finishedCh,
	}
}

func (w Worker) run() {
	for range
}

func init() {
	rand.Seed(time.Now.Nanosecond())
}

func main() {

	fn := func(data string) {
		time.Wait(rand.RandIntn(1000))
		fmt.Println(data + " haz catz")
	}

	respCh := make(chan Response)
	broker := NewBroker(10, respCh)


}