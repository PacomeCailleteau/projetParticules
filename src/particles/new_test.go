package particles
import ("testing";"project-particles/config"/*;"fmt"*/)


func Test_Vide_Creation_Particules(t *testing.T){
	s := NewSystem()
	if len(s.Content) != 0{
		t.Error("NewSystem est vide.")
	}
}

func Test_Creation_Particules(t *testing.T){
	config.General.InitNumParticles = 5
	s := NewSystem()
	if len(s.Content) != 5{
		t.Error("Il n'y a pas le bon nombre de particules crées")
	}
}

func Test_deplacement(t *testing.T){
	var particules_temoin []Particle = NewSystem().Content
	var particules_test []Particle = deplacement(NewSystem().Content)
	var changed bool
	for i := 0; i < len(particules_test); i++ {
		if particules_test[i].PositionX != particules_temoin[i].PositionX ||
		particules_test[i].PositionY != particules_temoin[i].PositionY{
			changed = true
		}
	}
	if changed == false{
		t.Error("Les position n'ont pas été changées.")
	} 
}
/*
func Test_deplacement_v2(t *testing.T){
	var PositionX float64 = float64(800)/2
	var PositionY float64 = float64(600)/2
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 1, 10)
	fmt.Println(particules_temoin[0].PositionX, particules_temoin[0].PositionY)
	var particules_test []Particle = particules_temoin
	fmt.Println(particules_temoin[0].PositionX, particules_temoin[0].PositionY)
	particules_test = deplacement(particules_test)
	fmt.Println(particules_temoin[0].PositionX, particules_temoin[0].PositionY)
	if particules_test[0].PositionX == particules_temoin[0].PositionX &&
	particules_test[0].PositionY == particules_temoin[0].PositionY{
		t.Error("La particule n'a pas bougé.", particules_test[0].PositionX, particules_temoin[0].PositionX, particules_test[0].PositionY, particules_temoin[0].PositionY)
	}
}
*/
/*
func Test_deplacement_v3(t *testing.T){
	var PositionX float64 = 400
	var PositionY float64 = 300
	var particules_temoin []Particle = ajout([]Particle{},PositionX,PositionY,1,1)
	fmt.Println("--------")
	fmt.Println(particules_temoin)
	var affichage Particle = ajout([]Particle{},PositionX,PositionY,1,1)[0]
	fmt.Println(affichage)
	fmt.Println("//////////////////")
	fmt.Println("affichage.PositionX :",affichage.PositionX)
	fmt.Println("affichage.PositionY :",affichage.PositionY)
	fmt.Println("affichage.Rotation :",affichage.Rotation)
	fmt.Println("affichage.ScaleX :",affichage.ScaleX)
	fmt.Println("affichage.ScaleY :",affichage.ScaleY)
	fmt.Println("affichage.SpeedX :",affichage.SpeedX)
	fmt.Println("affichage.SpeedY :",affichage.SpeedY)
	fmt.Println("Color RGB :",affichage.ColorRed,affichage.ColorGreen,affichage.ColorBlue)
	fmt.Println("affichage.Opacity :",affichage.Opacity)
	fmt.Println("//////////////////")
	var particules_test []Particle = particules_temoin
	fmt.Println("--------")
	fmt.Println(particules_test)
	particules_test = deplacement(particules_test)
	fmt.Println("--------")
	fmt.Println(particules_test)
	fmt.Println(particules_temoin)
	fmt.Println(deplacement(particules_test))
	fmt.Println("--------")
	fmt.Println(len(particules_test))
	for i := 0; i < len(particules_test); i++ {
		if particules_test[i].PositionX == particules_temoin[i].PositionX &&
		particules_test[i].PositionY == particules_temoin[i].PositionY{
			t.Error("Les positions n'ont été changées", particules_temoin[i].PositionX,particules_test[i].PositionX)
		}
	}
}
*/

func Test_deplacement_v4(t *testing.T){
	var PositionX float64 = 400
	var PositionY float64 = 300
	var particules []Particle = ajout([]Particle{},PositionX,PositionY,1,1)
	particules = deplacement(particules)
	if particules[0].PositionX == 400 && particules[0].PositionY == 300{
		t.Error("Les particules n'ont bougées.")
	}
}

