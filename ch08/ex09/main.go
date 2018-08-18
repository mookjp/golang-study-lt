package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

func main() {

	// 対象ディレクトリ名取得
	flag.Parse()
	roots := flag.Args()
	// ディレクトリの指定がなかったら current directory とする
	if len(roots) == 0 {
		roots = []string{"."}
	}

	for _, root := range roots {
		// root ごとの root, fileSize chan, sync.WaitGroup を作成
		fileSize := make(chan int64)
		var n sync.WaitGroup
		n.Add(1)
		// walkDir の実行
		go walkDir(root, &n, fileSize)
		// root ごとの 終了監視
		go func() {
			n.Wait()
			close(fileSize)
		}()

		// fileSize
		go func() {
			var tick <-chan time.Time
			if *vFlag {
				tick = time.Tick(500 * time.Millisecond)
			}
			var nfiles, nbytes int64
		loop:
			for {
				select {
				// fileSizesからサイズを取得する
				case size, ok := <-fileSize:
					// channel が閉じられていたらループを終了する
					if !ok {
						break loop
					}
					nfiles++
					nbytes += size
					// 500ms ごとに サイズを表示する
				case <-tick:
					printDiskUsageForRoot(root, nfiles, nbytes)
				}
			}
			// 最終的なディスク利用量を表示する
			printDiskUsageForRoot(root, nfiles, nbytes)
		}()
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %d bytes\n", nfiles, nbytes)
}
func printDiskUsageForRoot(root string, nfiles, nbytes int64) {
	fmt.Printf("root %s: %d files %d bytes\n", root, nfiles, nbytes)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()

	for _, entry := range dirents(dir) {
		// ディレクトリの場合は waitGroup に1追加して
		// そのサブディレクトリを walkDir にわたす
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
			// ファイルの場合はサイズをfileSizesに送信する
		} else {
			fileSizes <- entry.Size()
		}

	}
}

// 並列処理の対象資源のうち現在利用可能な数を表す
var sema = make(chan struct{}, 20)

// dirents は、ディレクトリ中のファイルエントリを返す
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // セマフォを加算する
	defer func() { <-sema }() // dirents 終了時にセマフォから減算する

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
