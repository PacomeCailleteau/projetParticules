package particles

import ("project-particles/config";"math/rand";"time")

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	rand.Seed(time.Now().UnixNano())
	var S System = System{Content: []Particle{}}
	if config.General.RandomSpawn{
		for i := 0; i < config.General.InitNumParticles; i++ {
			rand.Seed(time.Now().UnixNano())
			spdX := rand.Float64()
			if rand.Float64() > 0.5{
				spdX = -spdX
			}
			pX := rand.Float64()* float64(config.General.WindowSizeX)
			pY := rand.Float64()* float64(config.General.WindowSizeY)
			p := Particle{
				PositionX: pX,
				PositionY: pY,
				ScaleX: 1-(pY*100/float64(config.General.WindowSizeY)),
				ColorRed:1, ColorGreen:1, ColorBlue:1,
				Opacity:1,
				SpeedX: spdX, SpeedY:1,
			}
			S.Content = append(S.Content, p)
		}
	}else{
		for i := 0; i < config.General.InitNumParticles; i++ {
			spdX := rand.Float64()
			if rand.Float64() > 0.5{
				spdX= -spdX
			}
			p := Particle{
				PositionX: float64(config.General.SpawnX),
				PositionY: float64(config.General.SpawnY),
				ScaleX: 1- (float64(config.General.SpawnY)),
				ColorRed:1, ColorGreen:1, ColorBlue:1,
				Opacity:1,
				SpeedX: spdX, SpeedY:1,
			}
			S.Content = append(S.Content, p)
		}
	}
	return S
}