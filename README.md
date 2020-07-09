# firebase-idtoken-gen

## Installation
```
go get -u github.com/yuzuy/firebase-idtoken-gen/cmd/firebase-idtoken-gen
```

## Usage

```
firebase-idtoken-gen -p {PROJECT_ID} -credfile {CREDENTIALS_FILE_PATH} -apikey {API_KEY} {UID}
// Print ID token
```
If you not specify -p -credfile -apikey flag, firebase-idtoken-gen uses default value got from your environment.
