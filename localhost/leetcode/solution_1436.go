package solutions

/*
class Solution(object):

	def destCity(self, paths):
	    """
	    :type paths: List[List[str]]
	    :rtype: str
	    """
	    G = {}

	    for start, end in paths:
	        G[start] = end

	    ans = paths[0][0]

	    while True:
	        if ans in G:
	            ans = G[ans]
	            continue
	        return ans
*/
func destCity(paths [][]string) string {
	G := make(map[string]string)

	for _, item := range paths {
		G[item[0]] = item[1]
	}

	ans := paths[0][0]

	for true {
		if _, ok := G[ans]; ok {
			ans = G[ans]
			continue
		}

		return ans
	}

	return ans
}
