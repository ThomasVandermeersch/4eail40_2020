package piece

import (
	"fmt"

	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/player"
)

type ChessPiece struct {
	playerColor player.Color
	name        string
}

func NewChessPiece(playerColor player.Color, name string) ChessPiece {
	return ChessPiece{playerColor, name}
}

func (chessPiece ChessPiece) Moves(isCapture bool) map[coord.ChessCoordinates]bool {
	//TO DO
	return nil
}

func (chessPiece ChessPiece) Color() player.Color {
	//TO DO
	return chessPiece.playerColor
}

func (chessPiece ChessPiece) String() string {
	return fmt.Sprintf(chessPiece.name)
}

func (chessPiece ChessPiece) Name() string {
	return chessPiece.name
}
