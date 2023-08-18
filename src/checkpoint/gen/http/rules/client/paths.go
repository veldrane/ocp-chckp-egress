// Code generated by goa v3.7.2, DO NOT EDIT.
//
// HTTP request path constructors for the rules service.
//
// Command:
// $ goa gen checkpoint/design

package client

// EgressListRulesPath returns the URL path to the rules service egressList HTTP endpoint.
func EgressListRulesPath() string {
	return "/v1/egress"
}

// IngressListRulesPath returns the URL path to the rules service ingressList HTTP endpoint.
func IngressListRulesPath() string {
	return "/v1/ingress"
}