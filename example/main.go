// Copyright 2022 clavinjune/ononoki-engine
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// this is an example package of using ononoki rule engine
package main

import (
	"log"

	"github.com/clavinjune/ononoki"
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

	r4 := ononoki.NewRuleLeaf(f4, c2)
	r3 := ononoki.NewRuleLeaf(f3, c1)
	r2 := ononoki.NewRule([]ononoki.Concluder{r3, r4}, f2, nil)
	r1 := ononoki.NewRuleLeaf(f1, c3)
	root := ononoki.NewRuleRoot([]ononoki.Concluder{r1, r2})

	data := []map[string]any{
		{
			"height": int64(100), // procompsognathus
		},
		{
			"height": int64(101), // property not found when checking "type"
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
