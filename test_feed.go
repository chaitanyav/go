package main

import (
	"fmt"
	"github.com/iand/feedparser"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func parseFeed(feedURL string, completed chan bool) {
	response, err := http.Get(feedURL)
	if err != nil {
		fmt.Println(err)
	} else {
		responseBytes, error := ioutil.ReadAll(response.Body)
		if error != nil {
			fmt.Println(error)
		} else {
			reader := strings.NewReader(string(responseBytes))
			feed, _ := feedparser.NewFeed(reader)
			starString := strings.Repeat("*", 80)
			fmt.Println(starString)
			fmt.Println(feed.Title)
			fmt.Println(feed.Link)
			fmt.Println("Num Posts/Items: ", len(feed.Items))
			fmt.Println(starString)

			// uncomment the lines below to the individual posts/items
			/*
			     for _, item := range feed.Items {
			     fmt.Println("Title: ", item.Title)
			     fmt.Println("Description: ", item.Description)
			   }
			*/
		}
	}
	completed <- true
}

func main() {

	categoryByFeedURL := map[string]string{
		"http://cdnjs.com/rss.xml":                                                   "developer",
		"https://blog.twitter.com/api/blog.rss?name=engineering&alt=rss":             "developer",
		"http://rubyrogues.com/feed/":                                                "developer",
		"https://tech.dropbox.com/feed/":                                             "developer",
		"http://news.ycombinator.com/rss":                                            "developer",
		"http://feeds.feedburner.com/HighScalability":                                "developer",
		"http://feeds.feedburner.com/html5rocks":                                     "developer",
		"http://feeds.feedburner.com/TechCrunch/":                                    "tech",
		"http://venturebeat.com/feed/":                                               "tech",
		"https://github.com/blog.atom":                                               "developer",
		"http://weblog.rubyonrails.org/feed/atom.xml":                                "developer",
		"http://podcasts.thoughtbot.com/giantrobots.xml":                             "developer",
		"http://feeds.feedburner.com/railscasts":                                     "developer",
		"http://feeds.feedburner.com/nvidia/parallelforall?format=xml":               "developer",
		"http://feeds.feedburner.com/codinghorror":                                   "developer",
		"http://rss.slashdot.org/Slashdot/slashdot":                                  "tech",
		"http://www.kickstarter.com/projects/feed.atom":                              "tech",
		"http://feeds.mashable.com/Mashable":                                         "tech",
		"http://feeds.gawker.com/gizmodo/full":                                       "tech",
		"http://feeds.feedburner.com/ommalik":                                        "tech",
		"http://feeds.feedburner.com/OfficialAndroidBlog":                            "developer",
		"http://vimcasts.org/feeds/ogg":                                              "developer",
		"http://www.engadget.com/rss.xml":                                            "tech",
		"http://blogs.computerworld.com/rss.xml":                                     "tech",
		"http://feeds.feedburner.com/cnet/NnTv":                                      "tech",
		"http://feeds.wired.com/wired/index":                                         "tech",
		"http://www.anandtech.com/rss/":                                              "tech",
		"http://appleinsider.com.feedsportal.com/c/33975/f/616168/index.rss":         "tech",
		"http://feeds.arstechnica.com/arstechnica/index?format=xml":                  "tech",
		"http://www.pcworld.com/index.rss":                                           "tech",
		"http://feeds.searchengineland.com/searchengineland?format=xml":              "tech",
		"http://feeds.feedburner.com/UnwiredView":                                    "tech",
		"http://www.theregister.co.uk/headlines.atom":                                "tech",
		"http://www.technologyreview.com/topnews.rss":                                "tech",
		"http://feeds.feedburner.com/tedtalks_video":                                 "tech",
		"http://feeds.feedburner.com/avc":                                            "tech",
		"http://feed.onstartups.com/onstartups":                                      "tech",
		"http://feeds.feedburner.com/IeeeSpectrumPodcasts":                           "tech",
		"http://feeds.feedburner.com/IeeeSpectrumVideo":                              "tech",
		"http://feeds.feedburner.com/IeeeSpectrumFullText":                           "tech",
		"http://feeds.feedburner.com/GoogleOpenSourceBlog":                           "developer",
		"http://feeds.feedburner.com/37signals/beMH":                                 "developer",
		"http://www.joelonsoftware.com/rss.xml":                                      "developer",
		"http://martinfowler.com/feed.atom":                                          "developer",
		"http://feeds.hanselman.com/scotthanselman":                                  "developer",
		"http://www.randsinrepose.com/index.xml":                                     "developer",
		"http://bokardo.com/feed/":                                                   "developer",
		"http://blog.stackoverflow.com/feed/":                                        "developer",
		"http://feeds.feedburner.com/thoughtworks-blogs":                             "developer",
		"http://www.artima.com/weblogs/feeds/weblogs.rss":                            "developer",
		"http://feeds.feedburner.com/blogspot/RLXA":                                  "developer",
		"http://feeds.feedburner.com/blogspot/SbSV":                                  "developer",
		"http://feeds.feedburner.com/InsideSearch":                                   "developer",
		"http://feeds.feedburner.com/GDBcode":                                        "developer",
		"http://feeds2.feedburner.com/blogspot/Egta":                                 "developer",
		"http://feeds.feedburner.com/GoogleOnlineSecurityBlog":                       "developer",
		"http://feeds.feedburner.com/GoogleStudentBlog":                              "developer",
		"http://feeds.feedburner.com/GoogleTranslateBlog":                            "developer",
		"http://techblog.netflix.com/feeds/posts/default":                            "developer",
		"http://tech-book-store.amazon.com/blog/feed/recentPosts.rss":                "developer",
		"http://feeds.feedburner.com/Ruby5?format=xml":                               "developer",
		"http://feeds.feedburner.com/PetrMitrichev?format=xml":                       "developer",
		"http://xkcd.com/rss.xml":                                                    "developer",
		"http://www.perfplanet.com/rss.xml":                                          "developer",
		"http://feed.dilbert.com/dilbert/daily_strip?format=xml":                     "developer",
		"http://www.stanford.edu/group/edcorner/uploads/podcast/EducatorsCorner.xml": "developer",
		"http://www.zdnet.com/blog/rss.xml":                                          "tech",
		"http://feeds.reuters.com/reuters/topNews":                                   "news",
		"http://rss.cnn.com/rss/cnn_topstories.rss":                                  "news",
		"http://feeds.abcnews.com/abcnews/topstories":                                "news",
		"http://www.npr.org/rss/rss.php?id=1001":                                     "news",
		"http://feeds.harvardbusiness.org/harvardbusiness?format=xml":                "business",
		"http://feeds.thestreet.com/tsc/feeds/rss/latest-stories":                    "business",
		"http://www.forbes.com/real-time/feed2/":                                     "business",
		"http://feeds.feedburner.com/freakonomics":                                   "business",
		"http://podcast.cnbc.com/mmpodcast/lightninground.xml":                       "business",
		"http://www.economist.com/feeds/print-sections/77/business.xml":              "business",
		"http://feeds.feedburner.com/freakonomicsradio?format=xml":                   "business",
		"http://feeds.feedburner.com/fastcompany/headlines":                          "business",
		"http://feeds.fool.com/usmf/foolwatch":                                       "business",
	}
	numFeedURLs := len(categoryByFeedURL)
	completed := make(chan bool, numFeedURLs)

	for feedURL, _ := range categoryByFeedURL {
		go parseFeed(feedURL, completed)
	}

	for _ = range categoryByFeedURL {
		select {
		case <-completed:
			fmt.Println("Finished Parsing a feed at: ", time.Now())
		case <-time.After(1 * time.Minute):
			fmt.Println("Time out!!, lets move on")
			completed <- false
		}
	}

	fmt.Println("Finished fetching all feeds")
}
