# ssl-gen

`ssl-gen` is a small CLI utility for generating and trusting SSL Certificates on macOS.

You can generate a certificate with multiple domains if you desire:

```
ssl-gen new [cert-name] [domains...]
```

Example:

```
ssl-gen new ashsmith ashsmith.io ashsmith.co
Generating certificate called: ashsmith with [ashsmith.io ashsmith.co] domains, will be saved to: ./
Generating a 2048 bit RSA private key
...............................................................................................................................+++
............................+++
writing new private key to './ashsmith.key'
-----
Adding the certificate into Keychain so it is trusted by your machine! [requires sudo]
All done! 🎉
```

After running you'll have two files: `[certname].key` and `[certname].crt`.

You can then include these in your projects that require SSL.