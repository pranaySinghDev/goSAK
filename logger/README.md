# Logger

Golang supports inbuilt logs and a lot of third-party loggers and the number of tools available for logging is overwhelming and spamming.
This library gives you access to an easy and maintainable set of logging interfaces to play with.

### Design

Follows factory design pattern to adhere to your logging needs.
Can be configured to work with Zap and Logrous currently more integration on the way.

### Performance vs Stability
Logrus is a widely used logging tool and has been around for a long time but it's also in the maintenance phase.
* Logrus also uses an unstructured logging approach which makes performance weaker as compared to other tools (zap,zerolog).

Zap on the other hand uses multi-level optimizations
* It uses a strongly typed structure instead of interfaces which allows it to be reflection-free and allocated lesser memory.
* Refer to benchmarking section. [Benchmarks](https://github.com/uber-go/zap)


### Best Practises
* Always Log error on the main routine, try to trickle errors down to the main process.
* Info logs can be used anywhere in the code.
* Use golLangs internal log library to handle errors while building log module or any other module used at the time of booting up the application.
* Always try to log in standard output (stdout, stderr), since this will allow all the log aggregators to pick your logs from standard predefined locations.
* Avoid writing logs to custom files, unless necessary.
* Standardize your logs with a set of predefined fields and when possible use JSON formatter for better compatibility with modern log aggregators.
* Generate and add transaction IDs for connected applications, especially in the case of distributed microservices as it gives better visibility of logs.
* Do not log sensitive information like credit card number, IP address, location, etc, since it creates issues in GDPR and other compliance.


### Run the code
* Try Running test cases for available products of log factory:
```
{"caller":"github.com/pranaySinghDev/goSAK/logger.TestLogrusLoggerFactory","level":"info","line":"logger_test.go/logger:20","msg":"info logrus","ts":"2021-03-29T00:38:48.9304952+05:30"}
{"level":"info","ts":"2021-03-29T00:38:52.720277+05:30","line":"/mnt/d/workspace/goSAK/logger/logger_test.go:33","caller":"github.com/pranaySinghDev/goSAK/logger.TestZapLoggerFactory","msg":"info zap"}

```
* How to add this logger to your module

Create your config    
```
cfg := &config.LogConfig{
		Type:     config.Logrus,
		Level:    "DEBUG",
		Detailed: true,
	}
```
It will build your logger with standard configuration
```
	myLogger, err := Build(cfg)
	if err != nil {
		log.Fatalf("Init Logrus failed %v", err)
	}
```
Inject this logger as a dependency to your service or use it directly
```
myLogger.Infof("May the logs be with you !")
```