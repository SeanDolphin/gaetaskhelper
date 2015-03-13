package gaetaskhelper

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/module"
	"google.golang.org/appengine/taskqueue"
)

var defaultLevels map[string]string = map[string]string{}

func SetModule(ctx context.Context, task *taskqueue.Task, moduleName string) error {
	hostName, ok := defaultLevels[moduleName]
	if !ok {
		version, err := module.DefaultVersion(ctx, moduleName)
		if err != nil {
			return err
		}

		hostName, err = appengine.ModuleHostname(ctx, moduleName, version, "")
		if err != nil {
			return err
		}
		defaultLevels[moduleName] = hostName
	}

	task.Header.Set("Host", hostName)
	return nil
}
