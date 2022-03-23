package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hello from here!")
}

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
//1.2 Command-Line Arguments

// Echo1 prints its command-line arguments.
func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(os.Args[0])
	fmt.Println(s)
}

//////////////////////////////////////////////////////////////////

func Echo2() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println(i, arg)
	}
	fmt.Println(s)
}

//////////////////////////////////////////////////////////////////

// Echo3 - this would be more efficient for large data because it doesn't have to deal with the garbage-collecting of s
func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

//////////////////////////////////////////////////////////////////

func Echo31() {
	fmt.Println(os.Args[1:])
}

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
//1.3 Finding Duplicate Lines

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
// Ctrl Z + enter will end the input stream in windows command line, in linux is Ctrl D
func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//////////////////////////////////////////////////////////////////

// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
func Dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		CountLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			CountLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func CountLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//////////////////////////////////////////////////////////////////

func Dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
//1.4 Animated Gifs

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

//func main() {
//	lissajous(os.Stdout)
//}

// Lissajous generates GIF animations of random Lissajous figures.
// this one is not working, and I don't know why
func Lissajous(cycles float64, out io.Writer) {
	const (
		// cycles = 5 // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [size..+ size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	if cycles == 0 {
		cycles = 5
	}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
//1.5 Fetching URL

// Fetch prints the content found at a URL.
func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading%s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

//////////////////////////////////////////////////////////////////

// Fetch2 uses instead of ReadAll, io.Copy to copy directly the resp.Body to os.Stdout
// This is ex1.7, ex1.8 and ex1.9
func Fetch2() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading%s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("With status code %s\n", resp.Status)
	}
}

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
//1.6 Fetching URLs Concurrently

//func main() {
//	start := time.Now()
//	ch := make(chan string)
//	for _, url := range os.Args[1:] {
//		go Fetch3(url, ch) // start a goroutine
//	}
//	for range os.Args[1:] {
//		fmt.Println(<-ch) // receive from channel ch
//	}
//	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
//}

// Fetch3 fetches URLs in parallel and reports their times and sizes.
func Fetch3(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
//1.7 A Web Server

// Server1 is a minimal "echo" server.
func Server1() {
	http.HandleFunc("/hello", handler) //each request calls handler
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

//////////////////////////////////////////////////////////////////

var mu sync.Mutex
var count int

// Server2 is a minimal "echo" and counter server.
func Server2() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

// handler2 echoes the Path component of the requested URL.
func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n %v", r.URL.Path, r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mu.Unlock()
}

//////////////////////////////////////////////////////////////////

func Server3() {
	http.HandleFunc("/home", handler3)
	http.HandleFunc("/lis", handlerLis)
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

// handler3 echoes the HTTP request.
func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func handlerLis(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles, _ := strconv.ParseFloat(r.Form["cycles"][0], 64)
	Lissajous(cycles, w)
}
