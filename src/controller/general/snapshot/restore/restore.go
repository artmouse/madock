package restore

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"github.com/faradey/madock/src/controller/general/rebuild"
	"github.com/faradey/madock/src/helper/configs"
	"github.com/faradey/madock/src/helper/docker"
	"github.com/faradey/madock/src/helper/logger"
	"github.com/faradey/madock/src/helper/paths"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func Execute() {
	projectName := configs.GetProjectName()
	projectConf := configs.GetCurrentProjectConfig()

	dbsPath := paths.GetExecDirPath() + "/projects/" + projectName + "/backup/snapshot"
	var snapshotNames []string
	if paths.IsFileExist(dbsPath) {
		snapshotNames = paths.GetDirs(dbsPath)
		if len(snapshotNames) == 0 {
			fmt.Println("No snapshots")
		}
		for index, snapshotName := range snapshotNames {
			fmt.Println(strconv.Itoa(index+1) + ") " + filepath.Base(snapshotName))
		}
	}

	fmt.Println("Choose one of the offered variants")
	buf := bufio.NewReader(os.Stdin)
	sentence, err := buf.ReadBytes('\n')
	selected := strings.TrimSpace(string(sentence))
	selectedInt := 0
	if err != nil {
		logger.Fatalln(err)
	} else {
		selectedInt, err = strconv.Atoi(selected)

		if err != nil || selectedInt > len(snapshotNames) {
			logger.Fatal("The item you selected was not found")
		}
	}

	if projectConf["platform"] != "pwa" {
		selectedFile, err := os.Open(dbsPath + "/" + snapshotNames[selectedInt-1] + "/db.tar.gz")
		if err != nil {
			logger.Fatal(err)
		}
		defer selectedFile.Close()

		containerName := docker.GetContainerName(projectConf, projectName, "snapshot")
		docker.Down(false)
		docker.UpSnapshot(projectName)

		cmd := exec.Command("docker", "exec", "-i", "-u", "root", containerName, "bash", "-c", "rm -rf /var/www/mysql/* && cd /var/www && tar -zxf -")
		out, err := gzip.NewReader(selectedFile)
		if err != nil {
			logger.Fatal(err)
		}
		cmd.Stdin = out
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			logger.Fatal(err, containerName)
		}
		docker.StopSnapshot(projectName)
		rebuild.Execute()
	}
	fmt.Println("Snapshot restored successfully")
}
