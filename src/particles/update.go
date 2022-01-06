package particles

import ("math/rand";"time";"project-particles/config";"log")
// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

func (s *System) Update() {
	for p,_ := range s.Content {
		s.Content[p].PositionX += s.Content[p].SpeedX
		s.Content[p].PositionY += s.Content[p].SpeedY
		if s.Content[p].PositionX >= float64(config.General.WindowSizeX){
			s.Content = remove(s.Content, p)
			break
		}
		if s.Content[p].PositionX <= -10{
			s.Content = remove(s.Content, p)
			break
		}
		if s.Content[p].PositionY >= float64(config.General.WindowSizeY){
			s.Content = remove(s.Content, p)
			break
		}
		if s.Content[p].PositionY <= -10 || s.Content[p].PositionY >= float64(config.General.WindowSizeY){
			s.Content = remove(s.Content, p)
			break
		}
		s.Content[p].ScaleX -= 0.01
		s.Content[p].ScaleY -= 0.01 
		if s.Content[p].ScaleX <= 0{
			remove(s.Content, p)
			break
		}
	}
	
	rand.Seed(time.Now().UnixNano()) 
	spdX := rand.Float64()
	spdX -= 0.5
	spdY := rand.Float64()
	spdY -= 0.5
	s.reste += float64(config.General.SpawnRate)
	//SpawnRate
	for s.reste >=1{
		s.Content = append(s.Content, Particle{
			PositionX: rand.Float64()*float64(config.General.WindowSizeX),
			PositionY: rand.Float64()*float64(config.General.WindowSizeY),
			ScaleX: 2, 
			ScaleY: 2,
			ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
			Opacity: 1,
			SpeedX: spdX *10, 
			SpeedY: spdY *10,
		})
		log.Print(len(s.Content))

		s.reste -=1
	}
}

func remove(slice []Particle, s int) []Particle {
	return append(slice[:s], slice[s+1:]...)
}