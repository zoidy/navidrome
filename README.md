GoSonic
=======

[![Build Status](https://travis-ci.org/deluan/gosonic.svg?branch=master)](https://travis-ci.org/deluan/gosonic) [![Go Report Card](https://goreportcard.com/badge/github.com/deluan/gosonic)](https://goreportcard.com/report/github.com/deluan/gosonic)

__This is still a work in progress, and has no releases available__

GoSonic is an application that implements the [Subsonic API](http://www.subsonic.org/pages/api.jsp), but instead of
having its own music library like the original [Subsonic application](http://www.subsonic.org), it interacts directly
with your iTunes library.

The project's main goals are:

* Be fully compatible with available [Subsonic clients](http://www.subsonic.org/pages/apps.jsp)
  (actively being tested with
    [DSub](http://www.subsonic.org/pages/apps.jsp#dsub), 
    [SubFire](http://www.subsonic.org/pages/apps.jsp#subfire) and 
    [Jamstash](http://www.subsonic.org/pages/apps.jsp#jamstash))
* Use all metadata from iTunes, so that you can keep using iTunes to manage your music
* Keep iTunes stats (play counts, last played dates, ratings, etc..) updated, at least on Mac OS X. 
  This allows smart playlists to be used in Subsonic Clients 
* Help me learn Go ;) [![Gopher](https://blog.golang.org/favicon.ico)](https://golang.org)


###  Supported Subsonic API version

I'm currently trying to implement all functionality from API v1.8.0, with some exceptions.

Check the (almost) up to date [compatibility chart](https://github.com/deluan/gosonic/wiki/Compatibility) for what is working.

### Development Environment

You will need to install [Go 1.6](https://golang.org/dl/)

Then install dependencies:
```
$ go get github.com/beego/bee   
$ go get github.com/gpmgo/gopm
$ gopm get -v -g
```  

From here it's a normal [BeeGo](http://beego.me) development cycle. Some useful commands:

```bash
# Start local server (with hot reload)
$ bee run

# Start test runner on the browser
$ NOLOG=1 goconvey --port 9090

# Run all tests
$ go test ./... -v
```


### Copying

GoSonic - Copyright (C) 2016  Deluan Cotts Quintao

The source code is licensed under GPL v3. License is available [here](/LICENSE)