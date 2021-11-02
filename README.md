To produce logs:
- `cd producer` (in new terminal window/tab)      
- `go run .`  This will produce random number of logs/sec in range of 1 to 10

---
To view stats:
- `go run .`  (in new terminal window/tab). This will show logs stats in last 10 sec duration

---
To view Apache logs in terminal (produced by `producer`)
- `tail -f /var/log/apache2/access_log`  (in new terminal window/tab)
  OR  
-`tail -f /usr/local/var/log/httpd/access_log` 