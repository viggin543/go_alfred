pipeline {
    agent any
    environment {
        BANANA = 'banana'
    }
    stages {
        stage('Build') {
            steps {
                sh 'echo $BANANA'
                sh 'ls -la'
                sh 'pwd'
            }
        }
        stage('Test') {
            steps {
                echo 'this is a message $BANANA'
            }
        }
    }
}