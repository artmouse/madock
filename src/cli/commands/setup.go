package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/faradey/madock/src/cli/attr"
	"github.com/faradey/madock/src/cli/fmtc"
	"github.com/faradey/madock/src/configs"
	"github.com/faradey/madock/src/docker/builder"
	"github.com/faradey/madock/src/docker/scripts"
	"github.com/faradey/madock/src/paths"
	"github.com/faradey/madock/src/versions"
)

func Setup() {
	projectName := configs.GetProjectName()
	hasConfig := configs.IsHasConfig(projectName)
	continueSetup := true
	if hasConfig {
		fmtc.WarningLn("File env is already exist in project " + projectName)
		fmt.Println("Do you want to continue? (y/N)")
		fmt.Print("> ")

		buf := bufio.NewReader(os.Stdin)
		sentence, err := buf.ReadBytes('\n')
		selected := strings.TrimSpace(string(sentence))
		if err != nil {
			log.Fatal(err)
		} else if selected != "y" {
			if !attr.Options.Download && !attr.Options.Install {
				log.Fatal("Exit")
			}
			continueSetup = false
		}
	}

	if strings.Contains(projectName, ".") || strings.Contains(projectName, " ") {
		fmtc.ErrorLn("The project folder name cannot contain a period or space")
		return
	}

	fmtc.SuccessLn("Start set up environment")

	envFile := paths.MakeDirsByPath(paths.GetExecDirPath()+"/projects/"+projectName) + "/env.txt"
	var projectConfig map[string]string
	if _, err := os.Stat(envFile); !os.IsNotExist(err) {
		projectConfig = configs.GetProjectConfig(projectName)
	} else {
		projectConfig = configs.GetGeneralConfig()
	}

	toolsDefVersions := versions.GetVersions("")

	mageVersion := ""
	if toolsDefVersions.Php == "" {
		fmt.Println("")
		fmtc.Title("Specify Magento version: ")
		mageVersion, _ = waiter()
		if mageVersion != "" {
			toolsDefVersions = versions.GetVersions(mageVersion)
		}
	}

	if continueSetup {
		fmt.Println("")
		fmtc.Title("Your Magento version is " + toolsDefVersions.Magento)

		setupPhp(&toolsDefVersions.Php)
		setupDB(&toolsDefVersions.Db)
		setupComposer(&toolsDefVersions.Composer)
		setupSearchEngine(&toolsDefVersions.SearchEngine)
		if toolsDefVersions.SearchEngine == "Elasticsearch" {
			setupElastic(&toolsDefVersions.Elastic)
		} else {
			setupOpenSearch(&toolsDefVersions.OpenSearch)
		}

		setupRedis(&toolsDefVersions.Redis)
		setupRabbitMQ(&toolsDefVersions.RabbitMQ)
		setupHosts(projectName, &toolsDefVersions.Hosts, projectConfig)

		configs.SetEnvForProject(projectName, toolsDefVersions, projectConfig)
		paths.MakeDirsByPath(paths.GetExecDirPath() + "/projects/" + projectName + "/backup/db")

		fmtc.SuccessLn("\n" + "Finish set up environment")
		fmtc.ToDoLn("Optionally, you can configure SSH access to the development server in order ")
		fmtc.ToDoLn("to synchronize the database and media files. Enter SSH data in ")
		fmtc.ToDoLn(paths.GetExecDirPath() + "/projects/" + projectName + "/env.txt")
	}

	builder.Down(attr.Options.WithVolumes)
	builder.Start(attr.Options.WithChown)

	if attr.Options.Download {
		downloadMagento(projectName, mageVersion)
	}

	if attr.Options.Install {
		builder.InstallMagento(projectName, toolsDefVersions.Magento)
	}
}

