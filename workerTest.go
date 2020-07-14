package main

/*
func worker(wg *sync.WaitGroup, id int, jobs <-chan int, results chan<- bool) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j, "len: ", cap(jobs))
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- false
	}
	fmt.Println("worker", id, "Completed");
	wg.Done()
}

func allocate(noOfJobs int, jobs chan <- int) {
	for i := 1; i <= noOfJobs; i++ {
		fmt.Println("Submitting job: ", i)
		jobs <- i
	}
	close(jobs)
}

func result(done chan bool, results <-chan  bool) {
	for result := range results {
		fmt.Println("Got Result for: ", result)
	}
	done <- true
}

func createWorkers(numWorkers int, jobs <- chan int, results chan <- bool) {
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		//fmt.Println("Created worker: ", i)
		go worker(&wg, i, jobs, results)
	}
	wg.Wait()
	close(results)
}

func main() {

	const numJobs = 100
	const numWorkers = 5
	jobs := make(chan int,2)
	results := make(chan bool)
	done := make(chan bool)


	go result(done, results)


	go createWorkers(numWorkers, jobs, results)

	allocate(numJobs, jobs)
	<- done

}
 */