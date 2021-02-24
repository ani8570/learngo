# learngo
## 참고자료
예제로 배우는 Go 프로그래밍 [link](http://golang.site/Go/Basics)

노마드코더 [link](https://nomadcoders.co/go-for-beginners)
#
## 사용법
- 주의사항 및 에러


    c드라이브 아래 GO
    ex) c:\Go
    GOPATH, GOROOT, GOBIN 설정 해줌 좋음


    go.mod -> go mod init 이름
    #
    error : cgo: exec gcc: exec: gcc: executable file not found in %PATH%
        
    * [링크1](https://m.blog.naver.com/kh2un/222037244807)
    [링크2](https://sourceforge.net/projects/mingw-w64/files/mingw-w64/)


### 
* package 이름 작성
main은 컴파일러로 인식해서 한번만 , package 이름은 폴더로

       package main 

        fuc main() {
            ...
        }


    "fmt" : package for formatting 

    fmt.Println 대문자로 시작하는 함수는 export 가능

- 자료형

        var name string = "Lee"

        name := "Lee"

    this is sameting. but if you declared outside func, you have to declarared like this 
        
        "var name string = "Lee"


- slice선언 :  

        var URLs []string // []안에 숫자가 없는게 포인트
        URLs = append(URLs, res) //res값 뒤로 추가.
        URLs = append([]string{res}, URLs...) // res값 앞에 추가



- Go can return more than one.

    ex) totalLength, upperName := lenAndUpper("Lee")

    두개를 반환하는데 하나만 있는 건 안됨. 

        totalLength := lenAndUpper("Lee")

    이건 됨

         totalLength, _ := lenAndUpper("Lee")

    함수가 밑과 같다

        repeatMe("nico", "lynn", "dal", "marl", "flynn")

- ...string으로 여러개 받는다고 알림.
     
        func repeatMe(words ...string) {
            fmt.Println(words)
        }



- 밑에 두 함수는 같음...
2번째 함수를 naked return이라 함 ( naked : 적나라한)
    
        func lenAndUpper(name string) (int, string) {
            return len(name), strings.ToUpper(name)
        }

        func lenAndUpper2(name string) (lenght int, uppercase string) {
            lenght = len(name)
            uppercase = strings.ToUpper(name)
            return
        }

- 함수 끝나고 실행시킬 것이 있으면 defer를 사용한다.
ex) 이미지 제거, 저장, api로 요청을 보낸다거나 등등

        func lenAndUpper3(name string) (lenght int, uppercase string) {
            defer fmt.Println("I'm done")
            lenght = len(name)
            uppercase = strings.ToUpper(name)
            return
        }

    반복문 range ->index나옴

        func superAdd(numbers ...int) int {
            fmt.Println(numbers)
            for index, number := range numbers {
                fmt.Println(index, number)
            }
            return 1
            //range
        }

- if문 자체에서 변수 선언 가능

        func canIDrink(age int) bool {
        if koreanAge := age+2 ; koreanAge < 18 {
            return false
        }	
        return true
       }

    switch문에서도 같음

        switch koreanAge := age + 2; {
            case koreanAge < 18:
                return false
            case koreanAge == 18:
                return true

            }
            return false

- Go에서도 포인터와 래퍼런스 사용 가능

- 배열도 사용 가능, 똑같이 0부터 접근
    
      names := [5]string{"nico", "lynn", "dal", "marl", "flynn"}

- map도 사용 가능

    Lee := map[string]string{"name": "L", "age": "12"}
	fmt.Println(Lee)

    for key, value := range Lee {
        fmt.Println(key, value)
    }

    key값에 따라서 제거할 때
        
        delete(m, key) // m : map , key : string

    함수에서 map을 받아오면 *가 항상 붙어있다 생각하면 됨

        results := map[string]string

    밑에 처럼 적고 추가를 하면 panic 됨 --> map => nil
        
        var results = map[string]string 

    올바른 예
        
        var results = map[string]string {}
        var results = make(map[string]string) // make -> map만드는 함수





- 구조체 사용가능
(외부에서 구조체 사용시 사용가능하게 하려면 대문자)
(외부에서 구조체 내부 변수 사용시 사용가능하게 하려면 대문자)

        type Person struct {
            Name    string
            Age     int
            FavFood []string
        }
        type person struct {
            name    string
            age     int
            favFood []string
        }

        favFood := []string{"kimchi", "ramen"}
        l := person{"Lee", 25, favFood}	//can
        i := person{name: "Lee", age: 18, favFood: favFood}	//can
        i := person{name: "Lee", 18, favFood: favFood}	//can't

- receiver를 가질 수 있음. 
receiver 규칙 : 구조체 이름의 첫번째 글자 소문자로 해서 받음.

        func (a Account) Deposit(amount int) {
            ...
        }


- go routines : 메인 함수가 시작 되는 동안에 병렬 처리함.
 메인 함수가 끝나면? 다 안되어도 종료시킴

        func main() {
            go sexyCount("nico")
            sexyCount("lynn") //go sexyCount("lynn") --> start and end at the same time
        }
        func sexyCount(name string) {
            for i := 0; i < 10; i++ {
                fmt.Println(name, " is sexy", i)
                time.Sleep(time.Second)
            }
        }

- channel선언


        func main() {
            c := make(chan bool)
            people := [2]string{"nico", "flynn"}
            for _, person := range people {
                go isSexy(person, c)
            }
            result := <-c 
            fmt.Print(result)

        }
- channel에서 주고받을 자료형 선언
        c <- true로 값 전달
        c <- person + " is sexy"
        func isSexy(person string, c chan bool) {
            //time.Sleep(time.Second * 5)
            for i := 0; i < 100; i++ {
                fmt.Println(person, " is sexy", i)
                //time.Sleep(time.Second)
            }
            c <- true
        }

        func isSexy(person string, c chan string) {
            c <- person + " is sexy"
        }



- go get github.com/PuerkitoBio/goquery를  실행시켜 다운받아
HTML 내부를 들여다 볼 수 있게 해줌

        strconv.Itoa --> 숫자를 문자열로
        func strings.Join(elems []string, sep string) string // sep으로 elems를 연결
        func strings.Fields(s string) []string// 공백 단위로 구분
        func strings.TrimSpace(s string) string// 공백 제거

