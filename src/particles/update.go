package particles

import ("math/rand";"time";"project-particles/config";"log";"math")
// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

func (s *System) Update() {

	var particules []Particle = s.Content

	if config.General.Duree_Vie > 0{
		for i := 0; i < len(particules); i++ {
			particules[i].Duree_Vie++
			if particules[i].Duree_Vie > config.General.Duree_Vie{
				particules = suppression(particules,i)
			}
		}
	}

	//Actualisation de la position des particules dans le tableau en fonction de leur vitesse
	particules = deplacement(particules)//mise à jour de la position

	if config.General.Rebond_bords{particules = rebond_bords(particules)}

	if config.General.Rebond_particules{particules = rebond_particules(particules)}
	
	if config.General.Gravite != 0{
		particules = gravite(particules)
	}


	if config.General.Gravitation != 0{
		particules = gravitation(particules)
	}
	if config.General.Acceleration != 1{
		for i := 0; i < len(particules); i++ {
			particules[i] = acceleration(particules[i])
		}
	}

	//Actualisation de la taille des particules en fonction de deux paramètres dans la fonction
	particules = grossissement(particules,true,1)//mise à jour de la taille des particules

	//Suppression des particules à oublier
	particules = condition_suppression(particules, float64(config.General.ExterieurDeLecranXmin), float64(config.General.ExterieurDeLecranXmax), float64(config.General.ExterieurDeLecranYmin), float64(config.General.ExterieurDeLecranYmax), 0, -1)

	//Calcul et ajout du bon nombre de particules à afficher durant l'appel de la fonction update dépandemment du SpawnRate
	s.reste += float64(config.General.SpawnRate)//ajouter la valeur de demande du nombre de particules à ajouter
	//SpawnRate
	for s.reste >=1{//tant que des particules sont à faire apparaitre
		rand.Seed(time.Now().UnixNano())//permet de générer des nombres aléatoire grâce à une graine étant fonction de l'heure/date/...
		//initialisation des variables de position
		var x float64
		var y float64
		//RandomSpawn
		if config.General.RandomSpawn{
			x = rand.Float64()*float64(config.General.WindowSizeX)
			y = rand.Float64()*float64(config.General.WindowSizeY)
		}else{ //Non RandomSpawn
			x = float64(config.General.SpawnX)
			y = float64(config.General.SpawnY)
		}
		var taille float64 = config.General.InitSizeParticles//Initialisation de la taille de la particule
		var vitesse float64 = config.General.InitVitesseParticles //Initialisation de la vitesse de la particule
		//utilisation des variables
		particules = ajout(particules, x, y, taille, vitesse)

		s.reste -=1//noter qu'une particule a été ajoutée
	}
	s.Content = particules
	log.Print(len(particules)," énergie = ",energie(particules))
}

/*La fonction suppression sert à supprimer une particule contenue
dans un tableau de particule. Elle prend en entrée un tableau de 
particule ainsi que l'indice de celle à supprimer et en ressort
un tableau de particules dépourvue de celle supprimée.
Exemple : 
Départ : tableau = [particule1, particule2, particule3]
tableau = suppression(tableau,1)
Arrivée : tableau = [particule1, particule3]*/
func suppression(particules []Particle, i int) []Particle {
	return append(particules[:i], particules[i+1:]...)
}

/*La fonction deplacement sert à actualiser la position de chaque particule
du tableau de particules en fonction du temps et de sa vitesse.
Elle prend en entrée un tableau de particules et en ressort un tableau
de particules dont, pour chaques particules, les coordonnées x et y de celle-ci
ont été incrémentées de sa vitesse x et y.
Example :
Départ : tableau = [particule1(avec position x = a,position y = b, avec vitesse x = c et vitesse y = d),particule2(avec position x = e,position y = f, avec vitesse x = g et vitesse y = h)]
tableau = deplacement(tableau)
Arrivée : tableau = [particule1(avec position x = a+c,position y = b+d, avec vitesse x = c et vitesse y = d),particule2(avec position x = e+g,position y = f+h, avec vitesse x = g et vitesse y = h)]*/
func deplacement(particules []Particle) []Particle{
	for i := 0; i < len(particules); i++ {
		particules[i].PositionX += particules[i].SpeedX
		particules[i].PositionY += particules[i].SpeedY
	}
	return particules
}

