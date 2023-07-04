package helper

import (
	"archive/zip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

// Get operating system
func GetOperatingSystem() string {
	tmp := runtime.GOOS
	return tmp
}

// Manifest JSON structure
type Manifest struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	// ManifestVersion string `json:"manifest_version"`
}

type myCloser interface {
	Close() error
}

// closeFile is a helper function which streamlines closing
// with error checking on different file types.
func closeFile(f myCloser) {
	err := f.Close()
	check(err)
}

// readAll is a wrapper function for ioutil.ReadAll. It accepts a zip.File as
// its parameter, opens it, reads its content and returns it as a byte slice.
func readAll(file *zip.File) []byte {
	fc, err := file.Open()
	check(err)
	defer closeFile(fc)

	content, err := ioutil.ReadAll(fc)
	check(err)

	return content
}

// check is a helper function which streamlines error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Parse json
type jsonResult struct {
	name             string
	version          string
	description      string
	manifest_version string
}

func parseJsonFromContent(content []byte) jsonResult {

	var manifest Manifest
	json.Unmarshal(content, &manifest)

	result := jsonResult{
		name:        manifest.Name,
		version:     manifest.Version,
		description: manifest.Description,
		// manifest_version: manifest.ManifestVersion,
	}

	return result
}

// Parse JSON file from filename
func parseJsonFromFile(filename string) {
	fileContent, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var res map[string]interface{}
	json.Unmarshal([]byte(byteResult), &res)

}

// Google Chrome, Brave browsers
func FindChromeBrowserExtensions(to_save string, to_check_online bool, is_save_all bool) {
	// Directories to search
	BrowserPaths := GetBrowserPaths()
	rootFolder := BrowserPaths["chrome"]

	FindChromeBrowserExtensionsFromPath(rootFolder, to_save, to_check_online, false, false, is_save_all)
}

