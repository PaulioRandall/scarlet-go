package logging

/*
func writeLexemes(w io.Writer, head *lexeme.Lexeme) error {

	for lex := head; lex != nil; lex = lex.Next {

		b := []byte(lex.Raw)
		_, e := w.Write(b)

		if e != nil {
			return e
		}
	}

	return nil
}

/*
func logPhase(c config, ext string, head *lexeme.Lexeme) error {

	if !c.log {
		return nil
	}

	f := c.logFilename(ext)
	return writeLexemeFile(f, head)
}

/*
func writeLexemeFile(filename string, head *lexeme.Lexeme) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()
	return lexeme.PrintAll(f, head)
}

/*
func writeInstPhaseFile(filename string, ins []inst.Instruction) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()
	return inst.PrintAll(f, ins)
}

/*
func (b config) logFilename(ext string) string {
	f := filepath.Base(b.script)
	f = strings.TrimSuffix(f, filepath.Ext(f))
	return filepath.Join(b.logFile, f+ext)
}
*/
