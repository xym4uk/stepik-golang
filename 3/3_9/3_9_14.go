func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
res := make(chan int)

go func() {
var sum int
defer close(res)
for {
select {
case i := <-arguments:
sum += i
case <-done:
res <- sum
return
}
}
}()

return res
}
