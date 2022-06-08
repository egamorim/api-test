pipeline {
    agent {
        label 'docker' 
    }
    tools {
        go 'go'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }

    stages {
        stage('Compile') {
            steps {
                sh 'go build cmd/api/.'
            }
        }
        stage('Build') {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'docker build . -t egamorim/api-test'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}