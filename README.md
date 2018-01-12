expect-ct
=====

Simple application that listens for expect-ct errors (well, really any data) to
be posted to it via http, as per
[https://tools.ietf.org/html/draft-ietf-httpbis-expect-ct-01](https://tools.ietf.org/html/draft-ietf-httpbis-expect-ct-01)
and prints to stdout (if you use the included unit file for systemd, it goes
straight into journald).

Eventually, I'll add support for parsing json data (if anyone has a valid
expect-ct report, I'd love to see it, since I haven't recieved any, yet).

