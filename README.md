##### Example:

```ssh ubuntu@$(up -h mysite.com -p 22)```

Will automatically SSH when the host is available on port 22

When it successfully connects it returns the hostname so you can pipe it or use it in another command like the example.
