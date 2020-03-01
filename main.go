package main

import (
	"archive/zip"
	"flag"
	"io"
	"log"
	"os"
	"sync"
)

var(
	seq = flag.Bool("seq", false, "seq zip")
	conc = flag.Bool("conc", false, "conc zip")
)

func main() {
	flag.Parse()
	if *seq {
		handleSeq(os.Args[2:])
		return
	}
	if *conc {
		handleConc(os.Args[2:])
		return
	}
}

func handleSeq(strings []string) {
	for _, string := range strings {
		ZipFile("seq/"+string+".zip", string)
	}
}

func handleConc(strings []string) {
	wg := sync.WaitGroup{}
	for _, str := range strings {
		wg.Add(1)
		go func(wg *sync.WaitGroup, str string) {
			defer wg.Done()
			ZipFile("conc/"+str+".zip", str)
		}(&wg, str)
	}
	wg.Wait()
	//fmt.Printf("done")
}

func ZipFile(filename string, file string) {

	newZipFile, err := os.Create(filename)
	if err != nil {
		log.Printf("can't create file %s err: %v", filename, err)
		return
	}
	defer func() {
		err = newZipFile.Close()
		if err != nil {
			log.Printf("can't close newZipFile err: %v", err)
		}
	}()

	zipWriter := zip.NewWriter(newZipFile)
	defer func() {
		err = zipWriter.Close()
		if err != nil {
			log.Printf("can't close zipWriter err: %v", err)
		}
	}()
	zipfile, err := os.Open(file)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	defer func() {
		err = zipfile.Close()
		if err != nil {
			log.Printf("can't close zipfile err: %v", err)
		}
	}()
	info, err := zipfile.Stat()
	if err != nil {
		log.Printf("can't get zipfile stat err: %v", err)
		return
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		log.Printf("can't get FileInfoHeader err: %v", err)
		return
	}
	header.Name = file

	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		log.Printf("can't create zipWriter err: %v", err)
		return
	}
	if _, err = io.Copy(writer, zipfile); err != nil {
		log.Printf("can't copy ti zipfile err: %v", err)
		return
	}
	//log.Println("Zipped File: " + filename)
}
