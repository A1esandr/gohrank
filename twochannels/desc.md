You need to write the function 
```
func Merge2Channels(f func (int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) 
```
in package main.

Description of her work:

do the following n times

* read one number from each of the two channels in1 and in2, let's call them x1 and x2.
* calculate f (x1) + f (x2)
* write the received value to out

The Merge2Channels function must be non-blocking and return immediately.

The function f can run for a long time, waiting for something or performing calculations.

Input format

* The number of iterations is passed through the n argument.
* Integers are supplied through the channel arguments in1 and in2.
* The function for processing numbers before addition is passed through the f argument.

Output format

The channel for outputting the results is passed through the out argument.