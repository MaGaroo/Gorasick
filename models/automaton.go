package models

type Automaton struct {
	entryNode *Node
}

func NewAutomaton() *Automaton {
	return &Automaton{
		entryNode: NewNode(),
	}
}

func (automata *Automaton) AddPattern(pattern string) {
	var currentNode *Node
	currentNode = automata.entryNode
	for _, char := range pattern {
		currentNode = currentNode.GetChild(char, true)
	}
}

func (automata *Automaton) Build() {

}
