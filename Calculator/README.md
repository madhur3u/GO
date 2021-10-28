# Calculator Using GO And FYNE

<img src="https://user-images.githubusercontent.com/89251393/139284579-441f4be3-1069-4981-b879-b96bf5fd6ad6.png" width="150"/> <img src="https://user-images.githubusercontent.com/89251393/139284585-4beb4fba-52a6-4a16-abc0-a0ca439558ab.png" width="60"/> <img src="https://user-images.githubusercontent.com/89251393/139284590-a15400c8-6c0c-4703-aff1-876cc5c67e10.png" width="70"/>


This is a basic **Calculator** application made using GO Language and FYNE Module. It can perform all basic calculations and can take input using keyboard or calculator's own virtual keyboard. It can also show detailed history of previous calculations done in the present run of the application.

Below are the screenshots of application in both light and dark mode.

<img src="https://user-images.githubusercontent.com/89251393/139278465-40bb5c62-1367-4438-8ec7-3f9d36e3322b.png" width="250"/> <img src="https://user-images.githubusercontent.com/89251393/139278451-c28ff391-2ac8-4c77-9d1d-6c3de19c2eac.png" width="250"/> 

## Setup

To build this application yourself for your operating system follow the steps given below.
1. Install [GO](https://golang.org/doc/install) in your system.
2. Install [FYNE](https://developer.fyne.io/started/).
3. After completing the installation process download these files.

      1. [main.go](https://github.com/madhur3u/GO/blob/main/Calculator/main.go)
      2. [go.mod](https://github.com/madhur3u/GO/blob/main/Calculator/go.mod)
      3. [go.sum](https://github.com/madhur3u/GO/blob/main/Calculator/go.sum)
4. Move all downloaded files in a folder.
5. Open terminal in that folder and execute the following command.
```
go run main.go
```
6. The above command will run the app. To make an executable file run the following command and it will make an executable file which you can use without opening terminal or installing go or fyne in other systems.
```
go build main.go
```
7. To change the theme of application you need to change Line 21 in [main.go](https://github.com/madhur3u/GO/blob/main/Calculator/main.go).
```go
a.Settings().SetTheme(theme.LightTheme()) // for light theme
a.Settings().SetTheme(theme.DarkTheme())  // for dark theme
```
## Contact Details

I hope you like this. If you want you can modify or add features in the application according to your preferences.
If you have any query, suggestion, doubts etc., you can contact me on the given links below. 


<img src="https://user-images.githubusercontent.com/89251393/138821704-5538f667-ca94-4d9f-ad49-b3c48e1cdb0c.png" width="15"/> [LinkedIn](https://www.linkedin.com/in/madhur3u/)<br> <img src="https://user-images.githubusercontent.com/89251393/138821710-7b7585e0-4766-49ba-8543-c116d4da82c4.png" width="15"/> [Instagram](https://www.instagram.com/madhur3u/) <br><img src="https://user-images.githubusercontent.com/89251393/138821715-eab2496c-e895-4113-a26b-96c087a83d9b.png" width="15"/> m3333u@gmail.com <br><img src="https://user-images.githubusercontent.com/89251393/138822281-9aaf6bdc-2fe0-469a-bd31-ed43bc96dcc2.png" width="15"/> [@madhur3u](https://web.telegram.org/)

<!-- <img src="" width="123"/>  -->

