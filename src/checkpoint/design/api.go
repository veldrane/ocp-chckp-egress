package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("checkpoint", func() {
	Title("Checkpoint Openshfit API")
	Description("Backend presenting rules for Checkpoint based on ocp definition")
	Server("checkpoint", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

var CheckpointRuleSet = Type("CheckpointRuleSet", func() {
	Attribute("version", String, "Version of the object", func() {
		MaxLength(8)
		Default("1.0")
	})
	Attribute("description", String, "Desription of the object", func() {
		MaxLength(100)
		Default("")
	})
})

var CheckpointRule = Type("CheckpointRule", func() {
	Attribute("name", String, "Name of the object", func() {
		MaxLength(100)
		Default("generic-object")
	})
	Attribute("id", String, "Uuid", func() {
		MaxLength(100)
		Default("")
	})
	Attribute("description", String, "Description", func() {
		MaxLength(100)
		Default("")
	})
})

var StoredCheckpointRule = ResultType("application/vnd.goa.rule", func() {
	Description("A StoredCheckpointRules describes a one rules")
	TypeName("StoredCheckpointRule")
	Reference(CheckpointRule)

	Attribute("name", String, "Return the name of the record")
	Attribute("id", String, "uuid of the object")
	Attribute("description", String, "description of the object")
	Attribute("ranges", ArrayOf(String), "Ip adresss")
	Extend(CheckpointRule)

	View("default", func() {
		Attribute("name")
		Attribute("id")
		Attribute("description")
		Attribute("ranges")
	})
})

var StoredCheckpointRuleSet = ResultType("application/vnd.goa.ruleset", func() {
	Description("A StoredCheckpointRule describes a one rules")
	TypeName("StoredCheckpointRuleSet")
	Reference(CheckpointRuleSet)

	Attribute("version", String, "version of the ruleset")
	Attribute("description", String, "description of the object")
	Attribute("objects", ArrayOf(StoredCheckpointRule), "Rulesets")

	View("default", func() {
		Attribute("version")
		Attribute("description")
		Attribute("objects")
	})
})
