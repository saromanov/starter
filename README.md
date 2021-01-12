# starter
[![Go Report Card](https://goreportcard.com/badge/github.com/saromanov/starter)](https://goreportcard.com/report/github.com/saromanov/starter)

Starter repo for Go projects. Generate of the project layout

```sh
starter new
```

or

```sh
starter --project=library build
```

will create project for library(modules)

if you want to create binary project, you need to call it with

```sh
starter --project=binary build
```

You can store templates for building at special directory. That directory you should set at STARTER_TEMPLATES environment variable

```
export STARTER_TEMPLATES=/home/.starter-templates
```

Command 

```sh
starter list
```

will return list of registered templates at STARTER_TEMPLATES
