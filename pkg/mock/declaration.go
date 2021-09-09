package mock

import "github.com/flemay/envvars/pkg/envvars"

type DeclarationReader struct {
	ReadFunc func() (*envvars.Declaration, error)
}

func (r DeclarationReader) Read() (*envvars.Declaration, error) {
	if r.ReadFunc == nil {
		return nil, nil
	}
	return r.ReadFunc()
}