func downloadMagento(projectName, mageVersion string) {
	fmt.Println("")
	fmtc.TitleLn("Specify Magento version: ")
	fmt.Println("1) Community (default)")
	fmt.Println("2) Enterprise")
	edition, _ := waiter()
	edition = strings.TrimSpace(edition)
	if edition != "1" && edition != "2" && edition != "" {
		fmtc.ErrorLn("The specified edition '" + edition + "' is incorrect.")
		return
	}
	if edition == "1" || edition == "" {
		edition = "community"
	} else if edition == "2" {
		edition = "enterprise"
	}
	builder.DownloadMagento(projectName, edition, mageVersion)
}

func InstallMagento() {
	projectName := configs.GetProjectName()
	vers := versions.GetVersions("")
	if vers.Magento == "" {
		fmtc.ErrorLn("File composer.json is not exist")
		return
	}
	builder.InstallMagento(projectName, vers.Magento)
}

func setupPhp(defVersion *string) {
	setTitleAndRecommended("PHP", defVersion)

	availableVersions := []string{"Custom", "8.2", "8.1", "8.0", "7.4", "7.3", "7.2", "7.1", "7.0"}

	prepareVersions(availableVersions)
	invitation(defVersion)
	waiterAndProceed(defVersion, availableVersions)
}

func setupDB(defVersion *string) {
	setTitleAndRecommended("DB", defVersion)

	availableVersions := []string{"Custom", "10.6", "10.4", "10.3", "10.2", "10.1", "10.0"}

	prepareVersions(availableVersions)
	invitation(defVersion)
	waiterAndProceed(defVersion, availableVersions)
}

func setupComposer(defVersion *string) {
	setTitleAndRecommended("Composer", defVersion)

	availableVersions := []string{"Custom", "1", "2"}

	prepareVersions(availableVersions)
	invitation(defVersion)
	waiterAndProceed(defVersion, availableVersions)
}

func setupSearchEngine(defVersion *string) {
	setTitleAndRecommended("Search Engine", defVersion)

	availableVersions := []string{"", "OpenSearch", "Elasticsearch", "Do not use"}

	prepareVersions(availableVersions)
	invitation(defVersion)
	waiterAndProceed(defVersion, availableVersions)
}

func setupElastic(defVersion *string) {
	setTitleAndRecommended("Elasticsearch", defVersion)

	availableVersions := []string{"Custom", "8.5.3", "8.4.3", "7.17.5", "7.16.3", "7.10.1", "7.9.3", "7.7.1", "7.6.2", "6.8.20", "5.1.2"}

	prepareVersions(availableVersions)
	invitation(defVersion)
	waiterAndProceed(defVersion, availableVersions)
}

func setupOpenSearch(defVersion *string) {
	setTitleAndRecommended("OpenSearch", defVersion)

	availableVersions := []string{"Custom", "2.5.0", "1.2.0"}

	prepareVersions(availableVersions)
	invitation(defVersion)
	waiterAndProceed(defVersion, availableVersions)
}

func setupRedis(defVersion *string) {
	setTitleAndRecommended("Redis", defVersion)

	availableVersionsShort := []string{"Custom", "7.0", "6.2", "6.0", "5.0"}
	availableVersions := []string{"Custom", "7.0 (Magento version > 2.4.6)", "6.2 (Magento version > 2.4.3-p3)", "6.0 (Magento version <= 2.4.3-p3)", "5.0 (Magento version <= 2.3.2)"}

	prepareVersions(availableVersions)
	invitation(defVersion)
	waiterAndProceed(defVersion, availableVersionsShort)
}

func setupRabbitMQ(defVersion *string) {
	setTitleAndRecommended("RabbitMQ", defVersion)

	availableVersionsShort := []string{"Custom", "3.9", "3.8", "3.7"}
	availableVersions := []string{"Custom", "3.9 (Magento version > 2.4.3-p3)", "3.8 (Magento version <= 2.4.3-p3)", "3.7 (Magento version <= 2.3.4)"}

	prepareVersions(availableVersions)
	invitation(defVersion)
	waiterAndProceed(defVersion, availableVersionsShort)
}

