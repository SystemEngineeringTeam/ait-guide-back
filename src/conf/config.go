package conf

import "github.com/spf13/viper"

var m, n *viper.Viper

func init() {
	m = viper.New()
	m.SetConfigType("yaml")
	m.SetConfigName("mysql")
	m.AddConfigPath("conf/environments/")

	n = viper.New()
	n.SetConfigType("yaml")
	n.SetConfigName("neo4j")
	n.AddConfigPath("conf/environments/")
}

func GetMysqlConfig() *viper.Viper {
	if err := m.ReadInConfig(); err != nil {
		return nil
	}
	return m
}

func GetNeo4jConfig() *viper.Viper {
	if err := n.ReadInConfig(); err != nil {
		return nil
	}
	return n
}
