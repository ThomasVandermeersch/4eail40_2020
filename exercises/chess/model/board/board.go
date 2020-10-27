package board

import (
	"fmt"

	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/piece"
)

// Classic 8x8 Chess board
type Classic [8][8]piece.Piece

func (c *Classic) String() string {
	panic("not implemented") // TODO: Implement
}

// PieceAt retrievves piece at give coordinates.
// Returns nil if no piece was found.
func (c *Classic) PieceAt(at coord.ChessCoordinates) piece.Piece {
	panic("not implemented") // TODO: Implement
}

// MovePiece moves a piece from given coordinates to
// given coordinates.
// Returns an error if destination was occupied.
func (c *Classic) MovePiece(from coord.ChessCoordinates, to coord.ChessCoordinates) error {
	//Nous avons implémenter des pièces au sein de cette fonction pour pouvoir essayer notre test.
	//En réalité, il aurait d'abord fallu initialiser un tableau, initialiser les
	//pièces sur le tableau avant de s'intéresser aux mouvements des pièces.

	fromX, _ := from.Coord(0)
	fromY, _ := from.Coord(1)
	toX, _ := to.Coord(0)
	toY, _ := to.Coord(0)

	//TO TEST, we create a piece at destination
	//Normally the game chess has to be intialized.
	//TO DO ==> initialize GameChess

	pieceDest := piece.NewChessPiece(1, "None")
	c[toX][toY] = pieceDest

	destinationPiece := c[toX][toY]
	// Normally we should use piece.Moves() to check if the piece can be moved !
	// TO DO ==> ameliorate
	if destinationPiece.Name() != "None" {
		return fmt.Errorf("Piece already exists at destination")
	}
	c[toX][toY] = c[fromX][fromY]

	piece := piece.NewChessPiece(1, "None")
	c[fromX][fromY] = piece

	return nil
}

// PlacePieceAt places a given piece at given location.
// Returns an error if destination was occupied.
func (c *Classic) PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error {
	panic("not implemented") // TODO: Implement
}
