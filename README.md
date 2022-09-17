# whoami

Actually, I did this project to make something in Go and Docker.

I'm very new to GoLang. There are many mistakes and no tests.

## Deployment

```bash
docker run -d -p 5000:5000 tatoglu/whoami:latest
```

## Endpoints

There are 3 endpoints.

### Regular Usage

It can be used to learn hostname of the container.

```bash
curl http://localhost:5000/
```

```json
{"hostname": "89f02ed1a4dd"}
```

### Date

It can be used to learn current date and time of the container.

```bash
curl http://localhost:5000/date
```

```json
{"hostname": "89f02ed1a4dd", "date": "2022-09-17T14:02:52.397349799Z"}
```

### Fun Part

It can use to have fun. The endpoint get a joke from [I can haz dad joke](https://icanhazdadjoke.com/) website. Everytime a new joke will get.

```bash
curl http://localhost:5000/joke
```

```json
{"hostname": "89f02ed1a4dd", "joke":"People saying 'boo! to their friends has risen by 85% in the last year.... That's a frightening statistic."}
```

## Objectives

- [X] Develop an API.
- [X] Call a 3rd party website from the API.
- [ ] Move all the response models into a folder or module.
- [ ] Move all the handlers into a folder or module.
- [ ] Add GitHub Actions for deployment to Docker Hub.
