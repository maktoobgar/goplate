package app

import (
	"fmt"

	g "service/global"
	"service/pkg/colors"
)

func runCronJobs() {
	// "* * * * *" Means every minute. Reference:
	// https://crontab.guru/every-minute
	//
	// Just implemented for test purposes but useful
	// g.Cron.AddJob("* * * * *", extra_middlewares.IpRateLimitGarbageCollector)
}

func Info() {
	fmt.Println(colors.Cyan, fmt.Sprintf("\n==%sSystem Info%s==%s\n", colors.Yellow, colors.Cyan, colors.Reset))
	fmt.Printf("Name:\t\t\t%s%s%s\n", colors.Blue, g.Name, colors.Reset)
	fmt.Printf("Version:\t\t%s%s%s\n", colors.Blue, g.Version, colors.Reset)
	mainOrTest := "test"
	mainOrTestColor := colors.Red + mainOrTest + colors.Reset
	if !g.CFG.Debug {
		mainOrTest = "main"
		mainOrTestColor = colors.Green + mainOrTest + colors.Reset
	}

	if g.CFG.Database.Type == "sqlite3" {
		fmt.Printf("Main Database:\t\t%v, %v (%v)\n", g.CFG.Database.Type, g.CFG.Database.DbName, mainOrTestColor)
	} else {
		fmt.Printf("Main Database:\t\t%v, %v, %v:%v (%v)\n", g.CFG.Database.Type, g.CFG.Database.DbName, g.CFG.Database.Host, g.CFG.Database.Port, mainOrTestColor)
	}

	if g.CFG.Debug {
		fmt.Printf("Debug:\t\t\t%s%v%s\n", colors.Red, g.CFG.Debug, colors.Reset)
	} else {
		fmt.Printf("Debug:\t\t\t%s%v%s\n", colors.Green, g.CFG.Debug, colors.Reset)
	}
	fmt.Printf("Documentation:\t\t%s%s/swagger/index.html%s\n", colors.Green, g.CFG.Domain, colors.Reset)
	fmt.Printf("Address:\t\thttp://%s:%s\n", g.CFG.IP, g.CFG.Port)
	fmt.Printf("Allowed Origins:\t%v\n", g.CFG.AllowOrigins)
	if g.CFG.AllowHeaders != "" {
		fmt.Printf("Extra Allowed Headers:\t%v\n", g.CFG.AllowHeaders)
	}
	fmt.Print(colors.Cyan, "===============\n\n", colors.Reset)
}
