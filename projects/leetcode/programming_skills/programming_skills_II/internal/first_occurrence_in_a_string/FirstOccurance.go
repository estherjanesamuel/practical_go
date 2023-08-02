package firstoccurrenceinastring

/*
func strStr(haystack, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; i++ {
			if j == len(needle) {
				return i
			}
			if i + j == len(haystack) {
				return -1
			}
			if []rune(needle)[j] != []rune(haystack)[i + j] {
				break
			}
		}
	}
}
*/

func buildKmp(n string) []int{
    table := make([]int,len(n),len(n))
    for i,j := 0,1; j < len(n); {
        switch{
            case n[j] == n[i]:
                table[j] = i + 1
                i++
                j++
            case i == 0:
                 j++
            default:
                i = table[i - 1]
        }
    }
    
    return table
}

func strStr(haystack string, needle string) int {
    if len(needle) == 0{
        return 0
    }
    
    if len(haystack) == 0{
        return -1
    }

    t :=  buildKmp(needle)
    res := -1 
    
    for i,j := 0,0; i <= len(haystack);{
        if j == len(needle){
            res =  i - len(needle)
            break
        }
        
        if i == len(haystack){
            break
        }
        
        if haystack[i] == needle[j] {
            i++
            j++
        }else{
            if j == 0{
                i++
            }else{
                j = t[j - 1]
            }
        }
    }
    
    return res
}


func strStrII(heystack string, needle string) int {
    dict := make([]int,128)
    a := heystack
    b := needle

    for i := 0; i < 26; i++ {
        dict[i] = -1
    }

    for i := 0; i < len(b); i++ {
        dict[b[i]] = i
    }

    i, j := 0, 0
    for i <= len(a) - len(b) {
        j = 0
        for j < len(b) {
            if a[i] == b[j] {
                i++
                j++
            } else {
                index := i + len(b) - j
                if index >= len(a) {
                    return -1
                }
                if dict[a[index]] == -1 {
                    i = index + 1
                } else {
                    i = index - dict[a[index]]
                }
                break
            }
        }
        if j == len(b) {
            return i - len(b)
        }
    }
    return -1
}