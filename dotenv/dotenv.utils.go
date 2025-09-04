package dotenv

func parseLine(line string) (string, string) {
	var key, value string
	for i, char := range line {
		if char == '=' {
			key = line[:i]
			value = line[i+1:]
			break
		}
	}
	return key, value
}
