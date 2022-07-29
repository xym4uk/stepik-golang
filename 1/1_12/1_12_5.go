var workArray [10]uint8
var arr [6]uint8
var temp uint8
for i := 0; i < 10; i++ {
fmt.Scan(&workArray[i])
}

for i := 0; i < 6; i++ {
fmt.Scan(&arr[i])
}

for i := 0; i < 6; i += 2 {
temp = workArray[arr[i]]
workArray[arr[i]] = workArray[arr[i+1]]
workArray[arr[i+1]] = temp
}

for i := 0; i < 10; i++ {
fmt.Print(workArray[i], " ")
}