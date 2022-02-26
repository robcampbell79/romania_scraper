package main

import (
    "log"
    "github.com/gocolly/colly"
    "encoding/json"
    "io/ioutil"
)

type Crossing struct {
    ID              string  `json:"id"`
    Checkpoint      string  `json:"checkpoint"`
    Communication   string  `json:"communication"`
    Transportation  string  `json:"transportation"`
    Region          string  `json:"region"`
}

func main() {
    allCrossings := make([]Crossing, 0)

    c := colly.NewCollector(
        colly.AllowedDomains("dpsu.gov.ua"),
    )

    c.OnHTML("table", func(e *colly.HTMLElement) {
        e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
            crossing := Crossing {
                ID:             el.ChildText("td:nth-child(1)"),
                Checkpoint:     el.ChildText("td:nth-child(2)"),
                Communication:  el.ChildText("td:nth-child(3)"),
                Transportation: el.ChildText("td:nth-child(4)"),
                Region:         el.ChildText("td:nth-child(5)"),
            }

            allCrossings = append(allCrossings, crossing)
        })
    })

    c.Visit("https://dpsu.gov.ua/en/AT-THE-BORDER-WITH-ROMANIA/")

    writeJSON(allCrossings)
}

func writeJSON(data []Crossing) {
    file, err := json.MarshalIndent(data, "", " ")
    if err != nil {
        log.Println("Can't create JSON file")
    }

    _ = ioutil.WriteFile("romcrossings.json", file, 0644)
}