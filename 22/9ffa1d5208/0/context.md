# Session Context

## User Prompts

### Prompt 1

현 프로젝트에 오픈소스 기여자가 되고싶은데 나는 go 언어를 처음 봐 어떻게 어디서 부터 시작해야할까

### Prompt 2

내가 java kotlin jvm 진영의 서버개발자로 경력이 3년인데 너무 기본부터 배울 필욘 없을 거 같아

### Prompt 3

같이 코드 리딩하자 go.mod 부터

### Prompt 4

응 main.go 보자

### Prompt 5

main.go 는 그러면 enable rewind, enable status 같은 명령어를 입력 받는 곳인거야?

### Prompt 6

먼저 문법좀 파악하고 가자 := 이건 뭐야?

### Prompt 7

[Request interrupted by user]

### Prompt 8

먼저 문법좀 파악하고 가자 := 이건 뭐야? 그리고 go 언어에 대한 특징도 설명해줘

### Prompt 9

고루틴도 코루틴과 마찬가지로 컴파일 시 코드가 변형되어 추가적인 코드가 생겨서 스레드가 <-sigChan 호출하고 다른 일을 하러 가는거야?

### Prompt 10

go 는 인터프리터 언어야?

### Prompt 11

<-sigChan 은 뭔데 suspend 같은건가?

### Prompt 12

go 언어는 os를 직접 사용하는게 아니지? 인터페이스를 통해 실행하는거고 시스템 콜이 따로 있지?

### Prompt 13

go 런타임은 어떻게 스레드를 대기시키고 기다리고 있는거지 내부 로직이 궁금해

### Prompt 14

@cmd/entire/main.go#L32-46 여기 문법도 설명해줘

### Prompt 15

고루틴은 go 를 실행 시키는 엔진에 내부적으로 큐가 있어서 경량 스레드 큐를 사용한다는거지?

### Prompt 16

P 논리 프로세서가 어떤 의미를 갖나보네? 논리 프로세서가 뭔데? 하드웨어 스레드 말하는거야?

### Prompt 17

P 를 go 언어에서 뭐라고 부르는데?

### Prompt 18

gomaxprocs 라는 개념이 필요하게 된 배경이 와닿지가 않아 예시를 들어 설명해줘

### Prompt 19

응 보자

### Prompt 20

응 enable.go 보자

### Prompt 21

뭐 이해할 수가 없네 모든 함수들이 뭘 뜻하는건지 알 수가 없어 프로젝트 파악할 때 원래 이런거야?

### Prompt 22

사용은 이미 해봤어   hookAgent, ok := ag.(agent.HookSupport) 난 여기서 .(...) 이 문법이 타입 캐스팅을 하는 문법인거야?

### Prompt 23

그러면 좀 더 go 문법을 이해하면서 이 프로젝트를 파악할 수 있도록 나를 가이드 하도록 CLAUDE.md 에 작성해줘

### Prompt 24

어딨어 .claude 에 안보이는데?

### Prompt 25

아니 애초에 CLAUDE.md 파일이 하나밖에 없어

### Prompt 26

왜 안보이는거지 .claude 내부에 없어 reload from disk 했는데도 안보이는데

### Prompt 27

왜 홈디렉터리에 작성한거야? 프로젝트의 .claude 내부에 작성하는건 안돼?

### Prompt 28

근데 홈 디렉터리 .claude 에 추가한 이유가 뭐고 프로젝트의 .claude 에 추가한것과 어떤 차이가 있는지 알려줘

### Prompt 29

응 정리해줘

### Prompt 30

그럼 프로젝트의 루트에 있는 CLAUDE.md 와 .claude/CLAUDE.md 는 어떤 차이가 있는지 알려줘

### Prompt 31

그럼 claude 는 어떻게 각각의 CLAUDE.md 를 어떻게 사용하는지 알려줘

### Prompt 32

그럼 컨텍스트 window 가 너무 커지는거 아니야?

### Prompt 33

.claude/CLAUDE.md .gitignore 에 추가하고 커밋 및 푸시까지 진행해

