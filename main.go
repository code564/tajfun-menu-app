package main

import (
    "fmt"
	"time"
    "io/ioutil"

    "github.com/getlantern/systray"
	"github.com/go-co-op/gocron"
)

var aMealMenu0 *systray.MenuItem
var aMealMenu1 *systray.MenuItem
var bMealMenu0 *systray.MenuItem
var bMealMenu1 *systray.MenuItem

func getWeekDayIndex() int {
	weekdayIndex := int(time.Now().Weekday())-1
	if weekdayIndex > 4 || weekdayIndex < 0 {
		weekdayIndex = 0
	}
	return weekdayIndex
}

func initScheduler() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Seconds().Do(updateMenu)
	s.StartAsync()
	s.RunAll()
}

func updateMenu() {
	weekdayIndex := getWeekDayIndex()
	systray.SetTooltip("Tájfun Étterem Menü" + " " + time.Now().Format(time.RFC3339))
	menu := fetchMeals()
	dailyMenu := menu[weekdayIndex]
    
	if len(dailyMenu.mealA) > 0 {
		aMealMenu0.SetTitle(dailyMenu.mealA[0]);
	}
	if len(dailyMenu.mealA) > 1 {
		aMealMenu1.SetTitle(dailyMenu.mealA[1]);
	}
	if len(dailyMenu.mealB) > 0 {
		bMealMenu0.SetTitle(dailyMenu.mealB[0]);
	}
	if len(dailyMenu.mealB) > 1 {
		bMealMenu1.SetTitle(dailyMenu.mealB[1]);
	}
}

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	weekdayIndex := getWeekDayIndex()
	menu := fetchMeals()
	dailyMenu := menu[weekdayIndex]
    systray.SetTemplateIcon(iconImageByteData, iconImageByteData)
    systray.SetTooltip("Tájfun Étterem Menü")

	aMealMenu := systray.AddMenuItem("A menü", "")
	if len(dailyMenu.mealA) > 0 {
		aMealMenu0 = aMealMenu.AddSubMenuItem(dailyMenu.mealA[0], "")
	}
	if len(dailyMenu.mealA) > 1 {
		aMealMenu1 = aMealMenu.AddSubMenuItem(dailyMenu.mealA[1], "")
	}

	bMealMenu := systray.AddMenuItem("B menü", "")
	if len(dailyMenu.mealB) > 0 {
		bMealMenu0 = bMealMenu.AddSubMenuItem(dailyMenu.mealB[0], "")
	}
	if len(dailyMenu.mealB) > 1 {
		bMealMenu1 = bMealMenu.AddSubMenuItem(dailyMenu.mealB[1], "")
	}
	systray.AddSeparator()
	exitMenu := systray.AddMenuItem("Kilépés", "")
	go func() {
		<-exitMenu.ClickedCh
		fmt.Println("Tájfun Étterem Menü - kilépés")
		systray.Quit()
	}()
	initScheduler()
}

func onExit() {
    // Cleaning stuff here.
}

func getIcon(s string) []byte {
    b, err := ioutil.ReadFile(s)
    if err != nil {
        fmt.Print(err)
    }
    return b
}