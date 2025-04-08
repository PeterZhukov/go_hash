## Сравнение результатов хеширования разных функций

words.txt, words_alpha.txt взяты из https://github.com/dwyl/english-words/tree/master

A text file containing over 466k English words.
* words.txt contains all words. 
* words_alpha.txt contains only [[:alpha:]] words (words that only have letters, no numbers or symbols). If you want a quick solution choose this.

### Результаты: `crc32.ChecksumIEEE(data)`

```text
Phone Numbers:
total= 50000000
bucket= 0  nums= 4998363
bucket= 1  nums= 5000154
bucket= 2  nums= 5000697
bucket= 3  nums= 5000691
bucket= 4  nums= 5002783
bucket= 5  nums= 4998486
bucket= 6  nums= 4999356
bucket= 7  nums= 4999224
bucket= 8  nums= 4998801
bucket= 9  nums= 5001445

Words:
total= 370105
bucket= 0  nums= 37025
bucket= 1  nums= 37174
bucket= 2  nums= 36873
bucket= 3  nums= 37026
bucket= 4  nums= 37375
bucket= 5  nums= 36851
bucket= 6  nums= 37057
bucket= 7  nums= 37059
bucket= 8  nums= 36868
bucket= 9  nums= 36797

Emails:
total= 3701050
bucket= 0  nums= 370077
bucket= 1  nums= 370703
bucket= 2  nums= 369841
bucket= 3  nums= 369848
bucket= 4  nums= 370746
bucket= 5  nums= 370000
bucket= 6  nums= 369193
bucket= 7  nums= 369906
bucket= 8  nums= 370596
bucket= 9  nums= 370140
```
### Результаты: `murmur3.Sum32(data)`
```text
Phone Numbers:
total= 50000000
bucket= 0  nums= 5000071
bucket= 1  nums= 5003289
bucket= 2  nums= 5002011
bucket= 3  nums= 5001245
bucket= 4  nums= 5001642
bucket= 5  nums= 4997175
bucket= 6  nums= 4997554
bucket= 7  nums= 4997353
bucket= 8  nums= 4999585
bucket= 9  nums= 5000075

Words:
total= 370105
bucket= 0  nums= 36989
bucket= 1  nums= 36958
bucket= 2  nums= 37255
bucket= 3  nums= 37189
bucket= 4  nums= 36848
bucket= 5  nums= 37018
bucket= 6  nums= 37173
bucket= 7  nums= 36703
bucket= 8  nums= 36957
bucket= 9  nums= 37015

Emails:
total= 3701050
bucket= 0  nums= 370481
bucket= 1  nums= 370989
bucket= 2  nums= 370275
bucket= 3  nums= 370380
bucket= 4  nums= 369511
bucket= 5  nums= 370222
bucket= 6  nums= 369635
bucket= 7  nums= 369422
bucket= 8  nums= 369389
bucket= 9  nums= 370746
```