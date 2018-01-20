#!/bin/bash

LC_CTYPE=C tr -dc A-Za-z0-9_\!\@\#\$\%\^\&\*\(\)-+= < /dev/urandom | head -c 32 | tr -d "1Il0O"  | xargs
#LC_CTYPE=C cat /dev/urandom | tr -dc "[:graph:]" | tr -d "1Il0O~^" | fold -w 30 | head -n 1
