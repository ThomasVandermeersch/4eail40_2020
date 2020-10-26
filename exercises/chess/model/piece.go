package model

//TODO Implement game pieces here

type Piece interface{
	Initialize()
}

type Piece struct{
	Name string
	Couleur string
}

func (piece Piece) Initialize() Piece {

}