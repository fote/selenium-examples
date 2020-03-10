package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

func main() {
	selenium.SetDebug(true)
	// Setting capabilities
	caps := selenium.Capabilities{"platform": "LINUX", "browserName": "chrome", "version": "79.0"}

	// Point test to the local Selenoid
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:4444/wd/hub")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Open URL
	if err := wd.Get("https://duckduckgo.com/"); err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	elem, err := wd.FindElement(selenium.ByCSSSelector, "input#search_form_input_homepage")
	if err != nil {
		panic(err)
	}

	// Enter some new text in search box.
	err = elem.SendKeys("selenium")
	if err != nil {
		panic(err)
	}

	// // Click the find button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "input#search_button_homepage")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

	// Taking a screenshot
	screenshot, err := wd.Screenshot()
	if err != nil {
		panic(err)
	}

	t := time.Now()
	filename := "selenoid_" + t.Format("20060102150405") + ".png"

	err = ioutil.WriteFile(filename, screenshot, 0644)
	if err != nil {
		panic(err)
	}
	log.Println("SCREENSHOT  -   file:///" + filename)

}
