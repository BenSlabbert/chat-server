package cmd

import (
	"chat-server/chat"
	"context"
	"github.com/spf13/cobra"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chat",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {},
}

//var chatClient chat.ChatClient

type client struct {
	call     chat.ChatClient
	callOpts []grpc.CallOption

	limiter *rate.Limiter
}

var chatClient *client

// wrappedStream  wraps around the embedded grpc.ClientStream, and intercepts the RecvMsg and
// SendMsg method call.
type wrappedStream struct {
	grpc.ClientStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("Receive a message (Type: %T) at %v\n", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("Send a message (Type: %T) at %v\n", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

// streamInterceptor is an example stream interceptor.
func streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStream(s), nil
}

func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("unaryInterceptor before invoker: %s\n", method)
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	end := time.Now()
	log.Printf("unaryInterceptor after invoker: %s\n", method)
	log.Printf("RPC: %s, start time: %s, end time: %s, err: %v\n", method, start.Format(time.RFC3339), end.Format(time.RFC3339), err)

	if err != nil {
		// this is a naive retry
		// add rate liming and retry limits/backoff etc
		switch status.Code(err) {
		case codes.Unavailable:
			log.Println("service unavailable, retry...")
			return unaryInterceptor(ctx, method, req, reply, cc, invoker, opts...)
		default:
			log.Println("non retry error")
		}
	}

	return err
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(unaryInterceptor), grpc.WithStreamInterceptor(streamInterceptor), grpc.WithBlock())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer func(c *grpc.ClientConn) {
		log.Println("closing server connection")
		if e := c.Close(); e != nil {
			log.Printf("failed to close connection: %v", e)
		}
	}(conn)

	var callOpts []grpc.CallOption
	callOpts = append(callOpts, grpc.WaitForReady(true))

	chatClient = &client{
		call:     chat.NewChatClient(conn),
		callOpts: callOpts,
		limiter:  rate.NewLimiter(5, 1),
	}

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.newApp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".newApp" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".newApp")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}
