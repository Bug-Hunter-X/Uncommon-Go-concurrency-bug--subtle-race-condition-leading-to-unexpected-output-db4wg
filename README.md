# Uncommon Go Concurrency Bug: Subtle Race Condition

This repository demonstrates a subtle race condition in a Go program that uses goroutines, mutexes, and channels. The bug is not immediately obvious and can lead to unexpected or incorrect output.

## Bug Description

The program aims to concurrently process numbers from 0 to 9, using goroutines to send these numbers through a channel protected by a mutex.  A race condition occurs because `wg.Wait()` is called in the goroutine which closes the channel, before all goroutines have finished writing to the channel. This can lead to some numbers being missed or the program hanging.  The core issue lies in the timing of closing the channel.  Because the `wg.Wait()` is in a separate goroutine, there's no guarantee all numbers have been sent before the channel is closed.

## Solution

The solution ensures that the channel is closed only after all goroutines have completed sending their values. This is achieved by moving the `close(ch)` statement within the main function after the `wg.Wait()`. This guarantees that the channel is closed only after all writes are complete.

## How to Run

1. Clone the repository.
2. Navigate to the repository directory.
3. Run `go run bug.go` to see the buggy behavior.
4. Run `go run bugSolution.go` to see the corrected behavior. 
