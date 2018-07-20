package main

import (
  "bufio"
  "compress/gzip"
  "flag"
  "log"
  "os"
  "strconv"
  "strings"
  "time"

  "github.com/firstrow/tcp_server"
  "github.com/mmcloughlin/geohash"
)

var serverHost = flag.String("host", "localhost", "Hostname to bind to. eg: 192.168.1.10, default:localhost")
var serverPort = flag.String("port", "3333", "Port eg: 8080, default: 3333")
var dataPath = flag.String("dataPath", "./data/master.csv.gz", "Path to data file. eg: ./data/file.gz, default: ./data/master.csv.gz")

func main() {
  flag.Parse()
  connString := *serverHost + ":" + *serverPort
  dataFilePath := *dataPath

  server := tcp_server.New(connString)
  geohashToCountryMapping := getGeohashToCountryMapping(dataFilePath)

  server.OnNewMessage(func(c *tcp_server.Client, message string) {
    startNano := time.Now().UnixNano()
    reply := messageHandler(message, geohashToCountryMapping)
    c.Send(reply + "\n")
    log.Println("Response time: " + strconv.FormatInt((time.Now().UnixNano()-startNano), 10))
  })
  server.Listen()
}

func messageHandler(message string, geohashToCountryMapping map[string]string) string {
  message = strings.TrimSpace(message)
  response := ""
  startNano := time.Now().UnixNano()
  if message != "" {
    lat, lon := parseLatLonFromMessage(message)
    geohash6 := generateGeohash(lat, lon)
    country := getCountryFromGeohashToCountryMapping(geohash6, geohashToCountryMapping)
    timeTaken := time.Now().UnixNano() - startNano
    response = country + ", " + strconv.FormatInt(timeTaken, 10)
  }

  log.Println(message + " => " + response)
  return response
}

func getCountryFromGeohashToCountryMapping(geohash6 string, geohashToCountryMapping map[string]string) string {
  country := ""
  for geohashLength := 6; geohashLength > 1; geohashLength-- {
    country = geohashToCountryMapping[geohash6[0:geohashLength]]
    if country != "" {
      break
    }
  }
  return country
}

func parseLatLonFromMessage(message string) (float64, float64) {
  coords := strings.Split(message, ",")
  lat, lon := 0.0, 0.0
  if len(coords) == 2 {
    lat, _ = strconv.ParseFloat(strings.TrimSpace(coords[0]), 64)
    lon, _ = strconv.ParseFloat(strings.TrimSpace(coords[1]), 64)
  }
  return lat, lon
}

func generateGeohash(lat float64, lon float64) string {
  return geohash.EncodeWithPrecision(lat, lon, 6)
}

func getGeohashToCountryMapping(filePath string) map[string]string {
  geohashToCountryMapping := make(map[string]string)
  startNano := time.Now().UnixNano()
  log.Println("Loading data.")
  file, err := os.Open(filePath)
  defer file.Close()
  checkError(err)

  gzipReader, err := gzip.NewReader(file)
  defer gzipReader.Close()
  checkError(err)

  scanner := bufio.NewScanner(gzipReader)
  for scanner.Scan() {
    line := strings.Split(scanner.Text(), ",")
    geohashToCountryMapping[line[0]] = line[1]
  }

  checkError(scanner.Err())
  
  log.Println("Loading complete.")
  log.Println("Total Entries: " + strconv.Itoa(len(geohashToCountryMapping)))
  timeTaken := (time.Now().UnixNano() - startNano) / 1000000000
  log.Println("Boot time: " + strconv.FormatInt(timeTaken, 10) + " seconds")
  return geohashToCountryMapping
}

func checkError(e error) {
  if e != nil {
    log.Println(e)
  }
}
