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
                expression {
                    return env.CHANGE_TARGET == 'master';
                }
            }
            steps {
                echo "Running api tests..."
            }
        }

        stage('Clean Up') {
            steps {
                echo 'Complete!'
            }
        }
    }
}