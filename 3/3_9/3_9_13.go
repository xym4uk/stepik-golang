func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
res := make(chan int)
var i int

go func() {
defer close(res)
select {
case i = <-firstChan:
res <- i * i
case i = <-secondChan:
res <- i * 3
case <-stopChan:
break
}
}()

return res
}
