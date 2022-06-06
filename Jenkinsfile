pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
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