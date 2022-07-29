func vote(x int, y int, z int) int {
arr := [3]int{x, y, z}

zero := 0
one := 0
for _, value := range arr {
if value == 0 {
zero++
} else {
one++
}
}

if zero > one {
return 0
}

return 1
}