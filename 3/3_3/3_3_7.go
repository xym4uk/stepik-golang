fn := func(x uint) uint {
str := strconv.FormatUint(uint64(x), 10)
runed := make([]rune, 0, len(str))

for _, value := range str {
if value%2 == 0 && value != 48 {
runed = append(runed, value)
}
}
str = string(runed)

num, _ := strconv.Atoi(str)
if num == 0 {
return uint(100)
}

return uint(num)
}