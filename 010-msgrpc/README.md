# Test Run

[Metasploit User Manual](https://docs.rapid7.com/metasploit/rpc-api/)

```
msfconsole
msf6 > load msgrpc Pass=s3cr3t ServerHost=10.0.1.6

export MSFHOST=10.0.1.6:55552
export MSFPASS=s3cr3t