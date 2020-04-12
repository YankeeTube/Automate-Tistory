# Tistory Like Increment!
Tistory 블로그 게시물의 공감수를 증가시키는 `Simple`한 프로그램입니다.  
이 프로그램을 사용하여 블로그의 `문제가 발생하거나 정지사유`가 된다면 그 어떠한 책임도 지지 않습니다. 모든 책임은 `사용자`에게 있습니다.   
타 블로그의 게시물에 사용시에는 `DaumKakao`측 에서 로그인한 계정의 IP를 판단하여 처리할 수 있으니 악용을 금지합니다.  

  
## Quick Start of binary
Windows용 다운로드: `https://github.com/YankeeTube/Automate-Tistory/releases/download/0.1.0/windows.zip`  
Unix(Linux/Mac)용 다운로드: `https://github.com/YankeeTube/Automate-Tistory/releases/download/0.1.0/linux_or_mac.tar.xz`  
  
  
### Binary Usage  
[상세보기](https://gmyankee.tistory.com/292)
  
  
## Quick Start of source code
```go
git clone https://github.com/YankeeTube/Automate-Tistory.git tistory && cd tistory && \
go run main.go -blog=<BlogName> -post=<PostID> -target=<LikeCount>
```  
  
### Source Code Usage
`-blog` 증가할 블로그 명을 입력합니다.   
```
ex) -blog=gmyankee  // gmyankee.tistory.com
```
  
  
`-post` 공감 수를 증가시킬 PostID(URL)를 지정합니다. 반드시 숫자여야 가능합니다.  
```
ex) -post=291 // https://gmyankee.tistory.com/291
```
  
  
`-target` 증가 시킬 공감수의 수를 지정합니다.  
```
ex) -target=30 // 공감수 30회를 목표로 함
```

  
  
## License
AutomateTistory is licensed under [MIT License](https://github.com/YankeeTube/Automate-Tistory/blob/master/LICENSE).
