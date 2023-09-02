package GetFolders

import (
	"fmt"
	"os"
	"path/filepath"
)

func List(root string) []string {
	var folders []string
	//folders = append(folders, root)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			folders = append(folders, path)
		}
		//fmt.Printf("Path: %s, Size: %d bytes\n", path, info.Size())
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return folders
}
