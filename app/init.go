package app

import (
	"time"

	"github.com/hemanrnjn/family_tree_challenge/models"
)

var king models.Member
var queen models.Member

func init() {
	king.Name = "King Shan"
	king.Gender = "male"
	king.Spouse = &queen
	king.TimeAdded = time.Now()

	queen.Name = "Queen Anga"
	queen.Gender = "female"
	queen.Spouse = &king
	queen.TimeAdded = time.Now()

	initFamilyTree()
}

func initFamilyTree() {
	m := []models.Children{
		{
			Name:   "Chit",
			Gender: "male",
		},
		{
			Name:   "Amba",
			Gender: "female",
		},
		{
			Name:   "Ish",
			Gender: "male",
		},
		{
			Name:   "Vich",
			Gender: "male",
		},
		{
			Name:   "Aras",
			Gender: "male",
		},
		{
			Name:   "Satya",
			Gender: "female",
		},
	}

	for _, v := range m {
		queen.AddChildren(v)
	}

	chit := queen.FindMember("Chit")
	amba := chit.AddSpouse("Amba", "female")
	vich := queen.FindMember("Vich")
	lika := vich.AddSpouse("Lika", "female")
	aras := queen.FindMember("Aras")
	chitra := aras.AddSpouse("Chitra", "female")
	satya := queen.FindMember("Satya")
	_ = satya.AddSpouse("Vyan", "male")

	m = []models.Children{
		{
			Name:   "Dritha",
			Gender: "female",
		},
		{
			Name:   "Tritha",
			Gender: "female",
		},
		{
			Name:   "Vritha",
			Gender: "male",
		},
	}

	for _, v := range m {
		amba.AddChildren(v)
	}

	m = []models.Children{
		{
			Name:   "Vila",
			Gender: "female",
		},
		{
			Name:   "Chika",
			Gender: "female",
		},
	}

	for _, v := range m {
		lika.AddChildren(v)
	}

	m = []models.Children{
		{
			Name:   "Jnki",
			Gender: "female",
		},
		{
			Name:   "Ahit",
			Gender: "male",
		},
	}

	for _, v := range m {
		chitra.AddChildren(v)
	}

	m = []models.Children{
		{
			Name:   "Asva",
			Gender: "male",
		},
		{
			Name:   "Vyas",
			Gender: "male",
		},
		{
			Name:   "Atya",
			Gender: "female",
		},
	}

	for _, v := range m {
		satya.AddChildren(v)
	}

	dritha := amba.FindMember("Dritha")
	_ = dritha.AddSpouse("Jaya", "male")
	jnki := chitra.FindMember("Jnki")
	_ = jnki.AddSpouse("Arit", "male")
	asva := satya.FindMember("Asva")
	satvy := asva.AddSpouse("Satvy", "female")
	vyas := satya.FindMember("Vyas")
	krpi := vyas.AddSpouse("Krpi", "female")

	m = []models.Children{
		{
			Name:   "Yodhan",
			Gender: "male",
		},
	}

	for _, v := range m {
		dritha.AddChildren(v)
	}

	m = []models.Children{
		{
			Name:   "Laki",
			Gender: "male",
		},
		{
			Name:   "Lavnya",
			Gender: "female",
		},
	}

	for _, v := range m {
		jnki.AddChildren(v)
	}

	m = []models.Children{
		{
			Name:   "Vasa",
			Gender: "male",
		},
	}

	for _, v := range m {
		satvy.AddChildren(v)
	}

	m = []models.Children{
		{
			Name:   "Kriya",
			Gender: "male",
		},
		{
			Name:   "Krithi",
			Gender: "female",
		},
	}

	for _, v := range m {
		krpi.AddChildren(v)
	}
}
