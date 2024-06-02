pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                powershell './build.sh'
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