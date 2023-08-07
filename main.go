package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// demonstration of requirements
func main() {
	flightBytes, err := os.ReadFile("flights.json")
	if err != nil {
		log.Fatalf("failed to read flights.json: %v\n", err)
	}

	var flightData FlightData
	if err := json.Unmarshal(flightBytes, &flightData); err != nil {
		log.Fatalf("failed to unmarshal flight data %v\n", err)
	}

	fmt.Println(airplane)

	// add remove flights from the network

	planner := NewPlanner()
	header("Adding Flights...")
	for _, v := range flightData.Flights {

		err := planner.Add(v.Source, v.Destination, v.Duration)
		if err != nil {
			fmt.Printf("❌ Failed to add flight %s -> %s\n", v.Source, v.Destination)
		}

		fmt.Printf("✅ Succesfully added flight %s -> %s\n", v.Source, v.Destination)

	}
	fmt.Println("Done!")

	// shortest path between airports
	header("Searching for fastest flight plan for JFK -> LAX...")
	path, duration := planner.ShortestTravelPath("JFK", "LAX")
	fmt.Printf(
		"\nFastest travel plan has a duration %d with a flight path of %s\n",
		duration,
		strings.Join(path, " -> "),
	)

	// remove flights from the network
	header("Update: order to cancel flight JFK -> ATL received")
	err = planner.Remove("JFK", "ATL")
	if err != nil {
		fmt.Println("❌ Cancellation failed")
	}
	fmt.Printf("✅ Flight cancelled")

	// show all flights from given airport
	header("Latest Flight Schedule")
	for i, v := range flightData.Airports {
		flights, err := planner.Flights(v)
		if err != nil {
			fmt.Printf("❌Failed To lookup airport %s", v)
		}

		var flightsStr string
		if len(flights) == 0 {
			flightsStr = "None"
		} else {
			for _, f := range flights {
				flightsStr += fmt.Sprintf(" %s(%d)", f.To, f.Duration)
			}
		}

		fmt.Printf("(%d): Airport: %s Destinations: %s\n\n", i+1, v, flightsStr)
	}
}

type FlightData struct {
	Airports []string `json:"airports"`
	Flights  []struct {
		Source      string `json:"source"`
		Destination string `json:"destination"`
		Duration    int    `json:"duration"`
	} `json:"flights"`
}

func header(msg string) {
	bar := "#################################################"
	fmt.Printf("\n\n%s\n%s\n%s\n\n", bar, msg, bar)
}

var airplane = `
Flight Planner Demo
                                    |
                                    |
                                  .-'-.
                                 ' ___ '
                       ---------'  .-.  '---------
       _________________________'  '-'  '_________________________
        ''''''-|---|--/    \==][^',_m_,'^][==/    \--|---|-''''''
                      \    /  ||/   H   \||  \    /
                       '--'   OO   O|O   OO   '--'

ASCII art by Jon Hyatt (whatfer@u.washington.edu)`
