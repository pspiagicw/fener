package compile

type SymbolScope string

const (
	GLOBAL SymbolScope = "GLOBAL"
)

type Symbol struct {
	Index int
	Name  string
	Scope SymbolScope
}

type SymbolTable struct {
	symbols map[string]Symbol
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		symbols: make(map[string]Symbol),
	}
}

func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{
		Name:  name,
		Index: len(s.symbols),
		Scope: GLOBAL,
	}
	s.symbols[name] = symbol
	return symbol
}
func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	symbol, ok := s.symbols[name]
	return symbol, ok
}