func rebond_bords(particules []Particle) []Particle{
	var TailleX float64 = float64(config.General.WindowSizeX)
	var TailleY float64 = float64(config.General.WindowSizeY)
	for i := 0; i < len(particules); i++ {
		//bord gauche
		if particules[i].PositionX < 0{
		particules[i].PositionX -= 2*particules[i].PositionX
		particules[i].SpeedX = -particules[i].SpeedX
		}
		//bord droit
		if particules[i].PositionX > TailleX-10*particules[i].ScaleX{//-particules[i].ScaleX
			particules[i].PositionX -= 2*(particules[i].PositionX-(TailleX-10*particules[i].ScaleX))
			particules[i].SpeedX = -particules[i].SpeedX
		}
		//bord haut
		if particules[i].PositionY < 0{
			particules[i].PositionY -= 2*particules[i].PositionY
			particules[i].SpeedY = -particules[i].SpeedY
		}
		//bord bas
		if particules[i].PositionY > TailleY-10*particules[i].ScaleY{//-particules[i].ScaleY
			particules[i].PositionY -= 2*(particules[i].PositionY-(TailleY-10*particules[i].ScaleY))
			particules[i].SpeedY = -particules[i].SpeedY
		}
	}
	return particules
}

func gravite(particules []Particle) []Particle{
	var gravite = config.General.Gravite
	for i := 0; i < len(particules); i++ {
		particules[i].SpeedY += gravite
	}
	return particules
}


/*La fonction condition_suppression sert à supprimer une particule en fonction de certaines conditions étants sa position et sa taille.
Elle prend en entrée un tableau de particules et en ressort un tableau de particules dont certaines ont pu être supprimées selon des conditions étant l'appartenance à un intervalle de position x, un de y et un de sa taille.
Exemple : une particule du tableau de particules avec coordonnée x et y et taille z se fera supprimer
si x n'appartient pas à [xmin;xmax] 
ou si y n'appartient pas à [ymin;ymax] 
ou si z n'appartient pas à [zmin;zmax] 
*/
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

/*Actualisation de la taille de la particule
Additionne size à la taille de la particule si produit vaut false
et multiplie size à la taille de la particule si produit vaut true.
Exemple 1 :
Départ : tableau(particule1(tailleX = a, tailleY = b),particule1(tailleX = c, tailleY = d))
grossissement(tableau,false,taille)
Arrivée : tableau(particule1(tailleX = a + size, tailleY = b + size ),particule1(tailleX = c + size, tailleY = d + size )
Exemple 2 :
Départ : tableau(particule1(tailleX = a, tailleY = b),particule1(tailleX = c, tailleY = d))
grossissement(tableau,true,taille)
Arrivée : tableau(particule1(tailleX = a * size, tailleY = b * size ),particule1(tailleX = c * size, tailleY = d * size )*/
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

/*Ajoute au tableau de particules une particule avec sa position x, sa position y ,sa taille x = sa taille y, et sa vitesse pouvant être choisies,
ainsi que sa couleur choisie aléatoirement et une oppacité de 100%.
Prend en entrée un tableau de particules en entrée et en ressort un autre avec une particule ajoutée avec des paramètres donnés.
Exemple : 
Départ : tableau()
ajout(tableau,PositionX, PositionY, taille, vitesse)
Arrivée : tableau(particule1(positionX = a, PositionY = b, tailleX = c, tailleY = c, Color = rand.Float, Oppacité = 1, vitesseX = math.Cos(rand.Float64()*2*math.Pi)*d, vitesse = math.Sin(rand.Float64()*2*math.Pi)*d))
*/
func ajout(particules []Particle, PositionX, PositionY, taille, mult_vitesse float64) []Particle{
    var angle float64 = rand.Float64()*2*math.Pi
    particules = append(particules, Particle{//ajout d'une particule dont...
        PositionX: PositionX,//sa position en X
        PositionY: PositionY,//sa position en Y
        ScaleX: taille,//sa taille en X
        ScaleY: taille,//sa taille en Y
        ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),//sa couleur aléatoire en RGB
        Opacity: 1,//son opacité de 100%
        SpeedX: math.Cos(angle)*mult_vitesse,
        SpeedY: math.Sin(angle)*mult_vitesse,
        /*SpeedX: 2*(rand.Float64()-0.5)*mult_vitesse,//sa vitesse est aléatoire entre -5 et 5 en X
        SpeedY: 2*(rand.Float64()-0.5)*mult_vitesse,//sa vitesse est aléatoire entre -5 et 5 en Y
        */
    })
    return particules
}


func abs(n float64) float64{
    if n < 0{return -n}else{return n}}


func collision(particule1, particule2 Particle) bool{
    if procheX(particule1,particule2) && procheY(particule1,particule2){return true}else{return false}}

func procheX(particule1, particule2 Particle)bool{
    var distanceX float64 = abs(particule2.PositionX - particule1.PositionX)
    if particule1.PositionX < particule2.PositionX{//X 1-->2
        if distanceX <= particule1.ScaleX*10{return true}
    }else{//X 2-->1
        if distanceX <= particule2.ScaleX*10{return true}}
    return false}

func procheY(particule1, particule2 Particle)bool{
    var distanceY float64 = abs(particule2.PositionY - particule1.PositionY)
    if particule1.PositionY < particule2.PositionY{//X 1-->2
        if distanceY <= particule1.ScaleY*10{return true}
    }else{//X 2-->1
        if distanceY <= particule2.ScaleY*10{return true}}
    return false}

