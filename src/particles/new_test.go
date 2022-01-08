package particles
import ("testing";"project-particles/config")


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
	config.General.SpawnX = 150
	config.General.SpawnY = 250
	s:=NewSystem()
	config.General.SpawnX = 150
	config.General.SpawnY = 250
	t:=NewSystem()
	*t.Update()
	if s.Content.PositionX == *t && s.Content.PositionY == *t{
		t.Error("La particule ne bouge pas")
	} 
}

func Test_suppression(t *testing.T){
	return 0
}

func Test_grossissement(t* testing.T){
	return 0
}