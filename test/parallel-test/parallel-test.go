package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
    fmt.Println(runtime.GOMAXPROCS(0))

    var jobs []int=makeIncrementingArray(10)

    const workers int=4
    var workgroup sync.WaitGroup
    var jobsChannel chan int=make(chan int)

    for i:=0;i<workers;i++ {
        workgroup.Add(1)
        go worker("worker"+fmt.Sprint(i),jobsChannel,&workgroup)
    }

    for i := range jobs {
        fmt.Printf("starting job: %v (%v/%v)\n",jobs[i],i,len(jobs))
        jobsChannel <- jobs[i]
    }

    close(jobsChannel)
    workgroup.Wait()
}

func worker(
    name string,
    jobs <-chan int,
    workgroup *sync.WaitGroup,
) {
	for {
        var job int
        var ok bool
        job,ok=<-jobs

        if !ok {
            workgroup.Done()
            return
        }

		fmt.Printf("%v: working on %v\n",name,job)

        time.Sleep(
            time.Duration(rand.Intn(5-3)+3)*
            time.Second,
        )
	}
}

func makeIncrementingArray(high int) []int {
    var res []int

    for i:=0;i<high;i++ {
        res=append(res,i)
    }

    shuffleArray[int](res)
    return res
}

/** shuffle an array */
func shuffleArray[T any](array []T) {
    rand.Shuffle(len(array),func (i int,j int) {
        (array)[i],(array)[j]=(array)[j],(array)[i]
    })
}