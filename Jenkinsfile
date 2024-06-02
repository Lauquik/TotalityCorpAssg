pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                bash './build.sh'
            }
        }
        stage('Test'){
            steps {
                bash 'go test ./... -v'
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