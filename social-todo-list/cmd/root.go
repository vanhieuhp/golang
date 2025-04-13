package cmd

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/200Lab-Education/go-sdk/plugin/storage/sdkgorm"
	"social-todo-list/common"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("social-todo-list"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.PluginDBMain)),
	)

	return service
}
