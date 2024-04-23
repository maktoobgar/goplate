package g

import (
	_ "embed"
	"mime/multipart"

	"service/config"

	db "service/pkg/database"
	"service/pkg/logging"
	media_manager "service/pkg/media"

	"github.com/kataras/iris/v12"
	"github.com/robfig/cron/v3"
)

type (
	FileType struct {
		Info      *multipart.FileHeader
		File      multipart.File
		Extension string
	}

	CronJob struct {
		Job func()
	}
)

func (c *CronJob) Run() {
	c.Job()
}

//go:embed version
var Version string

//go:embed name
var Name string

const (
	// Header
	AccessToken  = "Authorization"
	RefreshToken = "Authorization"
	TempToken    = "Authorization"
	TokenId      = "TokenId"

	// Url
	TranslateKey = "translate"

	// Context
	WriterLock   = "WriterLock"
	ClosedWriter = "ClosedWriter"

	RequestBody = "RequestBody"
	DbInstance  = "DbInstance"
	UserKey     = "User"
	JobKey      = "Job"

	// Maximum File sizes (in KB)

	MaximumProfilePicSize  = int64(500 * 1024)  // 500 KB
	MaximumCertificateSize = int64(5000 * 1024) // 5000 KB(5 MB)

	// Regex
	UuidRegex string = `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`
)

// Config
var CFG *config.Config = nil

// SecretKey in bytes
var SecretKeyBytes []byte

// Utilities
var Logger logging.Logger = nil

// App
var App *iris.Application = nil

// Default DB
var DB db.RelationalDatabaseFunction = nil

// Media manager for all medias
var Media media_manager.MediaManager = nil

var UsersMedia media_manager.MediaManager = nil

// Cron of the project
var Cron *cron.Cron = nil

var MediaServeAddress = "/media"
