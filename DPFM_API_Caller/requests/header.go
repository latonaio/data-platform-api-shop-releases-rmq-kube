package requests

type Header struct {
	Shop		int		`json:"Shop"`
	IsReleased	*bool	`json:"IsReleased"`
}
