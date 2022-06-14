package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"os/exec"
	"strings"
	"time"
)

// Chromium provides the ability to open or connect to an instance of a chromium based browser on localhost.
type Chromium struct {
	browserCtx context.Context
	browserCancel context.CancelFunc
	dpCtx context.Context
	dpCancel context.CancelFunc
}

// Start creates a chromium headless browser instance.
// Brave Browser on macOS is the only supported browser at this time.
func (c *Chromium) Start () {
	if _, err := exec.LookPath("/Applications/Brave Browser.app/Contents/MacOS/Brave Browser"); err != nil {
		panic("Brave Browser must be installed to generate LeetCode problem README files.")
	}

	c.browserCtx, c.browserCancel = context.WithCancel(context.Background())

	cmd := exec.CommandContext(c.browserCtx,
		"/Applications/Brave Browser.app/Contents/MacOS/Brave Browser",
		"--headless",
		"--disable-gpu",
		"--remote-debugging-port=9222")

	if err := cmd.Start(); err != nil {
		panic(err)
	}
}

// Run executes a given set of chromedp.Action's on a locally running chromium debugging instance.
func (c *Chromium) Run(actions...chromedp.Action) error {
	var err error

	for retryCount := 0; retryCount < 3; retryCount++ {
		c.dpCtx, c.dpCancel = chromedp.NewRemoteAllocator(context.Background(), "http://localhost:9222")
		c.dpCtx, c.dpCancel = chromedp.NewContext(c.dpCtx)

		if err = chromedp.Run(c.dpCtx, actions...); err != nil {
			if strings.HasPrefix(err.Error(), "could not dial") {
				time.Sleep(100 * time.Millisecond)
				continue
			}
		}

		break
	}

	return err
}

// Close disposes of any managed chromium instances and connections.
func (c *Chromium) Close() {
	if c.dpCancel != nil {
		c.dpCancel()
	}

	if c.browserCancel != nil {
		c.browserCancel()
	}
}
