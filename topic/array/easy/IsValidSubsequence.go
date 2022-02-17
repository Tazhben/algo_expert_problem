package easy

func IsValidSubsequence(array []int, sequence []int) bool {
	//brute force
	arrayB := make([]bool, len(array))
	startI := 0
	for j := 0; j < len(sequence); j++ {
		check := false
		for i := startI; i < len(array); i++ {
			if sequence[j] == array[i] && arrayB[i] == false {
				check = true
				startI = i
				arrayB[i] = true
				break
			}
		}
		if !check {
			return check
		}
	}
	return true
}

/*
array := []int{1, 1, 6, 1}
sequence := []int{1, 1, 1, 6}
fmt.Println(easy.IsValidSubsequence(array, sequence))*/