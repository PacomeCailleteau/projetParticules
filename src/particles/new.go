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
	if config.General.RandomSpawn == true{
		for i := 0; i < config.General.InitNumParticles; i++ {
			S.Content = append(S.Content,
				Particle{
					PositionX: float64(config.General.WindowSizeX)*rand.Float64(),
					PositionY: float64(config.General.WindowSizeY)*rand.Float64(),
					ScaleX:    1, ScaleY: 1,
					ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
					Opacity: 1,
				})
		}
	}
	return S
}