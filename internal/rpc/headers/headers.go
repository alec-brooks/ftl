package headers

import (
	"net/http"

	"github.com/alecthomas/errors"
	"github.com/alecthomas/types"

	"github.com/TBD54566975/ftl/schema"
)

// Headers used by the internal RPC system.
const (
	DirectRoutingHeader = "FTL-Direct"
	// VerbHeader is the header used to pass the module.verb of the current request.
	//
	// One header will be present for each hop in the request path.
	VerbHeader = "FTL-Verb"
)

func IsDirectRouted(header http.Header) bool {
	return header.Get(DirectRoutingHeader) != ""
}

func SetDirectRouted(header http.Header) {
	header.Set(DirectRoutingHeader, "1")
}

// GetCallers history from an incoming request.
func GetCallers(header http.Header) ([]*schema.VerbRef, error) {
	headers := header.Values(VerbHeader)
	if len(headers) == 0 {
		return nil, nil
	}
	refs := make([]*schema.VerbRef, len(headers))
	for i, header := range headers {
		ref, err := schema.ParseRef(header)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid %s header %q", VerbHeader, header)
		}
		refs[i] = (*schema.VerbRef)(ref)
	}
	return refs, nil
}

// GetCaller returns the module.verb of the caller, if any.
//
// Will return an error if the header is malformed.
func GetCaller(header http.Header) (types.Option[*schema.VerbRef], error) {
	headers := header.Values(VerbHeader)
	if len(headers) == 0 {
		return types.None[*schema.VerbRef](), nil
	}
	ref, err := schema.ParseRef(headers[len(headers)-1])
	if err != nil {
		return types.None[*schema.VerbRef](), errors.WithStack(err)
	}
	return types.Some((*schema.VerbRef)(ref)), nil
}

// AddCaller to an outgoing request.
func AddCaller(header http.Header, ref *schema.VerbRef) {
	refStr := ref.String()
	if values := header.Values(VerbHeader); len(values) > 0 {
		if values[len(values)-1] == refStr {
			return
		}
	}
	header.Add(VerbHeader, refStr)
}

func SetCallers(header http.Header, refs []*schema.VerbRef) {
	header.Del(VerbHeader)
	for _, ref := range refs {
		AddCaller(header, ref)
	}
}