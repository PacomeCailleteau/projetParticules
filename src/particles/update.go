package particles

import ("math/rand";"time";"project-particles/config";"log")
// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

func (s *System) Update() {

	var particules []Particle = s.Content

	particules = deplacement(particules)//mise à jour de la position
	particules = grossissement(particules,true,0.98)//mise à jour de la taille des particules
	particules = condition_suppression(particules, -10, float64(config.General.WindowSizeX), -10, float64(config.General.WindowSizeY), 0, -1)
	

	s.reste += float64(config.General.SpawnRate)//ajouter la valeur de demande du nombre de particules à ajouter
	//SpawnRate
	for s.reste >=1{//tant que des particules sont à faire apparaitre
		rand.Seed(time.Now().UnixNano())//permet de générer des nombres aléatoire grâce à une graine étant fonction de l'heure/date/...
		//déclaration des variables
		var x float64 = rand.Float64()*float64(config.General.WindowSizeX)
		var y float64 = rand.Float64()*float64(config.General.WindowSizeY)
		var taille float64 = 0.5
		var vitesse float64 = 5
		//utilisation des variables
		particules = ajout(particules, x, y, taille, vitesse)

		s.reste -=1//noter qu'une particule a été ajoutée
	}
	s.Content = particules
	log.Print(len(particules))
}

func suppression(particules []Particle, i int) []Particle {
	return append(particules[:i], particules[i+1:]...)
}

func deplacement(particules []Particle) []Particle{
	for i := 0; i < len(particules); i++ {
		particules[i].PositionX += particules[i].SpeedX
		particules[i].PositionY += particules[i].SpeedY
	}
	return particules
}

func condition_suppression(particules []Particle, xmin,xmax, ymin,ymax float64, taillemin, taillemax float64) []Particle{
	for i := 0; i < len(particules); i++ {
		//déclaration des variables
		var PositionX float64 = particules[i].PositionX
		var PositionY float64 = particules[i].PositionY
		var tailleX float64 = particules[i].ScaleX
		var tailleY float64 = particules[i].ScaleY
		//tests
		if PositionX < xmin || PositionX > xmax ||
		PositionY < ymin || PositionY > ymax ||
		tailleX < taillemin || tailleY < taillemin ||
		(tailleX > taillemax || tailleY > taillemax) && taillemax !=-1{
			particules = suppression(particules, i)}
	}
	return particules
}

func grossissement(particules []Particle, produit bool, size float64) []Particle{
	for i := 0; i < len(particules); i++ {//pour toutes les particules
		if !produit{//si on veux additionner ou soustraire une valeur à leur taille...
			particules[i].ScaleX += size//..les agrandir ou les rétrécir en X
			particules[i].ScaleY += size//..les agrandir ou les rétrécir en Y
		}else{//sinon (si on veut multiplier leur taille par une valeur)...
			particules[i].ScaleX *= size//...multiplier leur taille en X par la valeur
			particules[i].ScaleY *= size//...multiplier leur taille en Y par la valeur
	}}
	return particules
}

func ajout(particules []Particle, PositionX, PositionY, taille, mult_vitesse float64) []Particle{
	particules = append(particules, Particle{//ajout d'une particule dont...
			PositionX: PositionX,//sa position en X
			PositionY: PositionY,//sa position en Y
			ScaleX: taille,//sa taille en X
			ScaleY: taille,//sa taille en Y
			ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),//sa couleur aléatoire en RGB
			Opacity: 1,//son opacité de 100%
			SpeedX: 2*(rand.Float64()-0.5)*mult_vitesse,//sa vitesse est aléatoire entre -5 et 5 en X
			SpeedY: 2*(rand.Float64()-0.5)*mult_vitesse,//sa vitesse est aléatoire entre -5 et 5 en Y
	})
	return particules
}