func rebond_particules(particules []Particle) []Particle{
    for i := 0; i < len(particules); i++ {//pour toutes les particules
        for j := i+1; j < len(particules); j++ {//pour toutes les particules suivantes
            if i != j{
                if collision(particules[i],particules[j]){
                    particules[i].SpeedX,particules[j].SpeedX = particules[j].SpeedX,particules[i].SpeedX
                    particules[i].SpeedY,particules[j].SpeedY = particules[j].SpeedY,particules[i].SpeedY
    }}}}
    return particules}

func acceleration(particule Particle)Particle{
    particule.SpeedX *= config.General.Acceleration
    particule.SpeedY *= config.General.Acceleration
    return particule
}

func energie(particules []Particle)float64{
    var energie float64
    //centre X
    var centreX float64
    for i := 0; i < len(particules); i++ {
        centreX += particules[i].PositionX
    }
    centreX /= float64(len(particules))
    //centre Y
    var centreY float64
    for i := 0; i < len(particules); i++ {
        centreY += particules[i].PositionY
    }
    centreY /= float64(len(particules))
    //calcul énergie mécanique
    for i := 0; i < len(particules); i++ {//pour toutes les particules
        var energie_cinetique float64 = masse(particules[i])*(particules[i].SpeedX*particules[i].SpeedX+particules[i].SpeedY*particules[i].SpeedY)/2
        var distanceX float64 = particules[i].PositionX-centreX
        var distanceY float64 = particules[i].PositionY-centreY
        var energie_potentielle float64 = math.Sqrt(distanceX*distanceX + distanceY*distanceY)*masse(particules[i])
        energie+= energie_cinetique + energie_potentielle}
    return math.Log(energie)
}

func attraction(particule Particle, force float64, angle float64, position bool, x, y float64)Particle{
    if position{//si attraction vers un point
        var distanceX float64 = x - particule.PositionX
        var distanceY float64 = y - particule.PositionY
        angle = math.Atan(distanceY/distanceX)
    }else{
        angle *= (2*math.Pi)/360
    }
    if particule.PositionX < x{force=-force}
    particule.SpeedX += math.Cos(angle)*force
    particule.SpeedY += math.Sin(angle)*force
    return particule
}

func gravitation(particules []Particle)[]Particle{
    for i := 0; i < len(particules); i++ {
        for j := 0; j < len(particules); j++ {
            if i != j{
                var distance float64 = ((particules[i].PositionX+particules[i].ScaleX*5-particules[j].ScaleX*5) - (particules[j].PositionX+particules[j].ScaleX*5-particules[i].ScaleX*5))*((particules[i].PositionX+particules[i].ScaleX*5-particules[j].ScaleX*5) - (particules[j].PositionX+particules[j].ScaleX*5-particules[i].ScaleX*5)) + ((particules[i].PositionY+particules[i].ScaleY*5-particules[j].ScaleY*5) - (particules[j].PositionY+particules[j].ScaleY*5-particules[i].ScaleY*5))*((particules[i].PositionY+particules[i].ScaleY*5-particules[j].ScaleY*5) - (particules[j].PositionY+particules[j].ScaleY*5-particules[i].ScaleY*5))
                if (distance < particules[i].ScaleX*10 || distance < particules[j].ScaleX*10) && config.General.Merge{
                	particules[i] = growth(particules[i],particules[j].ScaleX*particules[j].ScaleY)
                	particules[i].SpeedX += particules[j].SpeedX*(masse(particules[j])/masse(particules[i]))
                	particules[i].SpeedY += particules[j].SpeedY*(masse(particules[j])/masse(particules[i]))
                	particules[i].SpeedX *= masse(particules[j])/(masse(particules[i])+masse(particules[j]))
                	particules[i].SpeedY *= masse(particules[j])/(masse(particules[i])+masse(particules[j]))
                	particules = suppression(particules,j)
           		}else{
                	if distance < 60{distance = 60}
           			particules[i] = attraction(particules[i], masse(particules[j])*config.General.Gravitation*(1/(distance)), 0, true, particules[j].PositionX+particules[j].ScaleX*5-particules[i].ScaleX*5, particules[j].PositionY+particules[j].ScaleY*5-particules[i].ScaleY*5)
           			particules[i] = attraction(particules[i], -100*masse(particules[j])*config.General.Gravitation*(1/(distance*distance)), 0, true, particules[j].PositionX+particules[j].ScaleX*5-particules[i].ScaleX*5, particules[j].PositionY+particules[j].ScaleY*5-particules[i].ScaleY*5)
                }
            }
        }
    }
    return particules
}

func masse(particule Particle)float64{return particule.ScaleX*particule.ScaleY}

func growth(particule Particle, x float64)Particle{
	particule.ScaleX = math.Sqrt(particule.ScaleX*particule.ScaleY+x)
	particule.ScaleY = math.Sqrt(particule.ScaleX*particule.ScaleY+x)
	return particule
}