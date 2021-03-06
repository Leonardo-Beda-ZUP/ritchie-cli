package envcredential

import (
	"strings"

	"github.com/ZupIT/ritchie-cli/pkg/credential"
)

type CredentialResolver struct {
	credential.Finder
}

// NewResolver creates a credential resolver instance of Resolver interface
func NewResolver(cf credential.Finder) CredentialResolver {
	return CredentialResolver{cf}
}

func (c CredentialResolver) Resolve(name string) (string, error) {
	s := strings.Split(name, "_")
	service := strings.ToLower(s[1])
	cred, err := c.Find(service)
	if err != nil {
		return "", err
	}

	k := strings.ToLower(s[2])
	return cred.Credential[k], nil
}
