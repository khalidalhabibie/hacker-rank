package warm

func JumpingOnCloud(c []int32) int32 {

	var counter int32
	lengthOfC := len(c)

	for i := 0; i < lengthOfC-1; {
		if (i+2) == lengthOfC || c[i+2] == 1 {
			i++
			counter++
		} else {
			i += 2
			counter++
		}
	}

	return counter

}
