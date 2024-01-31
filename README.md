# jamanthi
A new programming language and its ecosystem

I have been intrigued by the inner workings of compilers for quite a long time. 

As I restarted by developer journey learning Golang, I figured that I could explore the compilers by implementing it.

This project will be an iteration as I go on designing individual components which makes up a compiler/ interpreter.

The various components involved in a programming language are
1.  The Tokens which makes up the Source Code
2.  The Lexical Analyser/ Tokenizer/ Scanner which breaks each word in a source into tokens
3.  The Parser, which builds a grammer from the output of the tokenizer.


The supported tokens are
()
{}
=
+
-
*
/
%
+=
-=
*=
/+
>
>=
<
<=
==
!=
strings
