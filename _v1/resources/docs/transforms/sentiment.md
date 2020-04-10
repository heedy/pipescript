Sentiment is a very basic assesment of a text's sentiment from -1 to 1. It uses the [AFINN-111](http://www2.imm.dtu.dk/pubdb/views/publication_details.php?id=6010) wordlist,
performing [sentiment analysis](https://en.wikipedia.org/wiki/Sentiment_analysis) to determine whether the text is positive or negative:

The code is based upon [github.com/thisandagain/sentiment](https://github.com/thisandagain/sentiment).


```json
["I hate this!"]
```

Using the sentiment transform:
```
sentiment
```

```json
[-1]
```

Negative values represent negative sentiment, whereas positive values represent positive sentiment.

This is a *very* basic transform, for use in quick visualization of text data. It is not meant to be used as a reliable indicator of sentiment.
