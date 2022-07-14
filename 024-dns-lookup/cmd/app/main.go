package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"text/tabwriter"

	"github.com/miekg/dns"
)

func LookupA(fqdn, serverAddr string) ([]string, error) {
	var m dns.Msg
	var ips []string
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeA)
	in, err := dns.Exchange(&m, serverAddr)
	if err != nil {
		return ips, err
	}
	if len(in.Answer) < 1 {
		return ips, errors.New("no answer")
	}
	for _, a := range in.Answer {
		if ip, ok := a.(*dns.A); ok {
			ips = append(ips, ip.A.String())
		}
	}
	return ips, nil
}

func LookupCNAME(fqdn, serverAddr string) ([]string, error) {
	var m dns.Msg
	var fqdns []string
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeCNAME)
	in, err := dns.Exchange(&m, serverAddr)
	if err != nil {
		return fqdns, err
	}
	if len(in.Answer) < 1 {
		return fqdns, errors.New("no answer")
	}
	for _, a := range in.Answer {
		if n, ok := a.(*dns.CNAME); ok {
			fqdns = append(fqdns, n.Target)
		}
	}
	return fqdns, nil
}

type Result struct {
	IPAddr   string
	Hostname string
}

func Lookup(fqdn, serverAddr string) []Result {
	log.Printf("Lookup %s from %s\n", fqdn, serverAddr)
	results := make([]Result, 0)
	cfqdn := fqdn
	// recursively going through each lookedup item
	for {
		cnames, err := LookupCNAME(cfqdn, serverAddr)
		if err == nil && len(cnames) > 0 {
			cfqdn = cnames[0]
			continue // We have to process the next CNAME.
		}
		ips, err := LookupA(cfqdn, serverAddr)
		if err != nil {
			break // There are no A records for this hostname.
		}
		for _, ip := range ips {
			results = append(results, Result{IPAddr: ip, Hostname: fqdn})
		}
		break // We have processed all the results.
	}
	return results
}

func Worker(end chan bool, task chan string, gather chan []Result, serverAddr string) {
	for fqdn := range task {
		results := Lookup(fqdn, serverAddr)
		if len(results) > 0 {
			gather <- results
		}
	}
	end <- true
}

func main() {
	// parse command line args
	var (
		flDomain      = flag.String("domain", "", "The domain to perform guessing against.")
		flWordlist    = flag.String("wordlist", "", "The wordlist to use for guessing.")
		flWorkerCount = flag.Int("c", 100, "The amount of workers to use.")
		flServerAddr  = flag.String("server", "8.8.8.8:53", "The DNS server to use.")
	)

	flag.Parse()

	if *flDomain == "" || *flWordlist == "" {
		fmt.Println("-domain and -wordlist are required")
		os.Exit(1)
	}

	var results []Result
	var task = make(chan string, *flWorkerCount)
	var gather = make(chan []Result)
	var end = make(chan bool)
	var done = make(chan bool)

	fh, err := os.Open(*flWordlist)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fh)

	// first, launch all the workers
	for i := 0; i < *flWorkerCount; i++ {
		go Worker(end, task, gather, *flServerAddr)
	}

	// ready to gather all results
	go func() {
		for r := range gather {
			results = append(results, r...) // appending multiple results each time
		}
		done <- true
	}()

	// scan each line from the file and send as task
	for scanner.Scan() {
		task <- fmt.Sprintf("%s.%s", scanner.Text(), *flDomain)
	}
	close(task) // signal to works, there are no more tasks

	// count works, wait for all to end
	for i := 0; i < *flWorkerCount; i++ {
		<-end
	}
	close(end)

	// no more gather
	close(gather)
	<-done
	close(done)

	// process results
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, ' ', 0)
	for _, r := range results {
		fmt.Fprintf(w, "%s\t%s\n", r.Hostname, r.IPAddr)
	}
	w.Flush()

}
