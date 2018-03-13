package workerpool

import (
	"testing"
	"time"
	"math"
	"math/rand"
)

type testPayload struct {
	radius float64
	reps int
	result int
}


func monte_carlo_pi(payload interface{})interface{} {
	var task = payload.(testPayload)
	var reps=task.reps
	var radius=task.radius
	var x, y float64
	count := 0
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	for i := 0; i < reps; i++ {
		x = random.Float64() * radius
		y = random.Float64() * radius

		if num := math.Sqrt(x*x + y*y); num < radius {
			count++
		}
	}

	task.result = count
	return task
}

func TestPool_Run(t *testing.T) {
	var samples=1000000000
	var cores=4
	p:=NewPool(cores)
	p.Start()
	go func(){
		for i:=0;i<cores;i++{
			p.NewTask(monte_carlo_pi,testPayload{radius:100.0,reps:samples/cores})
		}
	}()
	total := 0
	for i := 0; i < cores; i++ {
		var ret=<-p.RetChan
		total += ret.(testPayload).result
	}

	pi := (float64(total) / float64(samples)) * 4
	t.Log(pi)
}
func BenchmarkPool_Run_2(b *testing.B) {
	var samples=1000000000
	var cores=2
	p:=NewPool(cores)
	p.Start()
	go func(){
		for i:=0;i<cores;i++{
			p.NewTask(monte_carlo_pi,testPayload{radius:100.0,reps:samples/cores})
		}
	}()
	total := 0
	for i := 0; i < cores; i++ {
		var ret=<-p.RetChan
		total += ret.(testPayload).result
	}

	pi := (float64(total) / float64(samples)) * 4
	b.Log(pi)
}
func BenchmarkPool_Run_4(b *testing.B) {
	var samples=1000000000
	var cores=4
	p:=NewPool(cores)
	p.Start()
	go func(){
		for i:=0;i<cores;i++{
			p.NewTask(monte_carlo_pi,testPayload{radius:100.0,reps:samples/cores})
		}
	}()
	total := 0
	for i := 0; i < cores; i++ {
		var ret=<-p.RetChan
		total += ret.(testPayload).result
	}

	pi := (float64(total) / float64(samples)) * 4
	b.Log(pi)
}
func BenchmarkPool_Run_8(b *testing.B) {
	var samples=1000000000
	var cores=8
	p:=NewPool(cores)
	p.Start()
	go func(){
		for i:=0;i<cores;i++{
			p.NewTask(monte_carlo_pi,testPayload{radius:100.0,reps:samples/cores})
		}
	}()
	total := 0
	for i := 0; i < cores; i++ {
		var ret=<-p.RetChan
		total += ret.(testPayload).result
	}

	pi := (float64(total) / float64(samples)) * 4
	b.Log(pi)
}
func BenchmarkPool_Run_16(b *testing.B) {
	var samples=1000000000
	var cores=16
	p:=NewPool(cores)
	p.Start()
	go func(){
		for i:=0;i<cores;i++{
			p.NewTask(monte_carlo_pi,testPayload{radius:100.0,reps:samples/cores})
		}
	}()
	total := 0
	for i := 0; i < cores; i++ {
		var ret=<-p.RetChan
		total += ret.(testPayload).result
	}

	pi := (float64(total) / float64(samples)) * 4
	b.Log(pi)
}

