package particles
import ("testing";"project-particles/config"/*;"fmt"*/)


func Test_Vide_Creation_Particules(t *testing.T){
	var systeme System = NewSystem()//création d'un nouveau système (avec aucune particule)
	if len(systeme.Content) != 0{//s'il y a au moins une particule...
		t.Error("NewSystem est vide.")
	}
}

func Test_Init_Particule(t *testing.T){
	config.General.InitNumParticles = 5//initialisation du nombre de particule à afficher au début à 5
	var systeme System = NewSystem()//création d'un nouveau système prenant les informations de NewSystem à son execution
	if len(systeme.Content) != 5{//si le nombre de particule du système est différent de 5...
		t.Error("Le nombre de particules est incorrect")
	}
}

func Test_Random_Spawn(t *testing.T){
	config.General.RandomSpawn = false//établir l'apparition aléatoire en position des particule comme désactivé
	config.General.SpawnX = 400//initialiser la position X à apparaitre des particules à 400 pixels de la gauche de la fenêtre 
	config.General.SpawnY = 300//initialiser la position Y à apparaitre des particules à 300 pixels du haut de la fenêtre
	var systeme System = NewSystem()//création d'un nouveau système prenant les informations de NewSystem à son execution
	for i := 0; i<len(systeme.Content); i++{//pour toutes les particules générées par la fonction NewSystem...
		if systeme.Content[i].PositionX != 400 || systeme.Content[i].PositionY != 300{//...si sa position est incorecte par rapport aux attente X et Y...
			t.Error("Ce n'est pas le bon endroit d'apparition")
		}
	}
}

func Test_deplacement(t *testing.T){
	var particules_temoin []Particle = NewSystem().Content//création d'un tableau de particules témoin qui servira de comparaison à celui auquel on appliquera la fonction
	var particules_test []Particle = deplacement(NewSystem().Content)//création d'un tableua de particules de test qui sert à se faire comparer à celui temoin afin de vérifier les changement effectué après application de la fonction
	for i := 0; i < len(particules_test); i++ {//pour toutes les particules du tableau test...
		if particules_test[i].PositionX == particules_temoin[i].PositionX &&//si la position X de la particule est égale à celle de correspondante du tableau temoin...
		particules_test[i].PositionY == particules_temoin[i].PositionY{//..et si la position Y de la particule est égale à celle de correspondante du tableau temoin...
			t.Error("Les position n'ont pas été changées.")
		}
	} 
}

func Test_deplacement_v4(t *testing.T){
	var PositionX float64 = 400//initialiser la position X à apparaitre des particules à 400 pixels de la gauche de la fenêtre 
	var PositionY float64 = 300//initialiser la position Y à apparaitre des particules à 300 pixels du haut de la fenêtre
	var particules []Particle = ajout([]Particle{},PositionX,PositionY,1,1)//création d'un tableau de particules auquel on en ajoute une avec une vitesse non-nulle
	particules = deplacement(particules)//déplacement de la particule grâce à sa vitesse
	if particules[0].PositionX == 400 && particules[0].PositionY == 300{//si les coordonnées de la particule sont les mêmes qu'à sa création...
		t.Error("Les particules n'ont bougées.")
	}
}

