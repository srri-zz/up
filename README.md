##### Example:

```ssh ubuntu@$(up -h mysite.com -p 22)```

Will automatically SSH when the host is available on port 22

When it successfully connects it returns the hostname so you can pipe it or use it in another command like the example.

added support for lists:
▶ ./up -host=api.deploypanel.com -ports=80,443 -pr=tcp,udp
OK! - api.deploypanel.com - 80 - tcp
OK! - api.deploypanel.com - 443 - udp
