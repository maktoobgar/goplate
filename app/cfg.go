package app

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/robfig/cron/v3"
	migrate "github.com/rubenv/sql-migrate"

	iconfig "service/config"
	g "service/global"
	"service/pkg/colors"
	"service/pkg/config"
	db "service/pkg/database"
	"service/pkg/logging"
	media_manager "service/pkg/media"
)

var (
	cfg = &iconfig.Config{}
)

// Set Project PWD
func setPwd() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	for parent := pwd; true; parent = filepath.Dir(parent) {
		if _, err := os.Stat(filepath.Join(parent, "go.mod")); err == nil {
			cfg.PWD = parent
			break
		}
	}
	os.Chdir(cfg.PWD)
}

// Initialization for config files in configs folder
func initializeConfigs() {
	// Loads default config, you just have to hard code it
	if err := config.ParseYamlBytes(iconfig.ConfigData, cfg); err != nil {
		log.Fatalln(err)
	}

	if err1, err2 := config.Parse(cfg.PWD+"/env.yaml", cfg, false), config.Parse(cfg.PWD+"/env.yml", cfg, false); err1 != nil || err2 != nil {
		if err1 != nil {
			log.Fatalln(err1)
		} else if err2 != nil {
			log.Fatalln(err2)
		}
	}

	cfg.Language.Path = path.Join(cfg.PWD, cfg.Language.Path)
	g.SecretKeyBytes = []byte(cfg.SecretKey)
	g.CFG = cfg
}

// Logger initialization
func initialLogger() {
	cfg.Logging.Path += "/" + g.Name
	k := cfg.Logging
	opt := logging.Option(k)
	l, err := logging.New(&opt, cfg.Debug)
	if err != nil {
		log.Fatalln(err)
	}
	g.Logger = l
}

// Run dbs
func initialDBs() {
	var err error
	g.DB, err = db.New(cfg.Database, cfg.Debug)
	if err != nil {
		log.Fatalln(err)
	}
}

func MigrateLatestChanges() {
	db, err := g.DB()
	if err != nil {
		panic(err)
	}
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/",
	}

	n, err := migrate.Exec(db, g.CFG.Database.Type, migrations, migrate.Up)
	if err != nil {
		log.Fatalln(err)
	}
	if n > 0 {
		fmt.Printf("\n%s==%sMigrations%s==%s\n\n", colors.Cyan, colors.Green, colors.Cyan, colors.Reset)
		fmt.Printf("Applied %s%d%s migrations!\n", colors.Red, n, colors.Reset)
	}
}

func DemigrateOneChange() {
	db, err := g.DB()
	if err != nil {
		panic(err)
	}
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/",
	}

	n, err := migrate.ExecMax(db, g.CFG.Database.Type, migrations, migrate.Down, 1)
	if err != nil {
		log.Fatalln(err)
	}
	if n > 0 {
		fmt.Printf("\n%s==%sMigrations%s==%s\n\n", colors.Cyan, colors.Green, colors.Cyan, colors.Reset)
		fmt.Printf("Reversed %s%d%s migrations!\n", colors.Red, n, colors.Reset)
	}
}

// func redis() {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})
// }

func initialMedia() {
	g.Media = media_manager.NewMediaManager(cfg.Media, filepath.Join(cfg.PWD, cfg.Media), true)
	g.UsersMedia, _ = g.Media.GoTo("users", true)
}

func initialCron() {
	g.Cron = cron.New(cron.WithSeconds())
	g.Cron.Start()
}

// Gateway Server initialization
func InitializeService() {
	initialDBs()
	initialLogger()
	initialMedia()
	initialCron()
}

// Server initialization
func init() {
	setPwd()
	initializeConfigs()
}
