package log

import (
	"base/pkg/app"
	"base/pkg/config"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/util/rogger"
)

var (
	SLOG *rogger.Logger
	conf logConfig
)

type logConfig struct {
	Remote bool `yaml:"remote"`
	Hour   int `yaml:"hour"`
}

// Start 启动日志服务
func Start() {
	err := config.Config().Bind(app.CfgName, "log", &conf, func() {
		if config.Config().Bind(app.CfgName, "log", &conf, nil) == nil {
			initStart()
		}
	})
	if err != nil {
		return
	}
	initStart()
	//设置日志 等级
	rogger.SetLevel(rogger.StringToLevel(app.Cfg.LogLevel))
	rogger.Colored()
}

func initStart() {
	if conf.Remote {
		SLOG = tars.GetRemoteLogger("TLOG")
		SLOG.Info("启动远程日志...")
	} else {
		SLOG = tars.GetLogger("TLOG")
		SLOG.Info("启动本地日志...")
	}
	if conf.Hour == 0 {
		// SLOG = tars.GetLogger("TLOG")
		_ = SLOG.SetFileRoller(app.Cfg.LogPath, 1, int(app.Cfg.LogSize))
	} else if conf.Hour%24 == 0 {
		// SLOG = tars.GetDayLogger("TLOG", logConfig.Hour/24)
		_ = SLOG.SetDayRoller(app.Cfg.LogPath, conf.Hour/24)
	} else if conf.Hour > 0 {
		// SLOG = tars.GetHourLogger("TLOG", logConfig.Hour)
		_ = SLOG.SetHourRoller(app.Cfg.LogPath, conf.Hour)
	}
	SLOG.SetConsole()
	app.Register("logger", SLOG)
}