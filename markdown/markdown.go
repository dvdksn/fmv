package markdown

import (
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type Parser struct {
	Handler parser.Parser
}

func New() *Parser {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.New(
				meta.WithStoresInDocument(),
			),
		),
	)

	return &Parser{Handler: markdown.Parser()}
}

func (p *Parser) GetMetadata(filename string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	reader := text.NewReader(data)
	document := p.Handler.Parse(reader)
	if err != nil {
		return nil, err
	}

	return document.OwnerDocument().Meta(), nil
}