func Test_suppression_Position_X_min(t *testing.T){
	var PositionX float64 = -50//initialisation d'une position de coordonnées X = -50, donc en dehors, à gauche, de la fenêtre...
	var PositionY float64 = float64(600)/2//et d'une coordonnée Y de 600/2, qui dans ce contexte est assimilable à son placement en hauteur du milieu de la fenêtre
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 1, 1)//création d'un tableau de particules contenant une particule créée avec les précédentes informations
	var particules_test []Particle = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 0, 100)//application de la fonction sous des conditions de suppression tel que la particule sera supprimée du fait d'être en dehors de la fenêtre de positions acceptées
	if len(particules_test) != 0{//si la particule n'a pas été supprimée...
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_suppression_Position_X_max(t *testing.T){
	var PositionX float64 = float64(800) + 50//initialisation d'une position de coordonnées X = (taille de la fenêtre dans ce contexte) + 50, donc en dehors, à droite, de la fenêtre...
	var PositionY float64 = float64(600)/2//et d'une coordonnée Y de 600/2, qui dans ce contexte est assimilable à son placement en hauteur du milieu de la fenêtre
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 1, 1)//création d'un tableau de particules contenant une particule créée avec les précédentes informations
	var particules_test []Particle = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 0, 100)//application de la fonction sous des conditions de suppression tel que la particule sera supprimée du fait d'être en dehors de la fenêtre de positions acceptées
	if len(particules_test) != 0{//si la particule n'a pas été supprimée...
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_suppression_Position_Y_min(t *testing.T){
	var PositionX float64 = float64(800)/2//initialisation d'une position de coordonnées X = 800/2, qui dans ce contexte est assimilable à son placement de droite à gauche au mili de la fenêtre...
	var PositionY float64 = -50//et d'une coordonnée Y = -50 donc en dehors, au dessus, de la fenêtre
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 1, 1)//création d'un tableau de particules contenant une particule créée avec les précédentes informations
	var particules_test []Particle = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 0, 100)//application de la fonction sous des conditions de suppression tel que la particule sera supprimée du fait d'être en dehors de la fenêtre de positions acceptées
	if len(particules_test) != 0{//si la particule n'a pas été supprimée...
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_suppression_Position_Y_max(t *testing.T){
	var PositionX float64 = float64(800)/2//initialisation d'une position de coordonnées X = 800/2, qui dans ce contexte est assimilable à son placement de droite à gauche au mili de la fenêtre...
	var PositionY float64 = float64(600) + 50//et d'une coordonnée Y = 50 donc en dehors, en dessous, de la fenêtre
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 1, 1)//création d'un tableau de particules contenant une particule créée avec les précédentes informations
	var particules_test []Particle = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 0, 100)//application de la fonction sous des conditions de suppression tel que la particule sera supprimée du fait d'être en dehors de la fenêtre de positions acceptées
	if len(particules_test) != 0{//si la particule n'a pas été supprimée...
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_suppression_Taille_min(t *testing.T){
	var PositionX float64 = float64(800)/2//initialisation d'une position de coordonnées X = 800/2, qui dans ce contexte est assimilable à son placement de droite à gauche au mili de la fenêtre...
	var PositionY float64 = float64(600)/2//et d'une coordonnée Y de 600/2, qui dans ce contexte est assimilable à son placement en hauteur du milieu de la fenêtre
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 2, 1)//création d'un tableau de particules contenant une particule créée avec les précédentes informations avec comme particularité une taille inférieure à l'exigence qui suit
	var particules_test []Particle = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 5, 100)//application de la fonction sous des conditions de suppression tel que la particule sera supprimée du fait d'être trop petite
	if len(particules_test) != 0{//si la particule n'a pas été supprimée...
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_suppression_Taille_max(t *testing.T){
	var PositionX float64 = float64(800)/2//initialisation d'une position de coordonnées X = 800/2, qui dans ce contexte est assimilable à son placement de droite à gauche au mili de la fenêtre...
	var PositionY float64 = float64(600)/2//et d'une coordonnée Y de 600/2, qui dans ce contexte est assimilable à son placement en hauteur du milieu de la fenêtre
	var particules_temoin []Particle = ajout([]Particle{}, PositionX, PositionY, 150, 1)//création d'un tableau de particules contenant une particule créée avec les précédentes informations avec comme particularité une taille supérieure à l'exigence qui suit
	var particules_test []Particle = condition_suppression(particules_temoin, 0,float64(800), 0,float64(600), 0, 100)//application de la fonction sous des conditions de suppression tel que la particule sera supprimée du fait d'être trop grande
	if len(particules_test) != 0{//si la particule n'a pas été supprimée...
		t.Error("La particule aurait dû être supprimée.")
	}
}

func Test_grossissement_Additif_Positif(t* testing.T){
	var taille float64 = 100//initialisation d'une variable de taille valant 100 compris en pixels
	var particules_temoin []Particle = ajout([]Particle{}, 0, 0, taille, 1)//création d'un tableau temoin de particules contenant une particule de taille précédemment défini
	var particules_test []Particle = grossissement(particules_temoin, false, 10)//application de la fonction de grossissement tel que la taille doit être incrémentée de 10 pixels
	if particules_test[0].ScaleX != float64(100)+10{//si la taille n'est pas celle attendue (110) (le float sert à ignorer l'imprecision des calculs entre différentes bases)...
		t.Error("La particule n'a pas grossi.")
	}
}

func Test_grossissement_Additif_Negatif(t* testing.T){
	var taille float64 = 100//initialisation d'une variable de taille valant 100 compris en pixels
	var particules_temoin []Particle = ajout([]Particle{}, 0, 0, taille, 1)//création d'un tableau temoin de particules contenant une particule de taille précédemment défini
	var particules_test []Particle = grossissement(particules_temoin, false, -10)//application de la fonction de grossissement tel que la taille doit être décrémentée de 10 pixels
	if particules_test[0].ScaleX != float64(100)-10{//si la taille n'est pas celle attendue (90) (le float sert à ignorer l'imprecision des calculs entre différentes bases)...
		t.Error("La particule n'a pas maigri.")
	}
}

func Test_grossissement_Multiplicatif_Positif(t* testing.T){
	var taille float64 = 100//initialisation d'une variable de taille valant 100 compris en pixels
	var particules_temoin []Particle = ajout([]Particle{}, 0, 0, taille, 1)//création d'un tableau temoin de particules contenant une particule de taille précédemment défini
	var particules_test []Particle = grossissement(particules_temoin, true, 1.1)//application de la fonction de grossissement tel que la taille doit être multipliée par 1,1
	if particules_test[0].ScaleX != float64(100)*1.1{//si la taille n'est pas celle attendue (110) (le float sert à ignorer l'imprecision des calculs entre différentes bases)...
		t.Error("La particule n'a pas grossi.")
	}
}

func Test_grossissement_Multiplicatif_Negatif(t* testing.T){
	var taille float64 = 100//initialisation d'une variable de taille valant 100 compris en pixels
	var particules_temoin []Particle = ajout([]Particle{}, 0, 0, taille, 1)//création d'un tableau temoin de particules contenant une particule de taille précédemment défini
	var particules_test []Particle = grossissement(particules_temoin, true, 0.9)//application de la fonction de grossissement tel que la taille doit être multipliée par 0,9
	if particules_test[0].ScaleX != float64(100)*0.9{//si la taille n'est pas celle attendue (90) (le float sert à ignorer l'imprecision des calculs entre différentes bases)...
		t.Error("La particule n'a pas grossi.")
	}
}

func Test_Ajout_Creation_1(t *testing.T){
	if len(ajout([]Particle{},0,0,0,0)) != 1{//si aucune particule ne se crée avec la fonction...
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
		t.Error("Taille X différente de 4.")}
	if test[len(test)-1].ScaleY != 4{
		t.Error("Taille Y différente de 4.")}
	if (test[len(test)-1].SpeedX)*(test[len(test)-1].SpeedX)+(test[len(test)-1].SpeedY)*(test[len(test)-1].SpeedY) >= float64(7*7)+1e-12 || 
		(test[len(test)-1].SpeedX)*(test[len(test)-1].SpeedX)+(test[len(test)-1].SpeedY)*(test[len(test)-1].SpeedY) <= float64(7*7)-1e-12{
		t.Error("Speed X différente de 7.",(test[len(test)-1].SpeedX)*(test[len(test)-1].SpeedX)+(test[len(test)-1].SpeedY)*(test[len(test)-1].SpeedY))}
}

func Test_Ajout_2(t *testing.T){
	var systeme System 
	systeme.Content = ajout(systeme.Content,10,3,4,7)
	if len(systeme.Content) != 1{
		t.Error("Il n'y a pas le bon nombre de particules crées")
	}
}

func Test_Creation_Particule_New_System(t *testing.T){
	var systeme System = NewSystem()
	if len(systeme.Content) == 0{
		t.Error("La particule n'a été ajoutée")
	}
}

func Test_Spawn_Rate1(t *testing.T){
	config.General.SpawnRate = 1
	var systeme System
	systeme.Update()
	if len(systeme.Content) != 1{
		t.Error("Le taux d(apparition est mauvais")
	}
}