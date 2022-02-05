package warm

func GetUniqeElementSlice(slices ...[]string) []string {

	checker := make(map[string]int)

	// append all data to new slices
	var newSlice []string
	for i := range slices {
		newSlice = append(newSlice, slices[i]...)

	}

	for i := range newSlice {
		checker[newSlice[i]] += 1

	}

	// get new keys and append
	result := []string{}
	result = append(result, newSlice...)

	return result

}
