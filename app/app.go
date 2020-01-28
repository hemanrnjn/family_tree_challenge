package app

import (
	"fmt"
)

func Landing() {
	fmt.Println(getPaternalUncle("Dritha"))
}

func getPaternalUncle(name string) (res []string) {
	member := queen.FindMember(name)
	fmt.Println("Found member", member)
	if member != nil {
		if (*member).Parent != nil && (*(*member).Parent).Parent != nil {
			fmt.Println("Enters first")
			for _, v := range (*(*(*member).Parent).Parent).Children {
				if v.Gender == "male" && v.Spouse.Name != name {
					res = append(res, v.Name)
				}
			}
		} else if (*member).Parent.Spouse != nil && (*(*(*member).Parent).Spouse).Parent != nil {
			fmt.Println("Enters second")
			for _, v := range (*(*(*(*member).Parent).Spouse).Parent).Children {
				fmt.Println(v)
				if (*v).Gender == "male" && (*v.Spouse).Name != name {
					res = append(res, v.Name)
				}
			}
		} else {
			fmt.Println("No Paternal Uncle")
			return
		}

	} else {
		fmt.Println("Member not found")
		return
	}

	return
}
