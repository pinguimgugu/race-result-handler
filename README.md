# Race result
 This project has as main function create an output file with classification race and
 some important metrics of laps based on file race_log.txt

Make sure you have installed
 - docker engine
 - docker compose

 # Build base image and get dependences
 ```
    ./development.sh build && ./development.sh dep
 ```

 # Running unit test
 ```
    ./development.sh run-unit
 ```

# Running classification race process
 ```
    ./development.sh run
 ```

After run just open result_classification.csv to viewed classification race detailed in currenty directory
