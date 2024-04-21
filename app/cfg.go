package app

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/robfig/cron/v3"
	migrate "github.com/rubenv/sql-migrate"
	"golang.org/x/text/language"

	"service/build"
	iconfig "service/config"
	g "service/global"
	"service/pkg/colors"
	"service/pkg/config"
	db "service/pkg/database"
	"service/pkg/logging"
	media_manager "service/pkg/media"
	"service/pkg/translator"
)

var (
	cfg       = &iconfig.Config{}
	languages = []language.Tag{language.English, language.Persian}
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
	if err := config.ParseYamlBytes(build.Config, cfg); err != nil {
		log.Fatalln(err)
	}

	if err1, err2 := config.Parse(cfg.PWD+"/env.yaml", cfg, false), config.Parse(cfg.PWD+"/env.yml", cfg, false); err1 != nil || err2 != nil {
		if err1 != nil {
			log.Fatalln(err1)
		} else if err2 != nil {
			log.Fatalln(err2)
		}
	}

	if cfg.ClonesCount < 0 {
		cfg.ClonesCount = runtime.GOMAXPROCS(0)
	}

	g.SecretKeyBytes = []byte(cfg.SecretKey)

	mainOrTest := "test"
	if !cfg.Debug {
		mainOrTest = "main"
	}
	for name, database := range cfg.Gateway.Databases {
		if name == mainOrTest {
			g.MainDatabaseType = database.Type
			break
		}
	}
	g.CFG = cfg
}

// Translator initialization
func initialTranslator() {
	t, err := translator.New(build.Translations, languages[0], languages[1:]...)
	if err != nil {
		log.Fatalln(err)
	}
	g.Translator = t
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
	g.AllSQLCons, g.DB, err = db.New(cfg.Gateway.Databases, cfg.Debug)
	if err != nil {
		log.Fatalln(err)
	}

	var ok bool = false
	if !g.CFG.Debug {
		_, ok = g.AllSQLCons["main"]
		if !ok {
			log.Fatalln(errors.New("'main' db is not defined (required)"))
		}
	} else {
		_, ok = g.AllSQLCons["test"]
		if !ok {
			log.Fatalln(errors.New("'test' db is not defined"))
		}
	}
}

func migrateLatestChanges() {
	db, err := g.DB()
	if err != nil {
		panic(err)
	}
	mainOrTest := "test"
	if !g.CFG.Debug {
		mainOrTest = "main"
	}
	migrations := &migrate.FileMigrationSource{
		Dir: fmt.Sprintf("migrations/%s/", mainOrTest),
	}

	n, err := migrate.Exec(db, g.CFG.Gateway.Databases[mainOrTest].Type, migrations, migrate.Up)
	if err != nil {
		log.Fatalln(err)
	}
	if n > 0 {
		fmt.Printf("\n%s==%sMigrations%s==%s\n\n", colors.Cyan, colors.Green, colors.Cyan, colors.Reset)
		fmt.Printf("Applied %s%d%s migrations!\n", colors.Red, n, colors.Reset)
	}
}

func initialMedia() {
	g.Media = media_manager.NewMediaManager(cfg.Media, true)

	g.UsersMedia, _ = g.Media.GoTo("users", true)
}

func initialCron() {
	g.Cron = cron.New(cron.WithSeconds())
	g.Cron.Start()
}

// Server initialization
func init() {
	setPwd()
	initializeConfigs()
	initialDBs()
	migrateLatestChanges()
	initialTranslator()
	initialLogger()
	initialMedia()
	initialCron()
}
