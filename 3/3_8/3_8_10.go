// Пакет и функция main уже объявлены.
// Выводить или вводить ничего не нужно!
func removeDuplicates(input chan string, output chan string) {
	defer close(output)
	var prev string
	prev = <-input
	output <- prev
	for val := range input {
		if prev == val {
		continue
		}
		output <- val
		prev = val
	}
}




