package main

import (
	"fmt"
	"strings"
	"bytes"
	"os/exec"
)

func Download(url string) (string, error) {
  toolName := "yt-dlp"
  cmd := exec.Command(toolName, "-g", "-f", "best[ext=mp4]/best", url) // first tool

  var out bytes.Buffer
  cmd.Stdout = &out

  err := cmd.Run()
  if err == nil && out.Len() > 0 {
    return strings.TrimSpace(out.String()), nil
  }

  cmd = exec.Command("gallery-dl", "-g", url) // second tool
  out.Reset() // resets stored output
  cmd.Stdout = &out 

  err = cmd.Run()
  if err != nil {
    return " ", err
  }
 return strings.TrimSpace(out.String()), nil
}

func Main() {
	url := "Your Link"
	a, b := Download(url)
	fmt.Print(a, b) // returns output, err
}
