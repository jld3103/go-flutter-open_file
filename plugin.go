package open_file

import (
	"github.com/pkg/errors"
	"github.com/skratchdot/open-golang/open"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "open_file"

// OpenFilePlugin implements flutter.Plugin and handles method.
type OpenFilePlugin struct{}

var _ flutter.Plugin = &OpenFilePlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *OpenFilePlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("open_file", p.openFile)
	return nil
}

func (p *OpenFilePlugin) openFile(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})

	filePath := argsMap["file_path"].(string)
	if filePath == "" {
		return nil, errors.New("file_path is empty")
	}

	open.Start(filePath)
	return nil, nil
}
