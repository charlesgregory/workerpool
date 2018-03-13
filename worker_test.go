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
type work struct{
	a,b,res int
}
func doWork(payload interface{})interface{}{
	var task=payload.(work)
	task.res=task.a+task.b
	return task
}
func TestSimple(t *testing.T){
	p:=NewPool(4)
	p.Start()
	for i:=0;i<100000;i++{
		p.NewTask(doWork,work{a:i,b:i+1})
	}
	p.Close()
	for task:=range p.RetChan.Recv{
		w:=task.(work)
		if !(w.a+w.b==w.res){
			t.Fail()
		}
	}
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
		var ret=<-p.RetChan.Recv
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
		var ret=<-p.RetChan.Recv
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
		var ret=<-p.RetChan.Recv
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
		var ret=<-p.RetChan.Recv
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
		var ret=<-p.RetChan.Recv
		total += ret.(testPayload).result
	}

	pi := (float64(total) / float64(samples)) * 4
	b.Log(pi)
}

