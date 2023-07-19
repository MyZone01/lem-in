package utils

func RangePaths(Paths [][]string) {
	for i := 0; i < len(Paths)-1; {
		if len(Paths[i]) > len(Paths[i+1]) {
			tmp := Paths[i]
			Paths[i] = Paths[i+1]
			Paths[i+1] = tmp
			i = 0
		} else {
			i++
		}

	}
}