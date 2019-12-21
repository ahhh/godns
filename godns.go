package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var args = []string{
	"https://dns.twnic.tw/dns-query",
	"https://doh-2.seby.io/dns-query",
	"https://dns.containerpi.com/dns-query",
	"https://cloudflare-dns.com/dns-query",
	"https://doh-fi.blahdns.com/dns-query",
	"https://doh-jp.blahdns.com/dns-query",
	"https://dns.dns-over-https.com/dns-query",
	"https://doh.securedns.eu/dns-query",
	"https://dns.rubyfish.cn/dns-query",
}

func main() {
	//DNSResolvers := setupResolvers()

	// Let's setup our flags and parse them
	resolverTarget := flag.String("resolver", "https://cloudflare-dns.com/dns-query", "File of targets to connect to (host:port).  Port is optional.")
	queryTarget := flag.String("q", "example.com", "The domain to query")
	queryType := flag.String("type", "A", "The type of DNS query to make; default is A")
	flag.Parse()
	response, err := BaseRequest(*resolverTarget, *queryTarget, *queryType)
	if err != nil {
		log.Printf("Error: %v", response)
	}
	log.Println(response)
}

// BaseRequest makes a DNS over HTTP (DOH) GET request for a specified query
func BaseRequest(server, query, qtype string) (string, error) {
	//encquery := base64.StdEncoding.EncodeToString([]byte(query))
	//encquery = url.QueryEscape(encquery)
	qurl := server + "?name=" + query + "&type=" + qtype
	client := &http.Client{}
	req, _ := http.NewRequest("GET", qurl, nil)
	req.Header.Set("accept", "application/dns-json")
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error getting the url")
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error getting the url")
		return "", err
	}
	return string(body), nil
}

// Add a post resolution ?
