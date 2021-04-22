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
            if (env.CHANGE_TARGET == 'master') {
                echo "Running api tests..."
            } else {
                echo "Skipping api tests are target is not master."
            }
        }

        stage('Clean Up') {
            steps {
                echo 'Complete!'
            }
        }
    }
}