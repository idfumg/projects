package game

import (
	"fmt"
	"math/rand"

	"hangman/utils"
)

type Game struct {
	target         string
	current        []rune
	charsPositions map[rune][]int
}

func NewGame(target string) *Game {
	current := getMaskedWord(target)
	return &Game{
		target:         target,
		current:        current,
		charsPositions: getCharsPositions(current, target),
	}
}

func changeOneCharacter(s []rune) []rune {
	for {
		if i := rand.Intn(len(s)); s[i] != '_' && s[i] != ' ' {
			s[i] = '_'
			break
		}
	}
	return s
}

func isEnoughUnderscores(s []rune) bool {
	return float64(utils.GetCount(s, '_')) >= float64(len(s))/1.9
}

func getMaskedWord(ss string) []rune {
	s := []rune(ss)
	for !isEnoughUnderscores(s) {
		s = changeOneCharacter(s)
	}
	return s
}

func getCharsPositions(s []rune, t string) map[rune][]int {
	ans := map[rune][]int{}
	for i := 0; i < len(s); i += 1 {
		if s[i] != '_' {
			continue
		}
		ans[rune(t[i])] = append(ans[rune(t[i])], i)
	}
	return ans
}

func (g *Game) PrintState() {
	fmt.Printf("%s\n", string(g.current))
}

func (g *Game) ChangeState(c rune) {
	if idxs, ok := g.charsPositions[c]; ok {
		i := idxs[len(idxs)-1]
		g.charsPositions[c] = g.charsPositions[c][0 : len(idxs)-1]
		if len(g.charsPositions[c]) == 0 {
			delete(g.charsPositions, c)
		}
		g.current[i] = rune(g.target[i])
	}
}

func (g *Game) IsDone() bool {
	return utils.GetCount(g.current, '_') == 0
}
