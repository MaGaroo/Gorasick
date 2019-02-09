package models

type PatternMatcher struct {
	patterns  []string
	automaton *Automaton
}
