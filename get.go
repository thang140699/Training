package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	trTags = regexp.MustCompile(`<tr[^>]*>(?:.|\n)*<\/tr>`)
	tagA   = regexp.MustCompile(`<\s*a[^>]*>(.*?)<\s*/\s*a>`)
)

func getStringURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	// params := resp.Request.FormValue("")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	StringURL := string(body)
	return StringURL
}

func getTagTrElement(url string) string {
	divs := trTags.FindAllStringSubmatch(url, -1)
	return divs[0][0]
}

func getTagAElement(fatherElement string) string {
	trTag := getTagTrElement(fatherElement)
	tags := tagA.FindAllStringSubmatch(trTag, -1)
	return tags[0][0]
}

func getLinkDomain(URL string) string {
	StringURL := getStringURL(URL)
	tags := getTagAElement(StringURL)
	links := strings.Split(tags, "\"")
	IndexIncludeLinkOnHref := 1
	return links[IndexIncludeLinkOnHref]
}
func rename(Original_path string, fileName string) {

	New_path := Original_path + "/Domain.txt"
	r := os.Rename(fileName, New_path)
	if r != nil {
		log.Fatal(r)
	}
}
func Unzip(pathFolder string) {
	dst := pathFolder
	archive, err := zip.OpenReader(pathFolder + "/test.zip")
	fmt.Println()
	if err != nil {
		log.Fatalln("z : ", err)
	}
	defer archive.Close()
	fmt.Println(archive.File)
	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)
		fmt.Println("unzipping file ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
			return
		}
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}
		rename(pathFolder, filePath)
		dstFile.Close()
		fileInArchive.Close()
	}
}
func handleTime() string {
	year, month, day := time.Now().Year(), time.Now().Month(), time.Now().Day()
	ymd := strconv.Itoa(year) + "/" + month.String() + "/" + strconv.Itoa(day)
	err := os.MkdirAll(ymd, 0755)
	if err != nil {
		fmt.Println("a")
	}
	return ymd
}

func main() {
	const URL = "https://www.whoisds.com/newly-registered-domains"
	link := getLinkDomain(URL)
	fmt.Println("Link: ", link)
	res, err := http.Get(link)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	//handleTime()
	pathFolder := handleTime()

	out, err := os.Create(pathFolder + "/test.zip")
	if err != nil {
		log.Fatalln(err)
	}

	defer out.Close()
	_, err = io.Copy(out, res.Body)
	if err != nil {
		panic(err)
	}

	Unzip(pathFolder)
	e := os.Remove(pathFolder + "/test.zip")
	if e != nil {
		log.Fatal(e)
	}
}
