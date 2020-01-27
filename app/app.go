package app

import (
	"fmt"
)

func Landing() {
	fmt.Println("Hello World")
	fmt.Println(getPaternalUncle("Abc"))
}

func getPaternalUncle(name string) (res []string) {
	member := queen.FindMember(name)

	fmt.Println("From App: ", member)

	// for _, v := range member.Parent.Parent.Children {
	// 	if v.Gender == "male" && v.Spouse.Name != name {
	// 		res = append(res, v.Name)
	// 	}
	// }

	return res
}
