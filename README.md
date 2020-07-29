# Search algorithms in Go

This repository shows an example implementation of the following graph search algorithms.

- BFS (Breadth-First Search)
- IDDFS (Iterative Deepening Depth-First Search)
- A\* (A-star Search)
- IDA\* (Iterative Deepening A-star Search)
- BFS (Best First Search)

## Usage

You can implement your own search problem as `base.Board` and `base.Operation`,
and just replace it in `main.go` to get it going.

## Example

An example problem (16-puzzle) is implemented in `./_16puzzle`.

References
- Wikipedia, "15 puzzle", Retrieved on 2020/07/29, https://en.wikipedia.org/wiki/15_puzzle
- Wikipedia, "Iterative Deepening A*", https://en.wikipedia.org/wiki/Iterative_deepening_A*
- takaken, "１５パズル自動解答プログラムの作り方", コンピュータ＆パズル, Published on 2001/04/02, Retrieved on 2020/07/07, http://www.ic-net.or.jp/home/takaken/nt/slide/solve15.html
