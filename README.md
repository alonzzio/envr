# envr
Simple env return framework

Used package : https://github.com/joho/godotenv

Os Package

## Usage

Import
```
go get github.com/alonzzio/envr/envr
```

   
```
	directoryPath := "" // Your env files direcory. please add path with out '/'.
	err = envr.LoadEnv("directoryPath", "environment.env","environment1.env")
	if err != nil {
		return err
	}
```
This loads the files from  the same location. Currently not supported from different folders/paths

```
key,_:=envr.GetString("JWTSecretKey")
fmt.Println(key)
```


