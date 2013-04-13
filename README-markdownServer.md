
<center>
markdownServer.go
=============
</center>

markdownServer is a minimalist webserver that serves markdown docs (*.md)
from a directory.  My version uses /www but thats for convenience rather
than necessity.

Setup
-----

1. Create directory to be served and fill it with your markdown docs.
1. write an index.html file in that directory something like shown below
1. Start the server with markdownServer &
1. Browse to localhost:8080/md locally or hostname:8080/md from a different system.
 
You should see the index.html file created in the above step. 

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

Go makes this embarrasingly simple. It's only 90 lines of code. 

Still, it works well and it may serve as an example
for those just getting started with go.  I also use it
to preview markdown docs before I upload them to github.


BUGS
----
? None known

TODO
----
Nothing planned

NOTES
-----
Requires blackfriday from https://github.com/russross/blackfriday

Thanks go out to John Gruber and Aaron Swartz (RIP) for the original 
idea of more readable web text and early implementations.

http://en.wikipedia.org/wiki/Markdown

history
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

// EOF README.md  (this is a markdown document and tested OK with blackfriday)
