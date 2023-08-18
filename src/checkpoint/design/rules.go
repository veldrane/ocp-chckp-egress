package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("rules", func() {
	Description("Service provides management for stored chekcpoint rules")

	Error("NotFound", func() {
		Description("Notfound is the error returned by the service methods when the id of the stuff is not found.")
	})

	Error("MissingItem", func() {
		Description("MissingItem is the error returned by the service methods when some field in POST is missing")
	})

	Error("Timeout", func() {
		Description("Operation timed out")
	})

	Error("InternalError", func() {
		Description("Internal Server Error")
	})

	Method("egressList", func() {

		Description("List all egress rules")
		Result(StoredCheckpointRuleSet, func() {
			View("default")
		})

		HTTP(func() {
			GET("/v1/egress")
			Response(StatusOK)
		})
	})

	Method("ingressList", func() {

		Description("List all ingress rules")
		Result(StoredCheckpointRuleSet, func() {
			View("default")
		})

		HTTP(func() {
			GET("/v1/ingress")
			Response(StatusOK)
		})
	})

})
