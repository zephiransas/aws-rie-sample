# aws-rie-sample

aws-lambda-runtime-interface-emulatorを使い、ローカルでAWS LambdaとAPI Gatewayを使った開発ができるようにするサンプル

- Go
  - 1.20.0
- Docker
- [aws-lambda-runtime-interface-emulator](https://github.com/aws/aws-lambda-runtime-interface-emulator)
- [local-api-gateway](https://github.com/brefphp/local-api-gateway)
  - aws-lambda-rieへのリクエストを、APIGatewayのプロキシ統合リクエストにラップする
