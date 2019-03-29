const (
	devNull = "/dev/null"
)

func (p *parser) ParseGitFileHeader() (*File, error) {
	const prefix = "diff --git "

	if !strings.HasPrefix(p.Line(0), prefix) {
		return nil, nil
	}
	header := p.Line(0)[len(prefix):]

		return nil, p.Errorf(0, "git file header: %v", err)
	f := &File{}
		end, err := parseGitHeaderData(f, p.Line(1), defaultName)
		if err != nil {
			return nil, p.Errorf(1, "git file header: %v", err)
		if err := p.Next(); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err

			return nil, p.Errorf(0, "git file header: missing filename information")
		return nil, p.Errorf(0, "git file header: missing filename information")
	return f, nil
func (p *parser) ParseTraditionalFileHeader() (*File, error) {
	const shortestValidFragHeader = "@@ -0,0 +1 @@\n"
	const (
		oldPrefix = "--- "
		newPrefix = "+++ "
	)

	oldLine, newLine := p.Line(0), p.Line(1)

	if !strings.HasPrefix(oldLine, oldPrefix) || !strings.HasPrefix(newLine, newPrefix) {
		return nil, nil
	}
	// heuristic: only a file header if followed by a (probable) fragment header
	if len(p.Line(2)) < len(shortestValidFragHeader) || !strings.HasPrefix(p.Line(2), "@@ -") {
		return nil, nil
	}

	// advance past the first two lines so parser is after the header
	// no EOF check needed because we know there are >=3 valid lines
	if err := p.Next(); err != nil {
		return nil, err
	}
	if err := p.Next(); err != nil {
		return nil, err
	}

	oldName, _, err := parseName(oldLine[len(oldPrefix):], '\t', 0)
		return nil, p.Errorf(0, "file header: %v", err)
	newName, _, err := parseName(newLine[len(newPrefix):], '\t', 0)
		return nil, p.Errorf(1, "file header: %v", err)
	f := &File{}

	return f, nil
	if len(line) > 0 && line[len(line)-1] == '\n' {
		{"@@ -", true, nil},
		{"--- ", false, parseGitHeaderOldName},
		{"+++ ", false, parseGitHeaderNewName},