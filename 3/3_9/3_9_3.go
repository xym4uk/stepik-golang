done := make(chan struct{})
go func(done chan struct{}) {
work()
close(done)
}(done)
<-done