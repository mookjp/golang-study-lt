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

type rootInfo struct {
	root   string
	nfiles int64
	nbytes int64
}

type fileInfo struct {
	root string
	size int64
}

func main() {

	// 対象ディレクトリ名取得
	flag.Parse()
	roots := flag.Args()
	// ディレクトリの指定がなかったら current directory とする
	if len(roots) == 0 {
		roots = []string{"."}
	}

	rootMap := make(map[string]*rootInfo)
	fileSizeChan := make(chan fileInfo)
	var n sync.WaitGroup
	for _, root := range roots {
		info := rootInfo{root, 0, 0}
		n.Add(1)
		// walkDir の実行
		go walkDir(root, root, &n, fileSizeChan)
		go func() {
			n.Wait()
			// TODO: panic: close of closed channel が発生する
			close(fileSizeChan)
		}()
		// マップに追加
		rootMap[root] = &info
	}

	// size
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	//var nfiles, nbytes int64
loop:
	for {
		select {
		// fileSizesからサイズを取得する
		case fileInfo, ok := <-fileSizeChan:
			// channel が閉じられていたらループを終了する
			if !ok {
				break loop
			}
			rootMap[fileInfo.root].nfiles = rootMap[fileInfo.root].nfiles + 1
			rootMap[fileInfo.root].nbytes = rootMap[fileInfo.root].nbytes + fileInfo.size
			fmt.Fprintf(os.Stdout, "verbose - %s: nfliles: %d, nbytes: %d\n", fileInfo.root, rootMap[fileInfo.root].nfiles, rootMap[fileInfo.root].nbytes)
		//500ms ごとに サイズを表示する
		case <-tick:
			for _, rootInfo := range rootMap {
				printDiskUsageForRoot(rootInfo.root, rootInfo.nfiles, rootInfo.nbytes)
			}
		}
	}
	for _, set := range rootMap {
		// 最終的なディスク利用量を表示する
		printDiskUsageForRoot(set.root, set.nfiles, set.nbytes)
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %d bytes\n", nfiles, nbytes)
}
func printDiskUsageForRoot(root string, nfiles, nbytes int64) {
	fmt.Printf("root %s: %d files %d bytes\n", root, nfiles, nbytes)
}

func walkDir(root, dir string, n *sync.WaitGroup, fileSizes chan<- fileInfo) {
	defer n.Done()

	for _, entry := range dirents(dir) {
		// ディレクトリの場合は waitGroup に1追加して
		// そのサブディレクトリを walkDir にわたす
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(root, subdir, n, fileSizes)
			// ファイルの場合はサイズをfileSizesに送信する
		} else {
			fileSizes <- fileInfo{root, entry.Size()}
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
