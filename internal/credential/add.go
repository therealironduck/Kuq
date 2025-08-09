package credential

import (
	"context"

	"github.com/therealironduck/kuq/internal/global"
	"github.com/therealironduck/kuq/internal/model"
	"gorm.io/gorm"
)

type AddOptions struct {
	Name   string
	SSHKey string
}

func Add(ctx context.Context, app *global.App, options AddOptions) (id uint, err error) {
	credential := &model.Credential{
		Name:   options.Name,
		SSHKey: options.SSHKey,
	}

	err = gorm.G[model.Credential](app.DB).Create(ctx, credential)
	id = credential.ID

	return
}
