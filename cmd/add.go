package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type problem struct {
	name       string
	difficulty string
	topic      string
	link       string
	path       string
}

func (p *problem) setField(field *string, prompt string) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		*field = sanitize(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (p *problem) setName()  { p.setField(&p.name, "Name: ") }
func (p *problem) setTopic() { p.setField(&p.topic, "Topic: ") }
func (p *problem) setDiff()  { p.setField(&p.difficulty, "Difficulty: ") }
func (p *problem) setLink()  { p.setField(&p.link, "Link: ") }
func (p *problem) setPath() {
	p.path = filepath.Join("./solutions/", snakeCase(p.name)+".go")
}

func (p *problem) makeFile() error {
	if _, err := os.Stat(p.path); !os.IsNotExist(err) {
		return fmt.Errorf("file already exists: %s", p.path)
	}
	f, err := os.Create(p.path)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func (p *problem) addToTable() error {
	f, err := os.OpenFile("README.md", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lineNum := -2 // ignore first two lines
	for scanner.Scan() {
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	row := fmt.Sprintf(
		"|%03d|[%s](%s)|%s|`%s`|[link](%s)\n", lineNum, p.name, p.path, p.difficulty, p.topic, p.link,
	)
	_, err = f.WriteString(row)
	if err != nil {
		return err
	}
	return nil
}

func sanitize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func snakeCase(str string) string {
	str = strings.Replace(str, " ", "_", -1)
	return str
}

func getInput(p *problem) {
	p.setName()
	p.setTopic()
	p.setDiff()
	p.setLink()
	p.setPath()
}

func main() {
	var p problem
	getInput(&p)
	if err := p.makeFile(); err != nil {
		log.Fatal(err)
	}
	if err := p.addToTable(); err != nil {
		log.Fatal(err)
	}
}
