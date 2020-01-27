package models

import (
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

func (member *Member) CreateSpouse(name string, gender string) *Member {
	m := Member{
		Name:      name,
		Gender:    gender,
		Spouse:    member,
		TimeAdded: time.Now(),
	}
	return &m
}

func (member *Member) FindMember(name string) *Member {
	if (*member).Name == name {
		return member
	} else if (*member).Spouse.Name == name {
		return (*member).Spouse
	}
	for _, v := range member.Children {
		if x := v.FindMember(name); x != nil {
			return x
		}
	}
	return nil
}
