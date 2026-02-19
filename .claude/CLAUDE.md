## Go 코드 리딩 가이드 (JVM 3년차 개발자용)

이 사용자는 Java/Kotlin JVM 서버 개발 경력 3년차이며, Go 언어를 처음 배우고 있다.
Entire CLI 오픈소스 프로젝트에 기여하는 것이 목표이다.

### 코드 설명 시 규칙

1. **Go 문법이 처음 나오면 반드시 Kotlin/Java 동일 표현을 같이 보여줄 것**
   - 예: `hookAgent, ok := ag.(agent.HookSupport)` → `val hookAgent = ag as? HookSupport`
   - 예: `make([]string, 0, cap)` → `ArrayList<String>(cap)`
   - 예: `defer f.Close()` → `.use { }` 또는 try-with-resources

2. **함수를 설명할 때 "이 프로그램이 뭘 하는지" 맥락을 먼저 설명할 것**
   - 코드 구조보다 "왜 이 함수가 필요한지"를 먼저
   - 사용자 시나리오 기반으로 설명 (entire enable 하면 → 이 함수가 호출되고 → 이런 일이 일어남)

3. **한 번에 너무 많은 코드를 보여주지 말 것**
   - 파일 전체를 한번에 분석하지 않고, 핵심 흐름만 먼저 짚기
   - 사용자가 질문하면 그때 깊이 들어가기

4. **Go 특유의 패턴은 JVM과의 차이점 위주로 설명**
   - error 처리: try-catch 없음 → `if err != nil` 패턴
   - interface: 암묵적 구현 (implements 없음)
   - 포인터: `*T`, `&v` → Java는 기본이 참조, Go는 기본이 값
   - goroutine/channel: 코루틴과 비슷하지만 런타임 스케줄러 기반
   - 대문자/소문자: public/private 접근 제어
   - `:=` 단축 선언, `_` blank identifier, `defer`, `make`, `append`

5. **Spring/JVM 아키텍처 용어로 비유**
   - Cobra Command → Spring Controller
   - strategy 패키지 → Service 레이어
   - settings 패키지 → Repository/Config
   - PersistentPostRun → HandlerInterceptor
   - RunE → @RequestMapping 핸들러 메서드

### 이미 학습한 Go 문법

- `:=` (단축 변수 선언)
- `*T` / `&v` (포인터/주소)
- `<-chan` (채널 수신)
- `go func(){}()` (고루틴)
- `switch {}` (조건부 switch = Kotlin when)
- `errors.As()` (타입 매칭 = instanceof)
- `.(Type)` (타입 단언 = as/as?)
- `_` (blank identifier)
- 대문자 = public, 소문자 = private
- GMP 모델 (goroutine 스케줄링) 이해 완료

### 코드 리딩 진행 상황

- [x] go.mod (의존성 파악)
- [x] main.go (진입점, 시그널 핸들링, 에러 처리)
- [x] root.go (Cobra 명령어 등록, 서브커맨드 라우팅)
- [x] setup.go (entire enable 명령어 - 구조만 파악)
- [ ] strategy/ (핵심 비즈니스 로직 - 체크포인트/세션 관리)
- [ ] agent/ (에이전트 연동 - Claude Code, Gemini)
- [ ] checkpoint/ (체크포인트 저장소)
- [ ] session/ (세션 상태 관리)
- [ ] commands/ (기타 명령어들)
