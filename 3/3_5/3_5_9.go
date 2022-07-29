func main() {
sum := 0
num := bufio.NewScanner(os.Stdin)

for num.Scan() {
currentNum, _ := strconv.Atoi(num.Text())
sum += currentNum
}

io.WriteString(os.Stdout, strconv.Itoa(sum))
}