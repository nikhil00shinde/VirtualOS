package main

import (
	"io/ioutil"
	"path"
	"strconv"
	"strings"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)


func main() {
	a := app.New()
	w := a.NewWindow("Image Viewer")
	w.Resize(fyne.NewSize(400, 400))
	a.Settings().SetTheme(theme.DarkTheme())

	imageArr := []string{}

  root_path := ""
	data, _ := ioutil.ReadDir(root_path)
  
  for _,file := range data{
	  if !file.IsDir() {
			extens := strings.Split(file.Name(), ".")[1]
			if extens == "png" || extens == "jpeg"{
			fileName := path.Join(root_path,file.Name())
			imageArr = append(imageArr, fileName)
		   }
		}
	}


		tabs := container.NewAppTabs()

		for i ,files := range imageArr {
			image := canvas.NewImageFromFile(files)
			image.FillMode = canvas.ImageFillContain
      tabs.Append(container.NewTabItem("Image"+strconv.Itoa(i+1), image))
		}
		
	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs)
	w.ShowAndRun()

}


// package main

// import (
// 	"fmt"
//   "os"
// 	"io/ioutil"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/theme"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("TabContainer Widget")
// 	var first string
//   fmt.Println("Enter File Path: ")
// 	fmt.Scanln(&first)
  
// 	if _, err := os.Stat(first); !os.IsNotExist(err) {
// 		  fmt.Println("valid path")
// 				data, _ := ioutil.ReadDir(first)
//        for _,file := range data{
// 	   	fmt.Println(file.Name(),file.IsDir())
// 			 }
// 	}
	


// 	tabs := container.NewAppTabs(
// 		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
// 		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
// 	)

// 	tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

// 	tabs.SetTabLocation(container.TabLocationLeading)

// 	myWindow.SetContent(tabs)
// 	myWindow.ShowAndRun()
// }