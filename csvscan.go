/*
Copyright Mike Hughes 2013 (intermernet AT gmail DOT com)

csvscan is a utility for extracting fields of CSV files into separate files.

LICENSE: BSD 3-Clause License (see http://opensource.org/licenses/BSD-3-Clause)
*/

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	in   string
	out  string
	pre  string
	suf  string
	i    int
	ni   int
	frag bool
)

const (
	ps      = os.PathSeparator
	example = "\nUsage example: \"csvscan -in=in.csv -out=outdirectory -pre=Invoice- -suf=-2013.txt -i=2 -ni=1 -frag=true\""
)

func init() {
	flag.StringVar(&in, "in", "", "CSV file to import (Must be supplied!)")
	flag.StringVar(&out, "out", ".", "Output directory (Defaults to the current directory)")
	flag.StringVar(&pre, "pre", "", "File prefix (defaults to nothing)")
	flag.StringVar(&suf, "suf", ".txt", "File suffix and extension (defaults to \".txt\")")
	flag.IntVar(&i, "i", 1, "Index of field to export (defaults to \"1\")")
	flag.IntVar(&ni, "ni", 0, "Name Index (field to name files by - must be unique and filename safe!, defaults to line number)")
	flag.BoolVar(&frag, "frag", false, "Overwrite files (defaults to \"false\")")
}

func writeFile(name string, s string, frag bool, c chan error) {
	var ow = os.O_EXCL
	if frag == true {
		ow = os.O_TRUNC
	}
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|ow, 0666)
	if err != nil {
		c <- err
	}
	defer f.Close()
	_, err = f.WriteString(s)
	if err != nil {
		c <- err
	}
	return
}

func main() {
	usage := flag.Usage
	flag.Usage = func() {
		usage()
		fmt.Println(example)
	}
	flag.Parse()
	if in == "" {
		fmt.Print("\nYou must specify an input file\n\"csvscan -in=somefile.csv\"\n")
		flag.PrintDefaults()
		fmt.Println(example)

		return
	}
	ps := string(ps)
	if !strings.HasSuffix(out, ps) {
		out = out + ps
	}
	file, err := os.Open(in)
	if err != nil {
		log.Printf("\nError opening file:\n%s", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	c := make(chan error)
	n := 1
	for {
		record, err := reader.Read()
		switch {
		case err == io.EOF:
			break
		case err != nil:
			log.Printf("\nCSV Decoding error on line %d: %s\n\n", n, err)
			n++
		case ni != 0:
			go writeFile(out+pre+record[ni]+suf, record[i], frag, c)
			select {
			case err := <-c:
				log.Printf("\nError writing to file on line %d\n %s\n\n", n, err)
				n++
			default:
				n++
			}
		default:
			go writeFile(out+pre+strconv.Itoa(i)+suf, record[i], frag, c)
			select {
			case err := <-c:
				log.Printf("\nError writing to file on line %d\n %s\n\n", n, err)
				n++
			default:
				n++
			}
		}
	}
}
