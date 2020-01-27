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
			Name:   "Lika",
			Gender: "female",
		},
		{
			Name:   "Aras",
			Gender: "male",
		},
		{
			Name:   "Chitra",
			Gender: "female",
		},
		{
			Name:   "Satya",
			Gender: "female",
		},
		{
			Name:   "Vyan",
			Gender: "male",
		},
	}

	for _, v := range m {
		queen.AddChildren(v)
	}

	chit := queen.GetChildren("Chit")
	amba := queen.GetChildren("Amba")
	vich := queen.GetChildren("Vich")
	lika := queen.GetChildren("Lika")
	aras := queen.GetChildren("Aras")
	chitra := queen.GetChildren("Chitra")
	satya := queen.GetChildren("Satya")
	vyan := queen.GetChildren("Vyan")

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

	chit.Spouse = amba
	amba.Spouse = chit
	vich.Spouse = lika
	lika.Spouse = vich
	aras.Spouse = chitra
	chitra.Spouse = aras
	satya.Spouse = vyan
	vyan.Spouse = satya

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
			Name:   "Arit",
			Gender: "male",
		},
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
			Name:   "Satvy",
			Gender: "female",
		},
		{
			Name:   "Asva",
			Gender: "male",
		},
		{
			Name:   "Krpi",
			Gender: "female",
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

	dritha := amba.GetChildren("Dritha")
	jaya := dritha.CreateSpouse("Jaya", "male")
	arit := queen.FindMember("Arit")
	jnki := chitra.GetChildren("Jnki")
	asva := satya.GetChildren("Asva")
	satvy := asva.CreateSpouse("Satvy", "female")
	vyas := satya.GetChildren("Vyas")
	krpi := vyas.CreateSpouse("Krpi", "female")

	m = []models.Children{
		{
			Name:   "Yodhan",
			Gender: "male",
		},
	}

	for _, v := range m {
		dritha.AddChildren(v)
	}

	dritha.Spouse = jaya
	jaya.Spouse = dritha
	arit.Spouse = jnki
	jnki.Spouse = arit

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
