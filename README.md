Projects used:

https://dev.to/osuka42/building-a-simple-system-tray-app-with-go-899

https://github.com/gocolly/colly/blob/master/_examples/hackernews_comments/hackernews_comments.go

https://golang.org/doc/tutorial/getting-started

Generate png icons from font-awesome icons: https://fa2png.app/ (utensils is the desired icon)

Everyting about GO packages:
https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc
https://golangbyexample.com/import-local-module-golang/ 

Scrape tables on pages with gocolly:
https://www.reddit.com/r/golang/comments/cj8l2n/gocolly_help_scraping_only_the_first_table_on_the/

!!!!!
https://jonathanmh.com/web-scraping-golang-goquery/
https://benjamincongdon.me/blog/2018/03/01/Scraping-the-Web-in-Golang-with-Colly-and-Goquery/
set encoding of scraped page at goquery
https://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks#handle-non-utf8-html-pages
goquery selector doc:
https://pkg.go.dev/github.com/PuerkitoBio/goquery#Selection.Html


how to use struct:
https://gobyexample.com/structs
how to use string array in struct:
https://stackoverflow.com/questions/19482612/go-golang-array-type-inside-struct-missing-type-composite-literal

Native macOs desktop app with react and go
https://medium.com/@koddr/how-to-create-a-native-macos-app-on-golang-and-react-js-with-full-code-protection-9162b8c25be5

compile your go app with assets included into the binary:
https://github.com/gobuffalo/packr

schedule function run with cron:
https://stackoverflow.com/questions/16466320/is-there-a-way-to-do-repetitive-tasks-at-intervals
https://github.com/go-co-op/gocron/blob/master/example_test.go
scheduled restart of app in macos:
https://www.reddit.com/r/mac/comments/afw5sn/scheduled_app_restart/

where is $GOPATH on macos or how to set it in bash_profile
https://stackoverflow.com/questions/21499337/cannot-set-gopath-on-mac-osx

generate go byte array from png images:
https://github.com/cratonica/2goarray 
$GOPATH/bin/2goarray iconImageByteData main < icon/icon.png > icon.go

buildmodes and other build options in go:
https://golang.org/pkg/cmd/go/

this is how I built my app for macOs:
go build -buildmode=exe -o=tajfunMenu.app

here is how to add an app to start automatically on macOs login:
https://www.makeuseof.com/tag/add-remove-delay-startup-items-mac/
