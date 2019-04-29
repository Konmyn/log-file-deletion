package log_walk

import (
	"flag"
	"fmt"
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

	flag.StringVar(&filePath, "path", "/logs", "file path to scan")
	flag.IntVar(&workHour, "work-hour", 10, "work hour in the day")
	flag.IntVar(&preserveHour, "preserve-hour", 36, "how long will the log file to preserve")
	flag.BoolVar(&dryRun, "dry-run", false, "do not remove the fie in real")
	flag.Parse()

	threshold := (time.Hour * time.Duration(preserveHour)).Seconds()
	snapTime := time.Minute * 30

	for {
		currentTime := time.Now()
		if currentTime.Hour() != workHour && currentTime.Minute() >= int(snapTime.Minutes()) {
			log.Println("not in working hour, continue sleeping")
		} else {
			var totalSize int64 = 0
			err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() || !strings.Contains(info.Name(), ".log") {
					return nil
				}
				if diff := currentTime.Sub(info.ModTime()).Seconds(); diff <= threshold {
					return nil
				}
				if dryRun {
					fmt.Printf("[%s] last modify time: <%s>, size: <%d> --dry-run", path, info.ModTime(), info.Size())
				} else {
					fmt.Printf("[%s] last modify time: <%s>, size: <%d>", path, info.ModTime(), info.Size())
				}
				totalSize += info.Size()
				return nil
			})
			if err != nil {
				fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
			}
			fmt.Println("deleted file total size: ", totalSize)
		}
		time.Sleep(snapTime)
	}
}
