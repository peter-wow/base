package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// 练习 7.8： 很多图形界面提供了一个有状态的多重排序表格插件：
// 主要的排序键是最近一次点击过列头的列，
// 第二个排序键是第二最近点击过列头的列，等等。
// 定义一个sort.Interface的实现用在这样的表格中。
// 比较这个实现方式和重复使用sort.Stable来排序的方式。

func main() {
	printTracks(tracks)

	click("title")
	printTracks(tracks)

	click("artist")
	printTracks(tracks)

	click("album")
	printTracks(tracks)

	click("year")
	printTracks(tracks)

	click("length")
	printTracks(tracks)
}

// 模拟点击事件
func click(t string) {
	switch t {
	case "title":
		sort.Stable(custom{tracks,
			func(x, y *Track) bool {
				return x.Title < y.Title
			},
			func(x, y *Track) {
				x.Title, y.Title = y.Title, x.Title
			}})
	case "artist":
		sort.Stable(custom{tracks,
			func(x, y *Track) bool {
				return x.Artist < y.Artist
			},
			func(x, y *Track) {
				x.Artist, y.Artist = y.Artist, x.Artist
			}})
	case "album":
		sort.Stable(custom{tracks,
			func(x, y *Track) bool {
				return x.Album < y.Album
			},
			func(x, y *Track) {
				x.Album, y.Album = y.Album, x.Album
			}})
	case "year":
		sort.Stable(custom{tracks,
			func(x, y *Track) bool {
				return x.Year < y.Year
			},
			func(x, y *Track) {
				x.Year, y.Year = y.Year, x.Year
			}})
	case "length":
		sort.Stable(custom{tracks,
			func(x, y *Track) bool {
				return int64(x.Length) < int64(y.Length)
			},
			func(x, y *Track) {
				x.Length, y.Length = y.Length, x.Length
			}})
	}
}

type custom struct {
	t    []*Track
	less func(x, y *Track) bool
	swap func(x, y *Track)
}

// 实现sort接口
// ? 为什么要去实现接口, 自己写不也可以吗
func (x custom) Len() int           { return len(x.t) }
func (x custom) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x custom) Swap(i, j int)      { x.swap(x.t[i], x.t[j]) }

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}
