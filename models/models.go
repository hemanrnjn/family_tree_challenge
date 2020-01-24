package models

import "time"

type Member struct {
	Name      string
	Gender    string
	Spouse    *Member
	Children  []*Member
	TimeAdded time.Time
}

type Children struct {
	Name   string
	Gender string
}

func (member *Member) AddChildren(m Children) {
	var child Member
	child.Name = m.Name
	child.Gender = m.Gender
	child.TimeAdded = time.Now()
	member.Children = append(member.Children, &child)
}

func (member *Member) GetChildren(name string) *Member {
	for _, v := range member.Children {
		if v.Name == name {
			return v
		}
	}
	return nil
}

func (member *Member) findMember(name string) {

}
