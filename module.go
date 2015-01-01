package gaetaskhelper

import (
	"appengine"
	"appengine/module"
	"appengine/taskqueue"
)

var defaultLevels map[string]string = map[string]string{}

func SetModule(context appengine.Context, task *taskqueue.Task, moduleName string) error {
	hostName, ok := defaultLevels[moduleName]
	if !ok {
		version, err := module.DefaultVersion(context, moduleName)
		if err != nil {
			return err
		}

		hostName, err = appengine.ModuleHostname(context, moduleName, version, "")
		if err != nil {
			return err
		}
		defaultLevels[moduleName] = hostName
	}

	task.Header.Set("Host", hostName)
	return nil
}
