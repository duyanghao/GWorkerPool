package main

import (
	"flag"
	"sync"

	"github.com/duyanghao/GWorkerPool/worker"
	"github.com/golang/glog"
)

// args
var (
	argConfigPath = flag.String("config", "", "The configuration filepath for server.")
)

type GWorkerPool struct {
	cfg     *Config
	workers []worker.Worker
	wg      *sync.WaitGroup
}

func NewGWorkerPool(c *Config) (*GWorkerPool, error) {
	gwp := &GWorkerPool{
		cfg:     c,
		workers: []worker.Worker{},
		wg:      &sync.WaitGroup{},
	}
	if c.PrintWorker != nil {
		printWorker, err := worker.NewPrintWorker(c.PrintWorker)
		if err != nil {
			return nil, err
		}
		gwp.workers = append(gwp.workers, printWorker)
	}
	// TODO: other workers ...
	return gwp, nil
}

func main() {
	flag.Parse()
	// load configuration
	glog.V(5).Infof("Starting Loading configuration: %s.", *argConfigPath)
	cfg, err := LoadConfig(*argConfigPath)
	if err != nil {
		glog.Errorf("Loading configuration error: %v", err)
		return
	}
	glog.V(5).Infof("Loading configuration done.")
	// NewGWorkerPool
	gwp, err := NewGWorkerPool(cfg)
	if err != nil {
		glog.Errorf("NewGWorkerPool failed: %v", err)
		return
	}
	// execute all workers
	glog.V(5).Infof("Starting Workers ...")
	for _, w := range gwp.workers {
		gwp.wg.Add(1)
		go w.Run(gwp.wg)
	}
	gwp.wg.Wait()
	glog.V(5).Infof("all Workers finished")
}
