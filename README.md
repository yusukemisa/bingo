## bingo
`bingo` prints bingo cards with randomized numbers from 1 to 25.

## How to use
### Install
```
$ go get github.com/yusukemisa/bingo
```

### Run

```
$ bingo
 ------------------------
|  B |  I |  N |  G |  O |
| -- | -- | -- | -- | -- |
| 16 | 10 | 13 | 19 | 12 |
|  5 | 20 |  3 | 21 | 11 |
|  4 | 24 |    | 15 | 22 |
| 23 |  1 | 17 |  8 |  2 |
|  9 | 25 |  7 | 18 | 14 |
 ------------------------

```

## TODO
一般的なビンゴカードはB列に1 - 15, I列に16 - 30, N列に31 - 45, G列に46 - 60, O列に61 - 75の番号が割当てられる。
この仕様を満たすビンゴカードを`-r --regular`オプションを指定することで作成されるようにする。
