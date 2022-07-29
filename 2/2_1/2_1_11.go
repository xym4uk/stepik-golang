func sumInt(a ...int) (int, int) {
count := len(a)
sum := 0
for _, value := range a {
sum += value
}

return count, sum
}