func FindChromeBrowserExtensionsFromPath(rootFolder string, to_save string, to_check_online bool, is_msedge bool, is_opera bool, is_save_all bool) {
	// Default extensions list
	DefaultExtensionsList := []string{
		"nmmhkkegccagdldgiimedpiccmgmieda", // Chrome Web Store Payments | 1.0.0.6"
		"enegjkbbakeegngfapepobipndnebkdk", // Rich Hints Agent | 1.0.9 - Opera
		"gojhcdgcpbpfigcaejpfhfegekdgiblk", // Opera Wallet | 1.14 - Opera
		"igpdmclhhlcpoindmhkhillbfhdgoegm", // Aria | 1.0.15 - Opera
		"jmjflgjpcpepeafmmgdpfkogkghcpiha", // Edge relevant text changes | 1.1.3 - MS Edge
	}

	// Init colors
	red := color.New(color.FgRed).SprintFunc()

	// Check rootFolder
	if _, err := os.Stat(rootFolder); os.IsNotExist(err) {
		fmt.Printf("%s\n", red("ERROR: "+rootFolder+" does not exist"))
		return
	}

	// Create an empty list
	mylist := make([]string, 0)

	// Start the recursive search
	err := filepath.Walk(rootFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return nil
		}

		// Check if the current item is a file named "manifest.json"
		if info.Mode().IsRegular() && info.Name() == "manifest.json" {
			// fmt.Printf("Found manifest.json at: %s\n", path)

			data, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Printf("Error reading file %s: %s\n", path, err)
				return nil
			}

			var manifest Manifest
			err = json.Unmarshal(data, &manifest)
			if err != nil {
				fmt.Printf("Error parsing JSON in file %s: %s\n", path, err)
				return nil
			}

			// Resolve the __MSG_xxxxxxxxx__
			appName := manifest.Name
			appDescription := manifest.Description

			if strings.Contains(manifest.Name, "__MSG_") {
				// Resolve messages
				has_messages := false

				messages_json_path := filepath.Dir(path) + "/_locales/en/messages.json"
				// fmt.Println(messages_json_path)

				if FileExists(messages_json_path) == false {
					messages_json_path = filepath.Dir(path) + "/_locales/en_US/messages.json"

					if FileExists(messages_json_path) == true {
						has_messages = true
					}
					// else {
					// 	has_messages = false
					// }

				} else {
					has_messages = true
				}

				if has_messages == true {
					// Removing string
					appName = strings.ReplaceAll(appName, "__MSG_", "")
					appName = strings.ReplaceAll(appName, "__", "")
					appName = strings.ToLower(appName)

					appDescription = strings.ReplaceAll(appDescription, "__MSG_", "")
					appDescription = strings.ReplaceAll(appDescription, "__", "")
					appDescription = strings.ToLower(appDescription)

					// fmt.Println("appName: " + appName)
					// fmt.Println("appDescription: " + appDescription)
					// fmt.Println("")

					if FileExists(messages_json_path) {
						b, err := os.ReadFile(messages_json_path)
						if err != nil {
							fmt.Print(err)
						}

						json_string := string(b)                   // convert content to a 'string'
						json_string = strings.ToLower(json_string) // Convert to lowercase

						// Resolve app_name
						app_name := gjson.Get(json_string, appName+".message").String()
						app_name = strings.Title(app_name)
						appName = app_name

						// Resolve app_name
						app_desc := gjson.Get(json_string, appDescription+".message").String()
						appDescription = app_desc

						// fmt.Println("=================")
						// fmt.Println(app_name)
						// fmt.Println("=================")

						// app_desc := gjson.Get(json_string, appDescription+".message")

						// fmt.Println("=================")
						// fmt.Println(app_desc.String())
						// fmt.Println("=================")

					} else {

					}

				}

			} else {
				appName = manifest.Name
			}

			// Encodes the app name and description
			appName = strings.ReplaceAll(appName, "\\\\", "\\")
			appDescription = strings.ReplaceAll(appDescription, "\n", "")

			// fmt.Println("==========================")
			// fmt.Printf("%s\n", appDescription)
			// fmt.Println("==========================")

			// Extract extension ID
			regex := regexp.MustCompile(`Extensions[\\/]([a-zA-Z0-9]+)[\\/]`)

			matches := regex.FindStringSubmatch(path)

			// Skip check online
			if to_check_online == false {
				if len(matches) > 1 {
					extensionID := matches[1]

					// Check if is default extension
					is_default_ext := CheckItemInList(DefaultExtensionsList, extensionID)

					if is_default_ext == false {
						fmt.Printf("%s | %s | %s | %s\n", extensionID, appName, manifest.Version, appDescription)

						line := fmt.Sprintf("%s | %s | %s | %s", extensionID, appName, manifest.Version, appDescription)
						mylist = append(mylist, line)

					}

				} else {
					// Do nothing

				}
			} else {
				// Check online
				if len(matches) > 1 {
					extensionID := matches[1]

					// Check if is default extension
					is_default_ext := CheckItemInList(DefaultExtensionsList, extensionID)

					if is_default_ext == false {
						online_result := 0

						if is_msedge == false {
							if is_opera == true {
								online_result = CheckOperaExtensionID(extensionID)
							} else {
								online_result = CheckBrowserExtensionID(extensionID)
							}
						} else {
							online_result = CheckMSEdgeExtensionID(extensionID)
						}

						// Print 200 for success
						if online_result == 200 {
							green := color.New(color.FgGreen).SprintFunc()
							fmt.Printf("%s | %s | %s | %s %s\n", extensionID, appName, manifest.Version, appDescription, green("200"))

							line := fmt.Sprintf("%s | %s | %s | %s | %s", extensionID, appName, manifest.Version, appDescription, "200")
							mylist = append(mylist, line)

						} else {
							// Print 404 for error
							red := color.New(color.FgRed).SprintFunc()
							fmt.Printf("%s | %s | %s | %s %s\n", extensionID, appName, manifest.Version, appDescription, red("404"))

							line := fmt.Sprintf("%s | %s | %s | %s | %s", extensionID, appName, manifest.Version, appDescription, "404")
							mylist = append(mylist, line)
						}
					}

				} else {
					// Do nothing

					// extensionID := "none"
					// fmt.Printf("%s | %s | %s | %s\n", extensionID, appName, manifest.Version, appDescription)
				}
			}

		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", rootFolder, err)
	}

	// Save to file
	if to_save != "" {

		// is_save_all
		// fmt.Println(is_save_all)

		if is_save_all == true {

			// fmt.Println("Append")

			f, err := os.OpenFile(to_save, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			writer := csv.NewWriter(f)
			defer writer.Flush()

			for _, item := range mylist {
				err := writer.Write([]string{item})

				if err != nil {
					fmt.Println("Error writing to file:", err)
					return
				}
			}
			// writer.Flush()

		} else {
			file, err := os.Create(to_save)

			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			defer file.Close()

			writer := csv.NewWriter(file)
			defer writer.Flush()

			for _, item := range mylist {
				err := writer.Write([]string{item})

				if err != nil {
					fmt.Println("Error writing to file:", err)
					return
				}
			}

			fmt.Printf("\nOutput is saved to %s!\n", to_save)
		}

	}
}

