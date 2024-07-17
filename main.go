package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/robfig/cron"
)

// 0 * * * * : Every 00 minutes
// 0 0 * * * :Every day at midnight i.e. 00:00:00 AM
// 0 0 1 * * : Every month on the 1st at midnight
// 0 0 1 1 * : Every year on January 1st at midnight 00:00:00 AM
// 0 12 1 * * : At noon on the 1st of every month
// 0 12 * * * : Every day at noon
// 15 10 * * * : Every day at 10:15 AM
// 0 10 1 * * : At 10 AM on the 1st of every month​
// 0 10 1 1 * : At 10 AM on January 1st every year
// 0 0 * * 1-5 : Every weekday at midnight (Monday through Friday) i.e. 00:00:00 AM
// 0 0 1,15 * * : At midnight on the 1st day and 15th day of every month
// */15 * * * * : Every 15th minute (0, 15, 30, 45)
// 0 0 1-7 * * : At midnight on the 1st through 7th of every month
// 0 12 1 1,4,7,10 * : At noon on the 1st day of every 3rd month (January, April, July, October)
// 0 12 * * 2,4,6 : At noon on every Tuesday, Thursday, and Saturday
// 0 12,13,14 * * * : Every day at noon, 1 PM, and 2 PM
// */30 8-18 * * * : Every 30 minutes between 8 AM and 6 PM
// @yearly (or @annually) : Every year on January 1st at midnight 00:00:00 AM
// @monthly : Every month on the 1st at midnight​
// @weekly : Every week at 00:00:00 AM on Sunday
// @daily : Every day at 00:00:00 AM
// @hourly : Every hour at 00 minute i.e. 17:00:00 PM , 18:00:00 PM
// @every 1h30m10s : Every 1 hour 30 minutes and 10 seconds

// sample
func sample() {
	fmt.Println("cron job")

	//creates a cron job instance
	c := cron.New()

	//add the cron schedule and task to cron job
	c.AddFunc("@every 00h00m05s", greetings)

	//start the cron job scheduler
	c.Start()

	//keep program running
	select {}
}

func greetings() {
	fmt.Println("Message after 1 seconds")
}

func main() {
	//Initialize the cron scheduler
	c := InitCronScheduler()

	//stop
	defer c.Stop()

	//start server
	StartServer()
}

func InitCronScheduler() *cron.Cron {
	//Create a new cron instance
	c := cron.New()

	//Add a cron job that runs every 10 seconds
	c.AddFunc("@every 00h00m10s", apiCall)

	//start the cron scheduler
	c.Start()

	fmt.Println("cron scheduler started in loc", c.Location())
	return c
}

func apiCall() {
	resp, err := http.Get("http://localhost:8000")
	if err != nil {
		fmt.Println("Error while calling the API:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	fmt.Println(string(body))
}

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "calling the API after every 10 seconds at %s", time.Now().Format("2006-01-02 15:04:05"))
	})

	fmt.Println("server starting on 8000:")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("error in starting server", err)
	}
}