func setupHosts(projectName string, defVersion *string, projectConfig map[string]string) {
	host := strings.ToLower(projectName + projectConfig["DEFAULT_HOST_FIRST_LEVEL"] + ":base")
	if val, ok := projectConfig["HOSTS"]; ok && val != "" {
		host = val
	}
	fmtc.TitleLn("Hosts")
	fmt.Println("Input format: a.example.com:x_website_code b.example.com:y_website_code")
	fmt.Println("Recommended host: " + host)
	availableVersions := []string{projectName + projectConfig["DEFAULT_HOST_FIRST_LEVEL"] + ":base", "loc." + projectName + ".com:base"}
	prepareVersions(availableVersions)
	fmt.Println("Choose one of the suggested options or enter your hostname")
	fmt.Print("> ")
	selected, _ := waiter()
	if selected == "" && host != "" {
		*defVersion = host
		fmtc.SuccessLn("Your choice: " + *defVersion)
	} else if selected != "" {
		selectedInt, err := strconv.Atoi(selected)
		if err == nil && len(availableVersions) >= selectedInt {
			*defVersion = availableVersions[selectedInt-1]
		} else {
			*defVersion = selected
		}
		fmtc.SuccessLn("Your choice: " + *defVersion)
	}
}

func prepareVersions(availableVersions []string) {
	for index, ver := range availableVersions {
		if ver != "" {
			fmt.Println(strconv.Itoa(index) + ") " + ver)
		}
	}
}

func setTitleAndRecommended(title string, recommended *string) {
	fmt.Println("")
	fmtc.TitleLn(title)
	if *recommended != "" {
		fmt.Println("Recommended version: " + *recommended)
	}
}

func invitation(ver *string) {
	if *ver != "" {
		fmt.Println("Enter the item number or press Enter to select the recommended version")
	} else {
		fmt.Println("Enter the item number")
	}

	fmt.Print("> ")
}

func waiterAndProceed(defVersion *string, availableVersions []string) {
	selected, repoAndVersion := waiter()
	setSelectedVersion(defVersion, availableVersions, selected, repoAndVersion)
}

func waiter() (selected, repoAndVersion string) {
	buf := bufio.NewReader(os.Stdin)
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		log.Fatalln(err)
	}
	selected = strings.TrimSpace(string(sentence))
	if selected == "0" {
		fmt.Println("Enter the version")
		fmt.Print("> ")
		buf = bufio.NewReader(os.Stdin)
		sentence, err = buf.ReadBytes('\n')
		if err != nil {
			log.Fatalln(err)
		}
		repoAndVersion = strings.TrimSpace(string(sentence))
	}

	return
}

func setSelectedVersion(defVersion *string, availableVersions []string, selected, repoAndVersion string) {
	selectedInt, err := strconv.Atoi(selected)
	if selected == "" && *defVersion != "" {
		fmtc.SuccessLn("Your choice: " + *defVersion)
	} else if selected == "0" {
		*defVersion = repoAndVersion
	} else if selected != "" && err == nil && len(availableVersions) >= selectedInt {
		*defVersion = availableVersions[selectedInt]
		fmtc.SuccessLn("Your choice: " + *defVersion)
	} else {
		fmtc.WarningLn("Choose one of the options offered")
		waiterAndProceed(defVersion, availableVersions)
	}
}

func copyFile(pathFrom, pathTo string) {
	b, err := os.ReadFile(pathFrom)
	if err != nil {
		log.Fatal(err)
	}
	pathToAsSlice := strings.Split(pathTo, "/")
	paths.MakeDirsByPath(strings.Join(pathToAsSlice[:len(pathToAsSlice)-1], "/"))
	err = os.WriteFile(pathTo, b, 0755)
	if err != nil {
		log.Fatalf("Unable to write file: %v", err)
	}
}

func SetupEnv() {
	envFile := paths.GetRunDirPath() + "/app/etc/env.php"
	if _, err := os.Stat(envFile); !os.IsNotExist(err) && !attr.Options.Force {
		log.Fatal("The env.php file is already exist.")
	} else {
		data, err := json.Marshal(configs.GetCurrentProjectConfig())
		if err != nil {
			log.Fatal(err)
		}
		scripts.CreateEnv(string(data), attr.Options.Host)
	}
}
