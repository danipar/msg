package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/danipar/msg"
	"github.com/danipar/msg/vendor/github.com/spf13/viper"
	"io"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

var pow []int = []int{1, 2, 4, 8, 16, 32, 64, 128}
var m map[string]string

type goArgs struct {
	loglevel string
	id       string
}

type fileSet struct {
	dir  string
	name string
	ext  string
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(reflect.TypeOf(w))
	io.WriteString(w, "hello world!")
}

func main() {
	msg.Msg()

	fmt.Println("Hello, World!")
	fmt.Println("This time is,", time.Now())
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random number is,", rand.Intn(10))
	fmt.Println(add(2, 3))
	fmt.Println(swap("yahoo", "google"))
	for i, v := range pow {
		fmt.Printf("i=%d, v=%d, 2**%d = %d\n", i, v, i, v)
	}
	m = make(map[string]string)
	m["yahoo"] = "ok"
	m["google"] = "good"
	m["apple"] = "ok"
	fmt.Println(m)
	v, ok := m["yahoo"]
	fmt.Println(v, ok)
	for k, v := range m {
		fmt.Println(k, v)
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.WarnLevel)
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Warn("A group of walrus emerges from the ocean")
	viper.AutomaticEnv()
	configDir := viper.GetString("fileset")
	f := &fileSet{configDir, "test", "json"}
	fmt.Println(f.name)
	viper.SetConfigName(f.name)
	viper.SetConfigType(f.ext)
	viper.AddConfigPath(f.dir)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(viper.GetString("host.name"))
	fmt.Println(viper.GetString("go_options"))
	for k, v := range strings.Split(viper.GetString("go_options"), " ") {
		fmt.Println(k, v)
	}
	fmt.Println(reflect.TypeOf(msg.Msg))

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
