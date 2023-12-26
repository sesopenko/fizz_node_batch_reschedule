package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	skippedFrames := 0
	if len(args) > 0 {
		parsedSkip, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Unable to parse argument.  must be numeric, got: %s", args[0])
		}
		skippedFrames = parsedSkip
	}
	fmt.Printf("Skipping %d frames\n", skippedFrames)
	schedules := []string{
		"a.txt",
		"b.txt",
		"c.txt",
		"d.txt",
		"main_sched.txt",
	}
	for _, filePath := range schedules {
		scheduleBatch(skippedFrames, filePath)
		fmt.Printf("Wrote %s\n", filePath)
	}

}

func scheduleBatch(skippedFrames int, filePath string) {
	newLine := "\r\n"
	if runtime.GOOS != "windows" {
		newLine = "\n"
	}

	mainRead, _ := os.ReadFile(filePath)
	mainSched := string(mainRead)
	mainLines := splitLines(mainSched)
	if len(mainLines) == 1 {
		writeJoinedLines(filePath, mainLines[0])
		return
	}
	const regexPattern = `^"?(\d+)"?\s*:\s*(.*)$`
	re := regexp.MustCompile(regexPattern)
	finalLines := []string{}
	firstLine := ""
	for _, line := range mainLines {
		matches := re.FindStringSubmatch(line)
		if len(matches) == 0 {
			log.Fatalf("unable to match lien: %s", line)
		}
		keyFrame, _ := strconv.Atoi(matches[1])
		adjustedKeyframe := keyFrame - skippedFrames

		if keyFrame < skippedFrames {
			if len(finalLines) == 0 {
				firstLine = fmt.Sprintf(`"%d":%s`, 0, matches[2])
			}
			continue
		} else {
			if firstLine != "" {
				finalLines = append(finalLines, firstLine)
				firstLine = ""
			}
		}
		finalLines = append(finalLines, fmt.Sprintf(`"%d":%s`, adjustedKeyframe, matches[2]))
	}
	if firstLine != "" && len(finalLines) == 0 {
		finalLines = append(finalLines, firstLine)
	}
	joined := strings.Join(finalLines, newLine)

	writeJoinedLines(filePath, joined)
}

func writeJoinedLines(filePath string, joined string) {
	outPath := fmt.Sprintf("out_%s", filePath)
	file, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(joined)
	if err != nil {
		log.Fatalf("Error writing file: %s", err)
	}
}

func splitLines(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == '\n' || r == '\r'
	})
}
