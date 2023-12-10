package main

import "fmt"
import "time"

//Step 1
// worker function 引數： id, a jobs channel, and a results channel

func worker(id int, jobs <-chan int, results chan<- int){
	// iterates over the jobs channel
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		// sends the result to the results channel
		results <- j * 2
	}
}

func main() {
	//宣告工作數目
	const numJobs = 5
	
	//Step 2 : 
	//建立job channel
	jobs := make(chan int, numJobs)
	//建立results channel
	results := make(chan int, numJobs)

	//Step 3 :
	//建立3個goroutines for the worker function
	for i:=1; i <= 3; i++{
		go worker(i, jobs, results)
	}
	
	//Step 4 : 
	//發送工作到jobs channel
	for j:=1; j<=numJobs; j++{
		fmt.Println("input : ", j)
		jobs <- j
	}
	
	//Step 5 :
	//關閉jobs channel
	close(jobs)
	fmt.Println("close : ")
	//Step 6 :
	//get data from results
	for a := 1; a <= numJobs; a++{
		value := <- results
		fmt.Println("result value : ", value)
	}
	
	//Step 7 :
	//關閉result channel
	close(results)


}