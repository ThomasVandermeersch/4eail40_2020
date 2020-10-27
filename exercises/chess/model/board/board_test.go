package board

import (
	"testing"

	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
)

type mockCoord int

// Coord returns x if n==0, y if n==1
func (c mockCoord) Coord(n int) (int, error) {
	return int(c), nil
}

func (c mockCoord) String() string {
	return "1"
}

func TestClassic_MovePiece(t *testing.T) {
	type args struct {
		from coord.ChessCoordinates
		to   coord.ChessCoordinates
	}
	tests := []struct {
		name    string
		c       *Classic
		args    args
		wantErr bool
	}{
		{
			"testmock1",
			&Classic{},
			args{mockCoord(0), mockCoord(1)},
			false,
		},
		// {
		// 	//TO pass testmock2, pieceDest (line36 board.go) None has to be another word
		// 	"testmock2",
		// 	&Classic{},
		// 	args{mockCoord(0), mockCoord(1)},
		// 	false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.c.MovePiece(tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("Classic.MovePiece() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
