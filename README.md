# Dynamodb connect test

This project's intent is to measure the time taken by AWS dynamodb sdk to connect with the dynamodb instance.
Corresponding [issue link](https://github.com/aws/aws-sdk-go/issues/2624)

## Setup

1. I am connecting from india to an aws dynamodb instance in ireland.
2. For local, I am using dynamodb local with the below config.
```bash
docker run --rm -d -p 4569:8000 amazon/dynamodb-local -jar DynamoDBLocal.jar -inMemory -sharedDb
```

## Result Summary

- Aws Instance   : ~11 seconds
- Local Instance :  ~4 seconds

## Result

### Try 1

Starting AWS instance first
```bash
2019/05/26 17:47:24 Using AWS instance
2019/05/26 17:47:24 Request 1: Start
2019/05/26 17:47:36 Request 1: Complete: Total no. of tables: 16
2019/05/26 17:47:36 Request 2: Start
2019/05/26 17:47:36 Request 2: Complete: Total no. of tables: 16
2019/05/26 17:47:36
2019/05/26 17:47:36 Using local instance
2019/05/26 17:47:36 Request 1: Start
2019/05/26 17:47:37 Request 1: Complete: Total no. of tables: 18
2019/05/26 17:47:37 Request 2: Start
2019/05/26 17:47:37 Request 2: Complete: Total no. of tables: 18
```

Starting local instance first
```bash
2019/05/26 17:48:31 Using local instance
2019/05/26 17:48:31 Request 1: Start
2019/05/26 17:48:35 Request 1: Complete: Total no. of tables: 18
2019/05/26 17:48:35 Request 2: Start
2019/05/26 17:48:36 Request 2: Complete: Total no. of tables: 18
2019/05/26 17:48:36
2019/05/26 17:48:36 Using AWS instance
2019/05/26 17:48:36 Request 1: Start
2019/05/26 17:48:39 Request 1: Complete: Total no. of tables: 16
2019/05/26 17:48:39 Request 2: Start
2019/05/26 17:48:39 Request 2: Complete: Total no. of tables: 16
```

### Try 2

Starting AWS instance first
```bash
2019/05/26 21:15:21 Using AWS instance
2019/05/26 21:15:21 Request 1: Start
2019/05/26 21:15:33 Request 1: Complete: Total no. of tables: 18
2019/05/26 21:15:33 Request 2: Start
2019/05/26 21:15:34 Request 2: Complete: Total no. of tables: 18
2019/05/26 21:15:34
2019/05/26 21:15:34 Using local instance
2019/05/26 21:15:34 Request 1: Start
2019/05/26 21:15:34 Request 1: Complete: Total no. of tables: 18
2019/05/26 21:15:34 Request 2: Start
2019/05/26 21:15:34 Request 2: Complete: Total no. of tables: 18
```

Starting local instance first
```bash
2019/05/26 21:13:17 Using local instance
2019/05/26 21:13:17 Request 1: Start
2019/05/26 21:13:20 Request 1: Complete: Total no. of tables: 18
2019/05/26 21:13:20 Request 2: Start
2019/05/26 21:13:20 Request 2: Complete: Total no. of tables: 18
2019/05/26 21:13:20
2019/05/26 21:13:20 Using AWS instance
2019/05/26 21:13:20 Request 1: Start
2019/05/26 21:13:24 Request 1: Complete: Total no. of tables: 18
2019/05/26 21:13:24 Request 2: Start
2019/05/26 21:13:24 Request 2: Complete: Total no. of tables: 18
```

### Try 3

Starting AWS instance first
```bash
2019/05/26 21:16:11 Using AWS instance
2019/05/26 21:16:11 Request 1: Start
2019/05/26 21:16:20 Request 1: Complete: Total no. of tables: 18
2019/05/26 21:16:20 Request 2: Start
2019/05/26 21:16:20 Request 2: Complete: Total no. of tables: 18
2019/05/26 21:16:20
2019/05/26 21:16:20 Using local instance
2019/05/26 21:16:20 Request 1: Start
2019/05/26 21:16:21 Request 1: Complete: Total no. of tables: 18
2019/05/26 21:16:21 Request 2: Start
2019/05/26 21:16:21 Request 2: Complete: Total no. of tables: 18
```

Starting local instance first
```bash
2019/05/26 21:14:34 Using local instance
2019/05/26 21:14:34 Request 1: Start
2019/05/26 21:14:38 Request 1: Complete: Total no. of tables: 18
2019/05/26 21:14:38 Request 2: Start
2019/05/26 21:14:38 Request 2: Complete: Total no. of tables: 18
2019/05/26 21:14:38
2019/05/26 21:14:38 Using AWS instance
2019/05/26 21:14:38 Request 1: Start
2019/05/26 21:14:41 Request 1: Complete: Total no. of tables: 18
2019/05/26 21:14:41 Request 2: Start
2019/05/26 21:14:41 Request 2: Complete: Total no. of tables: 18
```