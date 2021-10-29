# Weather App Using GO And FYNE 

<img src="https://user-images.githubusercontent.com/89251393/139450178-7a2b501b-1688-48af-be37-214ae2c2d0c4.png" width="60"/> <img src="https://user-images.githubusercontent.com/89251393/139284579-441f4be3-1069-4981-b879-b96bf5fd6ad6.png" width="150"/> <img src="https://user-images.githubusercontent.com/89251393/139284585-4beb4fba-52a6-4a16-abc0-a0ca439558ab.png" width="60"/> <img src="https://user-images.githubusercontent.com/89251393/139284590-a15400c8-6c0c-4703-aff1-876cc5c67e10.png" width="70"/> 


This is a simple **Weather** application made using GO Language and FYNE Module. Using this app you can get the temperature and other basic weather details about the location you have entered.

<img src="https://user-images.githubusercontent.com/89251393/139450848-619a9aad-2b5b-4431-9bad-28d289583e74.png" width="400"/>

## Setup


> **Note :** This application is developed and tested using Linux (Ubuntu 20.04). Some features might not work for other Operating Systems, you have to make changes in the code accordingly. In case of any problem contact me on the links given [here](#contact-details).

To build this application yourself for your operating system follow the steps given below.
1. Install [GO](https://golang.org/doc/install) in your system.
2. Install [FYNE](https://developer.fyne.io/started/).
3. After completing the installation process download all files in of [this folder](https://github.com/madhur3u/GO/tree/main/Weather%20App).
4. Move all downloaded files in a folder.
5. Open terminal in that folder and execute the following command.
```
go run main.go
```
6. The above command will run the app. To make an executable file run the following command and it will make an executable file which you can use without opening terminal or installing go or fyne in other systems.
```
go build main.go
```
7. To change the theme of application you need to add a line in [main.go](https://github.com/madhur3u/GO/blob/main/Weather%20App/main.go).
```go
a.Settings().SetTheme(theme.LightTheme()) // for light theme
a.Settings().SetTheme(theme.DarkTheme())  // for dark theme
```
## Contact Details

I hope you like this. If you want you can modify or add features in the application according to your preferences.
If you have any query, suggestion, doubts etc., you can contact me on the given links below.


<img src="https://user-images.githubusercontent.com/89251393/138821704-5538f667-ca94-4d9f-ad49-b3c48e1cdb0c.png" width="15"/> [LinkedIn](https://www.linkedin.com/in/madhur3u/)<br> <img src="https://user-images.githubusercontent.com/89251393/138821710-7b7585e0-4766-49ba-8543-c116d4da82c4.png" width="15"/> [Instagram](https://www.instagram.com/madhur3u/) <br><img src="https://user-images.githubusercontent.com/89251393/138821715-eab2496c-e895-4113-a26b-96c087a83d9b.png" width="15"/> m3333u@gmail.com <br><img src="https://user-images.githubusercontent.com/89251393/138822281-9aaf6bdc-2fe0-469a-bd31-ed43bc96dcc2.png" width="15"/> [@madhur3u](https://web.telegram.org/)

<!-- <img src="" width="123"/>  -->
