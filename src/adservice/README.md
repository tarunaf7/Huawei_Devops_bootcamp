

The Ad service provides advertisement based on context keys. If no context keys are provided then it returns random ads.



The Ad service uses gradlew to compile/install/distribute. Gradle wrapper is already part of the source code. To build Ad Service, run:

```
./gradlew installDist
```
It will create executable script src/adservice/build/install/hipstershop/bin/AdService


If you need to upgrade the version of gradle then run

```
./gradlew wrapper --gradle-version <new-version>
```



From `src/adservice/`, run:

```
docker build ./
```

