This is a simple worker pool package for go.

To install:

    go get github.com/charlesgregory/workerpool

Example of how to use:

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
