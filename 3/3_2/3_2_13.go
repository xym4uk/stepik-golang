func adding(x, y string) int64 {
str1, str2 := "", ""
var num1, num2 int
for _, value := range x {
if unicode.IsNumber(rune(value)) {
str1 += string(value)
}
}
for _, value := range y {
if unicode.IsNumber(rune(value)) {
str2 += string(value)
}
}

num1, _ = strconv.Atoi(str1)
num2, _ = strconv.Atoi(str2)

return int64(num1 + num2)
}