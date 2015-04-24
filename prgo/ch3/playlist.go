package main
import (
	"fmt"
	"strings"
	"os"
	"path/filepath"
	"strconv"
	"io/ioutil"
	"log"
)

type Song struct {
	Title string 
	Filename string
	Seconds int
}

func main() {
	// check xem co duoi la .m3u hay .pls khong neu khong thi phai dien ten duoi cho dung
	m3uExt := strings.HasSuffix(os.Args[1], ".m3u")
	plsExt := strings.HasSuffix(os.Args[1], ".pls")
	if len(os.Args) == 1 || (!m3uExt && !plsExt) {
		fmt.Printf("usage %s <file.m3u> or <file.pls>\n", filepath.Base(os.Args[0])) // tra ve ten cua file
		os.Exit(1);
	}
	// doc va phan loai file .m3u va .pls 
	// ioutil.Readfile doc file va tra ve toan bo gia byte cua file 
	if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else if m3uExt {
		songs := readM3uPlayList(string(rawBytes));
		writePlsPlayList(songs);
	} else if plsExt {
		songs := readPlsPlayList(string(rawBytes));
		writeM3uPlayList(songs)
	}

}

func readM3uPlayList(data string) (songs []Song){
	var song Song 
	for _, line := range strings.Split(data, "\n") { // strings.Split tach thah cac sau con 
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#EXTM3U") {
			continue // thuc hien vong lap tiep theo 
		}
		if strings.HasPrefix(line, "#EXTINF:") {
			song.Title, song.Seconds = parseExtinfLine(line)
		} else {
			song.Filename = strings.Map(mapPlatformDirSeparator, line); // tra ve mot copy cua string theo mapping function 
		}
		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song {}
		}
	}
	return songs; 
}
func parseExtinfLine(line string) (title string, seconds int) {
	if i := strings.IndexAny(line, "-0123456789"); i > -1 { // tra ve vi tri cua phan tu dau tien trong day so trong line
		const separator = ","
		line = line[i:]
		if j := strings.Index(line, separator); j > -1 { // lay vi tri separator trong line
			title = line[j+len(separator):]
			var err error 
			if seconds, err = strconv.Atoi(line[:j]); err != nil { // Atoi = ParseInt
				log.Printf("failed to read the duration for '%s': '%v\n'", title, err)
				seconds = -1
			}
		}
	}
	return title, seconds
}

func mapPlatformDirSeparator(char rune) rune {
	if char == '/' || char == '\\' {
		return filepath.Separator // tra ve path dung voi os dang dung
	}
	return char
}

func writePlsPlayList(songs []Song) {
	fmt.Println("[playlist]")
	for i, song := range songs {
		i++ 
		fmt.Printf("File%d=%s\n", i, song.Filename)
		fmt.Printf("Title%d=%s\n", i, song.Title)
		fmt.Printf("Length%d=%d\n", i, song.Seconds)
	}
	fmt.Printf("NumberOfEntries=%d\nVersion=2\n", len(songs))
}
func readPlsPlayList(data string) (songs []Song) {
	var song Song 
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "[playlist]") {
			continue
		}
		if j := strings.Index(line, "="); j != -1 {
			if strings.HasPrefix(line[:j], "File") {
				song.Filename = line[j+1:]
			} else if strings.HasPrefix(line[:j], "Title") {
				song.Title = line[j+1:]
			} else if strings.HasPrefix(line[:j], "Length") {
				if seconds, err := strconv.Atoi(line[j+1:]); err == nil {
					song.Seconds = seconds;
				}
			}
		}
		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs
}
func writeM3uPlayList(songs []Song) {
	fmt.Println("#EXTM3U")
	for _, song := range songs {
		fmt.Printf("#EXTINF:%d,%s\n", song.Seconds, song.Title)
		fmt.Println(song.Filename);
		// fmt.Printf("%s\n", song.Filename)
	} 
}