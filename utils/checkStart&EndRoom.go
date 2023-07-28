package utils

func CheckStartEndRoom(tab []string) bool {
	end := 0
	start := 0

	for _, val := range tab {
		if val == "##start" || val == "##end" {
			end++
			start++
		}
	}
	if end == 2 && start == 2 {
		return true
	}
	return false
}