package main

var broker = flag.String("broker", "0.0.0.0:1883", "the broker url")

func main() {
    // parse flags
    flag.Parse()

    // print info
    fmt.Printf("Starting hub with %s ...\n", *broker)

}
