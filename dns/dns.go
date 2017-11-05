package dns

import (
  "log"
  "os/exec"
  "strings"
)

type DNS struct {
  Domain string
  RecordTypes []string
}

func Dns(Domain string) DNS {
  d := DNS{}
  d.Domain = d.sanitizeDomainName(Domain)
  d.RecordTypes = []string {
    "A",
    "AAAA",
    "NS",
    "SOA",
    "MX",
    "TXT",
    "DNSKEY",
  }
  return d
}

func (dns DNS) sanitizeDomainName(url string) string {
  url = strings.Replace(url, "http://", "", -1)
  url = strings.Replace(url, "https://", "", -1)
  var domain = strings.Split(url, "/")
  return strings.ToLower(domain[0])
}

func (dns DNS) GetRecords() []string {
  var records []string

  for _, element := range dns.RecordTypes {
    rec := dns.GetRecordsOfType(element)
    records = append(records, rec)
  }

  return records
}

func (dns DNS) GetRecordsOfType(t string) string {
  cmd := exec.Command("dig", "+nocmd", dns.Domain, t, "+multiline", "+noall", "+answer")
  out, err := cmd.CombinedOutput()

  if err != nil {
    log.Fatal(err)
  }

  return string(out)
}
