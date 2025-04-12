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
total= 6661890
bucket= 0  nums= 333763
bucket= 1  nums= 333006
bucket= 2  nums= 1444062
bucket= 3  nums= 332502
bucket= 4  nums= 1442503
bucket= 5  nums= 333521
bucket= 6  nums= 332600
bucket= 7  nums= 1443145
bucket= 8  nums= 332722
bucket= 9  nums= 334066
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
total= 6661890
bucket= 0  nums= 332801
bucket= 1  nums= 333517
bucket= 2  nums= 1442861
bucket= 3  nums= 333017
bucket= 4  nums= 1443168
bucket= 5  nums= 333459
bucket= 6  nums= 332964
bucket= 7  nums= 1442645
bucket= 8  nums= 333729
bucket= 9  nums= 333729
```