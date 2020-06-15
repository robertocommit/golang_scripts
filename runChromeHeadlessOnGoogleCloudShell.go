/** Download and install Chromium Driver onGoogle Cloud
https://stackoverflow.com/questions/60991515/how-do-i-run-headless-chrome-in-a-shell-on-google-cloud-platform
export CHROME_BIN=/usr/bin/google-chrome
export DISPLAY=:99.0
sh -e /etc/init.d/xvfb start
sudo apt-get update
sudo apt-get install -y libappindicator1 fonts-liberation libasound2 libgconf-2-4 libnspr4 libxss1 libnss3 xdg-utils
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
sudo dpkg -i google-chrome*.deb
*/

package main

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var responseText string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("YOURWEBSITEURL"),
		chromedp.OuterHTML("#content", &responseText),
	); err != nil {
		panic(err)
	}
	fmt.Println(responseText)
}
