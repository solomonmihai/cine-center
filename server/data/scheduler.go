package data

import (
	"fmt"
	"os/exec"

	"github.com/go-co-op/gocron/v2"
)

func CreateScheduler() error {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return err
	}

	_, err = scheduler.NewJob(
		gocron.DailyJob(0, gocron.NewAtTimes(gocron.NewAtTime(0, 0, 0))),
		gocron.NewTask(
			UpdateData,
		),
	)

	if err != nil {
		return err
	}

	fmt.Println("started scheduler")
	scheduler.Start()

	return nil
}

func UpdateData() error {
	fmt.Println("fetching data")
	if err := runScraper(); err != nil {
		fmt.Printf("failed to run scraper %s", err)
		return err
	}

	fmt.Println("loading data")
	if err := loadData(); err != nil {
		fmt.Printf("failed to load data %s", err)
		return err
	}

	return nil
}

func runScraper() error {
	cmd := exec.Command("python3", "/app/scraper/main.py")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}

	return nil
}
