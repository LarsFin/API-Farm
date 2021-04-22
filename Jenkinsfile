pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh 'printenv'
                echo 'Building...'
            }
        }

        stage('Test') {
            steps {
                echo 'Testing...'
            }
        }

        stage('Prod Check') {
            when {
                branch 'ruby/sinatra_patch_env'
            }
            steps {
                echo 'Running API Tests!'
            }
        }

        stage('Clean Up') {
            steps {
                echo 'Complete!'
            }
        }
    }
}