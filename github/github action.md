github action 을 이용하면 git 동작에 따른 workflow 를 정할 수 있다.\
til 의 README.md 는 git 에 푸시를 했을 때 자동으로 만들도록 되어있다.

아래는 workflow 에 대한 yml 파일
```yaml
name: build readme
on:
  push:
    branches:
      - main
    paths-ignore:
      - README.md
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: checkout
        uses: actions/checkout@v2
        with:
          fetch-depth: 0
      - name: golang
        uses: actions/setup-go@v2
      - name: gen
        run:
          go run .github/til.go
      - name: push
        run: |-
          git diff
          git config --global user.email "stanleylee0707@gmail.com"
          git config --global user.name "leejh-77"
          git diff --quiet || (git add README.md && git commit -m "Updated README")
          git push
```