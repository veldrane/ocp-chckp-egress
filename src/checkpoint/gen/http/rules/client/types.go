// Code generated by goa v3.7.2, DO NOT EDIT.
//
// rules HTTP client types
//
// Command:
// $ goa gen checkpoint/design

package client

import (
	rulesviews "checkpoint/gen/rules/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// EgressListResponseBody is the type of the "rules" service "egressList"
// endpoint HTTP response body.
type EgressListResponseBody struct {
	// version of the ruleset
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// description of the object
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rulesets
	Objects []*StoredCheckpointRuleResponseBody `form:"objects,omitempty" json:"objects,omitempty" xml:"objects,omitempty"`
}

// IngressListResponseBody is the type of the "rules" service "ingressList"
// endpoint HTTP response body.
type IngressListResponseBody struct {
	// version of the ruleset
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// description of the object
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rulesets
	Objects []*StoredCheckpointRuleResponseBody `form:"objects,omitempty" json:"objects,omitempty" xml:"objects,omitempty"`
}

// StoredCheckpointRuleResponseBody is used to define fields on response body
// types.
type StoredCheckpointRuleResponseBody struct {
	// Return the name of the record
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// uuid of the object
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// description of the object
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Ip adresss
	Ranges []string `form:"ranges,omitempty" json:"ranges,omitempty" xml:"ranges,omitempty"`
}

// NewEgressListStoredCheckpointRuleSetOK builds a "rules" service "egressList"
// endpoint result from a HTTP "OK" response.
func NewEgressListStoredCheckpointRuleSetOK(body *EgressListResponseBody) *rulesviews.StoredCheckpointRuleSetView {
	v := &rulesviews.StoredCheckpointRuleSetView{
		Version:     body.Version,
		Description: body.Description,
	}
	if body.Objects != nil {
		v.Objects = make([]*rulesviews.StoredCheckpointRuleView, len(body.Objects))
		for i, val := range body.Objects {
			v.Objects[i] = unmarshalStoredCheckpointRuleResponseBodyToRulesviewsStoredCheckpointRuleView(val)
		}
	}

	return v
}

// NewIngressListStoredCheckpointRuleSetOK builds a "rules" service
// "ingressList" endpoint result from a HTTP "OK" response.
func NewIngressListStoredCheckpointRuleSetOK(body *IngressListResponseBody) *rulesviews.StoredCheckpointRuleSetView {
	v := &rulesviews.StoredCheckpointRuleSetView{
		Version:     body.Version,
		Description: body.Description,
	}
	if body.Objects != nil {
		v.Objects = make([]*rulesviews.StoredCheckpointRuleView, len(body.Objects))
		for i, val := range body.Objects {
			v.Objects[i] = unmarshalStoredCheckpointRuleResponseBodyToRulesviewsStoredCheckpointRuleView(val)
		}
	}

	return v
}

// ValidateStoredCheckpointRuleResponseBody runs the validations defined on
// StoredCheckpointRuleResponseBody
func ValidateStoredCheckpointRuleResponseBody(body *StoredCheckpointRuleResponseBody) (err error) {
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.ID != nil {
		if utf8.RuneCountInString(*body.ID) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.id", *body.ID, utf8.RuneCountInString(*body.ID), 100, false))
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 100, false))
		}
	}
	return
}