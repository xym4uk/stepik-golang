func minimumFromFour() int {
var a, b, c, d int
fmt.Scan(&a, &b, &c, &d)
arr := [4]int{a, b, c, d}

min := a
for _, value := range arr {
if value < min {
min = value
}
}
return min
}