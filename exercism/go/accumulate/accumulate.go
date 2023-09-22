package accumulate

func Accumulate(collection []string, converter func(string) string) []string {

	var output []string
	for _, value := range collection {
		output = append(output, converter(value))
	}
	return output
}
