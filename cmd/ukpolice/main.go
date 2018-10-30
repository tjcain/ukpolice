package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tjcain/ukpolice"
)

func main() {
	start := time.Now()

	customClient := http.Client{Timeout: time.Second * 120}

	client := ukpolice.NewClient(&customClient)

	searches, resp, err := client.StopAndSearch.GetStopAndSearchesByForce(context.Background(),
		ukpolice.WithForce("metropolitan"), ukpolice.WithDate("2017-06"))
	if err != nil {
		log.Printf("StopAndSearch.GetStopAndSearchesByForce returned error: '%s'", err)
	}

	fmt.Println(time.Since(start).Seconds(), resp.Status, len(searches))

	searches, resp, err = client.StopAndSearch.GetStopAndSearchesByForce(context.Background(),
		ukpolice.WithForce("metropolitan"), ukpolice.WithDate("2016-11"))
	if err != nil {
		log.Printf("StopAndSearch.GetStopAndSearchesByForce returned error: '%s'", err)
	}

	fmt.Println(time.Since(start).Seconds(), resp.Status, len(searches))

	// a, _, err := getAvaliable(client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ch := make(chan string)

	// for _, m := range a {
	// 	go fetchAndStore(client, m.Date, "metropolitan", ch)
	// }

	// for range a {
	// 	fmt.Println(<-ch)
	// }

	// for _, m := range a {
	// 	for _, f := range m.StopAndSearch {
	// 		go fetchAndStore(client, m.Date, f, ch)
	// 	}
	// }

	// for _, m := range a {
	// 	for range m.StopAndSearch {
	// 		fmt.Println(<-ch)
	// 	}

	// }

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func getAvaliable(client *ukpolice.Client) ([]ukpolice.AvaliabilityInfo, *ukpolice.Response, error) {
	return client.Avaliability.GetAvaliabilityInfo(context.Background())
}

func fetchAndStore(client *ukpolice.Client, date, force string, ch chan<- string) {
	start := time.Now()

	searches, resp, err := client.StopAndSearch.GetStopAndSearchesByForce(context.Background(),
		ukpolice.WithDate(date), ukpolice.WithForce(force))

	if err != nil {
		ch <- fmt.Sprintf("ERROR FROM %s-%s: %s", date, force, err) // send err to channel ch
		return
	}

	status := resp.Status
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%s\t%s-%s\tresponse length: %d", secs, status, date, force, len(searches))
}
