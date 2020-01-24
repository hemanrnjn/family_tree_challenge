package models

import (
	"fmt"
	"time"
)

type Member struct {
	Name      string
	Gender    string
	Spouse    *Member
	Children  []*Member
	Parent    *Member
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
	child.Parent = member
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

var m *Member

func (member *Member) FindMember(name string) *Member {

	if (*member).Name == name {
		fmt.Println("FROM Models: ", (*member).Name)
		return member
	}
	for _, v := range member.Children {
		m = v.FindMember(name)
	}
	return m
}
