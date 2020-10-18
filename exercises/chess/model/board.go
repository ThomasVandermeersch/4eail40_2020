// Package model contains the gameplay logic for the game of chess
package model

import "fmt"
import "strconv"
//TODO Implement type Board

type Board interface{
	Initialize() 
}

type Coord interface{

}

type Coord2D struct{
	X int
	Y int
}

type Board8 struct {
	board map[Coord2D]Piece
}

func (board8 *Board8) Initialize() {
	board8.board = make(map[Coord2D]Piece)
	for x:=0 ; x < 8 ; x++{
		for y:=0; y < 8 ; y++{
			coord := Coord2D{x,y}
			
			
			piece := Piece{"-","action"}
			
			//Placement des pions
			if y==1{
				piece = Piece{"P","Blanc"}
			} else if y==6 {
				piece = Piece{"P","Noir"}
			
			//Placement des tours
			} else if ((y == 0 && x == 0) || (y == 0 && x == 7)){
				piece = Piece{"T","Blanc"}
			
			} else if ((y == 7 && x == 0) || (y == 7 && x == 7)){
				piece = Piece{"T","Noir"}
			
			} else if ((y==0 && x==1) || (y==0 && x==6)){
				piece = Piece{"C","Blanc"}
			} else if ((y==7 && x==1) || (y==7 && x==6)){
				piece = Piece{"C","Noir"}
			
			} else if ((y==0 && x==2) || (y==0 && x==5)){
				piece = Piece{"F","Blanc"}
			} else if ((y==7 && x==2) || (y==7 && x==5)){
				piece = Piece{"F","Noir"}
			
			} else if (y==0 && x==3) {
				piece = Piece{"R","Blanc"}
			} else if (y==7 && x==3){
				piece = Piece{"R","Noir"}
			
			} else if (y==0 && x==4) {
				piece = Piece{"r","Blanc"}
			} else if (y==7 && x==4){
				piece = Piece{"r","Noir"}
			}

			board8.board[coord] = piece
		}
	}
}


func(board8 *Board8) String() {
	result := ""
	for y:=0 ; y < 8 ; y++{
		for x:=0; x < 8 ; x++{
			coord := Coord2D{x,y}
			a := board8.board[coord] 
			if a.Couleur == "Blanc"{
				result += "\033[31m"  //rouge
			}else if (a.Couleur == "Noir"){
				result += "\033[34m"  //bleu
			}
			
			result += a.Name
			result += "\033[0m"
			result += " | "
		}
		result += "\n"
	}
	fmt.Println( result)

}

func (board8 *Board8) Move(xx1, yy1 , xx2, yy2 string){
	x1, err := strconv.Atoi(xx1)
	y1, err := strconv.Atoi(yy1)
	x2, err := strconv.Atoi(xx2)
	y2, err := strconv.Atoi(yy2)
	if(err ==nil){	
		fromCoord := Coord2D{x1,y1}
		toCoord := Coord2D{x2,y2}

		piece := board8.board[toCoord]
		board8.board[toCoord] = board8.board[fromCoord]
		board8.board[fromCoord] = piece	
	}
}