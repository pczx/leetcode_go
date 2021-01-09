package solution38

func permutation(s string) []string {
	chars, res := []byte(s), []string{}

	var dfs func(i int)

	dfs = func(x int) {
		if x == len(chars)-1 {
			res = append(res, string(chars))
			return
		}

		set := make(map[byte]struct{})

		for i := x; i < len(chars); i++ {
			if _, ok := set[chars[i]]; ok {
				continue
			}
			set[chars[i]] = struct{}{}
			chars[i], chars[x] = chars[x], chars[i]
			dfs(x + 1)
			chars[i], chars[x] = chars[x], chars[i]
		}
	}

	dfs(0)

	return res
}
