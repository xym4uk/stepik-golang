var num int
myMap := make(map[int]int)
for i := 0; i < 10; i++ {
fmt.Scan(&num)
val := num
if _, ok := myMap[num]; !ok {
myMap[val] = work(val)
}
fmt.Print(myMap[val], " ")
}