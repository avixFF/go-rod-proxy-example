package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/launcher/flags"
)

func main() {
	// Start HTTP proxy using mitmproxy with the following command:
	// > mitmdump -p 8080 --proxyauth user:password

	// Create a browser launcher
	l := launcher.New()
	// Pass '--proxy-server=127.0.0.1:8081' argument to the browser on launch
	l = l.Set(flags.ProxyServer, "127.0.0.1:8080")
	// Launch the browser and get debug URL
	controlUrl, _ := l.Launch()

	// Connect to the newly launched browser
	b := rod.New().ControlURL(controlUrl).MustConnect()

	// Handle proxy authentication pop-up
	// Notice how HandleAuth returns a function that
	// must be started as a goroutine!
	go b.HandleAuth("user", "password")()

	// Ignore certificate errors since we are using local insecure proxy
	b.MustIgnoreCertErrors(true)

	// Navigate to the page that prints IP address
	page := b.MustPage("http://api.ipify.org")

	// IP address should be the same, since we are using local
	// proxy, however the response signals that the proxy works
	println(page.MustElement("html").MustText())
}
