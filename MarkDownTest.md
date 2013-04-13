<!-- this is a markdown document MarkDownTest.md -->

MarkDown Test 
=============

The above header done with two underlines  (using equal sign)
should be larger than the one below which only has one (using minus sign)

MarkDown Test 
-------------

Portions of this document were extracted from [JOHN GRUBER's Markdown: Basics]
(http://daringfireball.net/projects/markdown/basics)

Indented text avoids whitespace deletion but is placed flush left.  The html `<code>`
attribute
is used which **may** result in a monospaced non-serif font.  Three lines below
are indented in the raw markdown doc.

	Edit BIG_NUMBER	in main() - to finish faster reduce it
	Edit stress in cpu_hog() - true will push CPU, false it loafs along
	Edit loops_to_do in main() - set to more than NUM_CPUS to evoke throttle

Horizontal rule is created by a string of three or more asterisks. 
	
***

or a horizontal rule can be created by a string of three or more minus signs

---

* List item 1 starts with a star
* List item 2 also starts with a star
* List item 3 same
* List item 4 same

Something must separate two adjacent lists or it will confuse spacing.

* L1
* L2
	* L2-a sub lists are indented, still just one star
	* L2-b same as above
* L3 unindent gets back to previous list level 

-----

<!-- dashes or something for mandatory separator -->
* one  ( star begins)
+ two  ( plus sign begins)
- three( minus sign begins)  -- all three interchangeable

-----

<!-- dashes or something for mandatory separator -->
1. one  - numbered lists begin with numbers and period
1. two - begins with number 1.
1. three - also begins with number 1.

-----

> This is a block (starts with `>`)

> quote.  It will be indented
> but  internal    white   spacing will NOT be preserved
> (Other than newlines that is)


Block quotes are *indented* <- note italics from single stars.
Or even stronger emphasis with **bold** text that's surrounded by double stars.

Single underscores can be used to add _italic emphasis_ or __bold emphasis__ with 
double underscores.

Backticks can make code containing `< & > or other punctuation` symbols more readable.  Text inside
backticks is rendered in a non-serif mono-spaced font (at least on some systems it uses the `<code>` tags).

Strikeouts ~~can be~~ are done with double tildes around the stricken text.

-----

Mail addresses work as mailto links if surrounded by angle brackets `<hotei1352@gmail.com>` 
renders as <hotei1352@gmail.com>

Inline links can be  automatic when surrounded by `< and >` <http://tip.golang.org/>
or you can use parentheses immediately after the link text. The anchor (underlined part) should be
enclosed in `[brackets]` and the href in `(parens)` <br> `[the golang website](http://tip.golang.org/ "OptionalRollOverTextGoesHere")`
This is a link to [the golang website](http://tip.golang.org/ "OptionalRollOverTextGoesHere").

***

In the example below note how the anchors "The Google Search
Engine" and "Google" both refer to index [12] which maps to the final URL.

I get 10 times more traffic from [The Google Search Engine][12] than from
[Yahoo][22] or [MSN][130].  And I use [Google][12] more than any other search engine.

[12]: http://google.com/        "Google"
[22]: http://search.yahoo.com/  "Yahoo Search"
[130]: http://search.msn.com/    "MSN Search"

The (elsewhere located) reference list looks like this [index] [URL] [alt-text] :

>`[12]: http://google.com/        "Google"`

>`[22]: http://search.yahoo.com/  "Yahoo Search"`

>`[130]: http://search.msn.com/    "MSN Search"`

***

Images can be inlined easily with syntax similar to links... but there are no resizing options.<br>
`![alt text](/www/images/A-10_Thunderbolt_II_In-flight-2.jpg "A-10 WartHog")`

![alt text](/www/images/A-10_Thunderbolt_II_In-flight-2.jpg "A-10 WartHog")


***

blackfriday
===========

Note: blackfriday is the markdown to html translator writen in go by Russ Ross.
It can be downloaded from `<http://github.com/russross/blackfriday>`
<http://github.com/russross/blackfriday>

Fractions like 4/5 are rendered using super/subscript

Strikethrough using 2 tildes to bracket ~~unwanted text~~

This "unadorned URL" became an automatic link `-->` https://github.com .

<font color=red>blackfriday.MarkdownCommon() did not render a table the way I 
expected.  Was expecting to see it like it shows on blackfriday's github page</font>

The table example below rendered as text without any table borders.  The page source
shows no `<table></table>` tags.  The lack of borders is most probably a bug.

-----------------
| Name    | Age |
|---------|------
| Bob     | 27  |
| Alice   | 23  |
-----------------

> Copyright (c) 2012-2013 David Rook. All rights reserved.
 
> Redistribution and use in source and binary forms, with or without modification, are
> permitted provided that the following conditions are met:
> 
>    1. Redistributions of source code must retain the above copyright notice, this list of
>        conditions and the following disclaimer.
> 
>    2. Redistributions in binary form must reproduce the above copyright notice, this list
>        of conditions and the following disclaimer in the documentation and/or other materials
>        provided with the distribution.
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
