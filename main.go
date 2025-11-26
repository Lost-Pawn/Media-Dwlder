package main

import (
	"fmt"
	"strings"
	"bytes"
	"os/exec"
)

func download(URL string) (string, error) {
  toolName := "yt-dlp"
  cmd := exec.Command(toolName, "-g", "-f", "best[ext=mp4]/best", URL)

  var out bytes.Buffer
  cmd.Stdout = &out

  err := cmd.Run()
  if err == nil && out.Len() > 0 {
    return strings.TrimSpace(out.String()), err
  }

  cmd = exec.Command("gallery-dl", "-g", URL)
  out.Reset()
  cmd.Stdout = &out

  err = cmd.Run()
  if err != nil {
    return " ", err
  }
 return strings.TrimSpace(out.String()), err
}

func main() {
	URL := "Your MMS Link"
	a, b := download(URL)
	fmt.Print(a, b)
}
