# `dig`

```
sudo apt update
sudo apt upgrade
sudo apt-get install -y dnsutils

dig -v
```

## The Address (A) record 

```
dig google.com. a
```

## The Start of Authority (SOA) record

```
dig google.com. soa
```

## The Name Server (NS) record

```
dig google.com. ns
```

## The Canonical Name (CNAME) record

```
dig mail.google.com. a
```

## The Mail Exchange (MX) record

```
dig google.com. mx
```

## The Pointer (PTR) record

```
dig 4.4.8.8.in-addr.arpa. ptr
```

## The Text (TXT) record

```
dig google.com. txt
```

## Test Run

```
go run main.go

dig @localhost -p 5053 google.com. a
dig @localhost -p 5053 google.com. soa
dig @localhost -p 5053 google.com. ns
dig @localhost -p 5053 google.com. txt
```
