package particles

import ("math/rand";"time")
// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	for p,_ := range s.Content {
		rand.Seed(time.Now().UnixNano())
		x := rand.Float64()
		y := rand.Float64()
		if rand.Float64() > 0.5{
			x = -x
		}
		if rand.Float64() > 1{
			y = -y
		}
	}
}