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
	child := Member{
		Name:      m.Name,
		Gender:    m.Gender,
		Parent:    member,
		TimeAdded: time.Now(),
	}
	member.Children = append(member.Children, &child)
}

func (member *Member) AddSpouse(name string, gender string) *Member {
	m := Member{
		Name:      name,
		Gender:    gender,
		Spouse:    member,
		TimeAdded: time.Now(),
	}
	member.Spouse = &m
	return &m
}

func (member *Member) FindMember(name string) *Member {
	if (*member).Name == name {
		return member
	} else if (*member).Spouse != nil && (*member).Spouse.Name == name {
		return (*member).Spouse
	} else if len((*member).Children) > 0 {
		for _, v := range member.Children {
			if x := v.FindMember(name); x != nil {
				return x
			}
		}
	} else if (*member).Spouse != nil && len((*(*member).Spouse).Children) > 0 {
		for _, v := range member.Spouse.Children {
			if x := v.FindMember(name); x != nil {
				return x
			}
		}
	}
	return nil
}
