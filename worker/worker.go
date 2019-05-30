package worker

import "sync"

type Worker interface {
	// AllocateConsumers
	AllocateConsumers(noOfConsumers int)
	// AllocateProducer
	AllocateProducers(noOfProducers int)
	// Consumers
	Consumer(wg *sync.WaitGroup)
	// Producer
	Producer(wg *sync.WaitGroup)
	// Run
	Run(wg *sync.WaitGroup)
	// Result
	Result(done chan struct{})
	// Human-readable name of the worker.
	Name() string
}
