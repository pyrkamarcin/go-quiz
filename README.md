# go-quiz

## Register

```shell script
curl --location --request POST 'localhost:5000/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@email.com",
    "password": "Test1234!"
}'
```

```json
{
    "status": "ok"
}
```

## Login

```shell script
curl --location --request POST 'localhost:5000/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@email.com",
    "password": "Test1234!"
}'
```

```json
{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTc3ODUxNzUsInN1YiI6MX0.QpKvC_1t5OGnBlDm21fxO-YrBx6CcmQIyKFYeCpVkjw",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlczIyMjJAZW1haWwuY29tIiwiZXhwIjoxNTk3OTU3OTc1fQ.KHia-zwGDQmkid6pmLx6J7PlbDlg6WNwGO3Cg_1wV-E"
}
```

## Open endpoint

```shell script
curl --location --request GET 'localhost:5000/open'
```

## Close endpoint

```shell script
curl --location --request GET 'localhost:5000/close' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QuZW1haWwrMkBlbWFpbC5jb20iLCJleHAiOjE1OTc5NTUwNTN9.MjhPlH5M9JUQbq2UD9IybAYupEv_MxsN6zNYGUpnyIo'
```

```json
{
    "email": "test@email.com"
}
```
