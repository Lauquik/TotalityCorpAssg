pipeline {
    agent any
    stages {
        stage('Install') {
            steps {
                powershell 'mkdir ./api/pb'
                powershell 'protoc --proto_path=./api/proto --go_out=./api/pb --go_opt=module=github.com/lavquik/totality/api/pb --go-grpc_out=./api/pb --go-grpc_opt=module=github.com/lavquik/totality/api/pb ./api/proto/user.proto'
            }
        }
        stage('Build') {
            steps {
                powershell 'go build -o build/userService ./cmd'
            }
        }
        stage('Test'){
            steps {
                powershell 'go test ./... -v'
            }
        }
    }
    post {
        success {
            echo 'This will run only if successful'
        }
        failure {
            echo 'This will run only if failed'
        }
    }
}