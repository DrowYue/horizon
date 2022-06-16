package manager

import (
	"context"

	"g.hz.netease.com/horizon/pkg/cluster/tekton/metrics"
	"g.hz.netease.com/horizon/pkg/pipelinerun/pipeline/dao"
	"gorm.io/gorm"
)

type Manager interface {
	Create(ctx context.Context, results *metrics.PipelineResults) error
}

type manager struct {
	dao dao.DAO
}

func (m manager) Create(ctx context.Context, results *metrics.PipelineResults) error {
	return m.dao.Create(ctx, results)
}

func New(db *gorm.DB) Manager {
	return &manager{
		dao: dao.NewDAO(db),
	}
}
