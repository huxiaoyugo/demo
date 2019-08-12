package main

import "fmt"

func main() {

	req_skills := []string{
		"A","B","H","J",
	}

	people := [][]string{
		{"A","B"},
		{"A","C"},
		{"B","D"},
		{"Z","G"},
		{"H","J"},
		{"H","J"},
		{"H","J"},
		{"A","C","D","E"},

	}
	fmt.Println(smallestSufficientTeam(req_skills, people))
}


func smallestSufficientTeam(req_skills []string, people [][]string) []int {

	tMap := make(map[string]int, len(req_skills))

	res := make([]int, 0)
	moreSkills := make([]string, 0)

	for index, skills := range people {

		if len(skills) == 0 {
			continue
		}

		res = append(res, index)
		for _, s := range skills {
			tMap[s]++
			if tMap[s] == 2 {
				moreSkills = append(moreSkills, s)
			}
		}

		if len(moreSkills) == 0 {
			continue
		}

		needDelRes := make([]int, 0)
	out:
		for _, pNum := range res {

			for _, skill := range people[pNum] {

				flag := false
				for _, more := range moreSkills {
					if more == skill {
						flag = true
						break
					}
				}
				if !flag {
					continue out
				}
			}

			//

			needDel := make([]string, 0)
			for _, skill := range people[pNum] {
				tMap[skill] -= 1

				if tMap[skill] <= 1 {
					needDel = append(needDel, skill)
				}
			}

			moreSkills = DelSlice(moreSkills, needDel)

			needDelRes = append(needDelRes, pNum)
		}

		for _, item := range needDelRes {
			res = DelSliceInt(res, item)
		}

	}
	return res

}

func DelSlice(s []string, needDel []string) []string {

	res := make([]string, 0)
out:
	for _, item := range s {
		for _, del := range needDel {
			if item == del {
				continue out
			}
		}
		res = append(res, item)
	}
	return res
}

func DelSliceInt(s []int, num int) []int {
	res := make([]int, 0)
	for _, item := range s {
		if item == num {
			continue
		}
		res = append(res, item)
	}
	return res
}
