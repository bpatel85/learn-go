package algo

/*
Given an arbitrary ransom note string and another string containing letters from all the magazines,
write a function that will return true if the ransom note can be constructed from the magazines ;
otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.
*/

func toMap(str string) map[string]int {
	strMap := make(map[string]int)
	for _, c := range str {
		key := string(c)
		cnt, ok := strMap[key]
		if !ok {
			strMap[key] = 1
		} else {
			strMap[key] = cnt + 1
		}
	}

	return strMap
}

func canConstruct(ransomNote string, magazine string) bool {
	magMap := toMap(magazine)

	for _, c := range ransomNote {
		key := string(c)
		cnt, ok := magMap[key]
		if !ok || cnt <= 0 {
			return false
		} else {
			magMap[key] = cnt - 1
		}
	}

	return true
}
