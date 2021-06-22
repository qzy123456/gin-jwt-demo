package tests

import (
	"fmt"
	"encoding/json"
	"testing"
)

func TestMakeTree(t *testing.T)  {
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
	//获取了根上的科室
	fmt.Println("根上的科室有:",deptRoots)


	var node *dept
	node = &depts[0]
	MakeTree(pdepts,node)
	fmt.Println("the result we got is",pdepts)
	data, _ := json.Marshal(node)
	fmt.Printf("%s", data)
}