package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	pb "github.com/team-til/til/server/_proto"
	"github.com/team-til/til/server/datastore"
	"github.com/team-til/til/server/service"
	"google.golang.org/grpc"
)

var cfgFile string
var log = logrus.StandardLogger()

var rootCmd = &cobra.Command{
	Use:   "til",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: start,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PROJ_ROOT/.til.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("$GOPATH/src/github.com/team-til/til/server/")
		viper.SetConfigName(".til")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
	}
}

func start(cmd *cobra.Command, args []string) {
	logrusEntry := logrus.NewEntry(log)

	var dbConfig DbConfig
	viper.UnmarshalKey("datastore", &dbConfig)

	connStr := dbConfig.BuildDbConnectionStr()

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	notesDs := datastore.NewNotesDatastore(db)

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_logrus.StreamServerInterceptor(logrusEntry),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
		)),
	)

	ts := service.NewTILServer(notesDs)
	pb.RegisterTilServiceServer(s, ts)

	listenAddr := fmt.Sprintf(":%d", viper.GetInt("port"))
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal("Couldn't create tcp listener. Err: %+v", err.Error())
	}

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve. Err: %+v", err.Error())
	}
}
