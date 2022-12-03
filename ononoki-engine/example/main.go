package main

import (
	ononoki "github.com/clavinjune/ononoki-engine"
	"log"
)

func main() {
	c1 := ononoki.NewConclusion("t-rex")
	c2 := ononoki.NewConclusion("brontosaurus")
	c3 := ononoki.NewConclusion("procompsognathus")

	f1 := ononoki.NewFact("height", ononoki.ComparatorLE, int64(100),
		ononoki.FactWithName[int64]("height less than or equal to 100"),
	)

	f2 := ononoki.NewFact("height", ononoki.ComparatorGT, int64(100),
		ononoki.FactWithName[int64]("height greater than 100"),
	)

	f3 := ononoki.NewFact("type", ononoki.ComparatorEQ, "carnivores",
		ononoki.FactWithName[string]("carnivores type"),
	)

	f4 := ononoki.NewFact("type", ononoki.ComparatorEQ, "herbivores",
		ononoki.FactWithName[string]("herbivores type"),
	)

	r4 := ononoki.NewRule(nil, f4, c2)
	r3 := ononoki.NewRule(nil, f3, c1)
	r2 := ononoki.NewRule([]ononoki.Concluder{r3, r4}, f2, nil)
	r1 := ononoki.NewRule(nil, f1, c3)
	root := ononoki.NewRule([]ononoki.Concluder{r1, r2}, nil, nil)

	data := []map[string]any{
		{
			"height": int64(100), // procompsognathus
		},
		{
			"height": int64(101), // no conclusion
		},
		{
			"height": int64(101),
			"type":   "carnivores", // t-rex
		},
		{
			"height": int64(101),
			"type":   "herbivores", // brontosaurus
		},
		{
			"height": int64(101),
			"type":   "omnivores", // no conclusion
		},
	}

	for _, d := range data {
		conclusion, err := root.Conclude(d)
		if err != nil {
			log.Println("err", err)
		} else {
			log.Println("conclusion", conclusion.Name)
		}
	}
}
