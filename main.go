package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/vecosy/vecosy/v2/pkg/vecosy"
)

func main() {
	jwsToken := "eyJhbGciOiJSUzI1NiJ9.YXBwMQo.A98GFL-P3vtehn0r5GCO_a0OYb5h6trxg3a8WE9hOPDzJ40yOEGtZxyUM6_3Exk65c52-nzWEEc5P-QtgGrgJFOOZlKneKoa1bYBlWRONoysuq95UtSY0doEOMWGvI9AqB685OzmVPuW2UlHg_HlQuuTO6Re1uKc5gr1qZPlyyWEsfoVYTFbfidLoBKWPOuZTxpd8uRx0Rv3LrrmFEcGPHaMNQ2WiXAEJG6OaMTBtwKiynEFH3DU5Rx2WP9M98bH-emC_w7Zq1xKaCOsj2t09F00KohcGC49zSPgPVpp_TwF1qt6_0d0Mnh_Eqi_NHpobVvO85ZOLS05AyW9LQyA5A"
	vecosyCl, err := vecosy.NewClientBuilder("localhost:8081", "app1", "1.0.0", "dev").WithJWSToken(jwsToken).Build(nil)
	panicOnError(err)
	err = vecosyCl.WatchChanges()
	panicOnError(err)
	fmt.Printf("db.user:%s\n", viper.GetString("db.user"))
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
