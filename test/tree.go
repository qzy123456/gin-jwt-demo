package tests

import (
"fmt"
"encoding/json"
)

type dept struct {
	DeptId string `json:"deptId"`
	FrameDeptStr string `json:"frameDeptStr"`
	Child []*dept `json:"child"`
}
func main() {
	depts := make([]dept,0)
	var a dept
	a.DeptId = "1"
	a.FrameDeptStr = ""
	depts = append(depts,a)
	a.DeptId="3"
	a.FrameDeptStr = "1"
	depts = append(depts,a)
	a.DeptId="4"
	a.FrameDeptStr = "1"
	depts = append(depts,a)
	a.DeptId="5"
	a.FrameDeptStr = "13"
	depts = append(depts,a)
	a.DeptId="6"
	a.FrameDeptStr = "13"
	depts = append(depts,a)
	fmt.Println(depts)

	deptRoots := make([]dept,0)
	for _,v := range depts{
		if v.FrameDeptStr == ""{
			deptRoots= append(deptRoots,v)
		}
	}

	pdepts := make([]*dept,0)
	for i,_ := range depts{
		var a *dept
		a = &depts[i]
		pdepts = append(pdepts,a)
	}

	var node *dept
	node = &depts[0]
	MakeTree(pdepts,node)
	data, _ := json.Marshal(node)
	fmt.Printf("%s", data)

}

func has(v1 dept,vs []*dept) bool  {
	var has bool
	has = false
	for _,v2 := range vs {
		v3 := *v2
		if v1.FrameDeptStr+v1.DeptId == v3.FrameDeptStr{
			has = true
			break
		}
	}
	return has
}

func MakeTree(vs []*dept,node *dept) {
	childs := findChild(node,vs)
	for _,child := range childs{
		node.Child = append(node.Child,child)
		if has(*child,vs) {
			MakeTree(vs,child)
		}
	}
}

func findChild(v *dept,vs []*dept)(ret []*dept)  {
	for _,v2 := range vs{
		if v.FrameDeptStr+v.DeptId == v2.FrameDeptStr{
			ret= append(ret,v2)
		}
	}
	return
}