package spotlight

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func Start() {
	path := "/Users/Zen/AppData/Local/Packages/Microsoft.Windows.ContentDeliveryManager_cw5n1h2txyewy/LocalState/Assets"
	GetSpotlight(path)
}
func GetSpotlight(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		println(file.Name())
		out, err := os.Create("./temp/spotlight" + file.Name() + ".jpg")
		defer out.Close()
		if err != nil {
			log.Fatal(err)
		}
		in, err := os.Open(path + "/" + file.Name())
		defer in.Close()
		if err != nil {
			log.Fatal(err)
		}

		io.Copy(out, in)
	}

}
