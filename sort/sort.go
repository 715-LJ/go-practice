package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type CustomSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

type byArtist []*Track

func main() {
	//排序
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	//反转
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	//按照标题、时间、长度排序
	sort.Sort(CustomSort{
		t: tracks,
		less: func(x, y *Track) bool {
			if x.Title != y.Title {
				return x.Title < y.Title
			}
			if x.Year != y.Year {
				return x.Year < y.Year
			}
			if x.Length != y.Length {
				return x.Length < y.Length
			}
			return false
		},
	})
	printTracks(tracks)

	//验证是否正常排序
	vs := []int{1, 2, 3, 4, 5}
	fmt.Println(sort.IntsAreSorted(vs))

	sort.Sort(sort.Reverse(sort.IntSlice(vs)))
	fmt.Println(sort.IntsAreSorted(vs))
}

func printTracks(tracks []*Track) {
	fmt.Printf("\n\n")
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (c CustomSort) Len() int           { return len(c.t) }
func (c CustomSort) Less(i, j int) bool { return c.less(c.t[i], c.t[j]) }
func (c CustomSort) Swap(i, j int)      { c.t[i], c.t[j] = c.t[j], c.t[i] }
