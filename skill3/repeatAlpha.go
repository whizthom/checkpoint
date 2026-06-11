package main

func RepeatAlpha1(s string) string {
	
	result := []rune{}

	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			count := int(ch - 'A' + 1)
			for i := 0; i < count; i++ {
				result = append(result, ch)
			}
		} else if ch >= 'a' && ch <= 'z' {
			count := int(ch - 'a' + 1)
			for i := 0; i < count; i++ {
				result = append(result, ch)
			}
		} else {
			result = append(result, ch)
		}
	}

	return string(result)
}

func RepeatAlpha(s string) string {
	
	result := []rune{}

	for _, ch:= range s{
		if ch >= 'A' && ch <= 'Z'{
			count := int(ch - 'A' + 1)

			for i:=0; i < count; i++{
				result = append(result, ch)
			}
		}else if 
		ch >= 'a' && ch <= 'z'{
			count := int(ch - 'a' + 1)

			for i:=0; i < count; i++{
				result = append(result, ch)
			}
		}else{
			result = append(result, ch)
		}
	}
	return string(result)
}
