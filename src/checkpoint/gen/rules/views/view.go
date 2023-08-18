// Code generated by goa v3.7.2, DO NOT EDIT.
//
// rules views
//
// Command:
// $ goa gen checkpoint/design

package views

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// StoredCheckpointRuleSet is the viewed result type that is projected based on
// a view.
type StoredCheckpointRuleSet struct {
	// Type to project
	Projected *StoredCheckpointRuleSetView
	// View to render
	View string
}

// StoredCheckpointRuleSetView is a type that runs validations on a projected
// type.
type StoredCheckpointRuleSetView struct {
	// version of the ruleset
	Version *string
	// description of the object
	Description *string
	// Rulesets
	Objects []*StoredCheckpointRuleView
}

// StoredCheckpointRuleView is a type that runs validations on a projected type.
type StoredCheckpointRuleView struct {
	// Return the name of the record
	Name *string
	// uuid of the object
	ID *string
	// description of the object
	Description *string
	// Ip adresss
	Ranges []string
}

var (
	// StoredCheckpointRuleSetMap is a map indexing the attribute names of
	// StoredCheckpointRuleSet by view name.
	StoredCheckpointRuleSetMap = map[string][]string{
		"default": {
			"version",
			"description",
			"objects",
		},
	}
	// StoredCheckpointRuleMap is a map indexing the attribute names of
	// StoredCheckpointRule by view name.
	StoredCheckpointRuleMap = map[string][]string{
		"default": {
			"name",
			"id",
			"description",
			"ranges",
		},
	}
)

// ValidateStoredCheckpointRuleSet runs the validations defined on the viewed
// result type StoredCheckpointRuleSet.
func ValidateStoredCheckpointRuleSet(result *StoredCheckpointRuleSet) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredCheckpointRuleSetView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateStoredCheckpointRuleSetView runs the validations defined on
// StoredCheckpointRuleSetView using the "default" view.
func ValidateStoredCheckpointRuleSetView(result *StoredCheckpointRuleSetView) (err error) {
	if result.Version != nil {
		if utf8.RuneCountInString(*result.Version) > 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.version", *result.Version, utf8.RuneCountInString(*result.Version), 8, false))
		}
	}
	if result.Description != nil {
		if utf8.RuneCountInString(*result.Description) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.description", *result.Description, utf8.RuneCountInString(*result.Description), 100, false))
		}
	}
	for _, e := range result.Objects {
		if e != nil {
			if err2 := ValidateStoredCheckpointRuleView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStoredCheckpointRuleView runs the validations defined on
// StoredCheckpointRuleView using the "default" view.
func ValidateStoredCheckpointRuleView(result *StoredCheckpointRuleView) (err error) {
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 100, false))
		}
	}
	if result.ID != nil {
		if utf8.RuneCountInString(*result.ID) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.id", *result.ID, utf8.RuneCountInString(*result.ID), 100, false))
		}
	}
	if result.Description != nil {
		if utf8.RuneCountInString(*result.Description) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.description", *result.Description, utf8.RuneCountInString(*result.Description), 100, false))
		}
	}
	return
}