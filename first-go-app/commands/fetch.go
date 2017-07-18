package commands

import (
	"fmt"
	"os"
	"os/signal"

	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Feeds []string
	Port  int
}

var fetchCmd = &cobra.Command{
	Use:   "fecth",
	Short: "Fetch feeds",
	Long:  `Dagobah will fetch all feeds listed in config file`,
	Run:   fetchRun,
}

func init() {
	fetchCmd.Flags().Int("rsstimeout", 5, "Timeout (min) for RSS retrival")
	viper.BindPFlag("rsstimeout", fetchCmd.Flags().Lookup("rsstimeout"))
}

func fetchRun(cmd *cobra.Command, args []string) {
	Fetcher()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}

func Fetcher() {
	var config Config

	if err := viper.Marshal(&config); err != nil {
		fmt.Println(err)
	}

	for _, feed := range config.Feeds {
		go PollFeed(feed.Url)
	}
}

func PollFeed(uri string) {
	timeout := viper.GetInt("RSSTimeout")
	if timeout < 1 {
		timeout = 1
	}
	feed := rss.New(timeout, true, chanHandler, itemHandler)

	for {
		if err := feed.Fetch(uri, nil); err != nil {
			fmt.Fprintf(os.Stderr, "[e] %s: %s", uri, err)
			return
		}

		fmt.Printf("Sleeping for %d seconds on %s\n", feed.SecondsTillUpdate(), uri)
		time.sleep(time.Duration(feed.SecondsTillUpdate() * 1e9))
	}
}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	fmt.Printf("%d new channel(s) in %s\n", len(newchannels), feed.Url)
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	fmt.Printf("%d new item(s) in %s\n", len(newitems), feed.Url)
}
