//https://gist.github.com/arehmandev/736daba40a3e1ef1fbe939c6674d7da8
// getting jenkins dsl to work with idea

//https://st-g.de/2016/08/jenkins-pipeline-autocompletion-in-intellij
//http://localhost:8080/job/go_alfred/pipeline-syntax/gdsl << GET THE DSL

// https://confluence.jetbrains.com/display/GRVY/Scripting+IDE+for+DSL+awareness << creating custom groovy dsl
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