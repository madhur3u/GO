# Gallery App Using GO And FYNE 

<img src="https://user-images.githubusercontent.com/89251393/139354175-ab656e10-0c3c-44b1-83a5-44d988277fe1.png" width="60"/> <img src="https://user-images.githubusercontent.com/89251393/139284579-441f4be3-1069-4981-b879-b96bf5fd6ad6.png" width="150"/> <img src="https://user-images.githubusercontent.com/89251393/139284585-4beb4fba-52a6-4a16-abc0-a0ca439558ab.png" width="60"/> <img src="https://user-images.githubusercontent.com/89251393/139284590-a15400c8-6c0c-4703-aff1-876cc5c67e10.png" width="70"/> 


This is a simple **Gallery** application made using GO Language and FYNE Module. It has basic functionalities to display png, jpg and jpeg images and also has next and prevous navigation button. To view images you just need to input the directory location where the images are stored and navigate using the buttons. The name of the image will be displayed on the top.

<img src="https://user-images.githubusercontent.com/89251393/139354938-9891706c-ae05-4383-87d3-77f5ef37ff0a.png" width="600"/>

## Setup

To build this application yourself for your operating system follow the steps given below.
1. Install [GO](https://golang.org/doc/install) in your system.
2. Install [FYNE](https://developer.fyne.io/started/).
3. After completing the installation process download these files.

      1. [main.go](https://github.com/madhur3u/GO/blob/main/Gallery/main.go)
      2. [go.mod](https://github.com/madhur3u/GO/blob/main/Gallery/go.mod)
      3. [go.sum](https://github.com/madhur3u/GO/blob/main/Gallery/go.sum)
4. Move all downloaded files in a folder.
5. Open terminal in that folder and execute the following command.
```
go run main.go
```
6. The above command will run the app. To make an executable file run the following command and it will make an executable file which you can use without opening terminal or installing go or fyne in other systems.
```
go build main.go
```
7. To change the theme of application you need to add a line in [main.go](https://github.com/madhur3u/GO/blob/main/Gallery/main.go).
```go
a.Settings().SetTheme(theme.LightTheme()) // for light theme
a.Settings().SetTheme(theme.DarkTheme())  // for dark theme
```
## Contact Details

I hope you like this. If you want you can modify or add features in the application according to your preferences.
If you have any query, suggestion, doubts etc., you can contact me on the given links below.


<img src="https://user-images.githubusercontent.com/89251393/138821704-5538f667-ca94-4d9f-ad49-b3c48e1cdb0c.png" width="15"/> [LinkedIn](https://www.linkedin.com/in/madhur3u/)<br> <img src="https://user-images.githubusercontent.com/89251393/138821710-7b7585e0-4766-49ba-8543-c116d4da82c4.png" width="15"/> [Instagram](https://www.instagram.com/madhur3u/) <br><img src="https://user-images.githubusercontent.com/89251393/138821715-eab2496c-e895-4113-a26b-96c087a83d9b.png" width="15"/> m3333u@gmail.com <br><img src="https://user-images.githubusercontent.com/89251393/138822281-9aaf6bdc-2fe0-469a-bd31-ed43bc96dcc2.png" width="15"/> [@madhur3u](https://web.telegram.org/)

<!-- <img src="" width="123"/>  -->
