package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/jadilet/calculator/eval"
)

const (
	MaxGoroutine = 20
)

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func runInMainGoroutine(fileName string) {
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Fprintln(os.Stdout, eval.Eval(scanner.Text()))
	}

}

func runInMultipleGoroutine(fileName string) {
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	jobs := make(chan string)
	results := make(chan string)

	wg := new(sync.WaitGroup)

	for w := 1; w <= MaxGoroutine; w++ {
		wg.Add(1)
		go worker(jobs, results, wg)
	}

	go func() {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			jobs <- scanner.Text()
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for v := range results {
		fmt.Fprintln(os.Stdout, v)
	}

}

func worker(jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		results <- eval.Eval(j)
	}
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("provide input file path as argument")
	}

	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	lines, _ := lineCounter(f)

	if lines > 999 {
		runInMultipleGoroutine(os.Args[1])
	} else {
		runInMainGoroutine(os.Args[1])
	}

}