// FireFoxAddon struct which contains meta data of plugin
// type FireFoxAddonMeta struct {
// 	id          string `json:"id"`
// 	name        string `json:"name"`
// 	description string `json:"description"`
// 	version     string `json:"version"`
// 	sourceURI   string `json:"sourceURI"`
// }

// defining a map
var result map[string]interface{}

// FireFox
func FindFireFoxExtensions(to_save string, to_check_online bool, is_save_all bool) {
	BrowserPaths := GetBrowserPaths()
	rootFolder := BrowserPaths["firefox"]

	FindFireFoxExtensionsFromPath(rootFolder, to_save, to_check_online, is_save_all)
}

func FindFireFoxExtensionsFromPath(rootFolder string, to_save string, to_check_online bool, is_save_all bool) {
	// Default plugins list
	DefaultPluginsList := []string{
		"langpack-en-GB@firefox.mozilla.org", // English (GB) Language Pack | 115.0.20230629.134642 - FireFox
		"langpack-en-CA@firefox.mozilla.org", // English (CA) Language Pack | 115.0.20230629.134642 - FireFox
	}

	// Init colors
	red := color.New(color.FgRed).SprintFunc()

	// Check rootFolder
	if _, err := os.Stat(rootFolder); os.IsNotExist(err) {
		fmt.Printf("%s\n", red("ERROR: "+rootFolder+" does not exist"))
		return
	}

	// Create an empty list
	mylist := make([]string, 0)

	// fmt.Println("Plugin Name, Version, Description")

	// Start the recursive search
	err := filepath.Walk(rootFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return nil
		}

		// Check if the current item is a file named "manifest.json"
		filename := info.Name()
		// fmt.Printf("Filename: " + filename + "\n")

		if (filename != "") && (strings.Contains(filename, "addons.json")) {
			// fmt.Printf("Filename: " + filename + "\n")
			// fmt.Printf("Filepath: " + path + "\n")
			b, err := os.ReadFile(path)
			if err != nil {
				fmt.Print(err)
			}

			// fmt.Println(b) // print the content as 'bytes'

			json_string := string(b) // convert content to a 'string'

			addons := gjson.Get(json_string, "addons")

			addons.ForEach(func(key, value gjson.Result) bool {

				// fmt.Println(value.String())

				err := json.Unmarshal([]byte(value.String()), &result)

				if err != nil {
					fmt.Println(err)
				}

				id := result["id"]
				name := result["name"]
				// description := result["description"]
				version := result["version"]
				sourceURI := fmt.Sprintf("%s", result["sourceURI"])

				// Check if is default plugin
				id_str := fmt.Sprintf("%s", id)
				is_default_plugin := CheckItemInList(DefaultPluginsList, id_str)

				if is_default_plugin == false {
					if to_check_online == true {

						online_result := CheckFireFoxPlugin(sourceURI)

						// Print 200 for success
						if online_result == 200 {
							green := color.New(color.FgGreen).SprintFunc()
							fmt.Printf("%s | %s | %s | %s %s\n", id, name, version, sourceURI, green("200"))

							line := fmt.Sprintf("%s | %s | %s | %s | %s", id, name, version, sourceURI, "200")
							mylist = append(mylist, line)

						} else {
							// Print 404 for error
							red := color.New(color.FgRed).SprintFunc()
							fmt.Printf("%s | %s | %s | %s %s\n", id, name, version, sourceURI, red("404"))

							line := fmt.Sprintf("%s | %s | %s | %s | %s", id, name, version, sourceURI, "404")
							mylist = append(mylist, line)

						}

					} else {
						fmt.Printf("%s | %s | %s | %s\n", id, name, version, sourceURI)

						line := fmt.Sprintf("%s | %s | %s | %s", id, name, version, sourceURI)
						mylist = append(mylist, line)
					}
				}

				return true // keep iterating
			})

		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", rootFolder, err)
	}

	// Save to file
	if to_save != "" {

		// is_save_all
		// fmt.Println(is_save_all)

		if is_save_all == true {

			// fmt.Println("Append")

			f, err := os.OpenFile(to_save, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			writer := csv.NewWriter(f)
			defer writer.Flush()

			for _, item := range mylist {
				err := writer.Write([]string{item})

				if err != nil {
					fmt.Println("Error writing to file:", err)
					return
				}
			}
			// writer.Flush()

		} else {
			file, err := os.Create(to_save)

			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			defer file.Close()

			writer := csv.NewWriter(file)
			defer writer.Flush()

			for _, item := range mylist {
				err := writer.Write([]string{item})

				if err != nil {
					fmt.Println("Error writing to file:", err)
					return
				}
			}

			fmt.Printf("\nOutput is saved to %s!\n", to_save)
		}

	}
}

// Chromium
// ~/.config/chromium/Default/Extensions
func FindChromiumBrowserExtensions(to_save string, to_check_online bool, is_save_all bool) {
	// Directories to search
	BrowserPaths := GetBrowserPaths()
	rootFolder := BrowserPaths["chromium"]

	FindChromeBrowserExtensionsFromPath(rootFolder, to_save, to_check_online, false, false, is_save_all)
}

// MS Edge
func FindMSEdgeBrowserExtensions(to_save string, to_check_online bool, is_save_all bool) {
	// Directories to search
	BrowserPaths := GetBrowserPaths()
	rootFolder := BrowserPaths["msedge"]

	FindChromeBrowserExtensionsFromPath(rootFolder, to_save, to_check_online, true, false, is_save_all)
}

// Get User Profile Name
func GetUserProfileName() string {
	var user = os.Getenv("USER")
	return user
}

// BrowserPaths struct
type BrowserPaths struct {
	chrome   string
	chromium string
	firefox  string
	msedge   string
}

// Get browser paths based on operating system
func GetBrowserPaths() map[string]string {
	var osType = GetOperatingSystem()

	BrowserPaths := map[string]string{
		"chrome":   "",
		"chromium": "",
		"firefox":  "",
		"msedge":   "",
	}

	if osType == "linux" {
		var user = os.Getenv("USER")

		BrowserPaths = map[string]string{
			"chrome":   "/home/" + user + "/.config/google-chrome/Default",
			"chromium": "/home/" + user + "/.config/chromium/Default",
			"firefox":  "/home/" + user + "/.mozilla/firefox/",
			"msedge":   "/home/" + user + "/.config/microsoft-edge/Default",
			"opera":    "/home/" + user + "/.config/opera",
			"vivaldi":  "/home/" + user + "/.config/vivaldi/",
			"brave":    "/home/" + user + "/.config/BraveSoftware/Brave-Browser/Default",
		}

		return BrowserPaths

	} else if osType == "windows" {
		var USERPROFILE = os.Getenv("USERPROFILE")
		var LOCALAPPDATA = os.Getenv("LOCALAPPDATA")
		var APPDATA = os.Getenv("APPDATA")

		BrowserPaths = map[string]string{
			"chrome":   USERPROFILE + "\\AppData\\Local\\Google\\Chrome\\User Data\\Default",
			"chromium": LOCALAPPDATA + "\\Google\\Chrome\\User Data",
			"firefox":  APPDATA + "\\Mozilla\\Firefox\\Profiles\\",
			"msedge":   USERPROFILE + "\\AppData\\Local\\Microsoft\\Edge\\User Data\\Default",
			"opera":    USERPROFILE + "\\AppData\\Roaming\\Opera Software",
			"vivaldi":  USERPROFILE + "\\AppData\\Local\\Vivaldi\\User Data\\Default",
			"brave":    USERPROFILE + "\\AppData\\Local\\BraveSoftware\\Brave-Browser\\User Data\\Default",
			"sogou":    USERPROFILE + "\\AppData\\Local\\Sogou\\SogouExplorer\\User Data\\Default", // Windows only
			"360":      USERPROFILE + "\\AppData\\Roaming\\360se6\\User Data\\Default",             // Windows only
		}

		return BrowserPaths

	} else if osType == "darwin" {
		var user = os.Getenv("USER")

		BrowserPaths = map[string]string{
			"chrome":   "/home/" + user + "/.config/google-chrome/Default/Extensions/",
			"chromium": "/home/" + user + "/.config/chromium/Default",
			"firefox":  "/home/" + user + "/.mozilla/firefox/",
			"msedge":   "",
		}

		return BrowserPaths

	}

	return BrowserPaths
}

// Verify Chrome extensions ID
func CheckBrowserExtensionID(extension_id string) int {
	// https://chrome.google.com/webstore/detail/kbfnbcaeplbcioakkpcpgfkobkghlhen

	url := fmt.Sprintf("https://chrome.google.com/webstore/detail/" + extension_id)

	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		// os.Exit(1)
		return -1
	}

	// fmt.Printf("client: got response!\n")
	// fmt.Printf("client: status code: %d\n", res.StatusCode)

	return res.StatusCode
}

