package main

import (
    "fmt"
    "log"   
    "github.com/kelseyhightower/envconfig"
)
type KafkaConfig struct {
    Address    string `envconfig:"KAFKA_ADDRESS" required:"true"`
    Port       int    `envconfig:"KAFKA_PORT" default:"9092"`
    WriteTopic string `envconfig:"KAFKA_WRITE_TOPIC" default:"topic1"`
    ReadTopic  string `split_words:"true"`
}
func main() {
    r, err := kafkaConfig()
    if err != nil {
        log.Fatal("fail")    }
    log.Fatal(r)
}
func kafkaConfig() (*KafkaConfig, error) {
    kafkaConfig := &KafkaConfig{}
    err := envconfig.Process("KAFKA", kafkaConfig)
    if err != nil {
        fmt.Printf("kafka config parameter error: %v", err)
        return nil, err
    }
    fmt.Printf("kafka config: %+v\n", kafkaConfig)
    return kafkaConfig, nil
}