func Test_suppression_Position_X_min(t *testing.T){
	var PositionX float64 = -50
	var PositionY float64 = float64(600)/2
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 1, 1)
	var particules_test []Particle = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 0, 100)
	if len(particules_test) != 0{
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_suppression_Position_X_max(t *testing.T){
	var PositionX float64 = float64(800) + 50
	var PositionY float64 = float64(600)/2
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 1, 1)
	var particules_test []Particle = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 0, 100)
	if len(particules_test) != 0{
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_suppression_Position_Y_min(t *testing.T){
	var PositionX float64 = float64(800)/2
	var PositionY float64 = -50
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 1, 1)
	var particules_test []Particle = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 0, 100)
	if len(particules_test) != 0{
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_suppression_Position_Y_max(t *testing.T){
	var PositionX float64 = float64(800)/2
	var PositionY float64 = float64(600) + 50
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 1, 1)
	var particules_test []Particle = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 0, 100)
	if len(particules_test) != 0{
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_suppression_Taille_min(t *testing.T){
	var PositionX float64 = float64(800)/2
	var PositionY float64 = float64(600)/2
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 2, 1)
	var particules_test []Particle = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 5, 100)
	if len(particules_test) != 0{
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_suppression_Taille_max(t *testing.T){
	var PositionX float64 = float64(800)/2
	var PositionY float64 = float64(600)/2
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 150, 1)
	var particules_test []Particle = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 0, 100)
	if len(particules_test) != 0{
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_grossissement_Additif_Positif(t* testing.T){
	var taille float64 = 100
	var particules_temoin []Particle = ajout([]Particle{}, 0, 0, taille, 1)
	var particules_test []Particle = grossissement(particules_temoin, false, 10)
	if particules_test[0].ScaleX != float64(100)+10{
		t.Error("La particule n'a pas grossi.")
	}
}

func Test_grossissement_Additif_Negatif(t* testing.T){
	var taille float64 = 100
	var particules_temoin []Particle = ajout([]Particle{}, 0, 0, taille, 1)
	var particules_test []Particle = grossissement(particules_temoin, false, -10)
	if particules_test[0].ScaleX != float64(100)-10{
		t.Error("La particule n'a pas maigri.")
	}
}

func Test_grossissement_Multiplicatif_Positif(t* testing.T){
	var taille float64 = 100
	var particules_temoin []Particle = ajout([]Particle{}, 0, 0, taille, 1)
	var particules_test []Particle = grossissement(particules_temoin, true, 1.1)
	if particules_test[0].ScaleX != float64(100)*1.1{
		t.Error("La particule n'a pas grossi.")
	}
}

func Test_grossissement_Multiplicatif_Negatif(t* testing.T){
	var taille float64 = 100
	var particules_temoin []Particle = ajout([]Particle{}, 0, 0, taille, 1)
	var particules_test []Particle = grossissement(particules_temoin, true, 0.9)
	if particules_test[0].ScaleX != float64(100)*0.9{
		t.Error("La particule n'a pas grossi.")
	}
}

func Test_Ajout_Creation_1(t *testing.T){
	if len(ajout([]Particle{},0,0,0,0)) != 1{
		t.Error("Aucune particule créée.")
	}
}

func Test_Ajout_Creation_2(t *testing.T){
	var test []Particle = ajout([]Particle{},0,0,0,0)
	if test[0].PositionX != 0{
		t.Error("X différent de 0.")}
	if test[0].PositionY != 0{
		t.Error("Y différent de 0.")}
	if test[0].Rotation != 0{
		t.Error("Rotation non nulle")}
	if test[0].ScaleX != 0{
		t.Error("Taille X non nulle.")}
	if test[0].ScaleY != 0{
		t.Error("Taille Y non nulle.")}
}

func Test_Ajout_Creation_3(t *testing.T){
	var test []Particle = ajout([]Particle{},10,3,4,7)
	if test[0].PositionX != 10{
		t.Error("X différent de 10.")}
	if test[0].PositionY != 3{
		t.Error("Y différent de 3.")}
	if test[0].Rotation != 0{
		t.Error("Rotation non nulle")}
	if test[0].ScaleX != 4{
		t.Error("Taille X différente de 7.")}
	if test[0].ScaleY != 4{
		t.Error("Taille Y différente de 7.")}
}

func Test_Ajout_1(t *testing.T){
	var test []Particle = ajout(NewSystem().Content,10,3,4,7)
	if len(test) != len(NewSystem().Content)+1{
		t.Error("La particule n'a été ajoutée.")
	}
	if test[len(test)-1].PositionX != 10{
		t.Error("X différent de 10.")}
	if test[len(test)-1].PositionY != 3{
		t.Error("Y différent de 3.")}
	if test[len(test)-1].Rotation != 0{
		t.Error("Rotation non nulle")}
	if test[len(test)-1].ScaleX != 4{
		t.Error("Taille X différente de 7.")}
	if test[len(test)-1].ScaleY != 4{
		t.Error("Taille Y différente de 7.")}
}

/*func ajout(particules []Particle, PositionX, PositionY, taille, mult_vitesse float64) []Particle{
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
}*/