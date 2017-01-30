# imagechecker

Check image links on URLs.

Example:

```
anton@shell:imagechecker:0$ ./imagechecker https://www.antonlindstrom.com/2015/03/29/introduction-to-apache-mesos.html
+--------------------------------------------------------------------------------+--------------+----------+-----------------+
|                                      URL                                       | CONTENT-TYPE | RESPONSE |      ETAG       |
+--------------------------------------------------------------------------------+--------------+----------+-----------------+
| https://www.antonlindstrom.com/images/mesosmasterquorum-al-150324-w1024_1x.png | image/png    |      200 | "586c1334-627d" |
+--------------------------------------------------------------------------------+--------------+----------+-----------------+
```
