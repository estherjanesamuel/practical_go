package guess

import "fmt"

func PublicGuess(n int, pick int) int {
	return guess_pick(n,pick)	
}

func GuessNumber(n int) int {
	fmt.Println("n: ",n)
	low := 1
	high := n
	i := 0

	for low <= high {
		i++
		fmt.Println("count no: ",i)
		mid := low + (high - low) / 2
		fmt.Println("mid: ",mid)
		fmt.Println("low: ",low)
		fmt.Println("high: ",high)

		g := guess(mid)
	fmt.Println("g: ", g)

		if g == 1 { // Your guess is lower
			low = mid + 1
		} else if g == -1 { // Your guess is higher
			high = mid - 1
		} else {
			return mid // your guess is equal
		}
	}
	return 1
}


func guess(num int) int {
	pick := 6
	// -1: Your guess is higher than the number I picked (i.e. num > pick).
	if num > pick {
		return -1
		} 
	// 1: Your guess is lower than the number I picked (i.e. num < pick).
	if num < pick {
		return 1
	}
	// 0: your guess is equal to the number I picked (i.e. num == pick).
	return 0
}

func guessNumber(n int, pick int) int {
	low := 1
	high := n

	for low <= high {
		mid := low + (high - low) / 2
		// mid := (high / 2)
		g := guess_pick(mid, pick)

		if g == 1 {
			low = mid + 1
		} else if g == -1 {
			high = mid - 1
		} else {
			return mid
		}
	}
	return 1
}


func guess_pick(num int, pick int) int {
	if num > pick {
		return -1
	} 
	if num < pick {
		return 1
	}
	return 0
}

func Mid(high int) (mid int) {
	low := 1
	mid = low + (high - low) / 2
	return
}

func mid(high int) (mid int) {
	// mid := (high / 2)
	mid = high / 2
	return
}

/*

 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0

 func guessNumber(n int) int {
    left, right := 1, n
    for left <= right {
        num := left + (right - left)/ 2
        g := guess(num)
        if g == 1 {
            left = num + 1
        } else if g == -1 {
            right = num - 1
        } else {
            return num
        }
    }
    return 1
}
*/