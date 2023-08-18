package checkpoint

import (
	"checkpoint/gen/rules"
	"context"
	"log"

	checkpoint "bitbucket.org/veldrane/golibs/checkpoint"
	"github.com/ryanfowler/uuid"
)

type rulessrvc struct {
	logger *log.Logger
	rules  *checkpoint.RulesT
}

// NewRules returns the rules service implementation.
func NewRules(logger *log.Logger, rules *checkpoint.RulesT) rules.Service {
	return &rulessrvc{logger, rules}
}

// List all stored rules
func (s *rulessrvc) EgressList(ctx context.Context) (res *rules.StoredCheckpointRuleSet, err error) {

	var object rules.StoredCheckpointRule

	ruleset := rules.StoredCheckpointRuleSet{
		Version:     "1.0",
		Description: "Checkpoint rules for egress objects",
	}

	uuidns, _ := uuid.Parse([]byte("93f3dd48-96ba-43a9-84a0-26467336d731"))
	objects := make([]rules.StoredCheckpointRule, 0)

	i := 0
	s.rules.Locks.Egress.RLock()
	for k, v := range s.rules.Egress {

		object.Name = k
		object.Description = "Ocp custom rule " + k
		object.ID = uuid.NewV5(uuidns, []byte(k)).String()
		object.Ranges = v
		objects = append(objects, object)
		ruleset.Objects = append(ruleset.Objects, &objects[i])
		i++
	}
	s.rules.Locks.Egress.RUnlock()

	s.logger.Print("egress rules.list has been hit")
	return &ruleset, nil
}

func (s *rulessrvc) IngressList(ctx context.Context) (res *rules.StoredCheckpointRuleSet, err error) {

	var object rules.StoredCheckpointRule

	ruleset := rules.StoredCheckpointRuleSet{
		Version:     "1.0",
		Description: "Checkpoint rules for ingress objects",
	}

	uuidns, _ := uuid.Parse([]byte("93f3dd48-96ba-43a9-84a0-26467336d731"))
	objects := make([]rules.StoredCheckpointRule, 0)

	i := 0
	s.rules.Locks.Ingress.RLock()
	for k, v := range s.rules.Ingress {

		object.Name = k
		object.Description = "Ocp custom rule " + k
		object.ID = uuid.NewV5(uuidns, []byte(k)).String()
		object.Ranges = v
		objects = append(objects, object)
		ruleset.Objects = append(ruleset.Objects, &objects[i])
		i++
	}
	s.rules.Locks.Ingress.RUnlock()

	s.logger.Print("ingress rules.ingressList has been hit")
	return &ruleset, nil
}
