func main() {
var a, b int
_, err := fmt.Scan(&a, &b)
if err != nil {
fmt.Println("ошибка")
return
}
res, err := divide(a, b)
if err != nil {
fmt.Println("ошибка")
return
}
fmt.Println(res)
}