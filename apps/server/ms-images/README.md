# Ms Images

## Usages

| Endpoint                           | Body                                                                                                         | Response type | Description                                                       | Tests |
|:-----------------------------------|--------------------------------------------------------------------------------------------------------------|--------------:|:------------------------------------------------------------------|:-----:|
| POST /api/images                   | content-type: multipart/form-data <br/> body: `{file: File, fileName: string, types: string[], alt: string}` |        200 Ok | Uploads images, transform to desired types and stores it in webp. |       |
| GET /api/images/:id                | -                                                                                                            |        200 Ok | File info                                                         |       |
| GET /api/images/:id/free-transform | params: width(number), height(number), blur(number)                                                          |        200 Ok | Returns image processed to desired dimensions. Can add blur too.  |       |

possible types: original, resized (half size of original), logo, thumbnail, thumbnailZoom

## Dev

Packages used: 
- disintegration/imaging v1.6.2
- gin-gonic/gin v1.9.0
- gin-contrib/cors v1.4.0
- google/uuid v1.3.0
- spf13/viper v1.15.0
- samber/lo v1.37.0
- driver/postgres v1.4.8
- gorm v1.24.5
- [nickalie/go-webpbin](https://github.com/nickalie/go-webpbin) 


## Run

### Dev
```
$ nx serve ms-images
```


### PROD - DOCKER
Build image
```
$ docker build --build-arg build_env=prod -t ms-images .
```

Run image 
```
$ docker run -d -p 8080:8080 -v $PWD/public:/app/images ms-images
```


