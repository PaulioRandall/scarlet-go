package lexeme

func CopyAll(head *Lexeme) *Lexeme {

	var que Queue = &Container{}

	for lex := head; lex != nil; lex = lex.Next {
		que.Put(Copy(lex))
	}

	return que.Head()
}

func Copy(lex *Lexeme) *Lexeme {
	return &Lexeme{
		Props: lex.Props,
		Raw:   lex.Raw,
	}
}
