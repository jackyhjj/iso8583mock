package infrastructures

import (
	"path"
	"runtime"
	"os"
	"io/ioutil"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
)

// SetConfig location
func SetConfig(p string) {
	config.SetConfigName("App")
	config.SetConfigType("yaml")
	config.AddConfigPath(p)
	config.AddConfigPath("./configurations")
	config.AddConfigPath(GetDefaultConfigPath())

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("config error: ", err)
	}

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		log.Warn("Config file changed:", e.Name)
	})

	//env := config.GetString("app.env")

	log.AddHook(NewLogHook().
		SetFormatType(config.GetString("log.format_output")).
		SetLogLevel(config.GetInt("log.level")).
		SetRotateLog(config.GetString("log.rotate")))

	isDebug := config.GetBool("app.debug")

	switch isDebug {
	case true:
		//log.SetLevel(log.DebugLevel)
		log.SetOutput(os.Stdout)
	default:
		//log.SetLevel(log.WarnLevel)
		log.SetOutput(ioutil.Discard)
	}

	log.SetFormatter(&log.TextFormatter{})
}

// GetDefaultConfigPath location
func GetDefaultConfigPath() string {
	_, filename, _, ok := runtime.Caller(0)

	if ok == false {
		log.Fatal("err")
	}

	filePath := path.Join(path.Dir(filename), "../configurations/")

	return filePath
}
