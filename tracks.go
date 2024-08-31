package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)


type Track struct {
	Title string
	Performer string
	Time string
}

func (t *Track) DisplayName() string {
	if t.Performer != "" {
		return fmt.Sprintf("%s %s - %s", t.Time, t.Performer, t.Title)
	} else {
		return fmt.Sprintf("%s %s", t.Time, t.Title)
	}
}


func FilterTitles(tracks []Track, ignore []string) []Track {
	for i := range(tracks) {
		for _, toIgnore := range(ignore) {
			tracks[i].Title = strings.Replace(tracks[i].Title, toIgnore, "", -1)
		}
	}
	return tracks
}

func ParseCueFile(filename string) ([]Track, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

	tracks := make([]Track, 0)
	pos := -1
	inHeader := true

    titlePattern := regexp.MustCompile(`TITLE "(.*)"`)
    performerPattern := regexp.MustCompile(`PERFORMER "(.*)"`)
    timePattern := regexp.MustCompile(`INDEX 01 (\d{2}:\d{2}:\d{2})`)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())

		if strings.Contains(line, "TRACK") {
			tracks = append(tracks, Track{})
			pos = pos + 1
			inHeader = false
		}

		if inHeader {
			continue
		}

		if match:= matchGroupOrEmpty(line, titlePattern); match != "" {
			tracks[pos].Title = match
		}

		if match:= matchGroupOrEmpty(line, performerPattern); match != "" {
			tracks[pos].Performer = match
		}

		if match:= matchGroupOrEmpty(line, timePattern); match != "" {
			tracks[pos].Time = match
		}
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return tracks, nil
}

func matchGroupOrEmpty(line string, pattern *regexp.Regexp) string {
	match := pattern.FindStringSubmatch(line)
	if match != nil {
		// [] is group
		return match[1]
	} else {
		return ""
	}
}
