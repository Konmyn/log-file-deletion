package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	var filePath string
	var workHour int
	var preserveHour int
	var dryRun bool
	var napTime int

	flag.StringVar(&filePath, "path", "/logs", "file path to scan")
	flag.IntVar(&workHour, "work-hour", 10, "work hour in the day")
	flag.IntVar(&preserveHour, "preserve-hour", 72, "how long will the log file to preserve")
	flag.IntVar(&napTime, "nap-time", 30, "execute period for this program")
	flag.BoolVar(&dryRun, "dry-run", false, "do not remove the fie in real")
	flag.Parse()

	threshold := (time.Hour * time.Duration(preserveHour)).Seconds()
	napDuration := time.Minute * time.Duration(napTime)

	for {
		currentTime := time.Now()
		if currentTime.Hour() != workHour || currentTime.Minute() >= int(napDuration.Minutes()) {
			log.Printf("current time(%d) is not in working hour(%d), continue sleeping", currentTime.Hour(), workHour)
		} else {
			var totalSize int64 = 0
			err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				if !strings.Contains(info.Name(), ".log") {
					return nil
				}
				if diff := currentTime.Sub(info.ModTime()).Seconds(); diff <= threshold {
					return nil
				}
				if dryRun {
					log.Printf("[%s] last modify time: <%s>, size in bytes: <%d> --dry-run", path, info.ModTime(), info.Size())
				} else {
					log.Printf("[%s] last modify time: <%s>, size in bytes: <%d>", path, info.ModTime(), info.Size())
					if pathErr := os.Remove(path); pathErr != nil {
						log.Printf("error deleting file: %s ---> %v", path, pathErr)
					}
				}
				totalSize += info.Size()
				return nil
			})
			if err != nil {
				log.Printf("error walking the path %v\n", err)
			}
			log.Println("deleted file total size in bytes: ", totalSize)
		}
		log.Printf("start sleeping %f minutes", napDuration.Minutes())
		time.Sleep(napDuration)
	}
}
