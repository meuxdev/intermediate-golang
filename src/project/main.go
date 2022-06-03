package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker with id %d started...\n", w.Id)
				fib := Fibo(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d ended..\n Fib Result: %d\n--------", w.Id, fib)
			case <-w.QuitChan:
				fmt.Printf("Worker with di %d Stopped\n", w.Id)
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: worker,
		MaxWorkers: maxWorkers,
		JobQueue:   jobQueue,
	}
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8081"
	)
	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)
	fmt.Printf("Server Fibonacci Running: -- \n Port: %s\n http://localhost%s/fib", port, port)
	dispatcher.Run()
	// http://localhost:8081/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})
	// run the server
	log.Fatal(http.ListenAndServe(port, nil))
}

func Fibo(n int) int {
	if n <= 1 {
		return n
	}

	return Fibo(n-1) + Fibo(n-2)
}

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))

	if err != nil {
		http.Error(w, "Invalid Delay Time--", http.StatusBadRequest) // 400
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid Value type-- Integer cast error", http.StatusBadRequest)
	}

	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Invalid name type-- String cast", http.StatusBadRequest)
	}

	job := Job{Name: name, Number: value, Delay: delay}

	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}
