package config

import (
	"encoding/json"
	"flag"
	"io"
	"os"
)

type Settings_struct struct {
	Height    int    `json:"height"`
	Width     int    `json:"width"`
	Cell_Size int    `json:"cell_size"`
	Title     string `json:"title"`
	FPS       int    `json:"fps"`
	BG_Col    string `json:"bg_col"`
	FG_Col    string `json:"fg_col"`
}

var Settings Settings_struct
var setting_json string

func Init() {
	flag.StringVar(&setting_json, "settings", "settings.json", "Settings file")
	flag.Parse()

	json_file, err := os.Open(setting_json)
	if err != nil {
		panic(err)
	}

	byte_val, err := io.ReadAll(json_file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byte_val, &Settings)
	if err != nil {
		panic(err)
	}

	json_file.Close()

	if Settings.Height <= 0 || Settings.Width <= 0 || Settings.FPS <= 0 || Settings.Cell_Size <= 0 {
		panic("Settings file contains non-positive values")
	}

	if Settings.Title == "" {
		Settings.Title = "Gonway's Game of Life"
	}

	if Settings.BG_Col == "" {
		Settings.BG_Col = "000000"
	}

	if Settings.FG_Col == "" {
		Settings.FG_Col = "FFFFFF"
	}

}