// Verify MS Edge extensions ID
func CheckMSEdgeExtensionID(extension_id string) int {
	// https://microsoftedge.microsoft.com/addons/detail/dgbhmbogkcdheijkkdmfhodkamcaiheo

	url := fmt.Sprintf("https://microsoftedge.microsoft.com/addons/detail/" + extension_id)

	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		// os.Exit(1)
		return -1
	}

	return res.StatusCode
}

// Verify FireFox plugin ID
func CheckFireFoxPlugin(plugin_url string) int {

	url := fmt.Sprintf("https://microsoftedge.microsoft.com/addons/detail/")

	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		// os.Exit(1)
		return -1
	}

	return res.StatusCode
}

// Verify Opera extension ID
func CheckOperaExtensionID(extension_id string) int {

	// https://addons.opera.com/en/search/?query=hnjalnkldgigidggphhmacmimbdlafdo
	url := fmt.Sprintf("https://addons.opera.com/en/search/?query=" + extension_id)

	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		// os.Exit(1)
		return -1
	}

	defer res.Body.Close()

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)

	if strings.Contains(responseString, "No search results for") {
		return 404
	} else {
		return 200
	}

	// return res.StatusCode
}

// Save to CSV
func save_csv(filename string, mylist [][]string) {
	// Create the output file
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)

	// Write each row of data to the CSV file
	for _, row := range mylist {
		err := writer.Write(row)
		if err != nil {
			panic(err)
		}
	}

	// Flush any buffered data to the underlying writer
	writer.Flush()

	// Check for any errors that occurred during the write
	if err := writer.Error(); err != nil {
		panic(err)
	}
}

// Check file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Opera browsers
func FindOperaBrowserExtensions(to_save string, to_check_online bool, is_save_all bool) {
	// Directories to search
	BrowserPaths := GetBrowserPaths()
	rootFolder := BrowserPaths["opera"]

	FindChromeBrowserExtensionsFromPath(rootFolder, to_save, to_check_online, false, true, is_save_all)
}

// Check if an item is in the list
func CheckItemInList(items []string, input string) bool {
	for _, item := range items {
		if strings.Contains(item, input) {
			return true
		}
	}
	return false
}

// Vivaldi
func FindVivaldiBrowserExtensions(to_save string, to_check_online bool, is_save_all bool) {
	// Directories to search
	BrowserPaths := GetBrowserPaths()
	rootFolder := BrowserPaths["vivaldi"]

	FindChromeBrowserExtensionsFromPath(rootFolder, to_save, to_check_online, false, false, is_save_all)
}

// Brave
func FindBraveBrowserExtensions(to_save string, to_check_online bool, is_save_all bool) {
	// Directories to search
	BrowserPaths := GetBrowserPaths()
	rootFolder := BrowserPaths["brave"]

	FindChromeBrowserExtensionsFromPath(rootFolder, to_save, to_check_online, false, false, is_save_all)
}
