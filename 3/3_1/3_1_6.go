for number, cities := range groupCity {
if number == 100 {
continue
}

for _, city := range cities {
delete(cityPopulation, city)
}
}
 