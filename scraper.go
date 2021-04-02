package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/PuerkitoBio/goquery"
)

type dailyMeal struct {
	weekDay int
	mealA []string
	mealB []string
}

func getAvailableWeekIndex(domQuery *goquery.Selection) string {
	availableWeekIndex := "0"
	domQuery.Find("strong").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		content := strings.TrimSpace(s.Text())
		if len(content) > 0 {
			availableWeekIndex = content
			return false
		}
		return true
	})
	return availableWeekIndex
}

func stringInArray(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func buildWeeklyMealStructArray(domQuery *goquery.Selection) []dailyMeal {
	weeklyMeal := []dailyMeal{}
	domQuery.Find("table").First().Find("tr").Next().Each(func(i int, tr *goquery.Selection) {
		actualDailyMeal := dailyMeal{ weekDay: i }
		tr.Find("td").Next().EachWithBreak(func(j int, td *goquery.Selection) bool {
			if j > 1 {
				return false
			}
			meals := []string{}
			td.Find("p").Find("strong").Each(func(k int, strong *goquery.Selection) {
				strongText := strings.TrimSpace(strong.Text())
				if strings.Contains(strongText, "&nbsp;") || strongText == "" {
					return
				}
				if !stringInArray(strongText, meals) {
					meals = append(meals, strongText)
				}
			})
			if j == 0 {
				actualDailyMeal.mealA = meals
			} else {
				actualDailyMeal.mealB = meals
			}
			return true
		})
		weeklyMeal = append(weeklyMeal, actualDailyMeal)
	})
	for _, s := range weeklyMeal {
		fmt.Printf("%#v\n",s)
	}
	return weeklyMeal
}

func fetchMeals() []dailyMeal {
	weeklyMealStructArray := []dailyMeal{}
	c := colly.NewCollector(colly.DetectCharset())
	c.OnHTML(".feedboxbody", func(feedboxbody *colly.HTMLElement) {
		doc := feedboxbody.DOM
		
		availableWeekIndex := getAvailableWeekIndex(doc)
		fmt.Println(availableWeekIndex)
		weeklyMealStructArray = buildWeeklyMealStructArray(doc)
	})
	c.Visit("http://tajfunbiliard.hu/?page=25")
	return weeklyMealStructArray
}