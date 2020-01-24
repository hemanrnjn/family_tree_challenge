package app

import (
	"fmt"
	"time"

	"github.com/hemanrnjn/family_tree_challenge/models"
)

var king models.Member
var queen models.Member

func Landing() {
	fmt.Println("Hello World")
}

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

	amba := queen.GetChildren("Amba")
	lika := queen.GetChildren("Lika")
	chitra := queen.GetChildren("Chitra")
	satya := queen.GetChildren("Satya")

	m = []models.Children{
		{
			Name:   "Dritha",
			Gender: "female",
		},
		{
			Name:   "Jaya",
			Gender: "male",
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
	jnki := chitra.GetChildren("Jnki")
	satvy := satya.GetChildren("Satvy")
	krpi := queen.GetChildren("Krpi")

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

// func getRelationship(name string, rel string) []string {

// }
