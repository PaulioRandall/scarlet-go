package docs

import (
	"fmt"
	"strings"
)

//   docs [search term]		Show documentation.

func Docs(searchTerm string) (int, error) {

	if strings.HasPrefix(searchTerm, "@") {
		return searchSpellDocs(searchTerm)
	}

	switch strings.ToLower(searchTerm) {
	case "":
		printOverview()

	case "comment", "comments":
		return 0, fmt.Errorf("%q documentation is not yet supported", searchTerm)

	case "variable", "variables":
		printVariablesOverview()

	case "type", "types":
		printTypesOverview()

	case "spell":
		printSpellOverview()

	case "spells":
		return 0, fmt.Errorf("%q documentation is not yet supported", searchTerm)

	case "-":
		return 0, fmt.Errorf("%q documentation is not yet supported", searchTerm)

	default:
		return 1, fmt.Errorf("Unexpected documentation argument %q", searchTerm)
	}

	return 0, nil
}

func searchSpellDocs(searchTerm string) (int, error) {
	return 0, fmt.Errorf("spell documentation is not yet supported")
}
