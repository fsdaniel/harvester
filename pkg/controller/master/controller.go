package master

import (
	"context"

	"github.com/rancher/steve/pkg/server"

	"github.com/rancher/harvester/pkg/config"
	"github.com/rancher/harvester/pkg/controller/master/image"
	"github.com/rancher/harvester/pkg/controller/master/keypair"
	"github.com/rancher/harvester/pkg/controller/master/template"
	"github.com/rancher/harvester/pkg/controller/master/virtualmachine"
)

type registerFunc func(context.Context, *config.Management) error

var registerFuncs = []registerFunc{
	image.Register,
	keypair.Register,
	template.Register,
	virtualmachine.Register,
}

func register(ctx context.Context, server *server.Server, management *config.Management) error {
	for _, f := range registerFuncs {
		if err := f(ctx, management); err != nil {
			return err
		}
	}
	return nil
}
