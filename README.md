# Tistory Like Increment!

### Quick Start
```go
git clone https://github.com/YankeeTube/Automate-Tistory.git tistory && cd tistory && \
./tistory_like -blog=<BlogName> -post=<PostID> -target=<LikeCount>
```

### Usage
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