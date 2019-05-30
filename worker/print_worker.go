package worker

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/golang/glog"
)

type PrintWorkerConfig struct {
	NoOfConsumers int `yaml:"noOfConsumers,omitempty"`
	NoOfProducers int `yaml:"noOfProducers,omitempty"`
	BufferOfJobs  int `yaml:"bufferOfJobs,omitempty"`
	TotalOfJobs   int `yaml:"totalOfJobs,omitempty"`
}

func (c *PrintWorkerConfig) Validate() error {
	if c.NoOfConsumers <= 0 || c.NoOfProducers <= 0 || c.BufferOfJobs <= 0 || c.TotalOfJobs <= 0 {
		return fmt.Errorf("invalid noOfConsumers or noOfProducers or bufferOfJobs or totalOfJobs, please check.")
	}
	return nil
}

type PrintJob struct {
	randomno int
}

type PrintResult struct {
	job PrintJob
	err error
}

type PrintWorker struct {
	cfg     *PrintWorkerConfig
	jobs    chan PrintJob
	results chan PrintResult
}

func NewPrintWorker(c *PrintWorkerConfig) (*PrintWorker, error) {
	return &PrintWorker{
		cfg:     c,
		jobs:    make(chan PrintJob, c.BufferOfJobs),
		results: make(chan PrintResult, c.BufferOfJobs),
	}, nil
}

func (pw *PrintWorker) Name() string {
	return "Print"
}

func (pw *PrintWorker) Consumer(wg *sync.WaitGroup) {
	for job := range pw.jobs {
		var err error
		// do something
		pw.results <- PrintResult{job, err}
	}
	wg.Done()
}

func (pw *PrintWorker) AllocateConsumers(noOfConsumers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfConsumers; i++ {
		wg.Add(1)
		go pw.Consumer(&wg)
	}
	wg.Wait()
	close(pw.results)
}

func (pw *PrintWorker) Producer(wg *sync.WaitGroup) {
	for i := 0; i < pw.cfg.TotalOfJobs/pw.cfg.NoOfProducers; i++ {
		randomno := rand.Intn(999)
		job := PrintJob{randomno}
		pw.jobs <- job
	}
	wg.Done()
}

func (pw *PrintWorker) AllocateProducers(noOfProducers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfProducers; i++ {
		wg.Add(1)
		go pw.Producer(&wg)
	}
	wg.Wait()
	close(pw.jobs)
}

func (pw *PrintWorker) Result(done chan struct{}) {
	for result := range pw.results {
		if result.err != nil {
			glog.Errorf("Job: %v failed: %v", result.job, result.err)
		} else {
			glog.V(5).Infof("Job: %v finished successfully", result.job)
		}
	}
	done <- struct{}{}
}

func (pw *PrintWorker) Run(wg *sync.WaitGroup) {
	startTime := time.Now()
	go pw.AllocateProducers(pw.cfg.NoOfProducers)
	done := make(chan struct{})
	go pw.Result(done)
	go pw.AllocateConsumers(pw.cfg.NoOfConsumers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	glog.V(5).Infof("PrintWorker total time taken %f seconds", diff.Seconds())
	// done
	wg.Done()
}
