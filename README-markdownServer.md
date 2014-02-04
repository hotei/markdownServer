
<center>
markdownServer.go
=============
</center>

markdownServer is a minimalist webserver that serves markdown docs (*.md)
from a directory.  My version uses /www but thats for convenience rather
than necessity. Builds with go 1.2 (and probably 1.X). 

Comments can be sent through github or directly to <hotei1352@gmail.com> .

Setup
-----

1. Create directory to be served and fill it with your markdown docs.
	* for convenience in updating these documents can be links 
	* for that matter the directory could be empty except for index.html which
	would point back to the original file locations but that makes backup a bit
	harder.  Your choice.
1. write an index.html file in that directory something like shown below
1. get blackfriday
	* I usually download zips from github since I'm usually not interested in history
	or branches.  go get may also work.
	* unpack it into your gopath so it ends up at $GOPATH/github.com/blackfriday
	* if you have a multi-part $GOPATH go get will put it into the first part. Thats
	the recommended way if you do it manually as well.
1. Edit markdownServer.go if desired for possible customizations:
	* portNum : I used 8080.  If you change it, the invoking line must also change
	* virtualURL : This directory doesn't exist in the file tree, its what you
	browse to the server with.  md made sense to me but "markdown" or "docs" might
	be just as good.
	* serverRoot : This one has to exist but it can be whereever you want. 
	* mdDir : This is optional but recommended.  If your index.html file doesn't
	exist or is unreadable this will patch in an error message and possibly a doc file
	or two in either markdown or html.
1. Start the server with markdownServer &  Assuming no problems it will announce
	itself at portNum and start serving.
1. Browse to localhost:8080/md locally or hostname:8080/md from a different system.
 
In the browser you should see the list of documents from your index.html file.  Clicking
on a markdown document will convert it to html on the fly.

Sample index.html
-----------------

```
<html>

<head>
<title> Markdown</title>
</head>
<body>
<h2>Markdown Docs</h2>
<ul>
<li><a href="todo.md" >Todo</li>
<hr>
<li><a href="README-abc.md" >abc README</a></li>
<li><a href="README-def.md" >def README</a></li>
</ul>

</body>
</html>
```

Go makes this embarrasingly simple. It's only 90 lines of code. The real magic
happens in Russ Ross's blackfriday or in the go net/http library.  My contribution
is just a little glue.

Still, it works well and it may serve as an example
for those just getting started with go.  I also use it
to preview markdown docs before I upload them to github.


BUGS
----
None known

SECURITY
--------
* No input is requested from the user. If that is changed a more paranoid approach
may be necessary (templates perhaps).
* Opening markdown files in itself should be safe but they can have links that go
anywhere. Following those links requires the usual warning.
* Embedded images that are corrupted (intentionally or otherwise) could pose a risk.
  * This is a browser issue so keeping that updated is a good idea.
* Security issues of dependencies and the go standard library should be tracked
  * Watch for new releases of blackfriday on github
  * Watch for new releases of go on golang.org and/or golang-nuts forum.
*  If you don't personally control the files that go into the served directory it
would be wise to sanity check the files before reading them. (see todo)

TODO
----
* test IsRegularFile(fname) before ioutil.ReadFile(fname)
	
NOTES
-----
Requires blackfriday from https://github.com/russross/blackfriday

Thanks go out to John Gruber and Aaron Swartz (RIP) for the original 
idea of more readable web text and early implementations.

http://en.wikipedia.org/wiki/Markdown

History
-------
* 2013-04-12 publish on github.com/hotei/markdownServer
* 2012-04-02 started

License
-------
The 'markdownServer' go program is distributed under the Simplified BSD License:

> Copyright (c) 2012-2013 David Rook. All rights reserved.
> 
> Redistribution and use in source and binary forms, with or without modification, are
> permitted provided that the following conditions are met:
> 
>    1. Redistributions of source code must retain the above copyright notice, this list of
>       conditions and the following disclaimer.
> 
>    2. Redistributions in binary form must reproduce the above copyright notice, this list
>       of conditions and the following disclaimer in the documentation and/or other materials
>       provided with the distribution.
> 
> THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER ``AS IS'' AND ANY EXPRESS OR IMPLIED
> WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
> FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> OR
> CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
> CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
> SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
> ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
> NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
> ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// EOF README-markdownServer.md  (this is a markdown document and tested OK with blackfriday)
