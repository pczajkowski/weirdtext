weirdtext
=========

WeirdText encoding and decoding
*******************************

WeirdText is a text encoding.

It is not "encryption" because humans can usually read it quite easily.  But machines may find it difficult to read without the list of original words.  Except of having fun, there are real-world applications for this, e.g. if encryption is forbidden by law in your country, but you still don't want your email content to get automatically processed somehow.

Encoding
========

For each original word in the original text, leave the first and last character of it in that position, but shuffle (permutate) all the characters in the middle of the word.  If possible, the resulting "encoded" word MUST NOT be the same as the original word. Keep everything else (whitespace, punctuation, etc.) like in the original.  To make decoding by a machine possible, your encoder shall also output a sorted list of original words (only include words that got shuffled, not text that did not).

The composite output of the encoder (see example below) contains encoded text (WeirdText) and also the sorted list of original words.

Decoding
========

For decoding composite text, first do a simple check whether the text looks like composite output of your encoder. If not, raise some reasonable exception.

Then, use the encoded text and the words list to decode the text.

Your decoded output should, as far as possible, be identical to the original text. In case of ambiguities (some encoded word could have been multiple original words), decoding errors are acceptable.

Example
=======

Original Text (this is a single string formatted nicely for better viewing!)::
'This is a long looong test sentence,\n'
'with some big (biiiiig) words!'

Encoded Text (see comment above)::
'\n---weird---\n'
'Tihs is a lnog loonog tset sntceene,\n'
'wtih smoe big (biiiiig) wdros!'
'\n---weird---\n'
'long looong sentence some test This with words'

Decoded Text::
'This is a long looong test sentence,\n'
'with some big (biiiiig) words!'

Task
====

* implement encoder first, this is easy to medium difficulty
* you don't need to implement tests until after you have implemented (or tried to implement) the decoder
* decoder is medium to hard difficulty, implement that after encoder
* if you have the decoder, implement a reasonable amount of tests
* if you can't create the decoder and you have time left, implement some simple tests for the encoder (as far as possible).

Implementation hints
====================

You may find these hints/code fragments useful:
* separator at the start of encoded output shall be used as a "magic" value and checked by decoder
* separator is also used to separate the encoded text from the sorted original word list separator = '\n---weird---\n'
* `tokenize_re = re.compile(r'(\w+)', re.U)` find out for what exactly this is useful
* `import random` find out for what exactly this is useful

Scoring
=======

* coding style
* code docs/comments
* quality of code
* knowledge of python/stdlib
* working encoder
* working decoder
* working tests
