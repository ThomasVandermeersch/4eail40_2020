package main

// En gros, un package c'est un répertoire qui contient des fichiers/des bouts de code , qu'on va vouloir utiliser
//souvent, donc au lieu de réécrire les fonctions, on importe un package.
//Pour qu'un fichier soit exécutable tout seul, il doit se retrouver dans le package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"../../model"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
}

func runCommand(commandStr string) (e error) {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	args := strings.Fields(commandStr)
	switch args[0] {
	case "exit":
		os.Exit(0)
	// add another case here for custom commands.
	case "new":
		a := model.Board8{}
		a.Initialize()
		a.String()
		// TODO Create a new game on a classic 8x8 board.
		// TODO Display the board on console.
		break
	case "move":
		// TODO Move a piece. (syntax: move <from> <to>)
		// TODO The command line should be in the form of move A2 A4.
		// TODO     => meaning move piece from position A2 to A4
		// TODO Display the board on console.
		break
	default:
		e = fmt.Errorf("unknown command %s", args[0])
	}
	return
}
