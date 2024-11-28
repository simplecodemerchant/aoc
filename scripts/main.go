package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"text/template"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	TOKEN := os.Getenv("TOKEN")
	if TOKEN == "" {
		panic("TOKEN is not set")
	}

	app := &cli.App{
		Name: "utils",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "year",
				Aliases:  []string{"y"},
				Required: true,
			},
			&cli.IntFlag{
				Name:     "day",
				Aliases:  []string{"d"},
				Required: true,
			},
			&cli.BoolFlag{
				Name: "debug",
			},
		},
		Action: func(c *cli.Context) error {
			year := c.Int("year")
			day := c.Int("day")

			debug := c.Bool("debug")

			if debug {
				slog.SetLogLoggerLevel(slog.LevelDebug)
				slog.Debug("debug enabled")
			}

			content, err := getInput(year, day, TOKEN)
			if err != nil {
				slog.Debug("Error reading input")
				return err
			}

			projectDir, err := makeDir(year, day)
			if err != nil {
				slog.Debug("Error creating project dir")
				return err
			}

			inputFile := filepath.Join(projectDir, fmt.Sprintf("%d_%d.txt", year, day))
			goFile := filepath.Join(projectDir, "main.go")

			err = writeInputFile(inputFile, *content)
			if err != nil {
				slog.Debug("Error writing input file")
				return err
			}

			err = writeMainGoFile(goFile, year, day)
			if err != nil {
				slog.Debug("Error writing main.go file")
				return err
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("%+v\n", err)
	}
}

func makeDir(year int, day int) (string, error) {
	projectDir := fmt.Sprintf("%d/day%d", year, day)

	err := os.MkdirAll(projectDir, 0755)
	if err != nil {
		return "", err
	}

	return projectDir, nil
}

func getInput(year int, day int, token string) (*[]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: token})

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return &body, nil

}

func writeInputFile(fname string, data []byte) error {

	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func writeMainGoFile(fname string, year int, day int) error {
	vars := make(map[string]interface{})
	vars["year"] = year
	vars["day"] = day

	tmpl, err := template.ParseFiles("templates/main.go.tmpl")
	if err != nil {
		fmt.Println("Cannot make template")
		return err
	}

	f, err := os.Create(fname)
	if err != nil {
		fmt.Println("Cannot make main.go file")
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, vars)
	if err != nil {
		return err
	}
	return nil
}
