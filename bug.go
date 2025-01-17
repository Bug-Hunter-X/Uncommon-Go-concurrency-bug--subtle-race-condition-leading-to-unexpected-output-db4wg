func main() {
    var m sync.Mutex
    var wg sync.WaitGroup
    ch := make(chan int)

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            m.Lock()
            ch <- i
            m.Unlock()
        }(i)
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    for v := range ch {
        fmt.Println(v)
    }
}