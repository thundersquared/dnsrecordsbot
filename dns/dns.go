package dns

import (
  "log"
  "net/url"
  "os/exec"
  "strings"
)

type DNS struct {
  Domain string
  RecordTypes []string
}

func Dns(Domain string) (DNS, error) {
  d := DNS{}

  host, err := d.sanitizeDomainName(Domain)

  if err != nil {
    return d, err
  }

  d.Domain = host
  d.RecordTypes = []string {
    "A",
    "AAAA",
    "CNAME",
    "MX",
    "NS",
    "PTR",
    "SOA",
    "TXT",
    "DNSKEY",
  }

  return d, nil
}

func (dns DNS) sanitizeDomainName(s string) (string, error) {
  s = strings.Replace(s, "http://", "", -1)
  s = strings.Replace(s, "https://", "", -1)

  domain := strings.Split(s, "/")

  u, err := url.Parse(strings.ToLower(domain[0]))
  if err != nil {
    return "", err
  }

  if (u.Host == "") {
    return strings.ToLower(domain[0]), nil
  } else {
    return u.Host, nil
  }
}

func (dns DNS) GetRecords() []string {
  var records []string

  if dns.Domain != "" {
    for _, element := range dns.RecordTypes {
      rec := dns.GetRecordsOfType(element)
      records = append(records, strings.TrimSpace(rec))
    }
  }

  return records
}

func (dns DNS) GetRecordsFrom(from string) []string {
  var records []string

  for _, element := range dns.RecordTypes {
    rec := dns.GetRecordsOfTypeFrom(element, from)
    records = append(records, strings.TrimSpace(rec))
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

func (dns DNS) GetRecordsOfTypeFrom(t string, from string) string {
  cmd := exec.Command("dig", "+nocmd", from, dns.Domain, t, "+multiline", "+noall", "+answer")
  out, err := cmd.CombinedOutput()

  if err != nil {
    log.Fatal(err)
  }

  return string(out)
}
