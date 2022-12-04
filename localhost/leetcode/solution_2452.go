package solutions

/*
class Solution(object):
    def twoEditWords(self, queries, dictionary):
        """
        :type queries: List[str]
        :type dictionary: List[str]
        :rtype: List[str]
        """

        def compare(s1, s2):
            i = 0
            edits = 0
            while i < len(s1):
                if s1[i] != s2[i]:
                    edits += 1
                    if edits > 2:
                        return -1
                i += 1

            return edits

        ans = []
        for q in queries:
            flag = False
            for d in dictionary:
                q_cmp = compare(q, d)
                if q_cmp >= 0:
                    flag = True
            if flag:
                ans.append(q)

        return ans
*/

func compare(s1, s2 string) int {
	i := 0
	edits := 0
	for i < len(s1) {
		if s1[i] != s2[i] {
			edits++
			if edits > 2 {
				return -1
			}
		}
		i++
	}

	return edits
}

func twoEditWords(queries []string, dictionary []string) []string {
	ans := make([]string, 0)
	var flag bool
	for _, q := range queries {
		flag = false
		for _, d := range dictionary {
			q_cmp := compare(q, d)
			if q_cmp >= 0 {
				flag = true
			}
		}
		if flag {
			ans = append(ans, q)
		}
	}
	return ans
}
