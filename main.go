package AOCUtils

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Session struct {
	sessionId string
}

type PartFunction func(input string) string

func NewSession(sessionId string) *Session {
	return &Session{sessionId: sessionId}
}

func (s *Session) NewDay(year int, day int) *Day {
	return &Day{
		Day:     day,
		Year:    year,
		Session: s,
	}
}

func (d Day) getInput() (input string, err error) {
	path := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", d.Year, d.Day)
	header := http.Header{}
	header.Add("Cookie", fmt.Sprintf("session=%s", d.Session.sessionId))
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return "", err
	}
	request.Header = header
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	if bytes.HasPrefix(body, []byte("Please don't repeatedly request this endpoint before it unlocks!")) {
		time.Sleep(1 * time.Second)
		return d.getInput()
	}
	return string(body), nil
}

func (d Day) saveInput() (err error) {
	input, err := d.getInput()
	if err != nil {
		return
	}
	path := fmt.Sprintf("./.aoc/%d/%d", d.Year, d.Day)
	err = os.MkdirAll(path, 0755)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	path = filepath.Join(path, "input.txt")
	err = os.WriteFile(path, []byte(input), 0644)
	if err != nil {
		return err
	}
	return nil
}

type Day struct {
	Year    int
	Day     int
	Session *Session
}

func (d Day) Run(part1 PartFunction, part2 PartFunction) (string, string, error) {
	path := fmt.Sprintf("./.aoc/%d/%d/input.txt", d.Year, d.Day)
	var file []byte
	var ok error
	if file, ok = os.ReadFile(path); os.IsNotExist(ok) {
		err := d.saveInput()
		if err != nil {
			return "", "", err
		}
		file, err = os.ReadFile(path)
		if err != nil {
			return "", "", err
		}
	}
	input := string(file)
	part1Result := part1(input)
	part2Result := part2(input)
	return part1Result, part2Result, nil
}

func (d Day) SubmitPart(part string, output string) string {
	apiUrl := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", d.Year, d.Day)

	data := url.Values{}
	data.Set("level", part)
	data.Set("answer", output)

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Cookie", fmt.Sprintf("session=%s", d.Session.sessionId))

	resp, _ := client.Do(r)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Join(strings.Split(doc.Find("article").Text(), "\t\n "), " ")
}
